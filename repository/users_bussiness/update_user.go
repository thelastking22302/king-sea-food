package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model/req_users"
)

type NewUpdateUser interface {
	UpdateUser(ctx context.Context, user *req_users.UpdateUsers, data map[string]interface{}) error
}

type updateUserController struct {
	u NewUpdateUser
}

func NewUpdateUserController(u NewUpdateUser) *updateUserController {
	return &updateUserController{u: u}
}

func (u updateUserController) NewUpdateUser(ctx context.Context, user *req_users.UpdateUsers, id string) error {
	if err := u.u.UpdateUser(ctx, user, map[string]interface{}{"user_id": id}); err != nil {
		return errors.New("update user db fail")
	}
	return nil
}
