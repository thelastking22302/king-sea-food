package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/req_users"
)

type SignInService interface {
	SignIn(ctx context.Context, data *req_users.RequestSignIn) (*model.Users, error)
}

type SignInController struct {
	s SignInService
}

func NewSignInController(s SignInService) *SignInController {
	return &SignInController{s: s}
}
func (si *SignInController) NewSignIn(ctx context.Context, data *req_users.RequestSignIn) (*model.Users, error) {
	if data.Email == "" {
		return nil, errors.New("SignIn error")
	}
	dataUsers, err := si.s.SignIn(ctx, data)
	if err != nil {
		return nil, errors.New("SignIn failed")
	}
	return dataUsers, nil
}
