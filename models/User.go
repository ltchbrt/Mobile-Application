package models

import "gorm.io/gorm"

type User struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Number  string
	ENC     string
	Deleted gorm.DeletedAt
}
