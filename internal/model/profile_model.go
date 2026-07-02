package model

type ProfileResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	PhotoURL  string `json:"photo_url"`
	CreatedAt string `json:"created_at"`
}

type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UploadProfilePhotoResponse struct {
	PhotoURL string `json:"photo_url"`
}