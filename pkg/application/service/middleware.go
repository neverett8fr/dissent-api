package service

import (
	"context"
	application "dissent-api-service/pkg/application/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tok := r.Header.Get(headerAuth)

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

		// Add the parsed event information to the request context
		ctx := context.WithValue(r.Context(), ctxEventBody, eventInformation)
		r = r.WithContext(ctx)

		err = TokenProvider.CheckToken(tok, eventInformation.Organiser)
		if err != nil {
			body := application.NewResponse(nil, err)
			w.WriteHeader(http.StatusBadRequest)
			writeReponse(w, body)
			return
		}

		next.ServeHTTP(w, r)
	})
}
