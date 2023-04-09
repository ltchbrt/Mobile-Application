package models

import "gorm.io/gorm"

type Contact struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Number       string
	Relationship string
	Address      string
	User         User
	UserID       string
	Deleted      gorm.DeletedAt
}
