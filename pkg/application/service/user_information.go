package service

import (
	application "dissent-api-service/pkg/application/entities"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newUserInformation(r *mux.Router) {
	r.HandleFunc("/test/{text}", testHandler).Methods("GET")
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("test: %v", text))

	writeReponse(w, r, body)
}
