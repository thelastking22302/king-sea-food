package food_bussiness

import (
	"context"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
)

type OrderService interface {
	CreateOrderTable(ctx context.Context, data *food.Order) error
	GetOrderTable(ctx context.Context, id map[string]interface{}) (*food.Order, error)
	UpdateOrder(ctx context.Context, id map[string]interface{}, data *food.Order) error
	DeleteOrderTable(ctx context.Context, id map[string]interface{}) error
}

type orderController struct {
	o      OrderService
	logger logger.Logger
}

func NewOrderController(o OrderService) *orderController {
	return &orderController{
		o:      o,
		logger: logger.GetLogger(), // Sử dụng logger singleton
	}
}

func (c *orderController) NewCreateOrderTable(ctx context.Context, data *food.Order) error {
	if err := c.o.CreateOrderTable(ctx, data); err != nil {
		c.logger.Errorf("Failed to create order: %v", err)
		return err
	}
	c.logger.Infof("Order created successfully: %+v", data)
	return nil
}

func (c *orderController) NewGetOrderTable(ctx context.Context, id string) (*food.Order, error) {
	data, err := c.o.GetOrderTable(ctx, map[string]interface{}{"order_id": id})
	if err != nil {
		c.logger.Errorf("Failed to get order with ID %s: %v", id, err)
		return nil, err
	}
	c.logger.Infof("Retrieved order: %+v", data)
	return data, nil
}

func (c *orderController) NewUpdateOrder(ctx context.Context, id string, data *food.Order) error {
	if err := c.o.UpdateOrder(ctx, map[string]interface{}{"order_id": id}, data); err != nil {
		c.logger.Errorf("Failed to update order with ID %s: %v", id, err)
		return err
	}
	c.logger.Infof("Order with ID %s updated successfully: %+v", id, data)
	return nil
}

func (c *orderController) NewDeleteOrderTable(ctx context.Context, id string) error {
	if err := c.o.DeleteOrderTable(ctx, map[string]interface{}{"order_id": id}); err != nil {
		c.logger.Errorf("Failed to delete order with ID %s: %v", id, err)
		return err
	}
	c.logger.Infof("Order with ID %s deleted successfully", id)
	return nil
}
