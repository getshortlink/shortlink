package link

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LinkHandler struct {
	Router chi.Router
}

func NewHandler() *LinkHandler {
	r := chi.NewRouter()

	handler := &LinkHandler{
		Router: r,
	}

	r.Post("/", handler.createLink)
	r.Get("/{key}", handler.getLink)

	return handler
}

type createLinkRequest struct {
	// TODO
}

type createLinkResponse struct {
	// TODO
}

func (h *LinkHandler) createLink(w http.ResponseWriter, req *http.Request) {
	// TODO
}

type getLinkResponse struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func (h *LinkHandler) getLink(w http.ResponseWriter, req *http.Request) {
	// TODO
}
