package main

import (
	"dissent-api-service/cmd"
	application "dissent-api-service/pkg/application/service"
	"dissent-api-service/pkg/config"
	"log"
	"reflect"

	"github.com/deta/deta-go/service/base"
	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes(conf config.Config, baseArr map[string]*base.Base) *mux.Router {
	r := mux.NewRouter()
	application.NewServiceRoutes(r, conf, baseArr)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	serviceDB, err := cmd.OpenDB()
	if err != nil {
		log.Fatalf("error starting db, err %v", err)
		return
	}
	log.Printf("configuration of DB setup with key %v", serviceDB.ProjectKey)

	dbArr, err := cmd.GetBaseArr(serviceDB)
	if err != nil {
		log.Fatalf("error getting DB arr, %v", err)
		return
	}
	log.Printf("DB array returned with names %v", reflect.ValueOf(dbArr).MapKeys())

	router := getRoutes(*conf, dbArr)
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}
