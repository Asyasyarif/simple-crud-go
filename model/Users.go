package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
