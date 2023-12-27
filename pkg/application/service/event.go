package service

import (
	application "dissent-api-service/pkg/application/entities"
	"dissent-api-service/pkg/infra/entities"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newEventOperation(r *mux.Router) {

	eventRouter := r.PathPrefix("/event").Subrouter()

	eventRouterWithAuth := r.PathPrefix("/event").Subrouter() // temp
	eventRouterWithAuth.Use(AuthMiddleware)

	eventRouter.HandleFunc("/test/{text}", testEventHandler).Methods(http.MethodGet)
	eventRouterWithAuth.HandleFunc("", createEventHandler).Methods(http.MethodPost)
	eventRouterWithAuth.HandleFunc("/{query}", getEventsHandler).Methods(http.MethodGet)
}

func testEventHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("EventTest: %v", text))

	writeReponse(w, body)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve the event information from the request context
	eventInformation := r.Context().Value(ctxEventBody).(newEventIn)

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
		usr.Username,
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

	_, err = DBConn.CreateEvent(event)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(fmt.Sprintf("event created with title %v", eventInformation.Title), err)

	writeReponse(w, body)
}

func getEventsHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	query := params["query"]

	events, err := DBConn.GetEvents(query)

	body := application.NewResponse(fmt.Sprintf("events %v", events), err)

	writeReponse(w, body)
}
