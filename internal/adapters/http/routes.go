package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", healthCheck)
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Antixy API running",
	})
}
