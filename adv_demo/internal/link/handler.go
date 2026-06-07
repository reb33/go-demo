package link

import (
	"adv_demo/pkg/request"
	"adv_demo/pkg/response"
	"errors"
	"fmt"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("PATCH /link/{id}", handler.UpdateLink())
	router.HandleFunc("DELETE /link/{id}", handler.DeleteLink())
}

func (handler *LinkHandler) GoTo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) CreateLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, http.StatusCreated, createdLink)
	}
}

func (h *LinkHandler) UpdateLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateLink"))
	}
}

func (h *LinkHandler) DeleteLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
