package cmd

import (
	"fmt"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	eventsTable = "events"
	userTable   = "users"
)

func OpenDB() (*deta.Deta, error) {

	d, err := deta.New()
	if err != nil {
		return nil, fmt.Errorf("error initialising Deta instance, err %v", err)
	}

	return d, nil
}

func GetBaseArr(d *deta.Deta) (map[string]*base.Base, error) {

	dbUsers, err := base.New(d, userTable)
	if err != nil {
		return nil, fmt.Errorf("error getting db, err %v", err)
	}

	dbEvents, err := base.New(d, eventsTable)
	if err != nil {
		return nil, fmt.Errorf("error getting db, err %v", err)
	}

	return map[string]*base.Base{
		eventsTable: dbEvents,
		userTable:   dbUsers,
	}, nil
}
