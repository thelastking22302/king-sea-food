package req_users

type RequestSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password_user" validate:"required"`
}
