package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	usersvc "github.com/betosardinha/antixy-api/internal/application/user"
	domain "github.com/betosardinha/antixy-api/internal/domain/user"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type userResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func createUser(service *usersvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := domain.User{
			Name:           req.Name,
			Email:          req.Email,
			PasswordDigest: req.Password,
		}

		created, err := service.CreateUser(c, u)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, toUserResponse(created))
	}
}

func toUserResponse(u domain.User) userResponse {
	return userResponse{
		ID:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func registerUserRoutes(r *gin.Engine, service *usersvc.Service) {
	users := r.Group("/users")

	users.POST("/", createUser(service))
}
