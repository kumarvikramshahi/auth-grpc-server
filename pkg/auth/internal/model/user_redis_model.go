package model

type User struct {
	Name     string `redis:"name" validate:"required,min=3,max=100"`
	Email    string `redis:"email" validate:"required,min=3,max=100"`
	Password string `redis:"password" validate:"required,min=3,max=50"`
}
