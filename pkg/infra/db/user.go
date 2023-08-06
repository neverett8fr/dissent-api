package db

import (
	"dissent-api-service/pkg/infra/entities"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (conn *DBConn) CreateUser(user entities.User) (string, error) {

	if !user.CheckValid() {
		err := fmt.Errorf("error creating user, entity not valid")
		log.Println(err)
		return "", err
	}

	u, err := conn.ReadUserID(user.Username)
	if err == nil {
		err := fmt.Errorf("error, expected no results / error - username must already exist")
		log.Println(err)
		return "", err
	}
	if u.Username != "" {
		err := fmt.Errorf("error username exists")
		log.Println(err)
		return "", err
	}

	key, err := conn.Conn[userTableUsers].Put(user)
	if err != nil {
		err := fmt.Errorf("error inserting new user, err %v", err)
		log.Println(err)
		return "", err
	}

	return key, nil
}

func (conn *DBConn) ReadUserID(username string) (entities.User, error) {

	var user entities.User
	err := conn.Conn[userTableUsers].Get(username, &user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (conn *DBConn) VerifyPassword(username string, password string) error {

	u, err := conn.ReadUserID(username)
	if err != nil {
		log.Printf("error verifying password, err %v", err)
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(password))
}
