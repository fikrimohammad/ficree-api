package models

import "gorm.io/gorm"

// Skill Model
type Skill struct {
	gorm.Model
	Name        string
	Description string
	Rating      int
	UserID      int
	User        User
}
