package models

import (
	"time"

	"gorm.io/gorm"
)

type Type struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     JSONB          `json:"title"`
	Type      string         `json:"type"`
}
