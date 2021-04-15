package models

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" csv:"email" binding:"required,max=200"`
}
