package repoimpl

import "gorm.io/gorm"

type sql struct {
	db *gorm.DB
}

func NewSql(db *gorm.DB) *sql {
	return &sql{db: db}
}
