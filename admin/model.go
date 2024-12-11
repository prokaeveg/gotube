package admin

import "time"

type ListUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	SecondName string `json:"second_name"`
}
