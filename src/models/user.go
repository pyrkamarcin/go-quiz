package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model `json:"model"`
	Email      string `json:"email" json:"unique;not null"`
	Password   string `json:"password"`
}
