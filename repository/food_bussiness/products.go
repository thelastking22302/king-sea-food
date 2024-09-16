package food_bussiness

import (
	"context"
	"errors"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
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
	c      ProductsService
	logger logger.Logger
}

func NewProductsController(c ProductsService) *ProductsController {
	return &ProductsController{
		c:      c,
		logger: logger.GetLogger(), // Sử dụng logger singleton
	}
}

func (f *ProductsController) NewCreateProducts(ctx context.Context, data *food.Product) error {
	if err := f.c.CreateProducts(ctx, data); err != nil {
		f.logger.Errorf("Failed to create product: %v", err)
		return errors.New("create fail bussiness")
	}
	f.logger.Infof("Product created successfully: %+v", data)
	return nil
}

func (f *ProductsController) NewGetProducts(ctx context.Context, id string) (*food.Product, error) {
	data, err := f.c.GetProducts(ctx, map[string]interface{}{"product_id": id})
	if err != nil {
		f.logger.Errorf("Failed to get product with ID %s: %v", id, err)
		return nil, err
	}
	f.logger.Infof("Retrieved product: %+v", data)
	return data, nil
}

func (f *ProductsController) NewGetProductsList(ctx context.Context, pagging *common.Paggings) ([]food.Product, error) {
	listData, err := f.c.GetProductsList(ctx, pagging)
	if err != nil {
		f.logger.Errorf("Failed to get product list: %v", err)
		return nil, err
	}
	f.logger.Infof("Retrieved product list: %d products found", len(listData))
	return listData, nil
}

func (f *ProductsController) NewUpdateProducts(ctx context.Context, data *food.Product, id string) error {
	if err := f.c.UpdateProducts(ctx, data, map[string]interface{}{"product_id": id}); err != nil {
		f.logger.Errorf("Failed to update product with ID %s: %v", id, err)
		return err
	}
	f.logger.Infof("Product with ID %s updated successfully: %+v", id, data)
	return nil
}

func (f *ProductsController) NewDeleteProducts(ctx context.Context, id string) error {
	if err := f.c.DeleteProducts(ctx, map[string]interface{}{"product_id": id}); err != nil {
		f.logger.Errorf("Failed to delete product with ID %s: %v", id, err)
		return err
	}
	f.logger.Infof("Product with ID %s deleted successfully", id)
	return nil
}

func (f *ProductsController) NewGetProductByName(ctx context.Context, name string) (*food.Product, error) {
	data, err := f.c.GetProductByName(ctx, map[string]interface{}{"title": name})
	if err != nil {
		f.logger.Errorf("Failed to get product by name %s: %v", name, err)
		return nil, err
	}
	f.logger.Infof("Retrieved product by name: %+v", data)
	return data, nil
}
