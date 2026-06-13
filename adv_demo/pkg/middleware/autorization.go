package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHendler := r.Header.Get("Authorization")
		if authHendler == "" {
			log.Println("Authorization header is empty")
			return
		}
		token := strings.TrimPrefix(authHendler, "Bearer ")
		log.Println(token)
		next.ServeHTTP(w, r)
	})
}
