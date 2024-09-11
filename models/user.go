package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserName string `gorm:"unique; not null"`
	Email    string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null"`
}
