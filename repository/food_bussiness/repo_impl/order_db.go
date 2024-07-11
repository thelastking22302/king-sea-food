package repoimpl

import (
	"context"
	"thelastking/kingseafood/model/food"
)

func (s *sql) CreateOrderTable(ctx context.Context, data *food.Order) error {

	if err := s.db.Table("orders").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) GetOrderTable(ctx context.Context, id map[string]interface{}) (*food.Order, error) {
	var data *food.Order
	if err := s.db.Table("orders").Where(id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sql) UpdateOrder(ctx context.Context, id map[string]interface{}, data *food.Order) error {

	if err := s.db.Table("orders").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) DeleteOrderTable(ctx context.Context, id map[string]interface{}) error {
	var data *food.Order
	if err := s.db.Table("orders").Where(id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
