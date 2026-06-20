package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

var configNomorData = []model.ConfigNomorResponse{
	{
		ID:           1,
		TakahID:      1,
		TakahCode:    "SKET",
		CompanyCode:  "CBN",
		DivisionCode: "",
		ResetType:    "monthly",
		LastNumber:   0,
		CreatedAt:    "2026-06-18 10:00",
	},
}

func findTakahByID(id int) (model.TakahResponse, bool) {
	for _, takah := range takahData {
		if takah.ID == id {
			return takah, true
		}
	}
	return model.TakahResponse{}, false
}

func GetAllConfigNomor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil diambil",
		"data":    configNomorData,
	})
}

func GetConfigNomorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID config nomor tidak valid"})
		return
	}

	for _, config := range configNomorData {
		if config.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data config nomor surat berhasil diambil",
				"data":    config,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Data config nomor surat tidak ditemukan"})
}

func CreateConfigNomor(c *gin.Context) {
	var request model.ConfigNomorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data config nomor surat wajib diisi dengan benar"})
		return
	}

	takah, found := findTakahByID(request.TakahID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "Master takah tidak ditemukan"})
		return
	}

	response := model.ConfigNomorResponse{
		ID:           len(configNomorData) + 1,
		TakahID:      request.TakahID,
		TakahCode:    takah.Code,
		CompanyCode:  request.CompanyCode,
		DivisionCode: request.DivisionCode,
		ResetType:    request.ResetType,
		LastNumber:   0,
		CreatedAt:    "2026-06-18 10:00",
	}

	configNomorData = append(configNomorData, response)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data config nomor surat berhasil dibuat",
		"data":    response,
	})
}

func UpdateConfigNomor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID config nomor tidak valid"})
		return
	}

	var request model.ConfigNomorRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data config nomor surat wajib diisi dengan benar"})
		return
	}

	takah, found := findTakahByID(request.TakahID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "Master takah tidak ditemukan"})
		return
	}

	response := model.ConfigNomorResponse{
		ID:           id,
		TakahID:      request.TakahID,
		TakahCode:    takah.Code,
		CompanyCode:  request.CompanyCode,
		DivisionCode: request.DivisionCode,
		ResetType:    request.ResetType,
		LastNumber:   0,
		CreatedAt:    "2026-06-18 10:00",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil diupdate",
		"data":    response,
	})
}

func DeleteConfigNomor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID config nomor tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}