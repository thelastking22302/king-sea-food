package model

type Roles int

const (
	MEMBERS Roles = iota
	ADMIN
)

func (r Roles) String() string {
	return []string{"MEMBERS", "ADMIN"}[r]
}
