package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"takah-api/internal/database"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")

	var id int
	var name, email, role string
	var photoURL sql.NullString
	var createdAt time.Time

	err := database.DB.QueryRow(`
		SELECT id, name, email, role, photo_url, created_at
		FROM users
		WHERE id = ?
	`, userID).Scan(&id, &name, &email, &role, &photoURL, &createdAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile berhasil diambil",
		"data": gin.H{
			"id":         id,
			"name":       name,
			"email":      email,
			"role":       role,
			"photo_url":  photoURL.String,
			"created_at": createdAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetInt("user_id")

	var request struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama wajib diisi",
		})
		return
	}

	_, err := database.DB.Exec(`
		UPDATE users
		SET name = ?
		WHERE id = ?
	`, request.Name, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengubah profile",
			"error":   err.Error(),
		})
		return
	}

	GetProfile(c)
}

func UpdatePassword(c *gin.Context) {
	userID := c.GetInt("user_id")

	var request struct {
		OldPassword     string `json:"old_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data password wajib diisi",
		})
		return
	}

	if request.NewPassword != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Konfirmasi password tidak sesuai",
		})
		return
	}

	var currentPassword string

	err := database.DB.QueryRow(`
		SELECT password
		FROM users
		WHERE id = ?
	`, userID).Scan(&currentPassword)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	if request.OldPassword != currentPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password lama salah",
		})
		return
	}

	_, err = database.DB.Exec(`
		UPDATE users
		SET password = ?
		WHERE id = ?
	`, request.NewPassword, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengubah password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
}

func UploadProfilePhoto(c *gin.Context) {
	userID := c.GetInt("user_id")

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Foto wajib dipilih",
		})
		return
	}

	uploadDir := "uploads/profiles"

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat folder upload",
			"error":   err.Error(),
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("profile_%d_%d%s", userID, time.Now().Unix(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal upload foto",
			"error":   err.Error(),
		})
		return
	}

	photoURL := "/" + filepath.ToSlash(filePath)

	_, err = database.DB.Exec(`
		UPDATE users
		SET photo_url = ?
		WHERE id = ?
	`, photoURL, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menyimpan foto profile",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Foto profile berhasil diupload",
		"data": gin.H{
			"photo_url": photoURL,
		},
	})
}