package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

var monitoringData = []model.MonitoringResponse{
	{
		ID:            1,
		SuratKeluarID: 1,
		NomorSurat:    "001/SKET/CBN/062026",
		Status:        "approved",
		LastApprover:  "M Yogi Darusmawan",
		LastNotes:     "Disetujui",
		UpdatedAt:     "2026-06-18 10:00",
	},
}

func GetAllMonitoring(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil diambil",
		"data":    monitoringData,
	})
}

func GetMonitoringByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID monitoring tidak valid"})
		return
	}

	for _, monitoring := range monitoringData {
		if monitoring.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data monitoring berhasil diambil",
				"data":    monitoring,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Data monitoring tidak ditemukan"})
}

func CreateMonitoring(c *gin.Context) {
	var request model.MonitoringRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data monitoring wajib diisi dengan benar"})
		return
	}

	response := model.MonitoringResponse{
		ID:            len(monitoringData) + 1,
		SuratKeluarID: request.SuratKeluarID,
		NomorSurat:    request.NomorSurat,
		Status:        request.Status,
		LastApprover:  request.LastApprover,
		LastNotes:     request.LastNotes,
		UpdatedAt:     "2026-06-18 10:00",
	}

	monitoringData = append(monitoringData, response)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data monitoring berhasil dibuat",
		"data":    response,
	})
}

func UpdateMonitoring(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID monitoring tidak valid"})
		return
	}

	var request model.MonitoringRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data monitoring wajib diisi dengan benar"})
		return
	}

	response := model.MonitoringResponse{
		ID:            id,
		SuratKeluarID: request.SuratKeluarID,
		NomorSurat:    request.NomorSurat,
		Status:        request.Status,
		LastApprover:  request.LastApprover,
		LastNotes:     request.LastNotes,
		UpdatedAt:     "2026-06-18 10:00",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil diupdate",
		"data":    response,
	})
}

func DeleteMonitoring(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID monitoring tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil dihapus",
		"data":    gin.H{"id": id},
	})
}
