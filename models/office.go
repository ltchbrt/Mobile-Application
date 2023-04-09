package models

import "gorm.io/gorm"

type Office struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Number  string
	Address string
	User    User
	UserID  string
	Deleted gorm.DeletedAt
}
