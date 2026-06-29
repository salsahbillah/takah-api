package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllConfigNomor(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			c.id,
			c.takah_id,
			t.code,
			c.company_code,
			c.division_code,
			c.reset_type,
			c.last_number,
			c.created_at
		FROM config_nomor_surat c
		JOIN master_takah t ON c.takah_id = t.id
		ORDER BY c.id ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data config nomor surat",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var configList []model.ConfigNomorResponse

	for rows.Next() {
		var config model.ConfigNomorResponse
		var divisionCode sql.NullString
		var createdAt time.Time

		err := rows.Scan(
			&config.ID,
			&config.TakahID,
			&config.TakahCode,
			&config.CompanyCode,
			&divisionCode,
			&config.ResetType,
			&config.LastNumber,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data config nomor surat",
				"error":   err.Error(),
			})
			return
		}

		if divisionCode.Valid {
			config.DivisionCode = divisionCode.String
		} else {
			config.DivisionCode = ""
		}

		config.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		configList = append(configList, config)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil diambil",
		"data":    configList,
	})
}

func GetConfigNomorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID config nomor tidak valid",
		})
		return
	}

	var config model.ConfigNomorResponse
	var divisionCode sql.NullString
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			c.id,
			c.takah_id,
			t.code,
			c.company_code,
			c.division_code,
			c.reset_type,
			c.last_number,
			c.created_at
		FROM config_nomor_surat c
		JOIN master_takah t ON c.takah_id = t.id
		WHERE c.id = ?
	`, id).Scan(
		&config.ID,
		&config.TakahID,
		&config.TakahCode,
		&config.CompanyCode,
		&divisionCode,
		&config.ResetType,
		&config.LastNumber,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data config nomor surat tidak ditemukan",
		})
		return
	}

	if divisionCode.Valid {
		config.DivisionCode = divisionCode.String
	} else {
		config.DivisionCode = ""
	}

	config.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil diambil",
		"data":    config,
	})
}

func CreateConfigNomor(c *gin.Context) {
	var request model.ConfigNomorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data config nomor surat wajib diisi dengan benar",
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
		INSERT INTO config_nomor_surat
			(takah_id, company_code, division_code, reset_type, last_number)
		VALUES
			(?, ?, ?, ?, ?)
	`,
		request.TakahID,
		request.CompanyCode,
		request.DivisionCode,
		request.ResetType,
		0,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data config nomor surat",
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

	response := model.ConfigNomorResponse{
		ID:           int(id),
		TakahID:      request.TakahID,
		TakahCode:    takahCode,
		CompanyCode:  request.CompanyCode,
		DivisionCode: request.DivisionCode,
		ResetType:    request.ResetType,
		LastNumber:   0,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data config nomor surat berhasil dibuat",
		"data":    response,
	})
}

func UpdateConfigNomor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID config nomor tidak valid",
		})
		return
	}

	var request model.ConfigNomorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data config nomor surat wajib diisi dengan benar",
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
		UPDATE config_nomor_surat
		SET 
			takah_id = ?,
			company_code = ?,
			division_code = ?,
			reset_type = ?
		WHERE id = ?
	`,
		request.TakahID,
		request.CompanyCode,
		request.DivisionCode,
		request.ResetType,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data config nomor surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data config nomor surat tidak ditemukan",
		})
		return
	}

	response := model.ConfigNomorResponse{
		ID:           id,
		TakahID:      request.TakahID,
		TakahCode:    takahCode,
		CompanyCode:  request.CompanyCode,
		DivisionCode: request.DivisionCode,
		ResetType:    request.ResetType,
		LastNumber:   0,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil diupdate",
		"data":    response,
	})
}

func DeleteConfigNomor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID config nomor tidak valid",
		})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM config_nomor_surat
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data config nomor surat",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data config nomor surat tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data config nomor surat berhasil dihapus",
		"data": gin.H{
			"id": id,
		},
	})
}