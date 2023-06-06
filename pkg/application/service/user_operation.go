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

func newUserOperation(r *mux.Router) {

	r.HandleFunc("/user/test-operation/{text}", testUserOperationHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", createUserHandler).Methods(http.MethodPost)
}

func testUserOperationHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("UserOperationTest: %v", text))

	writeReponse(w, body)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}
	userInformation := newUserIn{}
	err = json.Unmarshal(bodyIn, &userInformation)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}
	user, err := entities.NewUser(userInformation.Username, userInformation.Password)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	err = DBConn.CreateUser(user)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(fmt.Sprintf("user created with username %v", userInformation.Username), err)

	writeReponse(w, body)
}
