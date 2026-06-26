package model

type TemplateSurat struct {
	ID           int    `json:"id"`
	TakahID      int    `json:"takah_id"`
	TemplateName string `json:"template_name"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at"`
}