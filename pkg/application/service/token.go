package service

import (
	application "dissent-api-service/pkg/application/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newTokenOperation(r *mux.Router) {

	r.HandleFunc("/token/test/{text}", testAuthHandler).Methods(http.MethodGet)

	r.HandleFunc("/token", createToken).Methods(http.MethodPost)
	r.HandleFunc("/token", checkToken).Methods(http.MethodGet)
}

func testAuthHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("TokenOperationTest: %v", text))

	writeReponse(w, body)
}

func createToken(w http.ResponseWriter, r *http.Request) {

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error creating token, err %v", err)
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}
	userInformation := newUserIn{}
	err = json.Unmarshal(bodyIn, &userInformation)
	if err != nil {
		log.Printf("error creating token, err %v", err)
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	tok, err := TokenProvider.NewToken(userInformation.Username, userInformation.Password)
	if err != nil {
		log.Printf("error creating token, err %v", err)
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tok,
	// 	Expires: time.Now().Add(time.Hour),
	// })

	body := application.NewResponse(tokenOut{Token: tok}, err)

	writeReponse(w, body)
}

func checkToken(w http.ResponseWriter, r *http.Request) {
	// check headers or vars
	reqToken := r.Header.Get(headerAuth)

	err := TokenProvider.CheckToken(reqToken)
	if err != nil {
		log.Printf("error checking token, err %v", err)
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusUnauthorized)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse("token is valid", err)
	writeReponse(w, body)
}
