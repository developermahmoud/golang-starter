package models

import "time"

type Token struct {
	ID        uint64    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
	UserID    uint64    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
