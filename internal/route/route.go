package route

import (
	"takah-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Takah berjalan",
		})
	})

	auth := api.Group("/auth")
	auth.POST("/login", handler.Login)

	surat := api.Group("/surat")
	surat.GET("", handler.GetAllSurat)
	surat.POST("", handler.CreateSurat)
	surat.GET("/:id", handler.GetSuratByID)
	surat.PUT("/:id", handler.UpdateSurat)
	surat.DELETE("/:id", handler.DeleteSurat)
}
