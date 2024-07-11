package repoimpl

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

func (s *sql) CreateTable(ctx context.Context, data *food.Table) error {
	if err := s.db.Table("tablefood").Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) GetTable(ctx context.Context, id map[string]interface{}) (*food.Table, error) {
	var table food.Table
	if err := s.db.Table("tablefood").Where(id).First(&table).Error; err != nil {
		return nil, err
	}
	return &table, nil
}
func (s *sql) GetListTable(ctx context.Context, pagging *common.Paggings) ([]food.Table, error) {
	var tablelist []food.Table
	if err := s.db.Table("tablefood").Count(&pagging.Total).Error; err != nil {
		return nil, err
	}
	if err := s.db.Table("tablefood").
		Order("table_id desc").
		Offset((pagging.Page - 1) * pagging.Limit).
		Limit(pagging.Limit).
		Find(&tablelist).Error; err != nil {
		return nil, err
	}
	return tablelist, nil
}
func (s *sql) UpdateTable(ctx context.Context, id map[string]interface{}, data *food.Table) error {
	if err := s.db.Table("tablefood").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) DeleteTable(ctx context.Context, id map[string]interface{}) error {
	var data food.Table
	if err := s.db.Table("tablefood").Where(id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
