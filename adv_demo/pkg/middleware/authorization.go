package middleware

import (
	"adv_demo/configs"
	"adv_demo/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHendler := r.Header.Get("Authorization")
		if authHendler == "" {
			log.Println("Authorization header is empty")
			return
		}
		token := strings.TrimPrefix(authHendler, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		log.Println(isValid)
		log.Println(data)
		next.ServeHTTP(w, r)
	})
}
