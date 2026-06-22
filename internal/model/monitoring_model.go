package model

type MonitoringRequest struct {
	SuratKeluarID int    `json:"surat_keluar_id" binding:"required"`
	NomorSurat    string `json:"nomor_surat" binding:"required"`
	Status        string `json:"status" binding:"required"`
	LastApprover  string `json:"last_approver" binding:"required"`
	LastNotes     string `json:"last_notes"`
}

type MonitoringResponse struct {
	ID            int    `json:"id"`
	SuratKeluarID int    `json:"surat_keluar_id"`
	NomorSurat    string `json:"nomor_surat"`
	Status        string `json:"status"`
	LastApprover  string `json:"last_approver"`
	LastNotes     string `json:"last_notes"`
	UpdatedAt     string `json:"updated_at"`
}