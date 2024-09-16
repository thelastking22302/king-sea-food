package req_users

type UpdateUsers struct {
	FullName string `json:"full_name" validate:"required"`
	Male     string `json:"male" validate:"required"`
}
