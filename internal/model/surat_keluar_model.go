package model

type SuratKeluarRequest struct {
	TemplateID      int               `json:"template_id" binding:"required"`
	TakahID         int               `json:"takah_id" binding:"required"`
	TujuanSurat     string            `json:"tujuan_surat" binding:"required"`
	Perihal         string            `json:"perihal" binding:"required"`
	Lampiran        string            `json:"lampiran"`
	TanggalSurat    string            `json:"tanggal_surat" binding:"required"`
	FileSurat       string            `json:"file_surat"`
	ParameterValues map[string]string `json:"parameter_values"`
}

type SuratKeluarResponse struct {
	ID               int               `json:"id"`
	NomorSurat       string            `json:"nomor_surat"`
	TemplateID       int               `json:"template_id"`
	TemplateName     string            `json:"template_name"`
	TakahID          int               `json:"takah_id"`
	TakahCode        string            `json:"takah_code"`
	TujuanSurat      string            `json:"tujuan_surat"`
	Perihal          string            `json:"perihal"`
	Lampiran         string            `json:"lampiran"`
	TanggalSurat     string            `json:"tanggal_surat"`
	FileSurat        string            `json:"file_surat"`
	GeneratedContent string            `json:"generated_content"`
	ParameterValues  map[string]string `json:"parameter_values"`
	Status           string            `json:"status"`
	CreatedBy        string            `json:"created_by"`
	CreatedAt        string            `json:"created_at"`
}