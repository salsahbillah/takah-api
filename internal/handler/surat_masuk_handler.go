package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

var suratMasukData = []model.SuratMasukResponse{
	{
		ID:           1,
		NomorSurat:  "001/EXT/VI/2026",
		Pengirim:    "PT Contoh Indonesia",
		Penerima:    "Admin",
		Perihal:     "Undangan Kerja Sama",
		FileSurat:   "-",
		TanggalSurat: "2026-06-18",
		Keterangan:  "Surat diterima melalui email",
		Status:      "received",
		CreatedBy:   "Admin",
		CreatedAt:   "2026-06-18 10:00",
	},
}

func GetAllSuratMasuk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil diambil",
		"data":    suratMasukData,
	})
}

func GetSuratMasukByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID surat masuk tidak valid"})
		return
	}

	for _, surat := range suratMasukData {
		if surat.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data surat masuk berhasil diambil",
				"data":    surat,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Data surat masuk tidak ditemukan"})
}

func CreateSuratMasuk(c *gin.Context) {
	var request model.SuratMasukRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat masuk wajib diisi dengan benar"})
		return
	}

	response := model.SuratMasukResponse{
		ID:           len(suratMasukData) + 1,
		NomorSurat:  request.NomorSurat,
		Pengirim:    request.Pengirim,
		Penerima:    request.Penerima,
		Perihal:     request.Perihal,
		FileSurat:   request.FileSurat,
		TanggalSurat: request.TanggalSurat,
		Keterangan:  request.Keterangan,
		Status:      "received",
		CreatedBy:   "Admin",
		CreatedAt:   "2026-06-18 10:00",
	}

	suratMasukData = append(suratMasukData, response)

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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data surat masuk wajib diisi dengan benar"})
		return
	}

	response := model.SuratMasukResponse{
		ID:           id,
		NomorSurat:  request.NomorSurat,
		Pengirim:    request.Pengirim,
		Penerima:    request.Penerima,
		Perihal:     request.Perihal,
		FileSurat:   request.FileSurat,
		TanggalSurat: request.TanggalSurat,
		Keterangan:  request.Keterangan,
		Status:      "received",
		CreatedBy:   "Admin",
		CreatedAt:   "2026-06-18 10:00",
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat masuk berhasil dihapus",
		"data": gin.H{"id": id},
	})
}