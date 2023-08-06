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
	Username       string `json:"key"`
	PasswordHashed string `json:"password_hashed"`
}

type Event struct {
	ID          string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
	Organiser   string `json:"organiser_id"`
}

func (e *Event) CheckValid() bool {

	return true
}

func NewEvent(organiser string, title string, desc string, loc string, date string) (Event, error) {
	event := Event{
		Title:       title,
		Description: desc,
		Location:    loc,
		Date:        date,
		Organiser:   organiser,
	}

	return event, nil
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
