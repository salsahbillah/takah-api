package model

type MonitoringRequest struct {
	SuratKeluarID int    `json:"surat_keluar_id"`
	SuratMasukID  int    `json:"surat_masuk_id"`
	NomorSurat    string `json:"nomor_surat" binding:"required"`
	Status        string `json:"status" binding:"required"`
	LastApprover  string `json:"last_approver"`
	LastNotes     string `json:"last_notes"`
	UpdatedBy     int    `json:"updated_by"`
}

type MonitoringResponse struct {
	ID            int    `json:"id"`
	SuratKeluarID int    `json:"surat_keluar_id"`
	SuratMasukID  int    `json:"surat_masuk_id"`
	NomorSurat    string `json:"nomor_surat"`
	Status        string `json:"status"`
	LastApprover  string `json:"last_approver"`
	LastNotes     string `json:"last_notes"`
	UpdatedBy     int    `json:"updated_by"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}