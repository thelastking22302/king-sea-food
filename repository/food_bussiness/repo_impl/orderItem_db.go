package repoimpl

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

func (s *sql) GetOrderItems(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error) {
	var dataItems food.OrderItem
	if err := s.db.Table("order_items").Where(id).First(&dataItems).Error; err != nil {
		return nil, err
	}
	return &dataItems, nil
}
func (s *sql) CreateOrderItem(ctx context.Context, data *food.OrderItem) error {
	// var dataOrder food.Order
	// if err := s.db.Table("orders").Where("order_id = ?", data.Order_id).First(&dataOrder).Error; err != nil {
	// 	return err
	// }
	// var dataFood food.Product
	// if err := s.db.Table("products").Where("product_id = ?", data.Food_id).Find(&dataFood).Error; err != nil {
	// 	return err
	// }
	if err := s.db.Table("order_items").Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) UpdateOrderItem(ctx context.Context, id map[string]interface{}, data *food.OrderItem) error {
	if err := s.db.Table("order_items").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) GetOrderItemsByOder(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error) {
	var data food.OrderItem
	if err := s.db.Table("order_items").Where(id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
func (s *sql) GetOrderItemsByProduct(ctx context.Context, id map[string]interface{}, pagging *common.Paggings) ([]food.OrderItem, error) {
	var data []food.OrderItem
	if err := s.db.Table("order_items").Count(&pagging.Total).Error; err != nil {
		return nil, err
	}
	if err := s.db.Table("order_items").
		Order("order_item_id desc").
		Offset((pagging.Page - 1) * pagging.Page).
		Limit(pagging.Limit).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
