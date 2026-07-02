package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func isAllowedSuratMasukFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	allowedExtensions := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}

	return allowedExtensions[ext]
}

func saveSuratMasukFile(c *gin.Context, fieldName string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", nil
	}

	if !isAllowedSuratMasukFile(file.Filename) {
		return "", fmt.Errorf("file harus berformat PDF, DOC, atau DOCX")
	}

	uploadDir := "uploads/surat-masuk"

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("surat_masuk_%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return "/" + filepath.ToSlash(filePath), nil
}

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
	nomorSurat := c.PostForm("nomor_surat")
	pengirim := c.PostForm("pengirim")
	penerima := c.PostForm("penerima")
	perihal := c.PostForm("perihal")
	tanggalSurat := c.PostForm("tanggal_surat")
	keterangan := c.PostForm("keterangan")

	if nomorSurat == "" || pengirim == "" || penerima == "" || perihal == "" || tanggalSurat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nomor surat, pengirim, penerima, perihal, dan tanggal surat wajib diisi"})
		return
	}

	fileSurat, err := saveSuratMasukFile(c, "file_surat")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdBy := c.GetInt("user_id")
	if createdBy == 0 {
		createdBy = 1
	}

	result, err := database.DB.Exec(`
		INSERT INTO surat_masuk
			(nomor_surat, pengirim, penerima, perihal, file_surat, tanggal_surat, keterangan, status, created_by)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		nomorSurat,
		pengirim,
		penerima,
		perihal,
		fileSurat,
		tanggalSurat,
		keterangan,
		"received",
		createdBy,
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
		NomorSurat:   nomorSurat,
		Pengirim:     pengirim,
		Penerima:     penerima,
		Perihal:      perihal,
		FileSurat:    fileSurat,
		TanggalSurat: tanggalSurat,
		Keterangan:   keterangan,
		Status:       "received",
		CreatedBy:    strconv.Itoa(createdBy),
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

	nomorSurat := c.PostForm("nomor_surat")
	pengirim := c.PostForm("pengirim")
	penerima := c.PostForm("penerima")
	perihal := c.PostForm("perihal")
	tanggalSurat := c.PostForm("tanggal_surat")
	keterangan := c.PostForm("keterangan")

	if nomorSurat == "" || pengirim == "" || penerima == "" || perihal == "" || tanggalSurat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nomor surat, pengirim, penerima, perihal, dan tanggal surat wajib diisi"})
		return
	}

	var oldFileSurat sql.NullString

	err = database.DB.QueryRow(`
		SELECT file_surat
		FROM surat_masuk
		WHERE id = ?
	`, id).Scan(&oldFileSurat)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat masuk tidak ditemukan"})
		return
	}

	fileSurat := ""
	if oldFileSurat.Valid {
		fileSurat = oldFileSurat.String
	}

	newFileSurat, err := saveSuratMasukFile(c, "file_surat")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if newFileSurat != "" {
		fileSurat = newFileSurat
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
		nomorSurat,
		pengirim,
		penerima,
		perihal,
		fileSurat,
		tanggalSurat,
		keterangan,
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
		NomorSurat:   nomorSurat,
		Pengirim:     pengirim,
		Penerima:     penerima,
		Perihal:      perihal,
		FileSurat:    fileSurat,
		TanggalSurat: tanggalSurat,
		Keterangan:   keterangan,
		Status:       "received",
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