package food

import "time"

type MenuFood struct {
	Menu_ID    string     `json:"menu_id" gorm:"column:menu_id;"`
	Name       *string    `json:"name" validate:"required" gorm:"column:name;"`
	Category   *string    `json:"category" validate:"required" gorm:"column:category;"`
	Start_Day  *time.Time `json:"start_day" gorm:"column:start_day;"`
	End_Day    *time.Time `json:"end_day" gorm:"column:end_day;"`
	Created_At *time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_At *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Products   []*Product `json:"products" gorm:"-"`
}
