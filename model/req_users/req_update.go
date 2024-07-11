package req_users

type UpdateUsers struct {
	FullName string `json:"full_name,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Male     string `json:"male,omitempty" validate:"required"`
}
