package service

import (
	application "dissent-api-service/pkg/application/entities"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tok := r.Header.Get(headerAuth)

		err := TokenProvider.CheckToken(tok)
		if err != nil {
			body := application.NewResponse(nil, err)
			w.WriteHeader(http.StatusBadRequest)
			writeReponse(w, body)
			return
		}

		next.ServeHTTP(w, r)
	})
}
