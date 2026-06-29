package model

type SuratMasukRequest struct {
	NomorSurat   string `json:"nomor_surat" binding:"required"`
	Pengirim     string `json:"pengirim" binding:"required"`
	Penerima     string `json:"penerima" binding:"required"`
	Perihal      string `json:"perihal" binding:"required"`
	FileSurat    string `json:"file_surat"`
	TanggalSurat string `json:"tanggal_surat" binding:"required"`
	Keterangan   string `json:"keterangan"`
}

type SuratMasukResponse struct {
	ID            int    `json:"id"`
	NomorSurat   string `json:"nomor_surat"`
	Pengirim     string `json:"pengirim"`
	Penerima     string `json:"penerima"`
	Perihal      string `json:"perihal"`
	FileSurat    string `json:"file_surat"`
	TanggalSurat string `json:"tanggal_surat"`
	Keterangan   string `json:"keterangan"`
	Status       string `json:"status"`
	CreatedBy    string `json:"created_by"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}