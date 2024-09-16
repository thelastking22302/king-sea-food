package food

import "time"

type Table struct {
	Table_id         string     `json:"table_id" gorm:"colum:table_id;"`
	UserID           *string    `json:"user_id" validate:"required" gorm:"column:user_id;"`
	Number_of_guests *string    `json:"number_of_guests" validate:"required" gorm:"colum:number_of_guests;"`
	Table_number     *int       `json:"table_number" validate:"required" gorm:"colum:table_guests;"`
	Created_at       *time.Time `json:"created_at" gorm:"colum:created_at;"`
	Updated_at       *time.Time `json:"updated_at" gorm:"colum:updated_at;"`
}
