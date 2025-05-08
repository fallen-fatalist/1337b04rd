package httpserver

import (
	"1337bo4rd/internal/core/domain"
	"1337bo4rd/internal/core/port"
	"database/sql"
	"errors"
	"html/template"
	"net/http"
)

type PostHandler struct {
	svc  port.PostService
	tmpl *template.Template
}

var msg = "Failed to load threads."
var statusCode = http.StatusInternalServerError

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
		if errors.Is(err, sql.ErrNoRows) {
			statusCode = http.StatusBadRequest
			msg = err.Error()
		}
		renderError(w, h.tmpl, statusCode, msg)
		return
	}
	h.tmpl.ExecuteTemplate(w, "catalog.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandleArchive(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ListPosts()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			statusCode = http.StatusBadRequest
			msg = err.Error()
		}
		renderError(w, h.tmpl, statusCode, msg)
		return
	}
	h.tmpl.ExecuteTemplate(w, "archive.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// add logic for r.Method(POST) Addd comments!

	posts, err := h.svc.GetPostWithCommentsById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, port.ErrInvalidPostId) {
			statusCode = http.StatusBadRequest
			msg = err.Error()
		}
		renderError(w, h.tmpl, statusCode, msg)
		return
	}

	data := struct {
		Posts *domain.PostComents
		User  *domain.User
	}{posts, getSession(r)}

	h.tmpl.ExecuteTemplate(w, "post.html", data)
}

// func (h *PostHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {

// }
