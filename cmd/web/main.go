package main

import (
	"takah-api/internal/database"
	"takah-api/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {

	// Koneksi ke database MySQL
	database.ConnectDatabase()

	router := gin.Default()

	route.SetupRoutes(router)

	router.Run(":8080")
}