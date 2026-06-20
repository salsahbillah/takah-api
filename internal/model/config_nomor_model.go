package model

type ConfigNomorRequest struct {
	TakahID      int    `json:"takah_id" binding:"required"`
	CompanyCode  string `json:"company_code" binding:"required"`
	DivisionCode string `json:"division_code"`
	ResetType    string `json:"reset_type" binding:"required"`
}

type ConfigNomorResponse struct {
	ID           int    `json:"id"`
	TakahID      int    `json:"takah_id"`
	TakahCode    string `json:"takah_code"`
	CompanyCode  string `json:"company_code"`
	DivisionCode string `json:"division_code"`
	ResetType    string `json:"reset_type"`
	LastNumber   int    `json:"last_number"`
	CreatedAt    string `json:"created_at"`
}