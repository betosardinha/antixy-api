package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	usersvc "github.com/betosardinha/antixy-api/internal/application/user"
)

type Handler struct {
	UserSevice *usersvc.Service
}

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.GET("/", healthCheck)

	registerUserRoutes(r, h.UserSevice)
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Antixy API running",
	})
}
