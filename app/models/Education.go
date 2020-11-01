package models

import (
	"time"

	"gorm.io/gorm"
)

// Education Model
type Education struct {
	gorm.Model
	InstitutionName    string
	InstitutionIconURL string
	InstitutionWebURL  string
	Description        string
	StartsAt           time.Time
	EndsAt             time.Time
	Degree             string
	StudyField         string
	UserID             int
}

// Educations represents array of Education
type Educations []Education
