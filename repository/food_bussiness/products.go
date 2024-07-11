package food_bussiness

import (
	"context"
	"errors"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

type ProductsService interface {
	CreateProducts(ctx context.Context, data *food.Product) error
	GetProducts(ctx context.Context, id map[string]interface{}) (*food.Product, error)
	GetProductByName(ctx context.Context, name map[string]interface{}) (*food.Product, error)
	GetProductsList(ctx context.Context, pagging *common.Paggings, morekeys ...string) ([]food.Product, error)
	UpdateProducts(ctx context.Context, data *food.Product, id map[string]interface{}) error
	DeleteProducts(ctx context.Context, id map[string]interface{}) error
}

type ProductsController struct {
	c ProductsService
}

func NewProductsController(c ProductsService) *ProductsController {
	return &ProductsController{c: c}
}

func (f ProductsController) NewCreateProducts(ctx context.Context, data *food.Product) error {
	if err := f.c.CreateProducts(ctx, data); err != nil {
		return errors.New("create fail bussiness")
	}
	return nil
}

func (f ProductsController) NewGetProducts(ctx context.Context, id string) (*food.Product, error) {
	data, err := f.c.GetProducts(ctx, map[string]interface{}{"product_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f ProductsController) NewGetProductsList(ctx context.Context, pagging *common.Paggings) ([]food.Product, error) {
	listData, err := f.c.GetProductsList(ctx, pagging)
	if err != nil {
		return nil, err
	}
	return listData, nil
}

func (f ProductsController) NewUpdateProducts(ctx context.Context, data *food.Product, id string) error {
	if err := f.c.UpdateProducts(ctx, data, map[string]interface{}{"product_id": id}); err != nil {
		return err
	}
	return nil
}

func (f ProductsController) NewDeleteProducts(ctx context.Context, id string) error {
	if err := f.c.DeleteProducts(ctx, map[string]interface{}{"product_id": id}); err != nil {
		return err
	}
	return nil
}
func (f ProductsController) NewGetProductByName(ctx context.Context, name string) (*food.Product, error) {
	data, err := f.c.GetProductByName(ctx, map[string]interface{}{"title": name})
	if err != nil {
		return nil, err
	}
	return data, nil
}
