package domain

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" `
}
