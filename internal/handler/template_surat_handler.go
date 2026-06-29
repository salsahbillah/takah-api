package handler

import (
	"net/http"
	"strconv"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllTemplateSurat(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			ts.id,
			ts.takah_id,
			mt.code,
			ts.template_name,
			ts.content,
			ts.created_at
		FROM template_surat ts
		JOIN master_takah mt ON ts.takah_id = mt.id
		ORDER BY ts.id ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data template surat",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var templateList []model.TemplateSuratResponse

	for rows.Next() {
		var template model.TemplateSuratResponse
		var createdAt time.Time

		err := rows.Scan(
			&template.ID,
			&template.TakahID,
			&template.TakahCode,
			&template.TemplateName,
			&template.Content,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data template surat",
				"error":   err.Error(),
			})
			return
		}

		template.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		templateList = append(templateList, template)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data template surat berhasil diambil",
		"data":    templateList,
	})
}

func GetTemplateSuratByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID template surat tidak valid",
		})
		return
	}

	var template model.TemplateSuratResponse
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			ts.id,
			ts.takah_id,
			mt.code,
			ts.template_name,
			ts.content,
			ts.created_at
		FROM template_surat ts
		JOIN master_takah mt ON ts.takah_id = mt.id
		WHERE ts.id = ?
	`, id).Scan(
		&template.ID,
		&template.TakahID,
		&template.TakahCode,
		&template.TemplateName,
		&template.Content,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data template surat tidak ditemukan",
		})
		return
	}

	template.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data template surat berhasil diambil",
		"data":    template,
	})
}

func CreateTemplateSurat(c *gin.Context) {
	var request model.TemplateSuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data template surat wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var takahCode string
	err := database.DB.QueryRow(`
		SELECT code FROM master_takah WHERE id = ?
	`, request.TakahID).Scan(&takahCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Master takah tidak ditemukan",
		})
		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO template_surat
			(takah_id, template_name, content)
		VALUES
			(?, ?, ?)
	`,
		request.TakahID,
		request.TemplateName,
		request.Content,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data template surat",
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

	response := model.TemplateSuratResponse{
		ID:           int(id),
		TakahID:      request.TakahID,
		TakahCode:    takahCode,
		TemplateName: request.TemplateName,
		Content:      request.Content,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data template surat berhasil dibuat",
		"data":    response,
	})
}

func UpdateTemplateSurat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID template surat tidak valid",
		})
		return
	}

	var request model.TemplateSuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data template surat wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var takahCode string
	err = database.DB.QueryRow(`
		SELECT code FROM master_takah WHERE id = ?
	`, request.TakahID).Scan(&takahCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Master takah tidak ditemukan",
		})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE template_surat
		SET 
			takah_id = ?,
			template_name = ?,
			content = ?
		WHERE id = ?
	`,
		request.TakahID,
		request.TemplateName,
		request.Content,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data template surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data template surat tidak ditemukan",
		})
		return
	}

	response := model.TemplateSuratResponse{
		ID:           id,
		TakahID:      request.TakahID,
		TakahCode:    takahCode,
		TemplateName: request.TemplateName,
		Content:      request.Content,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data template surat berhasil diupdate",
		"data":    response,
	})
}

func DeleteTemplateSurat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID template surat tidak valid",
		})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM template_surat
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data template surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data template surat tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data template surat berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}