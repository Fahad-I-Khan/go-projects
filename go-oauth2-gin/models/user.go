package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	FullName  string
	AvatarURL string
	Provider  string // like "google"
}
