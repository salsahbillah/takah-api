package handler

import (
	"net/http"
	"strconv"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

var approvalData = []model.ApprovalResponse{}

func GetAllApproval(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Data approval berhasil diambil",
		"data":    approvalData,
	})
}

func GetApprovalByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID approval tidak valid"})
		return
	}

	for _, approval := range approvalData {
		if approval.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data approval berhasil diambil",
				"data":    approval,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Data approval tidak ditemukan"})
}

func CreateApproval(c *gin.Context) {
	var request model.ApprovalRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data approval wajib diisi dengan benar"})
		return
	}

	response := model.ApprovalResponse{
		ID:             len(approvalData) + 1,
		SuratKeluarID:  request.SuratKeluarID,
		Approver:       request.Approver,
		ApprovalStatus: "pending",
		Notes:          request.Notes,
		ApprovedAt:     "-",
	}

	approvalData = append(approvalData, response)

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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data approval wajib diisi dengan benar"})
		return
	}

	response := model.ApprovalResponse{
		ID:             id,
		SuratKeluarID:  request.SuratKeluarID,
		Approver:       request.Approver,
		ApprovalStatus: "approved",
		Notes:          request.Notes,
		ApprovedAt:     "2026-06-18 10:00",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Approval surat berhasil diupdate",
		"data":    response,
	})
}