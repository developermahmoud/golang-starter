package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uint64         `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	EmailVerifiedAt  time.Time      `json:"email_verified_at"`
	MobileVerifiedAt time.Time      `json:"mobile_verified_at"`
	Email            string         `json:"email" gorm:"type:varchar(100);unique"`
	Mobile           string         `json:"mobile" gorm:"type:varchar(30);unique"`
	CountryCode      string         `json:"country_code" gorm:"type:char(2)"`
	Password         []byte         `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
