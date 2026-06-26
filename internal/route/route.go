package route

import (
	"takah-api/internal/handler"
	"takah-api/internal/middleware"

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

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	surat := protected.Group("/surat")
	surat.GET("", handler.GetAllSurat)
	surat.POST("", handler.CreateSurat)
	surat.GET("/:id", handler.GetSuratByID)
	surat.PUT("/:id", handler.UpdateSurat)
	surat.DELETE("/:id", handler.DeleteSurat)

	takah := protected.Group("/takah")
	takah.GET("", handler.GetAllTakah)
	takah.POST("", handler.CreateTakah)
	takah.GET("/:id", handler.GetTakahByID)
	takah.PUT("/:id", handler.UpdateTakah)
	takah.DELETE("/:id", handler.DeleteTakah)

	configNomor := protected.Group("/config-nomor")
	configNomor.GET("", handler.GetAllConfigNomor)
	configNomor.POST("", handler.CreateConfigNomor)
	configNomor.GET("/:id", handler.GetConfigNomorByID)
	configNomor.PUT("/:id", handler.UpdateConfigNomor)
	configNomor.DELETE("/:id", handler.DeleteConfigNomor)

	suratKeluar := protected.Group("/surat-keluar")
	suratKeluar.GET("", handler.GetAllSuratKeluar)
	suratKeluar.POST("", handler.CreateSuratKeluar)
	suratKeluar.GET("/:id", handler.GetSuratKeluarByID)
	suratKeluar.PUT("/:id", handler.UpdateSuratKeluar)
	suratKeluar.DELETE("/:id", handler.DeleteSuratKeluar)

	approval := protected.Group("/approval")
	approval.GET("", handler.GetAllApproval)
	approval.POST("", handler.CreateApproval)
	approval.GET("/:id", handler.GetApprovalByID)
	approval.PUT("/:id", handler.UpdateApproval)

	monitoring := protected.Group("/monitoring")
	monitoring.GET("", handler.GetAllMonitoring)
	monitoring.POST("", handler.CreateMonitoring)
	monitoring.GET("/:id", handler.GetMonitoringByID)
	monitoring.PUT("/:id", handler.UpdateMonitoring)
	monitoring.DELETE("/:id", handler.DeleteMonitoring)

	templateSurat := protected.Group("/template-surat")
	templateSurat.GET("", handler.GetAllTemplateSurat)
	templateSurat.POST("", handler.CreateTemplateSurat)
	templateSurat.GET("/:id", handler.GetTemplateSuratByID)
	templateSurat.PUT("/:id", handler.UpdateTemplateSurat)
	templateSurat.DELETE("/:id", handler.DeleteTemplateSurat)

	suratMasuk := protected.Group("/surat-masuk")
	suratMasuk.GET("", handler.GetAllSuratMasuk)
	suratMasuk.POST("", handler.CreateSuratMasuk)
	suratMasuk.GET("/:id", handler.GetSuratMasukByID)
	suratMasuk.PUT("/:id", handler.UpdateSuratMasuk)
	suratMasuk.DELETE("/:id", handler.DeleteSuratMasuk)
}