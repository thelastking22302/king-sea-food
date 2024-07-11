package repoimpl

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

func (s *sql) CreateProducts(ctx context.Context, data *food.Product) error {
	var menu food.MenuFood
	if err := s.db.Table("menus").Where("menu_id = ?", data.Menu_ID).First(&menu).Error; err != nil {
		return err
	}
	if err := s.db.Table("products").Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) GetProducts(ctx context.Context, id map[string]interface{}) (*food.Product, error) {
	var dataProduct food.Product
	if err := s.db.Table("products").Where(id).First(&dataProduct).Error; err != nil {
		return nil, err
	}
	return &dataProduct, nil
}
func (s *sql) GetProductsList(ctx context.Context, pagging *common.Paggings, morekeys ...string) ([]food.Product, error) {
	var data []food.Product
	if err := s.db.Table("products").Count(&pagging.Total).Error; err != nil {
		return nil, err
	}
	if err := s.db.Table("products").Where("status <> ?", "Deleted").
		Order("product_id desc").
		Offset((pagging.Page - 1) * pagging.Limit).Limit(pagging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sql) UpdateProducts(ctx context.Context, data *food.Product, id map[string]interface{}) error {
	if err := s.db.Table("products").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) DeleteProducts(ctx context.Context, id map[string]interface{}) error {
	if err := s.db.Table("products").Where(id).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) GetProductByName(ctx context.Context, name map[string]interface{}) (*food.Product, error) {
	var data food.Product
	if err := s.db.Table("products").Where(name).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
