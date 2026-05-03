package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Name           string
	Email          string
	PasswordDigest string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
