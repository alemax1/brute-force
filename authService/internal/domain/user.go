package domain

type User struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Username string `json:"username" validate:"required,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
