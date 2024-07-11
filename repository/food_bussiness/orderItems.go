package food_bussiness

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

type OrderItemService interface {
	GetOrderItems(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error)
	GetOrderItemsByOder(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error)
	GetOrderItemsByProduct(ctx context.Context, id map[string]interface{}, pagging *common.Paggings) ([]food.OrderItem, error)
	CreateOrderItem(ctx context.Context, data *food.OrderItem) error
	UpdateOrderItem(ctx context.Context, id map[string]interface{}, data *food.OrderItem) error
}
type orderItemController struct {
	o OrderItemService
}

func NewOrderItemController(o OrderItemService) *orderItemController {
	return &orderItemController{o: o}
}
func (c *orderItemController) NewGetOrderItems(ctx context.Context, id string) (*food.OrderItem, error) {
	data, err := c.o.GetOrderItems(ctx, map[string]interface{}{"order_item_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *orderItemController) NewCreateOrderItem(ctx context.Context, data *food.OrderItem) error {
	if err := c.o.CreateOrderItem(ctx, data); err != nil {
		return err
	}
	return nil
}
func (c *orderItemController) NewUpdateOrderItem(ctx context.Context, id string, data *food.OrderItem) error {
	if err := c.o.UpdateOrderItem(ctx, map[string]interface{}{"order_item_id": id}, data); err != nil {
		return err
	}
	return nil
}
func (c *orderItemController) NewGetOrderItemsByOder(ctx context.Context, id string) (*food.OrderItem, error) {
	data, err := c.o.GetOrderItemsByOder(ctx, map[string]interface{}{"order_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *orderItemController) NewGetOrderItemsByProduct(ctx context.Context, id string, pagging *common.Paggings) ([]food.OrderItem, error) {
	dataList, err := c.o.GetOrderItemsByProduct(ctx, map[string]interface{}{"product_id": id}, pagging)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}
