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

func GetAllApproval(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT 
			a.id,
			a.surat_keluar_id,
			s.nomor_surat,
			a.approver_id,
			a.approver_name,
			a.approval_status,
			a.notes,
			a.approved_at,
			a.created_at
		FROM approval_surat a
		JOIN surat_keluar s ON a.surat_keluar_id = s.id
		ORDER BY a.id DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data approval",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	var approvalList []model.ApprovalResponse

	for rows.Next() {
		var approval model.ApprovalResponse
		var approverID sql.NullInt64
		var notes sql.NullString
		var approvedAt sql.NullTime
		var createdAt time.Time

		err := rows.Scan(
			&approval.ID,
			&approval.SuratKeluarID,
			&approval.NomorSurat,
			&approverID,
			&approval.ApproverName,
			&approval.ApprovalStatus,
			&notes,
			&approvedAt,
			&createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gagal membaca data approval",
				"error":   err.Error(),
			})
			return
		}

		if approverID.Valid {
			approval.ApproverID = int(approverID.Int64)
		}
		if notes.Valid {
			approval.Notes = notes.String
		}
		if approvedAt.Valid {
			approval.ApprovedAt = approvedAt.Time.Format("2006-01-02 15:04:05")
		} else {
			approval.ApprovedAt = "-"
		}

		approval.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		approvalList = append(approvalList, approval)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data approval berhasil diambil",
		"data":    approvalList,
	})
}

func GetApprovalByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID approval tidak valid",
		})
		return
	}

	var approval model.ApprovalResponse
	var approverID sql.NullInt64
	var notes sql.NullString
	var approvedAt sql.NullTime
	var createdAt time.Time

	err = database.DB.QueryRow(`
		SELECT 
			a.id,
			a.surat_keluar_id,
			s.nomor_surat,
			a.approver_id,
			a.approver_name,
			a.approval_status,
			a.notes,
			a.approved_at,
			a.created_at
		FROM approval_surat a
		JOIN surat_keluar s ON a.surat_keluar_id = s.id
		WHERE a.id = ?
	`, id).Scan(
		&approval.ID,
		&approval.SuratKeluarID,
		&approval.NomorSurat,
		&approverID,
		&approval.ApproverName,
		&approval.ApprovalStatus,
		&notes,
		&approvedAt,
		&createdAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data approval tidak ditemukan",
		})
		return
	}

	if approverID.Valid {
		approval.ApproverID = int(approverID.Int64)
	}
	if notes.Valid {
		approval.Notes = notes.String
	}
	if approvedAt.Valid {
		approval.ApprovedAt = approvedAt.Time.Format("2006-01-02 15:04:05")
	} else {
		approval.ApprovedAt = "-"
	}

	approval.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data approval berhasil diambil",
		"data":    approval,
	})
}

func CreateApproval(c *gin.Context) {
	var request model.ApprovalRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data approval wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	status := request.ApprovalStatus
	if status == "" {
		status = "pending"
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memulai transaksi approval",
			"error":   err.Error(),
		})
		return
	}
	defer tx.Rollback()

	var nomorSurat string
	err = tx.QueryRow(`
		SELECT nomor_surat FROM surat_keluar WHERE id = ?
	`, request.SuratKeluarID).Scan(&nomorSurat)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Surat keluar tidak ditemukan",
		})
		return
	}

	result, err := tx.Exec(`
		INSERT INTO approval_surat
			(surat_keluar_id, approver_id, approver_name, approval_status, notes)
		VALUES
			(?, ?, ?, ?, ?)
	`,
		request.SuratKeluarID,
		request.ApproverID,
		request.ApproverName,
		status,
		request.Notes,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat data approval",
			"error":   err.Error(),
		})
		return
	}

	_, err = tx.Exec(`
		UPDATE surat_keluar
		SET status = ?
		WHERE id = ?
	`, "pending", request.SuratKeluarID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate status surat keluar",
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

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menyimpan transaksi approval",
			"error":   err.Error(),
		})
		return
	}

	response := model.ApprovalResponse{
		ID:             int(id),
		SuratKeluarID:  request.SuratKeluarID,
		NomorSurat:     nomorSurat,
		ApproverID:     request.ApproverID,
		ApproverName:   request.ApproverName,
		ApprovalStatus: status,
		Notes:          request.Notes,
		ApprovedAt:     "-",
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Surat berhasil dikirim untuk approval",
		"data":    response,
	})
}


func UpdateApproval(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID approval tidak valid"})
		return
	}

	var request model.ApprovalRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data approval wajib diisi dengan benar",
			"error":   err.Error(),
		})
		return
	}

	status := request.ApprovalStatus
	if status == "" {
		status = "approved"
	}

	if status != "approved" && status != "rejected" && status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status approval hanya boleh pending, approved, atau rejected",
		})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memulai transaksi approval", "error": err.Error()})
		return
	}
	defer tx.Rollback()

	var suratKeluarID int
	var nomorSurat string

	err = tx.QueryRow(`
		SELECT a.surat_keluar_id, s.nomor_surat
		FROM approval_surat a
		JOIN surat_keluar s ON a.surat_keluar_id = s.id
		WHERE a.id = ?
	`, id).Scan(&suratKeluarID, &nomorSurat)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data approval tidak ditemukan"})
		return
	}

	approvedAt := sql.NullTime{}
	if status == "approved" || status == "rejected" {
		approvedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	_, err = tx.Exec(`
		UPDATE approval_surat
		SET
			approver_id = ?,
			approver_name = ?,
			approval_status = ?,
			notes = ?,
			approved_at = ?
		WHERE id = ?
	`,
		request.ApproverID,
		request.ApproverName,
		status,
		request.Notes,
		approvedAt,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate data approval", "error": err.Error()})
		return
	}

	suratStatus := "pending"
	if status == "approved" {
		suratStatus = "approved"
	}
	if status == "rejected" {
		suratStatus = "rejected"
	}

	_, err = tx.Exec(`
		UPDATE surat_keluar
		SET status = ?
		WHERE id = ?
	`, suratStatus, suratKeluarID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengupdate status surat keluar", "error": err.Error()})
		return
	}

	_, err = tx.Exec(`
		INSERT INTO monitoring_surat
			(surat_keluar_id, nomor_surat, status, last_approver, last_notes, updated_by)
		VALUES
			(?, ?, ?, ?, ?, ?)
	`, suratKeluarID, nomorSurat, suratStatus, request.ApproverName, request.Notes, request.ApproverID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data monitoring", "error": err.Error()})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan transaksi approval", "error": err.Error()})
		return
	}

	approvedAtString := "-"
	if approvedAt.Valid {
		approvedAtString = approvedAt.Time.Format("2006-01-02 15:04:05")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Approval surat berhasil diupdate",
		"data": model.ApprovalResponse{
			ID:             id,
			SuratKeluarID:  suratKeluarID,
			NomorSurat:     nomorSurat,
			ApproverID:     request.ApproverID,
			ApproverName:   request.ApproverName,
			ApprovalStatus: status,
			Notes:          request.Notes,
			ApprovedAt:     approvedAtString,
			CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}