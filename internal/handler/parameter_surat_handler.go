package handler

import (
	"net/http"
	"strconv"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllParameterSurat(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			ps.id,
			ps.template_id,
			ts.template_name,
			ps.parameter_name,
			ps.parameter_key,
			ps.input_type,
			ps.is_required,
			ps.created_at
		FROM parameter_surat ps
		JOIN template_surat ts ON ps.template_id = ts.id
		ORDER BY ps.id ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data parameter surat",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var parameterList []model.ParameterSuratResponse

	for rows.Next() {
		var parameter model.ParameterSuratResponse
		var createdAt time.Time

		err := rows.Scan(
			&parameter.ID,
			&parameter.TemplateID,
			&parameter.TemplateName,
			&parameter.ParameterName,
			&parameter.ParameterKey,
			&parameter.InputType,
			&parameter.IsRequired,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data parameter surat",
				"error":   err.Error(),
			})
			return
		}

		parameter.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		parameterList = append(parameterList, parameter)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data parameter surat berhasil diambil",
		"data":    parameterList,
	})
}

func GetParameterSuratByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter surat tidak valid",
		})
		return
	}

	var parameter model.ParameterSuratResponse
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			ps.id,
			ps.template_id,
			ts.template_name,
			ps.parameter_name,
			ps.parameter_key,
			ps.input_type,
			ps.is_required,
			ps.created_at
		FROM parameter_surat ps
		JOIN template_surat ts ON ps.template_id = ts.id
		WHERE ps.id = ?
	`, id).Scan(
		&parameter.ID,
		&parameter.TemplateID,
		&parameter.TemplateName,
		&parameter.ParameterName,
		&parameter.ParameterKey,
		&parameter.InputType,
		&parameter.IsRequired,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data parameter surat tidak ditemukan",
		})
		return
	}

	parameter.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data parameter surat berhasil diambil",
		"data":    parameter,
	})
}

func CreateParameterSurat(c *gin.Context) {
	var request model.ParameterSuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data parameter surat wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var templateName string
	err := database.DB.QueryRow(`
		SELECT template_name FROM template_surat WHERE id = ?
	`, request.TemplateID).Scan(&templateName)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Template surat tidak ditemukan",
		})
		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO parameter_surat
			(template_id, parameter_name, parameter_key, input_type, is_required)
		VALUES
			(?, ?, ?, ?, ?)
	`,
		request.TemplateID,
		request.ParameterName,
		request.ParameterKey,
		request.InputType,
		request.IsRequired,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data parameter surat",
			"error":   err.Error(),
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Data berhasil dibuat, tetapi ID gagal dibaca",
			"error":   err.Error(),
		})
		return
	}

	response := model.ParameterSuratResponse{
		ID:            int(id),
		TemplateID:    request.TemplateID,
		TemplateName:  templateName,
		ParameterName: request.ParameterName,
		ParameterKey:  request.ParameterKey,
		InputType:     request.InputType,
		IsRequired:    request.IsRequired,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data parameter surat berhasil dibuat",
		"data":    response,
	})
}

func UpdateParameterSurat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter surat tidak valid",
		})
		return
	}

	var request model.ParameterSuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data parameter surat wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var templateName string
	err = database.DB.QueryRow(`
		SELECT template_name FROM template_surat WHERE id = ?
	`, request.TemplateID).Scan(&templateName)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Template surat tidak ditemukan",
		})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE parameter_surat
		SET
			template_id = ?,
			parameter_name = ?,
			parameter_key = ?,
			input_type = ?,
			is_required = ?
		WHERE id = ?
	`,
		request.TemplateID,
		request.ParameterName,
		request.ParameterKey,
		request.InputType,
		request.IsRequired,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data parameter surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data parameter surat tidak ditemukan",
		})
		return
	}

	response := model.ParameterSuratResponse{
		ID:            id,
		TemplateID:    request.TemplateID,
		TemplateName:  templateName,
		ParameterName: request.ParameterName,
		ParameterKey:  request.ParameterKey,
		InputType:     request.InputType,
		IsRequired:    request.IsRequired,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data parameter surat berhasil diupdate",
		"data":    response,
	})
}

func DeleteParameterSurat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID parameter surat tidak valid",
		})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM parameter_surat
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data parameter surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data parameter surat tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data parameter surat berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}