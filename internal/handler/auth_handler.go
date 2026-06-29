package handler

import (
	"database/sql"
	"net/http"
	"time"

	"takah-api/internal/database"
	"takah-api/internal/middleware"
	"takah-api/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var request model.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email dan password wajib diisi dengan benar",
		})
		return
	}

	var userID int
	var name string
	var email string
	var password string
	var role string

	err := database.DB.QueryRow(`
		SELECT id, name, email, password, role
		FROM users
		WHERE email = ?
	`, request.Email).Scan(
		&userID,
		&name,
		&email,
		&password,
		&role,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data user",
			"error":   err.Error(),
		})
		return
	}

	if request.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	claims := jwt.MapClaims{
		"id":    userID,
		"name":  name,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(middleware.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"data": gin.H{
			"token": tokenString,
			"user": gin.H{
				"id":    userID,
				"name":  name,
				"email": email,
				"role":  role,
			},
		},
	})
}