package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Email    string `json:"email"`
	password string `json:"-"`
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
