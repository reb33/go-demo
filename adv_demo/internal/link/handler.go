package link

import (
	"adv_demo/pkg/request"
	"adv_demo/pkg/response"
	"errors"
	"net/http"
	"strconv"
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
		var link = NewLink(body.Url)
		for {
			isExist, err := handler.LinkRepository.IsHashExist(link.Hash)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !isExist {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, http.StatusCreated, createdLink)
	}
}

func (handler *LinkHandler) UpdateLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(&UpdateLinkParams{
			ID:   id,
			Url:  body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, http.StatusOK, link)
	}
}

func (handler *LinkHandler) DeleteLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, isDeleted, err := handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !isDeleted {
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}
		response.Json(w, http.StatusOK, link)
	}
}
