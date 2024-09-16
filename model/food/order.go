package food

import "time"

type Order struct {
	Order_id     string    `json:"order_id" gorm:"column:order_id;"`
	UserID       *string   `json:"user_id" validate:"required" gorm:"column:user_id;"`
	Total_price  *float64  `json:"total_price" gorm:"total_price;"`
	Order_date   time.Time `json:"order_date" validate:"required" gorm:"column:order_date;"`
	Order_Status *string   `json:"order_status" gorm:"column:order_status;" validate:"eq=processing|eq=success|eq=defeat"`
	Created_at   time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at   time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
