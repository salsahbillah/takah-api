package model

type TakahRequest struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

type TakahResponse struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	CreatedBy   string `json:"created_by"`
	CreatedTime string `json:"created_time"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedTime string `json:"updated_time"`
}
