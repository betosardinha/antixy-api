package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/betosardinha/antixy-api/internal/adapters/http"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func New(database *gorm.DB) *App {
	router := gin.Default()

	http.RegisterRoutes(router)

	return &App{
		Router: router,
		DB:     database,
	}
}
