package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model"
)

type ProfileUserService interface {
	ProfileUserByID(ctx context.Context, id map[string]interface{}) (*model.Users, error)
}

type profileController struct {
	p ProfileUserService
}

func NewProfileController(p ProfileUserService) *profileController {
	return &profileController{p: p}
}
func (c *profileController) NewProfileUserByID(ctx context.Context, id string) (*model.Users, error) {
	dataUser, err := c.p.ProfileUserByID(ctx, map[string]interface{}{"user_id": id})
	if err != nil {
		return nil, errors.New("loi tai profileUserById")
	}
	return dataUser, nil
}
