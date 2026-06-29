package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllTakah(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			id, code, name, description, sort_order, created_by, updated_by, created_at, updated_at
		FROM master_takah
		ORDER BY sort_order ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data takah",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var takahList []model.TakahResponse

	for rows.Next() {
		var takah model.TakahResponse
		var createdBy, updatedBy int
		var createdAt, updatedAt string

		err := rows.Scan(
			&takah.ID,
			&takah.Code,
			&takah.Name,
			&takah.Description,
			&takah.Order,
			&createdBy,
			&updatedBy,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data takah",
				"error":   err.Error(),
			})
			return
		}

		takah.CreatedBy = strconv.Itoa(createdBy)
		takah.UpdatedBy = strconv.Itoa(updatedBy)
		takah.CreatedTime = createdAt
		takah.UpdatedTime = updatedAt

		takahList = append(takahList, takah)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil diambil",
		"data":    takahList,
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

	var takah model.TakahResponse
	var createdBy, updatedBy int
	var createdAt, updatedAt string

	err = database.DB.QueryRow(`
		SELECT 
			id, code, name, description, sort_order, created_by, updated_by, created_at, updated_at
		FROM master_takah
		WHERE id = ?
	`, id).Scan(
		&takah.ID,
		&takah.Code,
		&takah.Name,
		&takah.Description,
		&takah.Order,
		&createdBy,
		&updatedBy,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data takah tidak ditemukan",
		})
		return
	}

	takah.CreatedBy = strconv.Itoa(createdBy)
	takah.UpdatedBy = strconv.Itoa(updatedBy)
	takah.CreatedTime = createdAt
	takah.UpdatedTime = updatedAt

	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil diambil",
		"data":    takah,
	})
}

func CreateTakah(c *gin.Context) {
	var request model.TakahRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data takah wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO master_takah 
			(code, name, description, sort_order, created_by, updated_by)
		VALUES 
			(?, ?, ?, ?, ?, ?)
	`,
		request.Code,
		request.Name,
		request.Description,
		request.Order,
		1,
		1,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data takah",
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data takah berhasil dibuat",
		"data": gin.H{
			"id":          id,
			"code":        request.Code,
			"name":        request.Name,
			"description": request.Description,
			"order":       request.Order,
			"created_by":  1,
			"updated_by":  1,
		},
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
			"error":   err.Error(),
		})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE master_takah
		SET code = ?, name = ?, description = ?, sort_order = ?, updated_by = ?
		WHERE id = ?
	`,
		request.Code,
		request.Name,
		request.Description,
		request.Order,
		1,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data takah",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data takah tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data takah berhasil diupdate",
		"data": gin.H{
			"id":          id,
			"code":        request.Code,
			"name":        request.Name,
			"description": request.Description,
			"order":       request.Order,
			"updated_by":  1,
		},
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

	result, err := database.DB.Exec(`
		DELETE FROM master_takah
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data takah",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data takah tidak ditemukan",
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