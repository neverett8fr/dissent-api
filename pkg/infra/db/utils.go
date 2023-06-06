package db

import "database/sql"

// INSERT INTO users (email, password) VALUES (
// 	'johndoe@mail.com',
// 	crypt('johnspassword', gen_salt('bf'))
// );

// bf=blowfish

// SELECT id
//   FROM users
//  WHERE email = 'johndoe@mail.com'
//    AND password = crypt('johnspassword', password);

const (
	userColumnID           = "id"
	userColumnUsername     = "username"
	userColumnPasswordHash = "password_hash"

	userTableUsers = "users"
)

type DBConn struct {
	Conn *sql.DB
}

func NewDBConnFromExisting(conn *sql.DB) *DBConn {
	return &DBConn{
		Conn: conn,
	}
}
