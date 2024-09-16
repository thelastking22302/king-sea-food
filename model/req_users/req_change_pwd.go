package req_users

type ChangePwd struct {
	Email         string `json:"email" gorm:"column:email;" validate:"required,email"`
	Password_user string `json:"password_user" gorm:"column:password_user;" validate:"required,min=8"`
}
