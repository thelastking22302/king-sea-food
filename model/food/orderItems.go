package food

import "time"

type OrderItem struct {
	Order_item_id string    `json:"order_item_id" gorm:"column:order_item_id;"`
	Order_id      string    `json:"order_id" validate:"required" gorm:"column:order_id;"`
	Quantity      *int      `json:"quantity" validate:"required" gorm:"column:quantity;"`
	Unit_price    *float64  `json:"unit_price" validate:"required" gorm:"column:unit_price;"`
	Created_at    time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at    time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Product_ID    string    `json:"product_id" validate:"required" gorm:"column:product_id;"`
}
