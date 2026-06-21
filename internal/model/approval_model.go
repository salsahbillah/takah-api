package model

type ApprovalRequest struct {
	SuratKeluarID int    `json:"surat_keluar_id" binding:"required"`
	Approver      string `json:"approver" binding:"required"`
	Notes         string `json:"notes"`
}

type ApprovalResponse struct {
	ID             int    `json:"id"`
	SuratKeluarID  int    `json:"surat_keluar_id"`
	Approver       string `json:"approver"`
	ApprovalStatus string `json:"approval_status"`
	Notes          string `json:"notes"`
	ApprovedAt     string `json:"approved_at"`
}