package httpserver

import (
	"1337bo4rd/internal/core/domain"
	"1337bo4rd/internal/core/port"
	"database/sql"
	"errors"
	"html/template"
	"log/slog"
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
	if r.URL.Path != "/" {
		statusCode = http.StatusNotFound
		msg = "No such url path >,<"
		renderError(w, h.tmpl, statusCode, msg)
		return
	}
	posts, _ := h.svc.ListActive()
	h.tmpl.ExecuteTemplate(w, "catalog.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandleArchive(w http.ResponseWriter, r *http.Request) {
	posts, _ := h.svc.ListPosts()
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

func (h *PostHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.tmpl.ExecuteTemplate(w, "create-post.html", nil)
		return
	}

	userSession := getSession(r)
	title := r.FormValue("title")
	content := r.FormValue("content")

	// CONTINUE WITH MINIIO!
	// var img io.Reader
	// if f, _, err := r.FormFile("image"); err == nil {
	// 	defer f.Close()
	// 	img = f
	// }

	post := &domain.Post{
		Title:      title,
		Content:    content,
		UserName:   userSession.Name,
		UserAvatar: userSession.Avatar,
	}

	err := h.svc.CreatePost(post)
	if err != nil {
		if errors.Is(err, port.ErrEmptyTitle) || errors.Is(err, port.ErrEmptyContent) {
			statusCode = http.StatusBadRequest
			msg = err.Error()
		}
		renderError(w, h.tmpl, statusCode, msg)
		return
	}
	slog.Info("created post")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
