package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name           string
	Email          string
	PhoneNumber    string
	ProfilePicture string
	GithubURL      string
	LinkedinURL    string
	TwitterURL     string
	Summary        string
	Title          string
}
