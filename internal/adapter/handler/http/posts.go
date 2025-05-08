package httpserver

import (
	"1337bo4rd/internal/core/domain"
	"1337bo4rd/internal/core/port"
	"html/template"
	"net/http"
)

type PostHandler struct {
	svc  port.PostService
	tmpl *template.Template
}

func NewPostHandler(svc port.PostService) *PostHandler {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	return &PostHandler{
		svc,
		tmpl,
	}
}

func (h *PostHandler) HandleCatalog(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ListActive()
	if err != nil {
		renderError(w, h.tmpl, http.StatusInternalServerError, "Failed to load threads.")
		return
	}
	h.tmpl.ExecuteTemplate(w, "catalog.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandleArchive(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ListPosts()
	if err != nil {
		renderError(w, h.tmpl, http.StatusInternalServerError, "Failed to load threads.")
		return
	}
	h.tmpl.ExecuteTemplate(w, "archive.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	post, err := h.svc.GetPostById(id)

	if err != nil {
		renderError(w, h.tmpl, http.StatusInternalServerError, "Failed to load threads.")
		return
	}
	h.tmpl.ExecuteTemplate(w, "post.html", struct{ Post *domain.Post }{post})
}
