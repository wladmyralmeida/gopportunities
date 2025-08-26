package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		// Define your API routes here
		api.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Get One Opening API",
			})
		})

		api.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "List Openings API",
			})
		})

		api.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Create Opening",
			})
		})

		api.DELETE("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, nil)
		})

		api.PUT("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Update Opening",
			})
		})
	}
}
