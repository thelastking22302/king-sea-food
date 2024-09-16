package food_bussiness

import (
	"context"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, data *food.InvoiceFood) error
	GetInvoice(ctx context.Context, id map[string]interface{}) (*food.InvoiceFood, error)
	UpdateInvoice(ctx context.Context, id map[string]interface{}, data *food.InvoiceFood) error
	DeleteInvoice(ctx context.Context, id map[string]interface{}) error
}

type invoiceController struct {
	o      InvoiceService
	logger logger.Logger
}

func NewInvoiceController(o InvoiceService) *invoiceController {
	return &invoiceController{
		o:      o,
		logger: logger.GetLogger(), // Sử dụng logger singleton
	}
}

func (c *invoiceController) NewCreateInvoice(ctx context.Context, data *food.InvoiceFood) error {
	if err := c.o.CreateInvoice(ctx, data); err != nil {
		c.logger.Errorf("Failed to create invoice: %v", err)
		return err
	}
	c.logger.Infof("Invoice created successfully: %+v", data)
	return nil
}

func (c *invoiceController) NewGetInvoiceTable(ctx context.Context, id string) (*food.InvoiceFood, error) {
	data, err := c.o.GetInvoice(ctx, map[string]interface{}{"invoice_id": id})
	if err != nil {
		c.logger.Errorf("Failed to get invoice with ID %s: %v", id, err)
		return nil, err
	}
	c.logger.Infof("Retrieved invoice: %+v", data)
	return data, nil
}

func (c *invoiceController) NewUpdateInvoice(ctx context.Context, id string, data *food.InvoiceFood) error {
	if err := c.o.UpdateInvoice(ctx, map[string]interface{}{"invoice_id": id}, data); err != nil {
		c.logger.Errorf("Failed to update invoice with ID %s: %v", id, err)
		return err
	}
	c.logger.Infof("Invoice with ID %s updated successfully: %+v", id, data)
	return nil
}

func (c *invoiceController) NewDeleteInvoiceTable(ctx context.Context, id string) error {
	if err := c.o.DeleteInvoice(ctx, map[string]interface{}{"invoice_id": id}); err != nil {
		c.logger.Errorf("Failed to delete invoice with ID %s: %v", id, err)
		return err
	}
	c.logger.Infof("Invoice with ID %s deleted successfully", id)
	return nil
}
