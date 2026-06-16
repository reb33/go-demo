package auth

import (
	"adv_demo/configs"
	"adv_demo/pkg/request"
	"adv_demo/pkg/response"
	"errors"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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
		token, err := handler.AuthService.Login(payload.Email, payload.Password)
		if err != nil {
			if errors.Is(err, ErrInvalidCredentials) {
				response.Json(w, http.StatusBadRequest, ErrResponse{Error: err.Error()})
				return
			}
			response.Json(w, http.StatusInternalServerError, ErrResponse{Error: err.Error()})
			return
		}

		response.Json(w, http.StatusOK, LoginResponse{Token: token})
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		token, err := handler.AuthService.Register(payload.Email, payload.Password, payload.Name)
		if err != nil {
			if errors.Is(err, ErrUserAlreadyExist) {
				response.Json(w, http.StatusBadRequest, ErrResponse{Error: err.Error()})
				return
			}
			response.Json(w, http.StatusInternalServerError, ErrResponse{Error: err.Error()})
			return
		}

		response.Json(w, http.StatusOK, RegisterResponse{Token: token})
	}
}
