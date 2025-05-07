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
	posts, _ := h.svc.ListActive()
	h.tmpl.ExecuteTemplate(w, "catalog.html", struct{ Posts []domain.Post }{posts})
}

func (h *PostHandler) HandleArchive(w http.ResponseWriter, r *http.Request) {

}

// func (h *PostHandler) HandlePosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case http.MethodPost:
// 		var post domain.Post
// 		decoder := json.NewDecoder(r.Body)
// 		err := decoder.Decode(&post)
// 		if err != nil {
// 			jsonErrorRespond(w, err, http.StatusInternalServerError)
// 			return
// 		}
// 		err = h.svc.CreatePost(&post)
// 		if err != nil {
// 			statusCode := http.StatusInternalServerError
// 			switch err {
// 			case port.ErrEmptyAvatar,
// 				port.ErrEmptyContent,
// 				port.ErrEmptyTitle:
// 				statusCode = http.StatusBadRequest
// 			}
// 			jsonErrorRespond(w, err, statusCode)
// 			return
// 		}
// 		jsonMessageRespond(w, "Post have created succesfully", http.StatusOK)
// 		return
// 	case http.MethodGet:
// 		posts, err := h.svc.ListPosts()
// 		if err != nil {
// 			if errors.Is(err, sql.ErrNoRows) {
// 				jsonMessageRespond(w, port.ErrNoPosts.Error(), http.StatusOK)
// 				return
// 			}
// 			jsonErrorRespond(w, err, http.StatusInternalServerError)
// 		}

// 		jsonPayload, err := json.MarshalIndent(posts, "", "   ")
// 		if err != nil {
// 			jsonErrorRespond(w, err, http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(jsonPayload)
// 		return
// 	default:
// 		w.Header().Set("Allow", "GET, POST")
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// }
