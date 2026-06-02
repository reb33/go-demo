package auth

import (
	"adv_demo/configs"
	"adv_demo/pkg/request"
	"adv_demo/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(payload)

		secret := handler.Config.Auth.Secret
		res := LoginResponse{
			Token: secret,
		}
		response.Json(w, http.StatusOK, res)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(payload)

		secret := handler.Config.Auth.Secret
		res := RegisterResponse{
			Status: "success",
			Token:  secret,
		}
		response.Json(w, http.StatusOK, res)
	}
}
