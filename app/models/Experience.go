package models

import (
	"time"

	"gorm.io/gorm"
)

// Experience Model
type Experience struct {
	gorm.Model
	PositionName string
	StartsAt     time.Time
	EndsAt       time.Time
	Description  string
	CompanyName  string
	Location     string
	UserID       int
	User         User
}
