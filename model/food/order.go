package food

import "time"

type Order struct {
	Order_id   string    `json:"order_id" gorm:"column:order_id;"`
	Table_id   *string   `json:"table_id" validate:"required" gorm:"column:table_id;"`
	Order_date time.Time `json:"order_date" validate:"required" gorm:"column:order_date;"`
	Created_at time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
