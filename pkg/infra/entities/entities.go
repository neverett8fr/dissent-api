package entities

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	PasswordHashed string `json:"password_hashed"`
}

func (u *User) CheckValid() bool {

	if len(u.Username) < 5 || u.PasswordHashed == "" {
		return false
	}
	return true
}

func NewUser(username string, password string) (User, error) {

	if len(username) < 5 {
		return User{}, fmt.Errorf("error username is too short")
	}
	if len(password) < 5 {
		return User{}, fmt.Errorf("error password is too short")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("error hashing password, err %v", err)
	}

	u := User{
		Username:       username,
		PasswordHashed: string(passwordHashed),
	}
	if !u.CheckValid() {
		return User{}, fmt.Errorf("error creating user, configuration is not valid")
	}

	return u, nil
}
