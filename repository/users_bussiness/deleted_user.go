package users_bussiness

import (
	"context"
	"errors"
)

type NewDeletedUserController interface {
	DeletedUser(ctx context.Context, data map[string]interface{}) error
}

type deletedUserService struct {
	d NewDeletedUserController
}

func NewDeletedUserService(d NewDeletedUserController) *deletedUserService {
	return &deletedUserService{d: d}
}

func (d *deletedUserService) NewDeletedUserByID(ctx context.Context, id string) error {
	if err := d.d.DeletedUser(ctx, map[string]interface{}{"user_id": id}); err != nil {
		return errors.New("deleted user faild")
	}
	return nil
}
