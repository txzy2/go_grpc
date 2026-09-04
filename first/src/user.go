package src

import (
	"errors"
)

type User struct {
	id   int    `json:"id"`
	name string `json:"name"`

	email    string `json:"email"`
	password string `json:"-"`
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("name must be not empty")
	}

	return &User{name: name}, nil
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetEmail() string { return u.email }
