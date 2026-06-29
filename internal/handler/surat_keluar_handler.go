package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/helper"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllSuratKeluar(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			sk.id, sk.nomor_surat, sk.takah_id, mt.code,
			sk.tujuan_surat, sk.perihal, sk.lampiran,
			sk.tanggal_surat, sk.file_surat, sk.status,
			sk.created_by, sk.created_at
		FROM surat_keluar sk
		JOIN master_takah mt ON sk.takah_id = mt.id
		ORDER BY sk.id DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data surat keluar", "error": err.Error()})
		return
	}
	defer rows.Close()

	var suratList []model.SuratKeluarResponse

	for rows.Next() {
		var surat model.SuratKeluarResponse
		var lampiran, fileSurat sql.NullString
		var createdBy sql.NullInt64
		var tanggalSurat time.Time
		var createdAt time.Time

		err := rows.Scan(
			&surat.ID,
			&surat.NomorSurat,
			&surat.TakahID,
			&surat.TakahCode,
			&surat.TujuanSurat,
			&surat.Perihal,
			&lampiran,
			&tanggalSurat,
			&fileSurat,
			&surat.Status,
			&createdBy,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca data surat keluar", "error": err.Error()})
			return
		}

		if lampiran.Valid {
			surat.Lampiran = lampiran.String
		}
		if fileSurat.Valid {
			surat.FileSurat = fileSurat.String
		}
		if createdBy.Valid {
			surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
		}

		surat.TanggalSurat = tanggalSurat.Format("2006-01-02")
		surat.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

		suratList = append(suratList, surat)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil diambil",
		"data":    suratList,
	})
}

func GetSuratKeluarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	var surat model.SuratKeluarResponse
	var lampiran, fileSurat sql.NullString
	var createdBy sql.NullInt64
	var tanggalSurat time.Time
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			sk.id, sk.nomor_surat, sk.takah_id, mt.code,
			sk.tujuan_surat, sk.perihal, sk.lampiran,
			sk.tanggal_surat, sk.file_surat, sk.status,
			sk.created_by, sk.created_at
		FROM surat_keluar sk
		JOIN master_takah mt ON sk.takah_id = mt.id
		WHERE sk.id = ?
	`, id).Scan(
		&surat.ID,
		&surat.NomorSurat,
		&surat.TakahID,
		&surat.TakahCode,
		&surat.TujuanSurat,
		&surat.Perihal,
		&lampiran,
		&tanggalSurat,
		&fileSurat,
		&surat.Status,
		&createdBy,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
		return
	}

	if lampiran.Valid {
		surat.Lampiran = lampiran.String
	}
	if fileSurat.Valid {
		surat.FileSurat = fileSurat.String
	}
	if createdBy.Valid {
		surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
	}

	surat.TanggalSurat = tanggalSurat.Format("2006-01-02")
	surat.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil diambil",
		"data":    surat,
	})
}

func CreateSuratKeluar(c *gin.Context) {
	var request model.SuratKeluarRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat keluar wajib diisi dengan benar", "error": err.Error()})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memulai transaksi", "error": err.Error()})
		return
	}
	defer tx.Rollback()

	var configID int
	var takahCode string
	var companyCode string
	var resetType string
	var lastNumber int

	err = tx.QueryRow(`
		SELECT c.id, mt.code, c.company_code, c.reset_type, c.last_number
		FROM config_nomor_surat c
		JOIN master_takah mt ON c.takah_id = mt.id
		WHERE c.takah_id = ?
		FOR UPDATE
	`, request.TakahID).Scan(
		&configID,
		&takahCode,
		&companyCode,
		&resetType,
		&lastNumber,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Config nomor surat untuk takah ini tidak ditemukan"})
		return
	}

	nomorSurat := helper.GenerateNomorSurat(lastNumber, takahCode, companyCode, resetType)
	nextNumber := lastNumber + 1

	_, err = tx.Exec(`
		UPDATE config_nomor_surat
		SET last_number = ?
		WHERE id = ?
	`, nextNumber, configID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate nomor terakhir", "error": err.Error()})
		return
	}

	result, err := tx.Exec(`
		INSERT INTO surat_keluar
			(nomor_surat, takah_id, tujuan_surat, perihal, lampiran, tanggal_surat, file_surat, status, created_by)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		nomorSurat,
		request.TakahID,
		request.TujuanSurat,
		request.Perihal,
		request.Lampiran,
		request.TanggalSurat,
		request.FileSurat,
		"draft",
		1,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data surat keluar", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data berhasil dibuat, tetapi ID gagal dibaca", "error": err.Error()})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan transaksi", "error": err.Error()})
		return
	}

	response := model.SuratKeluarResponse{
		ID:           int(id),
		NomorSurat:  nomorSurat,
		TakahID:      request.TakahID,
		TakahCode:    takahCode,
		TujuanSurat:  request.TujuanSurat,
		Perihal:      request.Perihal,
		Lampiran:     request.Lampiran,
		TanggalSurat: request.TanggalSurat,
		FileSurat:    request.FileSurat,
		Status:       "draft",
		CreatedBy:    "1",
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data surat keluar berhasil dibuat sebagai draft",
		"data":    response,
	})
}

func UpdateSuratKeluar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	var request model.SuratKeluarRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat keluar wajib diisi dengan benar", "error": err.Error()})
		return
	}

	var takahCode string
	err = database.DB.QueryRow(`
		SELECT code FROM master_takah WHERE id = ?
	`, request.TakahID).Scan(&takahCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Master takah tidak ditemukan"})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE surat_keluar
		SET 
			takah_id = ?,
			tujuan_surat = ?,
			perihal = ?,
			lampiran = ?,
			tanggal_surat = ?,
			file_surat = ?
		WHERE id = ?
	`,
		request.TakahID,
		request.TujuanSurat,
		request.Perihal,
		request.Lampiran,
		request.TanggalSurat,
		request.FileSurat,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate data surat keluar", "error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil diupdate",
		"data": gin.H{
			"id":            id,
			"takah_id":      request.TakahID,
			"takah_code":    takahCode,
			"tujuan_surat":  request.TujuanSurat,
			"perihal":       request.Perihal,
			"lampiran":      request.Lampiran,
			"tanggal_surat": request.TanggalSurat,
			"file_surat":    request.FileSurat,
		},
	})
}

func DeleteSuratKeluar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM surat_keluar
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data surat keluar", "error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil dihapus",
		"data":    gin.H{"id": id},
	})
}