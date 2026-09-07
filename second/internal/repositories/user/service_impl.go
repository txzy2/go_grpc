package user

import (
	"second/handler/http/v1/user"
	"second/internal/domain"
)

// TODO: ДОбавить репо
type service struct {
	//userRepo IUserRepo
}

func newUserService() IUserService {
	return &service{}
}

func (ser *service) CreateUser(dto user.CreateUserDTO) (*domain.User, error) {
	return nil, nil
}
