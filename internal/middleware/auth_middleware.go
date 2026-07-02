package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("takah-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		userName, _ := claims["name"].(string)
		userEmail, _ := claims["email"].(string)
		userRole, _ := claims["role"].(string)

		c.Set("user_id", int(userID))
		c.Set("user_name", userName)
		c.Set("user_email", userEmail)
		c.Set("user_role", userRole)

		c.Next()
	}
}