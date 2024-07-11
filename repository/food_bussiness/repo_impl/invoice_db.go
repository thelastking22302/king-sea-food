package repoimpl

import (
	"context"
	"thelastking/kingseafood/model/food"
)

func (s *sql) CreateInvoice(ctx context.Context, data *food.InvoiceFood) error {
	var dataOrder food.Order
	if err := s.db.Table("orders").Where("order_id = ?", data.Order_ID).First(&dataOrder).Error; err != nil {
		return err
	}
	if err := s.db.Table("invoicefood").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) GetInvoice(ctx context.Context, id map[string]interface{}) (*food.InvoiceFood, error) {
	var data *food.InvoiceFood
	if err := s.db.Table("invoicefood").Where(id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sql) UpdateInvoice(ctx context.Context, id map[string]interface{}, data *food.InvoiceFood) error {
	if err := s.db.Table("invoicefood").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) DeleteInvoice(ctx context.Context, id map[string]interface{}) error {
	var data *food.InvoiceFood
	if err := s.db.Table("invoicefood").Where(id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
