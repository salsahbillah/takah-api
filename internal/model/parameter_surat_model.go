package model

type ParameterSuratRequest struct {
	TemplateID     int    `json:"template_id" binding:"required"`
	ParameterName string `json:"parameter_name" binding:"required"`
	ParameterKey  string `json:"parameter_key" binding:"required"`
	InputType     string `json:"input_type" binding:"required"`
	IsRequired    bool   `json:"is_required"`
}

type ParameterSuratResponse struct {
	ID            int    `json:"id"`
	TemplateID    int    `json:"template_id"`
	TemplateName  string `json:"template_name"`
	ParameterName string `json:"parameter_name"`
	ParameterKey  string `json:"parameter_key"`
	InputType     string `json:"input_type"`
	IsRequired    bool   `json:"is_required"`
	CreatedAt     string `json:"created_at"`
}