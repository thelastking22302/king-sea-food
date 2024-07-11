package food_bussiness

import (
	"context"
	"errors"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

type MenuService interface {
	CreateMenu(ctx context.Context, data *food.MenuFood) error
	GetMenu(ctx context.Context, id map[string]interface{}) (*food.MenuFood, error)
	GetListMenu(ctx context.Context, pagging *common.Paggings) ([]food.MenuFood, error)
	UpdateFoodMenu(ctx context.Context, id map[string]interface{}, data *food.MenuFood) error
	DeleteFoodMenu(ctx context.Context, id map[string]interface{}) error
}

type menuController struct {
	m MenuService
}

func NewMenuController(m MenuService) *menuController {
	return &menuController{m: m}
}

func (f menuController) NewCreateMenu(ctx context.Context, data *food.MenuFood) error {
	if err := f.m.CreateMenu(ctx, data); err != nil {
		return errors.New("create fail bussiness")
	}
	return nil
}

func (f menuController) NewGetMenu(ctx context.Context, id string) (*food.MenuFood, error) {
	data, err := f.m.GetMenu(ctx, map[string]interface{}{"menu_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f menuController) NewGetListMenu(ctx context.Context, pagging *common.Paggings) ([]food.MenuFood, error) {
	listData, err := f.m.GetListMenu(ctx, pagging)
	if err != nil {
		return nil, err
	}
	return listData, nil
}

func (f menuController) NewUpdateFoodMenu(ctx context.Context, id string, data *food.MenuFood) error {
	if err := f.m.UpdateFoodMenu(ctx, map[string]interface{}{"menu_id": id}, data); err != nil {
		return err
	}
	return nil
}
func (f menuController) NewDeleteFoodMenu(ctx context.Context, id string) error {
	if err := f.m.DeleteFoodMenu(ctx, map[string]interface{}{"menu_id": id}); err != nil {
		return err
	}
	return nil
}
