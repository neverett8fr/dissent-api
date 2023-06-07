package service

import (
	application "dissent-api-service/pkg/application/entities"
	"dissent-api-service/pkg/infra/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func newEventOperation(r *mux.Router) {

	eventRouter := r.PathPrefix("/event").Subrouter()

	eventRouterWithAuth := r.PathPrefix("/event").Subrouter() // temp
	eventRouterWithAuth.Use(AuthMiddleware)

	eventRouter.HandleFunc("/test/{text}", testEventHandler).Methods(http.MethodGet)
	eventRouterWithAuth.HandleFunc("", createEventHandler).Methods(http.MethodPost)
}

func testEventHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("EventTest: %v", text))

	writeReponse(w, body)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// Parse the request body into an event struct
	eventInformation := newEventIn{}
	err = json.Unmarshal(bodyIn, &eventInformation)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// Get id from username
	usr, err := DBConn.ReadUserID(eventInformation.Organiser)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// Create a new event object
	event, err := entities.NewEvent(
		usr,
		eventInformation.Title,
		eventInformation.Description,
		eventInformation.Location,
		eventInformation.Date,
	)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	err = DBConn.CreateEvent(event)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(fmt.Sprintf("event created with title %v", eventInformation.Title), err)

	writeReponse(w, body)
}
