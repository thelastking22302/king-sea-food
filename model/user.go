package model

import "time"

type Users struct {
	UserID    string     `json:"user_id" gorm:"column:user_id;"`
	FullName  string     `json:"full_name" gorm:"column:full_name;"`
	Email     string     `json:"email" gorm:"column:email;"`
	Password  string     `json:"-" gorm:"column:password;"`
	Male      string     `json:"male" gorm:"column:male;"`
	Role      string     `json:"-" gorm:"column:role;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
