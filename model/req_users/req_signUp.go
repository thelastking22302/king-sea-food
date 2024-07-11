package req_users

type RequestSignUp struct {
	FullName string `json:"full_name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Male     string `json:"male,omitempty" validate:"required"`
}
