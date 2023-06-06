package service

import (
	application "dissent-api-service/pkg/application/entities"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newUserInformation(r *mux.Router) {
	r.HandleFunc("/user/test-information/{text}", testUseInformationHandler).Methods(http.MethodGet)
}

func testUseInformationHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("UserInformationTest: %v", text))

	writeReponse(w, body)
}
