package service

import (
	"database/sql"
	"dissent-api-service/pkg/config"
	"dissent-api-service/pkg/infra/auth"
	"dissent-api-service/pkg/infra/db"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	DBConn        *db.DBConn
	TokenProvider auth.TokenProvider
)

func NewServiceRoutes(r *mux.Router, conn *sql.DB, conf config.Config) {
	DBConn = db.NewDBConnFromExisting(conn)
	TokenProvider = auth.InitialiseTokenProvider(conf.Service.HMACSigningKey, DBConn)

	newUserInformation(r)
}

func writeReponse(w http.ResponseWriter, r *http.Request, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
