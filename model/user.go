package model

import "time"

type Users struct {
	UserID        string     `json:"user_id" gorm:"column:user_id;"`
	FullName      string     `json:"full_name" gorm:"column:full_name;" validate:"required"`
	Email         string     `json:"email" gorm:"column:email;" validate:"required,email"`
	Password_user string     `json:"password_user" gorm:"column:password_user;" validate:"required,min=8"`
	Male          string     `json:"male" gorm:"column:male;" validate:"required"`
	Role          string     `json:"-" gorm:"column:role_user;"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
