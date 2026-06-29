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

func GetAllMonitoring(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT
			id,
			surat_keluar_id,
			surat_masuk_id,
			nomor_surat,
			status,
			last_approver,
			last_notes,
			updated_by,
			created_at,
			updated_at
		FROM monitoring_surat
		ORDER BY id DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data monitoring",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var monitoringList []model.MonitoringResponse

	for rows.Next() {
		var monitoring model.MonitoringResponse
		var suratKeluarID sql.NullInt64
		var suratMasukID sql.NullInt64
		var lastApprover sql.NullString
		var lastNotes sql.NullString
		var updatedBy sql.NullInt64
		var createdAt time.Time
		var updatedAt time.Time

		err := rows.Scan(
			&monitoring.ID,
			&suratKeluarID,
			&suratMasukID,
			&monitoring.NomorSurat,
			&monitoring.Status,
			&lastApprover,
			&lastNotes,
			&updatedBy,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data monitoring",
				"error":   err.Error(),
			})
			return
		}

		if suratKeluarID.Valid {
			monitoring.SuratKeluarID = int(suratKeluarID.Int64)
		}
		if suratMasukID.Valid {
			monitoring.SuratMasukID = int(suratMasukID.Int64)
		}
		if lastApprover.Valid {
			monitoring.LastApprover = lastApprover.String
		}
		if lastNotes.Valid {
			monitoring.LastNotes = lastNotes.String
		}
		if updatedBy.Valid {
			monitoring.UpdatedBy = int(updatedBy.Int64)
		}

		monitoring.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		monitoring.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

		monitoringList = append(monitoringList, monitoring)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil diambil",
		"data":    monitoringList,
	})
}

func GetMonitoringByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID monitoring tidak valid",
		})
		return
	}

	var monitoring model.MonitoringResponse
	var suratKeluarID sql.NullInt64
	var suratMasukID sql.NullInt64
	var lastApprover sql.NullString
	var lastNotes sql.NullString
	var updatedBy sql.NullInt64
	var createdAt time.Time
	var updatedAt time.Time

	err = database.DB.QueryRow(`
		SELECT
			id,
			surat_keluar_id,
			surat_masuk_id,
			nomor_surat,
			status,
			last_approver,
			last_notes,
			updated_by,
			created_at,
			updated_at
		FROM monitoring_surat
		WHERE id = ?
	`, id).Scan(
		&monitoring.ID,
		&suratKeluarID,
		&suratMasukID,
		&monitoring.NomorSurat,
		&monitoring.Status,
		&lastApprover,
		&lastNotes,
		&updatedBy,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data monitoring tidak ditemukan",
		})
		return
	}

	if suratKeluarID.Valid {
		monitoring.SuratKeluarID = int(suratKeluarID.Int64)
	}
	if suratMasukID.Valid {
		monitoring.SuratMasukID = int(suratMasukID.Int64)
	}
	if lastApprover.Valid {
		monitoring.LastApprover = lastApprover.String
	}
	if lastNotes.Valid {
		monitoring.LastNotes = lastNotes.String
	}
	if updatedBy.Valid {
		monitoring.UpdatedBy = int(updatedBy.Int64)
	}

	monitoring.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	monitoring.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil diambil",
		"data":    monitoring,
	})
}

func CreateMonitoring(c *gin.Context) {
	var request model.MonitoringRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data monitoring wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var suratKeluarID interface{} = nil
	var suratMasukID interface{} = nil

	if request.SuratKeluarID != 0 {
		suratKeluarID = request.SuratKeluarID
	}
	if request.SuratMasukID != 0 {
		suratMasukID = request.SuratMasukID
	}

	result, err := database.DB.Exec(`
		INSERT INTO monitoring_surat
			(surat_keluar_id, surat_masuk_id, nomor_surat, status, last_approver, last_notes, updated_by)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)
	`,
		suratKeluarID,
		suratMasukID,
		request.NomorSurat,
		request.Status,
		request.LastApprover,
		request.LastNotes,
		request.UpdatedBy,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data monitoring",
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

	response := model.MonitoringResponse{
		ID:            int(id),
		SuratKeluarID: request.SuratKeluarID,
		SuratMasukID:  request.SuratMasukID,
		NomorSurat:    request.NomorSurat,
		Status:        request.Status,
		LastApprover:  request.LastApprover,
		LastNotes:     request.LastNotes,
		UpdatedBy:     request.UpdatedBy,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data monitoring berhasil dibuat",
		"data":    response,
	})
}

func UpdateMonitoring(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID monitoring tidak valid",
		})
		return
	}

	var request model.MonitoringRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data monitoring wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	var suratKeluarID interface{} = nil
	var suratMasukID interface{} = nil

	if request.SuratKeluarID != 0 {
		suratKeluarID = request.SuratKeluarID
	}
	if request.SuratMasukID != 0 {
		suratMasukID = request.SuratMasukID
	}

	result, err := database.DB.Exec(`
		UPDATE monitoring_surat
		SET
			surat_keluar_id = ?,
			surat_masuk_id = ?,
			nomor_surat = ?,
			status = ?,
			last_approver = ?,
			last_notes = ?,
			updated_by = ?
		WHERE id = ?
	`,
		suratKeluarID,
		suratMasukID,
		request.NomorSurat,
		request.Status,
		request.LastApprover,
		request.LastNotes,
		request.UpdatedBy,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data monitoring",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data monitoring tidak ditemukan",
		})
		return
	}

	response := model.MonitoringResponse{
		ID:            id,
		SuratKeluarID: request.SuratKeluarID,
		SuratMasukID:  request.SuratMasukID,
		NomorSurat:    request.NomorSurat,
		Status:        request.Status,
		LastApprover:  request.LastApprover,
		LastNotes:     request.LastNotes,
		UpdatedBy:     request.UpdatedBy,
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil diupdate",
		"data":    response,
	})
}

func DeleteMonitoring(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID monitoring tidak valid",
		})
		return
	}

	result, err := database.DB.Exec(`
		DELETE FROM monitoring_surat
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus data monitoring",
			"error":   err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data monitoring tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data monitoring berhasil dihapus",
		"data":    gin.H{"id": id},
	})
}