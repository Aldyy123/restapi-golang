package models

import (
	"gorm.io/gorm"
)

type Users struct {
	ID       int `gorm:"primaryKey;autoIncrement:true"`
	Username string
	Password string
	Token    string
	gorm.Model
}

type Books struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Title       string
	Description string
}
