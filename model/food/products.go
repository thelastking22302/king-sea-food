package food

import "time"

type Product struct {
	Product_ID  string     `json:"product_id" gorm:"column:product_id;"`
	Title       *string    `json:"title" validate:"required,min=2,max=100" gorm:"column:title;"`
	Image       *string    `json:"image" validate:"required" gorm:"column:image;"`
	Description string     `json:"description" validate:"required" gorm:"column:description;"`
	Price       *float64   `json:"price" validate:"required" gorm:"price;"`
	Status      *string    `json:"status" validate:"required" gorm:"status;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Menu_ID     string     `json:"menu_id" validate:"required" gorm:"column:menu_id;"`
}
