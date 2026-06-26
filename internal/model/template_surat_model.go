package model

type TemplateSuratRequest struct {
	TakahID      int    `json:"takah_id" binding:"required"`
	TemplateName string `json:"template_name" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

type TemplateSuratResponse struct {
	ID           int    `json:"id"`
	TakahID      int    `json:"takah_id"`
	TakahCode    string `json:"takah_code"`
	TemplateName string `json:"template_name"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at"`
}