package food_bussiness

import (
	"context"

	"thelastking/kingseafood/controller/common"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/pkg/logger"
)

type NewTableService interface {
	CreateTable(ctx context.Context, data *food.Table) error
	GetTable(ctx context.Context, id map[string]interface{}) (*food.Table, error)
	GetListTable(ctx context.Context, pagging *common.Paggings) ([]food.Table, error)
	UpdateTable(ctx context.Context, id map[string]interface{}, data *food.Table) error
	DeleteTable(ctx context.Context, id map[string]interface{}) error
}

type tableController struct {
	m      NewTableService
	logger logger.Logger
}

func NewTableController(m NewTableService) *tableController {
	return &tableController{
		m:      m,
		logger: logger.GetLogger(),
	}
}

func (t *tableController) NewCreateTable(ctx context.Context, data *food.Table) error {
	if err := t.m.CreateTable(ctx, data); err != nil {
		t.logger.Errorf("Failed to create table: %v", err)
		return err
	}
	t.logger.Infof("Table created successfully: %+v", data)
	return nil
}

func (t *tableController) NewGetTable(ctx context.Context, id string) (*food.Table, error) {
	data, err := t.m.GetTable(ctx, map[string]interface{}{"table_id": id})
	if err != nil {
		t.logger.Errorf("Failed to get table with ID %s: %v", id, err)
		return nil, err
	}
	t.logger.Infof("Retrieved table: %+v", data)
	return data, nil
}

func (t *tableController) NewGetListTable(ctx context.Context, pagging *common.Paggings) ([]food.Table, error) {
	dataList, err := t.m.GetListTable(ctx, pagging)
	if err != nil {
		t.logger.Errorf("Failed to get list of tables: %v", err)
		return nil, err
	}
	t.logger.Infof("Retrieved list of tables: %d tables found", len(dataList))
	return dataList, nil
}

func (t *tableController) NewUpdateTable(ctx context.Context, id string, data *food.Table) error {
	if err := t.m.UpdateTable(ctx, map[string]interface{}{"table_id": id}, data); err != nil {
		t.logger.Errorf("Failed to update table with ID %s: %v", id, err)
		return err
	}
	t.logger.Infof("Table with ID %s updated successfully: %+v", id, data)
	return nil
}

func (t *tableController) NewDeleteTable(ctx context.Context, id string) error {
	if err := t.m.DeleteTable(ctx, map[string]interface{}{"table_id": id}); err != nil {
		t.logger.Errorf("Failed to delete table with ID %s: %v", id, err)
		return err
	}
	t.logger.Infof("Table with ID %s deleted successfully", id)
	return nil
}
