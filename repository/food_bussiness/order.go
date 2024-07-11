package food_bussiness

import (
	"context"
	"thelastking/kingseafood/model/food"
)

type OrderService interface {
	CreateOrderTable(ctx context.Context, data *food.Order) error
	GetOrderTable(ctx context.Context, id map[string]interface{}) (*food.Order, error)
	UpdateOrder(ctx context.Context, id map[string]interface{}, data *food.Order) error
	DeleteOrderTable(ctx context.Context, id map[string]interface{}) error
}
type orderController struct {
	o OrderService
}

func NewOrderController(o OrderService) *orderController {
	return &orderController{o: o}
}
func (c *orderController) NewCreateOrderTable(ctx context.Context, data *food.Order) error {
	if err := c.o.CreateOrderTable(ctx, data); err != nil {
		return err
	}
	return nil
}
func (c *orderController) NewGetOrderTable(ctx context.Context, id string) (*food.Order, error) {
	data, err := c.o.GetOrderTable(ctx, map[string]interface{}{"order_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *orderController) NewUpdateOrder(ctx context.Context, id string, data *food.Order) error {
	if err := c.o.UpdateOrder(ctx, map[string]interface{}{"order_id": id}, data); err != nil {
		return err
	}
	return nil
}
func (c *orderController) NewDeleteOrderTable(ctx context.Context, id string) error {
	if err := c.o.DeleteOrderTable(ctx, map[string]interface{}{"order_id": id}); err != nil {
		return err
	}
	return nil
}
