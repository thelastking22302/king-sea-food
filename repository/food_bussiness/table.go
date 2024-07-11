package food_bussiness

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
)

type NewTableService interface {
	CreateTable(ctx context.Context, data *food.Table) error
	GetTable(ctx context.Context, id map[string]interface{}) (*food.Table, error)
	GetListTable(ctx context.Context, pagging *common.Paggings) ([]food.Table, error)
	UpdateTable(ctx context.Context, id map[string]interface{}, data *food.Table) error
	DeleteTable(ctx context.Context, id map[string]interface{}) error
}
type tableController struct {
	m NewTableService
}

func NewTableController(m NewTableService) *tableController {
	return &tableController{m: m}
}
func (t tableController) NewCreateTable(ctx context.Context, data *food.Table) error {
	if err := t.m.CreateTable(ctx, data); err != nil {
		return err
	}
	return nil
}

func (t tableController) NewGetTable(ctx context.Context, id string) (*food.Table, error) {
	data, err := t.m.GetTable(ctx, map[string]interface{}{"table_id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t tableController) NewGetListTable(ctx context.Context, pagging *common.Paggings) ([]food.Table, error) {
	dataList, err := t.m.GetListTable(ctx, pagging)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}

func (t tableController) NewUpdateTable(ctx context.Context, id string, data *food.Table) error {
	if err := t.m.UpdateTable(ctx, map[string]interface{}{"table_id": id}, data); err != nil {
		return err
	}
	return nil
}
func (t tableController) NewDeleteTable(ctx context.Context, id string) error {
	if err := t.m.DeleteTable(ctx, map[string]interface{}{"table_id": id}); err != nil {
		return err
	}
	return nil
}
