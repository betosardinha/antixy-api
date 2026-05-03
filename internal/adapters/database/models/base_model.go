package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = uuid.NewV7()
	return err
}
