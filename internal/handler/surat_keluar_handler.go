package handler

import (
	"net/http"
	"strconv"
	"takah-api/internal/helper"
	"takah-api/internal/model"
	"github.com/gin-gonic/gin"
	
)

var suratKeluarData = []model.SuratKeluarResponse{
	{
		ID:           1,
		NomorSurat:  "001/SKET/CBN/062026",
		TakahID:      1,
		TakahCode:    "SKET",
		TujuanSurat:  "Mahasiswa",
		Perihal:      "Surat Keterangan Magang",
		Lampiran:     "-",
		TanggalSurat: "2026-06-18",
		FileSurat:    "-",
		Status:       "draft",
		CreatedBy:    "Admin",
		CreatedAt:    "2026-06-18 10:00",
	},
}

func GetAllSuratKeluar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil diambil",
		"data":    suratKeluarData,
	})
}

func GetSuratKeluarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	for _, surat := range suratKeluarData {
		if surat.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data surat keluar berhasil diambil",
				"data":    surat,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Data surat keluar tidak ditemukan"})
}

func CreateSuratKeluar(c *gin.Context) {
	var request model.SuratKeluarRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat keluar wajib diisi dengan benar"})
		return
	}

	var selectedConfig model.ConfigNomorResponse
	configFound := false

	for i, config := range configNomorData {
		if config.TakahID == request.TakahID {
			configNomorData[i].LastNumber++
			selectedConfig = configNomorData[i]
			configFound = true
			break
		}
	}

	if !configFound {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Config nomor surat untuk takah ini tidak ditemukan",
		})
		return
	}

	nomorSurat := helper.GenerateNomorSurat(
		selectedConfig.LastNumber-1,
		selectedConfig.TakahCode,
		selectedConfig.CompanyCode,
		selectedConfig.ResetType,
	)

	response := model.SuratKeluarResponse{
		ID:           len(suratKeluarData) + 1,
		NomorSurat:  nomorSurat,
		TakahID:      request.TakahID,
		TakahCode:    selectedConfig.TakahCode,
		TujuanSurat:  request.TujuanSurat,
		Perihal:      request.Perihal,
		Lampiran:     request.Lampiran,
		TanggalSurat: request.TanggalSurat,
		FileSurat:    request.FileSurat,
		Status:       "draft",
		CreatedBy:    "Admin",
		CreatedAt:    "2026-06-18 10:00",
	}

	suratKeluarData = append(suratKeluarData, response)

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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat keluar wajib diisi dengan benar"})
		return
	}

	response := model.SuratKeluarResponse{
		ID:           id,
		NomorSurat:  "AUTO-GENERATE-NANTI",
		TakahID:      request.TakahID,
		TakahCode:    "-",
		TujuanSurat:  request.TujuanSurat,
		Perihal:      request.Perihal,
		Lampiran:     request.Lampiran,
		TanggalSurat: request.TanggalSurat,
		FileSurat:    request.FileSurat,
		Status:       "draft",
		CreatedBy:    "Admin",
		CreatedAt:    "2026-06-18 10:00",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil diupdate",
		"data":    response,
	})
}

func DeleteSuratKeluar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat keluar tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat keluar berhasil dihapus",
		"data": gin.H{"id": id},
	})
}