package food_bussiness

import (
	"context"
	"thelastking/kingseafood/model/food"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, data *food.InvoiceFood) error
	GetInvoice(ctx context.Context, id map[string]interface{}) (*food.InvoiceFood, error)
	UpdateInvoice(ctx context.Context, id map[string]interface{}, data *food.InvoiceFood) error
	DeleteInvoice(ctx context.Context, id map[string]interface{}) error
}
type invoiceController struct {
	o InvoiceService
}

func NewInvoiceController(o InvoiceService) *invoiceController {
	return &invoiceController{o: o}
}
func (c *invoiceController) NewCreateInvoice(ctx context.Context, data *food.InvoiceFood) error {
	if err := c.o.CreateInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}
func (c *invoiceController) NewGetInvoiceTable(ctx context.Context, id string) (*food.InvoiceFood, error) {
	data, err := c.o.GetInvoice(ctx, map[string]interface{}{"invoice_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *invoiceController) NewUpdateInvoice(ctx context.Context, id string, data *food.InvoiceFood) error {
	if err := c.o.UpdateInvoice(ctx, map[string]interface{}{"invoice_id": id}, data); err != nil {
		return err
	}
	return nil
}
func (c *invoiceController) NewDeleteInvoiceTable(ctx context.Context, id string) error {
	if err := c.o.DeleteInvoice(ctx, map[string]interface{}{"invoice_id": id}); err != nil {
		return err
	}
	return nil
}
