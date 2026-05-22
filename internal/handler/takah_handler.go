package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)


var takahData = []model.TakahResponse{
	{
		ID:          1,
		Code:        "SKET",
		Name:        "Surat Keterangan",
		Description: "Jenis surat keterangan",
		Order:       1,
		CreatedBy:   "M Yogi Darusmawan",
		CreatedTime: "15-07-2024 22:00",
		UpdatedBy:   "M Yogi Darusmawan",
		UpdatedTime: "15-07-2024 22:00",
	},
	{
		ID:          2,
		Code:        "SKK",
		Name:        "Surat Keterangan Kerja",
		Description: "Jenis surat keterangan kerja",
		Order:       2,
		CreatedBy:   "M Yogi Darusmawan",
		CreatedTime: "15-07-2024 22:00",
		UpdatedBy:   "M Yogi Darusmawan",
		UpdatedTime: "15-07-2024 22:00",
	},
}

func GetAllTakah(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil diambil",
		"data":    takahData,
	})
}

func GetTakahByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID takah tidak valid",
		})
		return
	}

	for _, takah := range takahData {
		if takah.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data takah berhasil diambil",
				"data":    takah,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Data takah tidak ditemukan",
	})
}

func CreateTakah(c *gin.Context) {
	var request model.TakahRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data takah wajib diisi dengan benar",
		})
		return
	}

	response := model.TakahResponse{
		ID:          len(takahData) + 1,
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
		CreatedBy:   "Admin",
		CreatedTime: "2026-05-04 15:00",
		UpdatedBy:   "Admin",
		UpdatedTime: "2026-05-04 15:00",
	}

	takahData = append(takahData, response)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data takah berhasil dibuat",
		"data":    response,
	})
}

func UpdateTakah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID takah tidak valid",
		})
		return
	}

	var request model.TakahRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data takah wajib diisi dengan benar",
		})
		return
	}

	response := model.TakahResponse{
		ID:          id,
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
		Order:       request.Order,
		CreatedBy:   "Admin",
		CreatedTime: "2026-05-04 15:00",
		UpdatedBy:   "Admin",
		UpdatedTime: "2026-05-04 15:00",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil diupdate",
		"data":    response,
	})
}

func DeleteTakah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID takah tidak valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}
