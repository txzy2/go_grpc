package user

import (
	"second/handler/http/v1/user"
	"second/internal/domain"
)

type IUserService interface {
	CreateUser(dto user.CreateUserDTO) (*domain.User, error)
}
