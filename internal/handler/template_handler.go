package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"takah-api/internal/model"
)

var templateData = []model.TemplateSurat{
	{
		ID:           1,
		TakahID:      1,
		TemplateName: "Template Surat Undangan",
		Content:      "Dengan hormat, kami mengundang...",
		CreatedAt:    "2026-06-18 10:00",
	},
}

func GetAllTemplateSurat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": templateData,
		"message": "Data template surat berhasil diambil",
	})
}
func GetTemplateSuratByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, template := range templateData {
		if template.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": template,
				"message": "Data template surat berhasil diambil",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Data template surat tidak ditemukan",
	})
}
func CreateTemplateSurat(c *gin.Context) {
	var input model.TemplateSurat

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	input.ID = len(templateData) + 1
	input.CreatedAt = "2026-06-18 10:00"

	templateData = append(templateData, input)

	c.JSON(http.StatusCreated, gin.H{
		"data": input,
		"message": "Data template surat berhasil dibuat",
	})
}
func UpdateTemplateSurat(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input model.TemplateSurat

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for i, template := range templateData {
		if template.ID == id {
			input.ID = id
			input.CreatedAt = template.CreatedAt

			templateData[i] = input

			c.JSON(http.StatusOK, gin.H{
				"data": input,
				"message": "Data template surat berhasil diupdate",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Data template surat tidak ditemukan",
	})
}
func DeleteTemplateSurat(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, template := range templateData {
		if template.ID == id {

			templateData = append(templateData[:i], templateData[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"data": gin.H{
					"id": id,
				},
				"message": "Data template surat berhasil dihapus",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Data template surat tidak ditemukan",
	})
}
