package food_bussiness

import (
	"context"
	"errors"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
)

type MenuService interface {
	CreateMenu(ctx context.Context, data *food.MenuFood) error
	GetMenu(ctx context.Context, id map[string]interface{}) (*food.MenuFood, error)
	GetListMenu(ctx context.Context, pagging *common.Paggings) ([]food.MenuFood, error)
	UpdateFoodMenu(ctx context.Context, id map[string]interface{}, data *food.MenuFood) error
	DeleteFoodMenu(ctx context.Context, id map[string]interface{}) error
	ViewProductFromMenu(ctx context.Context, id map[string]interface{}) ([]food.Product, error)
}

type menuController struct {
	m      MenuService
	logger logger.Logger
}

func NewMenuController(m MenuService) *menuController {
	return &menuController{
		m:      m,
		logger: logger.GetLogger(),
	}
}

func (f *menuController) NewCreateMenu(ctx context.Context, data *food.MenuFood) error {
	if err := f.m.CreateMenu(ctx, data); err != nil {
		f.logger.Errorf("Failed to create menu: %v", err)
		return errors.New("create fail business")
	}
	f.logger.Infof("Menu created successfully: %+v", data)
	return nil
}

func (f *menuController) NewGetMenu(ctx context.Context, id string) (*food.MenuFood, error) {
	data, err := f.m.GetMenu(ctx, map[string]interface{}{"menu_id": id})
	if err != nil {
		f.logger.Errorf("Failed to get menu with ID %s: %v", id, err)
		return nil, err
	}
	f.logger.Infof("Retrieved menu: %+v", data)
	return data, nil
}

func (f *menuController) NewGetListMenu(ctx context.Context, pagging *common.Paggings) ([]food.MenuFood, error) {
	listData, err := f.m.GetListMenu(ctx, pagging)
	if err != nil {
		f.logger.Errorf("Failed to get menu list: %v", err)
		return nil, err
	}
	f.logger.Infof("Retrieved menu list: %d menus found", len(listData))
	return listData, nil
}

func (f *menuController) NewUpdateFoodMenu(ctx context.Context, id string, data *food.MenuFood) error {
	if err := f.m.UpdateFoodMenu(ctx, map[string]interface{}{"menu_id": id}, data); err != nil {
		f.logger.Errorf("Failed to update menu with ID %s: %v", id, err)
		return err
	}
	f.logger.Infof("Menu with ID %s updated successfully: %+v", id, data)
	return nil
}

func (f *menuController) NewDeleteFoodMenu(ctx context.Context, id string) error {
	if err := f.m.DeleteFoodMenu(ctx, map[string]interface{}{"menu_id": id}); err != nil {
		f.logger.Errorf("Failed to delete menu with ID %s: %v", id, err)
		return err
	}
	f.logger.Infof("Menu with ID %s deleted successfully", id)
	return nil
}

func (f *menuController) NewViewProductFromMenu(ctx context.Context, id string) ([]food.Product, error) {
	listdata, err := f.m.ViewProductFromMenu(ctx, map[string]interface{}{"menu_id": id})
	if err != nil {
		f.logger.Errorf("Failed to view products from menu with ID %s: %v", id, err)
		return nil, err
	}
	f.logger.Infof("Retrieved products from menu ID %s: %d products found", id, len(listdata))
	return listdata, nil
}
