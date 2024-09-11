package models

import (
	"gorm.io/gorm"
)

type Store struct {
	*gorm.Model
	URL   string
	Short string `gorm:"unique"`
	Owner User   `gorm:"foreignKey:UserId"`
}
