package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model"
)

type SignUpStorages interface {
	SignUp(ctx context.Context, data *model.Users) (*model.Users, error)
}

type signUpController struct {
	s SignUpStorages
}

func NewSignUpController(s SignUpStorages) *signUpController {
	return &signUpController{s: s}
}

func (s *signUpController) NewSignUp(ctx context.Context, data *model.Users) (*model.Users, error) {
	if data.Email == "" {
		return nil, errors.New("khong hop le")
	}
	dataSignUp, err := s.s.SignUp(ctx, data)
	if err != nil {
		return nil, errors.New("failed to sign up")
	}
	return dataSignUp, nil
}
