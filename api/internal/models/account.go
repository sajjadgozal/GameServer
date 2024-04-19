package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Balance  int64  `json:"balance"`
	Password string `json:"password"`
	Email    string `json:"email"`
	OAuth    string `json:"oauth"`
}
