package user

import (
	"second/handler/http/v1/user"
	"second/internal/domain"
)

// TODO: ДОбавить репо
type UserService struct {
	//userRepo IUserRepo
}

func newUserService() IUserService {
	return &UserService{}
}

func (ser *UserService) CreateUser(dto user.CreateUserDTO) (*domain.User, error) {
	return nil, nil
}
