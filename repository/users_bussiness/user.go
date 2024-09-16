package users_bussiness

import (
	"context"
	"errors"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/food"
	"thelastking/kingseafood/model/req_users"
	"thelastking/kingseafood/pkg/logger"
)

type UserService interface {
	DeletedUser(ctx context.Context, data map[string]interface{}) error
	ProfileUserByID(ctx context.Context, id map[string]interface{}) (*model.Users, error)
	SignIn(ctx context.Context, data *req_users.RequestSignIn) (*model.Users, error)
	SignUp(ctx context.Context, data *model.Users) (*model.Users, error)
	UpdateUser(ctx context.Context, update *req_users.UpdateUsers, id map[string]interface{}) error
	HistoryPurchases(ctx context.Context, id map[string]interface{}) (*model.Users, []food.OrderItem, error)
	ChangePwdUser(ctx context.Context, id map[string]interface{}, upd *req_users.ChangePwd) error
}

type userController struct {
	u      UserService
	logger logger.Logger
}

func NewUserController(u UserService) *userController {
	return &userController{
		u:      u,
		logger: logger.GetLogger(),
	}
}
func (u *userController) NewChangePwdUser(ctx context.Context, id string, upd *req_users.ChangePwd) error {
	if err := u.u.ChangePwdUser(ctx, map[string]interface{}{"user_id": id}, upd); err != nil {
		u.logger.Errorf("Change password failed for user ID %s: %v", id, err)
		return errors.New("change password user db fail")
	}
	return nil
}
func (u *userController) NewUpdateUser(ctx context.Context, update *req_users.UpdateUsers, id string) error {
	if err := u.u.UpdateUser(ctx, update, map[string]interface{}{"user_id": id}); err != nil {
		u.logger.Errorf("Update user failed for user ID %s: %v", id, err)
		return errors.New("update user db fail")
	}
	u.logger.Infof("User updated successfully: %s", id)
	return nil
}
func (d *userController) NewDeletedUserByID(ctx context.Context, id string) error {
	if err := d.u.DeletedUser(ctx, map[string]interface{}{"user_id": id}); err != nil {
		d.logger.Errorf("Failed to delete user with ID %s: %v", id, err)
		return errors.New("deleted user failed")
	}
	d.logger.Infof("Deleted user with ID %s successfully", id)
	return nil
}
func (c *userController) NewProfileUserByID(ctx context.Context, id string) (*model.Users, error) {
	dataUser, err := c.u.ProfileUserByID(ctx, map[string]interface{}{"user_id": id})
	if err != nil {
		c.logger.Errorf("Error retrieving profile for user ID %s: %v", id, err)
		return nil, errors.New("error retrieving profile")
	}
	return dataUser, nil
}
func (si *userController) NewSignIn(ctx context.Context, data *req_users.RequestSignIn) (*model.Users, error) {
	if data.Email == "" {
		si.logger.Warnf("SignIn attempt with empty email")
		return nil, errors.New("SignIn error")
	}
	dataUsers, err := si.u.SignIn(ctx, data)
	if err != nil {
		si.logger.Errorf("SignIn failed for email %s: %v", data.Email, err)
		return nil, errors.New("SignIn failed")
	}
	si.logger.Infof("User signed in successfully: %s", data.Email)
	return dataUsers, nil
}
func (s *userController) NewSignUp(ctx context.Context, data *model.Users) (*model.Users, error) {
	if data.Email == "" {
		s.logger.Warnf("SignUp attempt with empty email")
		return nil, errors.New("invalid data")
	}
	dataSignUp, err := s.u.SignUp(ctx, data)
	if err != nil {
		s.logger.Errorf("SignUp failed for email %s: %v", data.Email, err)
		return nil, errors.New("failed to sign up")
	}
	s.logger.Infof("User signed up successfully: %s", data.Email)
	return dataSignUp, nil
}
func (u *userController) NewHistoryPurchases(ctx context.Context, id string) (*model.Users, []food.OrderItem, error) {
	dataUser, dataProduct, err := u.u.HistoryPurchases(ctx, map[string]interface{}{"user_id": id})
	if err != nil {
		u.logger.Errorf("Error retrieving purchase history for user ID %s: %v", id, err)
		return nil, nil, err
	}
	u.logger.Infof("Retrieved purchase history for user ID %s", id)
	return dataUser, dataProduct, nil
}
