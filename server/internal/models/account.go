package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name      string    `json:"name"`
	Balance   int64     `json:"balance"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	OAuth     string    `json:"oauth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
