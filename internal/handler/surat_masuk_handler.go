package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllSuratMasuk(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			id, nomor_surat, pengirim, penerima, perihal, file_surat,
			tanggal_surat, keterangan, status, created_by, created_at, updated_at
		FROM surat_masuk
		ORDER BY id DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data surat masuk", "error": err.Error()})
		return
	}
	defer rows.Close()

	var list []model.SuratMasukResponse

	for rows.Next() {
		var surat model.SuratMasukResponse
		var fileSurat, keterangan sql.NullString
		var createdBy sql.NullInt64
		var tanggalSurat, createdAt, updatedAt time.Time

		err := rows.Scan(
			&surat.ID,
			&surat.NomorSurat,
			&surat.Pengirim,
			&surat.Penerima,
			&surat.Perihal,
			&fileSurat,
			&tanggalSurat,
			&keterangan,
			&surat.Status,
			&createdBy,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca data surat masuk", "error": err.Error()})
			return
		}

		if fileSurat.Valid {
			surat.FileSurat = fileSurat.String
		}
		if keterangan.Valid {
			surat.Keterangan = keterangan.String
		}
		if createdBy.Valid {
			surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
		}

		surat.TanggalSurat = tanggalSurat.Format("2006-01-02")
		surat.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		surat.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

		list = append(list, surat)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil diambil",
		"data":    list,
	})
}

func GetSuratMasukByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat masuk tidak valid"})
		return
	}

	var surat model.SuratMasukResponse
	var fileSurat, keterangan sql.NullString
	var createdBy sql.NullInt64
	var tanggalSurat, createdAt, updatedAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			id, nomor_surat, pengirim, penerima, perihal, file_surat,
			tanggal_surat, keterangan, status, created_by, created_at, updated_at
		FROM surat_masuk
		WHERE id = ?
	`, id).Scan(
		&surat.ID,
		&surat.NomorSurat,
		&surat.Pengirim,
		&surat.Penerima,
		&surat.Perihal,
		&fileSurat,
		&tanggalSurat,
		&keterangan,
		&surat.Status,
		&createdBy,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat masuk tidak ditemukan"})
		return
	}

	if fileSurat.Valid {
		surat.FileSurat = fileSurat.String
	}
	if keterangan.Valid {
		surat.Keterangan = keterangan.String
	}
	if createdBy.Valid {
		surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
	}

	surat.TanggalSurat = tanggalSurat.Format("2006-01-02")
	surat.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	surat.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil diambil",
		"data":    surat,
	})
}

func CreateSuratMasuk(c *gin.Context) {
	var request model.SuratMasukRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat masuk wajib diisi dengan benar", "error": err.Error()})
		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO surat_masuk
			(nomor_surat, pengirim, penerima, perihal, file_surat, tanggal_surat, keterangan, status, created_by)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		request.NomorSurat,
		request.Pengirim,
		request.Penerima,
		request.Perihal,
		request.FileSurat,
		request.TanggalSurat,
		request.Keterangan,
		"received",
		1,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data surat masuk", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data berhasil dibuat, tetapi ID gagal dibaca", "error": err.Error()})
		return
	}

	response := model.SuratMasukResponse{
		ID:            int(id),
		NomorSurat:   request.NomorSurat,
		Pengirim:     request.Pengirim,
		Penerima:     request.Penerima,
		Perihal:      request.Perihal,
		FileSurat:    request.FileSurat,
		TanggalSurat: request.TanggalSurat,
		Keterangan:   request.Keterangan,
		Status:       "received",
		CreatedBy:    "1",
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data surat masuk berhasil dibuat",
		"data":    response,
	})
}

func UpdateSuratMasuk(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat masuk tidak valid"})
		return
	}

	var request model.SuratMasukRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat masuk wajib diisi dengan benar", "error": err.Error()})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE surat_masuk
		SET
			nomor_surat = ?,
			pengirim = ?,
			penerima = ?,
			perihal = ?,
			file_surat = ?,
			tanggal_surat = ?,
			keterangan = ?
		WHERE id = ?
	`,
		request.NomorSurat,
		request.Pengirim,
		request.Penerima,
		request.Perihal,
		request.FileSurat,
		request.TanggalSurat,
		request.Keterangan,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate data surat masuk", "error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat masuk tidak ditemukan"})
		return
	}

	response := model.SuratMasukResponse{
		ID:            id,
		NomorSurat:   request.NomorSurat,
		Pengirim:     request.Pengirim,
		Penerima:     request.Penerima,
		Perihal:      request.Perihal,
		FileSurat:    request.FileSurat,
		TanggalSurat: request.TanggalSurat,
		Keterangan:   request.Keterangan,
		Status:       "received",
		CreatedBy:    "1",
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil diupdate",
		"data":    response,
	})
}

func DeleteSuratMasuk(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat masuk tidak valid"})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM surat_masuk
		WHERE id = ?
	`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data surat masuk", "error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat masuk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil dihapus",
		"data":    gin.H{"id": id},
	})
}