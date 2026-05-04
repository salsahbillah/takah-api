package main

import (
	"takah-api/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	route.SetupRoutes(router)

	router.Run(":8080")
}