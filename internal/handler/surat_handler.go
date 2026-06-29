package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/database"

	"github.com/gin-gonic/gin"
)

func GetAllSurat(c *gin.Context) {
	var totalSuratKeluar int
	var totalSuratMasuk int
	var totalDraft int
	var totalPending int
	var totalApproved int
	var totalRejected int
	var totalCompleted int
	var totalMonitoring int

	err := database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar`).Scan(&totalSuratKeluar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat keluar", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_masuk`).Scan(&totalSuratMasuk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat masuk", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar WHERE status = 'draft'`).Scan(&totalDraft)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat draft", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar WHERE status = 'pending'`).Scan(&totalPending)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat pending", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar WHERE status = 'approved'`).Scan(&totalApproved)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat approved", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar WHERE status = 'rejected'`).Scan(&totalRejected)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat rejected", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM surat_keluar WHERE status = 'completed'`).Scan(&totalCompleted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung surat completed", "error": err.Error()})
		return
	}

	err = database.DB.QueryRow(`SELECT COUNT(*) FROM monitoring_surat`).Scan(&totalMonitoring)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghitung monitoring surat", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Summary data surat berhasil diambil",
		"data": gin.H{
			"total_surat_keluar": totalSuratKeluar,
			"total_surat_masuk":  totalSuratMasuk,
			"total_semua_surat":  totalSuratKeluar + totalSuratMasuk,
			"status_surat_keluar": gin.H{
				"draft":     totalDraft,
				"pending":   totalPending,
				"approved":  totalApproved,
				"rejected":  totalRejected,
				"completed": totalCompleted,
			},
			"total_monitoring": totalMonitoring,
		},
	})
}

func GetSuratByID(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Detail surat dipisahkan berdasarkan jenis surat",
		"data": gin.H{
			"surat_keluar": "/api/v1/surat-keluar/:id",
			"surat_masuk":  "/api/v1/surat-masuk/:id",
		},
	})
}

func CreateSurat(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "CRUD surat dipisahkan berdasarkan jenis surat",
		"data": gin.H{
			"buat_surat_keluar": "/api/v1/surat-keluar",
			"buat_surat_masuk":  "/api/v1/surat-masuk",
		},
	})
}

func UpdateSurat(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Update surat dipisahkan berdasarkan jenis surat",
		"data": gin.H{
			"update_surat_keluar": "/api/v1/surat-keluar/:id",
			"update_surat_masuk":  "/api/v1/surat-masuk/:id",
		},
	})
}

func DeleteSurat(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Delete surat dipisahkan berdasarkan jenis surat",
		"data": gin.H{
			"delete_surat_keluar": "/api/v1/surat-keluar/:id",
			"delete_surat_masuk":  "/api/v1/surat-masuk/:id",
		},
	})
}