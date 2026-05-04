package model

type SuratRequest struct {
	NomorSurat string `json:"nomor_surat" binding:"required"`
	Judul      string `json:"judul" binding:"required"`
	Pengirim   string `json:"pengirim" binding:"required"`
	Penerima   string `json:"penerima" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type SuratResponse struct {
	ID         int    `json:"id"`
	NomorSurat string `json:"nomor_surat"`
	Judul      string `json:"judul"`
	Pengirim   string `json:"pengirim"`
	Penerima   string `json:"penerima"`
	Status     string `json:"status"`
}
