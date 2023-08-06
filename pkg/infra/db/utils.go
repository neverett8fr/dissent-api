package db

import (
	"github.com/deta/deta-go/service/base"
)

const (
	userColumnID           = "id"
	userColumnUsername     = "username"
	userColumnPasswordHash = "password_hash"
	userTableUsers         = "users"

	eventColumnTitle     = "title"
	eventColumnDesc      = "description"
	eventColumnLocation  = "location"
	eventColumnDate      = "date"
	eventColumnOrganiser = "organiser_id"
	eventsTable          = "events"
)

type DBConn struct {
	Conn map[string]*base.Base
}

func NewDBConnFromExisting(baseArr map[string]*base.Base) *DBConn {
	return &DBConn{
		Conn: baseArr,
	}
}
