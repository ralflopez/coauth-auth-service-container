package models

type Role string

const (
	Member Role = "Member"
	Guest  Role = "Guest"
	Admin  Role = "Admin"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=8"`
	Role     Role   `json:"role"`
}
