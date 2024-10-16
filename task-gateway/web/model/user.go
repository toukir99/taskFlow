package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username"`
	Email    int    `json:"email" validate:"required"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
