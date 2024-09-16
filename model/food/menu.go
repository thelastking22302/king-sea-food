package food

import "time"

type MenuFood struct {
	Menu_ID    string     `json:"menu_id" gorm:"column:menu_id;"`
	Name       *string    `json:"name_menu" validate:"required" gorm:"column:name_menu;"`
	Category   *string    `json:"category" validate:"required" gorm:"column:category;"`
	Created_At *time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_At *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
