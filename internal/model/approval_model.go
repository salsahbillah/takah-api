package model

type ApprovalRequest struct {
	SuratKeluarID  int    `json:"surat_keluar_id" binding:"required"`
	ApproverID     int    `json:"approver_id"`
	ApproverName   string `json:"approver_name" binding:"required"`
	ApprovalStatus string `json:"approval_status"`
	Notes          string `json:"notes"`
}

type ApprovalResponse struct {
	ID             int    `json:"id"`
	SuratKeluarID  int    `json:"surat_keluar_id"`
	NomorSurat     string `json:"nomor_surat"`
	ApproverID     int    `json:"approver_id"`
	ApproverName   string `json:"approver_name"`
	ApprovalStatus string `json:"approval_status"`
	Notes          string `json:"notes"`
	ApprovedAt     string `json:"approved_at"`
	CreatedAt      string `json:"created_at"`
}