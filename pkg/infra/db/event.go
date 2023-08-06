package db

import (
	"dissent-api-service/pkg/infra/entities"
	"fmt"
	"log"
)

func (conn *DBConn) CreateEvent(event entities.Event) (string, error) {

	if !event.CheckValid() {
		err := fmt.Errorf("error creating user, entity not valid")
		log.Println(err)
		return "", err
	}

	if event.ID != "" {
		err := fmt.Errorf("error username exists")
		log.Println(err)
		return "", err
	}

	key, err := conn.Conn[eventsTable].Put(event)
	if err != nil {
		err := fmt.Errorf("error inserting new event, err %v", err)
		log.Println(err)
		return "", err
	}

	return key, nil
}
