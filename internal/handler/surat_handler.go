package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

// GET /surat
func GetAllSurat(c *gin.Context) {
	surat := []model.SuratResponse{
		{
			ID:          1,
			NomorSurat: "001/HR/2026",
			Judul:       "Surat Undangan Rapat",
			Pengirim:    "HRD",
			Penerima:    "Seluruh Karyawan",
			Status:      "masuk",
		},
		{
			ID:          2,
			NomorSurat: "002/ADM/2026",
			Judul:       "Surat Pemberitahuan",
			Pengirim:    "Administrasi",
			Penerima:    "Divisi Keuangan",
			Status:      "keluar",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat berhasil diambil",
		"data":    surat,
	})
}

// POST /surat
func CreateSurat(c *gin.Context) {
	var request model.SuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data surat wajib diisi dengan benar",
		})
		return
	}

	response := model.SuratResponse{
		ID:          3,
		NomorSurat: request.NomorSurat,
		Judul:       request.Judul,
		Pengirim:    request.Pengirim,
		Penerima:    request.Penerima,
		Status:      request.Status,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data surat berhasil dibuat",
		"data":    response,
	})
}

// GET /surat
func GetSuratByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID surat tidak valid",
		})
		return
	}

	suratList := []model.SuratResponse{
		{
			ID:          1,
			NomorSurat: "001/HR/2026",
			Judul:       "Surat Undangan Rapat",
			Pengirim:    "HRD",
			Penerima:    "Seluruh Karyawan",
			Status:      "masuk",
		},
		{
			ID:          2,
			NomorSurat: "002/ADM/2026",
			Judul:       "Surat Pemberitahuan",
			Pengirim:    "Administrasi",
			Penerima:    "Divisi Keuangan",
			Status:      "keluar",
		},
	}

	for _, surat := range suratList {
		if surat.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data surat berhasil diambil",
				"data":    surat,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Data surat tidak ditemukan",
	})
}

// PUT /surat
func UpdateSurat(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID surat tidak valid",
		})
		return
	}

	var request model.SuratRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data surat wajib diisi dengan benar",
		})
		return
	}

	response := model.SuratResponse{
		ID:          id,
		NomorSurat: request.NomorSurat,
		Judul:       request.Judul,
		Pengirim:    request.Pengirim,
		Penerima:    request.Penerima,
		Status:      request.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat berhasil diupdate",
		"data":    response,
	})
}

// DELETE /surat
func DeleteSurat(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID surat tidak valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data surat berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}