package service

import (
	application "dissent-api-service/pkg/application/entities"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// for UK & EU law we need a Privacy Policy
// UK one must include:
// tell people what rights they have in relation to use of their data
//   e.g. access, rectification, erasure, restriction, objection, data portability
// and tell them how we use the data, if we share it, etc..

// for quick reference:
//  for users we will only collect a username and a hashed password
//  for events we will only collect when an event is created,
//		the date of the event, description, location,
//		the organiser (username), and title.

func newLegal(r *mux.Router) {

	r.HandleFunc("/legal/test/{text}", testLegalHandler).Methods(http.MethodGet)
	r.HandleFunc("/legal/privacy", tempPrivacyHandler).Methods(http.MethodGet)
	r.HandleFunc("/legal/canary", tempCanaryPolicy).Methods(http.MethodGet)

}

func testLegalHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := application.NewResponse(fmt.Sprintf("LegalTest: %v", text))

	writeReponse(w, body)
}

func tempPrivacyHandler(w http.ResponseWriter, r *http.Request) {

	policy := "<html><head><title>Privacy</title></head><body>" +
		"We will never sell your personal information to any other 3rd party company at any time<br/>" +
		"for users we will store: a username, a hashed password<br/>" +
		"for each event we will store: the date/time the event is created, the date of the event, the description of the event, the location of the event, the organiser (their username) and the title of the event<br/>" +
		"</body></html>"

	_, err := w.Write([]byte(policy))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}

func tempCanaryPolicy(w http.ResponseWriter, r *http.Request) {

	policy := "<html><head><title>Canary Policy</title></head><body>" +
		"we positively confirm the following (12th November 2023):<br/>" +
		"  to the best of our knowledge, we have not been compromised or suffered a data breach,<br/>" +
		"  we have not disclosed any private encryption keys,<br/>" +
		"  we have not been forced to modify our system to allow access or information leaking to a third party.<br/><br/>" +
		"THIS MESSAGE HAS NOT BEEN SIGNED - this is coming soon." +
		"</body></html>"

	_, err := w.Write([]byte(policy))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
