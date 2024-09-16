package repouserimpl

import (
	"context"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/model/req_users"

	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSql(db *gorm.DB) *sql {
	return &sql{db: db}
}

func (s *sql) ChangePwdUser(ctx context.Context, id map[string]interface{}, upd *req_users.ChangePwd) error {
	if err := s.db.Table("users").Where(id).Updates(upd).Error; err != nil {
		return err
	}
	return nil
}

func (sql *sql) SignUp(ctx context.Context, data *model.Users) (*model.Users, error) {
	if err := sql.db.Table("users").FirstOrCreate(&data, &model.Users{Email: data.Email}).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (sql *sql) SignIn(ctx context.Context, data *req_users.RequestSignIn) (*model.Users, error) {
	var dataUsers model.Users
	if err := sql.db.Table("users").Where("email = ?", data.Email).First(&dataUsers).Error; err != nil {
		return nil, err
	}
	return &dataUsers, nil
}
func (sql *sql) ProfileUserByID(ctx context.Context, id map[string]interface{}) (*model.Users, error) {
	var data model.Users
	if err := sql.db.Table("users").Where(id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (sql *sql) UpdateUser(ctx context.Context, update *req_users.UpdateUsers, id map[string]interface{}) error {
	if err := sql.db.Table("users").Where(id).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (sql *sql) DeletedUser(ctx context.Context, data map[string]interface{}) error {
	if err := sql.db.Table("users").Where(data).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (sql *sql) HistoryPurchases(ctx context.Context, id map[string]interface{}) (*model.Users, []food.OrderItem, error) {
	var dataUsers *model.Users
	var dataProducts []food.OrderItem
	if err := sql.db.Table("users").
		Select("users.user_id, order_items.product_id").
		Joins("JOIN orders ON users.user_id = orders.user_id").
		Joins("JOIN order_items ON orders.order_id = order_items.order_id").
		Where("users.user_id = ?", id["user_id"]).
		Group("users.user_id, order_items.product_id").
		Scan(&dataProducts).Error; err != nil {
		return nil, nil, err
	}
	return dataUsers, dataProducts, nil
}
