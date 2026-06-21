package middleware

import (
	"adv_demo/configs"
	"adv_demo/pkg/jwt"
	"context"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHendler := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHendler, "Bearer"){
			writeUnauthorized(w)
			return
		}

		if authHendler == "" {
			log.Println("Authorization header is empty")
			return
		}
		token := strings.TrimPrefix(authHendler, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
