package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/betosardinha/antixy-api/internal/adapters/database/repositories"
	"github.com/betosardinha/antixy-api/internal/adapters/http"
	usersvc "github.com/betosardinha/antixy-api/internal/application/user"
)

type App struct {
	Router *gin.Engine
}

func New(db *gorm.DB) *App {
	router := gin.Default()

	userRepository := repositories.NewUserRepository(db)

	userService := usersvc.NewService(userRepository)

	handler := &http.Handler{
		UserSevice: userService,
	}

	http.RegisterRoutes(router, handler)

	return &App{
		Router: router,
	}
}
