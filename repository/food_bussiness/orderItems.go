package food_bussiness

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
)

type OrderItemService interface {
	GetOrderItems(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error)
	GetOrderItemsByOder(ctx context.Context, id map[string]interface{}) (*food.OrderItem, error)
	GetOrderItemsByProduct(ctx context.Context, id map[string]interface{}, pagging *common.Paggings) ([]food.OrderItem, error)
	CreateOrderItem(ctx context.Context, data *food.OrderItem) error
	UpdateOrderItem(ctx context.Context, id map[string]interface{}, data *food.OrderItem) error
}

type orderItemController struct {
	o      OrderItemService
	logger logger.Logger
}

func NewOrderItemController(o OrderItemService) *orderItemController {
	return &orderItemController{
		o:      o,
		logger: logger.GetLogger(), // Sử dụng logger singleton
	}
}

func (c *orderItemController) NewGetOrderItems(ctx context.Context, id string) (*food.OrderItem, error) {
	data, err := c.o.GetOrderItems(ctx, map[string]interface{}{"order_item_id": id})
	if err != nil {
		c.logger.Errorf("Failed to get order item with ID %s: %v", id, err)
		return nil, err
	}
	c.logger.Infof("Retrieved order item: %+v", data)
	return data, nil
}

func (c *orderItemController) NewCreateOrderItem(ctx context.Context, data *food.OrderItem) error {
	if err := c.o.CreateOrderItem(ctx, data); err != nil {
		c.logger.Errorf("Failed to create order item: %v", err)
		return err
	}
	c.logger.Infof("Order item created successfully: %+v", data)
	return nil
}

func (c *orderItemController) NewUpdateOrderItem(ctx context.Context, id string, data *food.OrderItem) error {
	if err := c.o.UpdateOrderItem(ctx, map[string]interface{}{"order_item_id": id}, data); err != nil {
		c.logger.Errorf("Failed to update order item with ID %s: %v", id, err)
		return err
	}
	c.logger.Infof("Order item with ID %s updated successfully: %+v", id, data)
	return nil
}

func (c *orderItemController) NewGetOrderItemsByOder(ctx context.Context, id string) (*food.OrderItem, error) {
	data, err := c.o.GetOrderItemsByOder(ctx, map[string]interface{}{"order_id": id})
	if err != nil {
		c.logger.Errorf("Failed to get order items by order ID %s: %v", id, err)
		return nil, err
	}
	c.logger.Infof("Retrieved order item by order ID %s: %+v", id, data)
	return data, nil
}

func (c *orderItemController) NewGetOrderItemsByProduct(ctx context.Context, id string, pagging *common.Paggings) ([]food.OrderItem, error) {
	dataList, err := c.o.GetOrderItemsByProduct(ctx, map[string]interface{}{"product_id": id}, pagging)
	if err != nil {
		c.logger.Errorf("Failed to get order items by product ID %s: %v", id, err)
		return nil, err
	}
	c.logger.Infof("Retrieved order items by product ID %s: %d items found", id, len(dataList))
	return dataList, nil
}
