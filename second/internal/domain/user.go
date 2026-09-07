package domain

import (
	"second/pkg/enums"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Email    string `json:"email"`
	password string `json:"-"`

	Status enums.StatusState `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) SetPassword(rawPassword string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashed)
	return nil
}

func (u *User) CheckPassword(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(rawPassword))
	return err == nil
}
