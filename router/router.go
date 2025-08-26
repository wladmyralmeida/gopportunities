package router

import "github.com/gin-gonic/gin"

func SetupRouter() {
	// Initialize Router
	r := gin.Default()

	// Define Routes
	initializeRoutes(r)

	// Run the server
	r.Run(":8080")
}
