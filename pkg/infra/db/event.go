package db

import (
	"dissent-api-service/pkg/infra/entities"
	"fmt"
	"log"
)

func (conn *DBConn) CreateEvent(event entities.Event) error {

	if !event.CheckValid() {
		err := fmt.Errorf("error creating user, entity not valid")
		log.Println(err)
		return err
	}

	if event.ID != 0 {
		err := fmt.Errorf("error username exists")
		log.Println(err)
		return err
	}

	_, err := conn.Conn.Exec(
		fmt.Sprintf(
			"INSERT INTO %s(%s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5)",
			eventsTable, eventColumnTitle, eventColumnDesc, eventColumnDate, eventColumnLocation, eventColumnOrganiser,
		),
		event.Title, event.Description, event.Date, event.Location, event.Organiser,
	)
	if err != nil {
		err := fmt.Errorf("error inserting new event, err %v", err)
		log.Println(err)
		return err
	}

	return nil
}
