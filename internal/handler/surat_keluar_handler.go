package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/helper"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func generateContent(templateContent string, nomorSurat string, request model.SuratKeluarRequest) string {
	content := templateContent

	replacements := map[string]string{
		"nomor_surat":   nomorSurat,
		"tujuan_surat":  request.TujuanSurat,
		"perihal":       request.Perihal,
		"lampiran":      request.Lampiran,
		"tanggal_surat": request.TanggalSurat,
	}

	for key, value := range request.ParameterValues {
		replacements[key] = value
	}

	for key, value := range replacements {
		content = strings.ReplaceAll(content, "{{"+key+"}}", value)
	}

	return content
}

func parseParameterValues(value sql.NullString) map[string]string {
	result := map[string]string{}

	if !value.Valid || value.String == "" {
		return result
	}

	_ = json.Unmarshal([]byte(value.String), &result)
	return result
}

func GetAllSuratKeluar(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			sk.id,
			sk.nomor_surat,
			sk.template_id,
			COALESCE(ts.template_name, ''),
			sk.takah_id,
			mt.code,
			sk.tujuan_surat,
			sk.perihal,
			sk.lampiran,
			sk.tanggal_surat,
			sk.file_surat,
			sk.parameter_values,
			sk.generated_content,
			sk.status,
			sk.created_by,
			sk.created_at
		FROM surat_keluar sk
		JOIN master_takah mt ON sk.takah_id = mt.id
		LEFT JOIN template_surat ts ON sk.template_id = ts.id
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
		var templateID sql.NullInt64
		var templateName sql.NullString
		var lampiran, fileSurat, parameterValues, generatedContent sql.NullString
		var createdBy sql.NullInt64
		var tanggalSurat time.Time
		var createdAt time.Time

		err := rows.Scan(
			&surat.ID,
			&surat.NomorSurat,
			&templateID,
			&templateName,
			&surat.TakahID,
			&surat.TakahCode,
			&surat.TujuanSurat,
			&surat.Perihal,
			&lampiran,
			&tanggalSurat,
			&fileSurat,
			&parameterValues,
			&generatedContent,
			&surat.Status,
			&createdBy,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca data surat keluar", "error": err.Error()})
			return
		}

		if templateID.Valid {
			surat.TemplateID = int(templateID.Int64)
		}
		if templateName.Valid {
			surat.TemplateName = templateName.String
		}
		if lampiran.Valid {
			surat.Lampiran = lampiran.String
		}
		if fileSurat.Valid {
			surat.FileSurat = fileSurat.String
		}
		if generatedContent.Valid {
			surat.GeneratedContent = generatedContent.String
		}
		if createdBy.Valid {
			surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
		}

		surat.ParameterValues = parseParameterValues(parameterValues)
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
	var templateID sql.NullInt64
	var templateName sql.NullString
	var lampiran, fileSurat, parameterValues, generatedContent sql.NullString
	var createdBy sql.NullInt64
	var tanggalSurat time.Time
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			sk.id,
			sk.nomor_surat,
			sk.template_id,
			COALESCE(ts.template_name, ''),
			sk.takah_id,
			mt.code,
			sk.tujuan_surat,
			sk.perihal,
			sk.lampiran,
			sk.tanggal_surat,
			sk.file_surat,
			sk.parameter_values,
			sk.generated_content,
			sk.status,
			sk.created_by,
			sk.created_at
		FROM surat_keluar sk
		JOIN master_takah mt ON sk.takah_id = mt.id
		LEFT JOIN template_surat ts ON sk.template_id = ts.id
		WHERE sk.id = ?
	`, id).Scan(
		&surat.ID,
		&surat.NomorSurat,
		&templateID,
		&templateName,
		&surat.TakahID,
		&surat.TakahCode,
		&surat.TujuanSurat,
		&surat.Perihal,
		&lampiran,
		&tanggalSurat,
		&fileSurat,
		&parameterValues,
		&generatedContent,
		&surat.Status,
		&createdBy,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
		return
	}

	if templateID.Valid {
		surat.TemplateID = int(templateID.Int64)
	}
	if templateName.Valid {
		surat.TemplateName = templateName.String
	}
	if lampiran.Valid {
		surat.Lampiran = lampiran.String
	}
	if fileSurat.Valid {
		surat.FileSurat = fileSurat.String
	}
	if generatedContent.Valid {
		surat.GeneratedContent = generatedContent.String
	}
	if createdBy.Valid {
		surat.CreatedBy = strconv.FormatInt(createdBy.Int64, 10)
	}

	surat.ParameterValues = parseParameterValues(parameterValues)
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

	var templateName string
	var templateContent string
	var templateTakahID int

	err = tx.QueryRow(`
		SELECT template_name, content, takah_id
		FROM template_surat
		WHERE id = ?
	`, request.TemplateID).Scan(&templateName, &templateContent, &templateTakahID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Template surat tidak ditemukan"})
		return
	}

	if templateTakahID != request.TakahID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Template surat tidak sesuai dengan jenis surat"})
		return
	}

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

	generatedContent := generateContent(templateContent, nomorSurat, request)

	parameterValuesJSON, err := json.Marshal(request.ParameterValues)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memproses parameter surat", "error": err.Error()})
		return
	}

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
			(nomor_surat, takah_id, template_id, tujuan_surat, perihal, lampiran, tanggal_surat, file_surat, parameter_values, generated_content, status, created_by)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		nomorSurat,
		request.TakahID,
		request.TemplateID,
		request.TujuanSurat,
		request.Perihal,
		request.Lampiran,
		request.TanggalSurat,
		request.FileSurat,
		string(parameterValuesJSON),
		generatedContent,
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
		ID:               int(id),
		NomorSurat:       nomorSurat,
		TemplateID:       request.TemplateID,
		TemplateName:     templateName,
		TakahID:          request.TakahID,
		TakahCode:        takahCode,
		TujuanSurat:      request.TujuanSurat,
		Perihal:          request.Perihal,
		Lampiran:         request.Lampiran,
		TanggalSurat:     request.TanggalSurat,
		FileSurat:        request.FileSurat,
		GeneratedContent: generatedContent,
		ParameterValues:  request.ParameterValues,
		Status:           "draft",
		CreatedBy:        "1",
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
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

	var nomorSurat string
	err = database.DB.QueryRow(`
		SELECT nomor_surat
		FROM surat_keluar
		WHERE id = ?
	`, id).Scan(&nomorSurat)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
		return
	}

	var templateName string
	var templateContent string
	var templateTakahID int
	var takahCode string

	err = database.DB.QueryRow(`
		SELECT ts.template_name, ts.content, ts.takah_id, mt.code
		FROM template_surat ts
		JOIN master_takah mt ON ts.takah_id = mt.id
		WHERE ts.id = ?
	`, request.TemplateID).Scan(&templateName, &templateContent, &templateTakahID, &takahCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Template surat tidak ditemukan"})
		return
	}

	if templateTakahID != request.TakahID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Template surat tidak sesuai dengan jenis surat"})
		return
	}

	generatedContent := generateContent(templateContent, nomorSurat, request)

	parameterValuesJSON, err := json.Marshal(request.ParameterValues)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memproses parameter surat", "error": err.Error()})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE surat_keluar
		SET 
			takah_id = ?,
			template_id = ?,
			tujuan_surat = ?,
			perihal = ?,
			lampiran = ?,
			tanggal_surat = ?,
			file_surat = ?,
			parameter_values = ?,
			generated_content = ?
		WHERE id = ?
	`,
		request.TakahID,
		request.TemplateID,
		request.TujuanSurat,
		request.Perihal,
		request.Lampiran,
		request.TanggalSurat,
		request.FileSurat,
		string(parameterValuesJSON),
		generatedContent,
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
		"data": model.SuratKeluarResponse{
			ID:               id,
			NomorSurat:       nomorSurat,
			TemplateID:       request.TemplateID,
			TemplateName:     templateName,
			TakahID:          request.TakahID,
			TakahCode:        takahCode,
			TujuanSurat:      request.TujuanSurat,
			Perihal:          request.Perihal,
			Lampiran:         request.Lampiran,
			TanggalSurat:     request.TanggalSurat,
			FileSurat:        request.FileSurat,
			GeneratedContent: generatedContent,
			ParameterValues:  request.ParameterValues,
		},
	})
}

func SubmitSuratKeluarApproval(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memulai transaksi", "error": err.Error()})
		return
	}
	defer tx.Rollback()

	var nomorSurat string
	var status string

	err = tx.QueryRow(`
		SELECT nomor_surat, status
		FROM surat_keluar
		WHERE id = ?
	`, id).Scan(&nomorSurat, &status)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Surat keluar tidak ditemukan"})
		return
	}

	if status != "draft" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Surat hanya bisa diajukan jika status masih draft"})
		return
	}

	_, err = tx.Exec(`
		INSERT INTO approval_surat
			(surat_keluar_id, approver_id, approver_name, approval_status, notes)
		VALUES
			(?, ?, ?, ?, ?)
	`, id, 1, "Admin Takah", "pending", "")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data approval", "error": err.Error()})
		return
	}

	_, err = tx.Exec(`
		UPDATE surat_keluar
		SET status = 'pending'
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate status surat", "error": err.Error()})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan transaksi", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Surat berhasil diajukan untuk approval",
		"data": gin.H{
			"id":          id,
			"nomor_surat": nomorSurat,
			"status":      "pending",
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