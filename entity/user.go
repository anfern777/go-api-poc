package entity

import "gorm.io/gorm"

type Role string

type User struct {
	gorm.Model
	Privilege Role   `json:"privilege" binding:"required" validate:"is-role"`
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  string `json:"-" binding:"required" validate:"is-valid-password"`
}
