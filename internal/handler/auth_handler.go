package handler

import (
	"net/http"

	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request model.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email dan password wajib diisi dengan benar",
		})
		return
	}

	if request.Email != "admin@takah.com" || request.Password != "password123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"data": gin.H{
			"token": "dummy-token",
		},
	})
}
