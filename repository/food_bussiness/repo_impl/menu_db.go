package repoimpl

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

func (s *sql) CreateMenu(ctx context.Context, data *food.MenuFood) error {
	if err := s.db.Table("menus").Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) GetMenu(ctx context.Context, id map[string]interface{}) (*food.MenuFood, error) {
	var data food.MenuFood
	if err := s.db.Table("menus").Where(id).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *sql) GetListMenu(ctx context.Context, pagging *common.Paggings) ([]food.MenuFood, error) {
	var data []food.MenuFood
	if err := s.db.Table("menus").Count(&pagging.Total).Error; err != nil {
		return nil, err
	}
	if err := s.db.Table("menus").
		Order("menu_id desc").
		Offset((pagging.Page - 1) * pagging.Limit).
		Limit(pagging.Limit).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sql) UpdateFoodMenu(ctx context.Context, id map[string]interface{}, data *food.MenuFood) error {
	if err := s.db.Table("menus").Where(id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sql) DeleteFoodMenu(ctx context.Context, id map[string]interface{}) error {
	var data food.MenuFood
	if err := s.db.Table("menus").Where(id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sql) ViewProductFromMenu(ctx context.Context, id map[string]interface{}) ([]food.Product, error) {
	var data []food.Product
	if err := s.db.Table("menus").Select("products.*").
		Joins("JOIN products ON products.menu_id = menus.menu_id ").
		//Khi sử dụng JOIN, bạn cần chỉ định rõ ràng tên bảng trước tên cột trong mệnh đề WHERE
		//để tránh lỗi "ambiguous column reference"
		Where("menus.menu_id = ?", id["menu_id"]). // Chỉ định rõ ràng menus.menu_id
		Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
