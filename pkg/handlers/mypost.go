package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"Forum/models"
)

var TemplateMyPost = "templates/html/myposts.html"

func (h *Handler) MyPost(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplateMyPost)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	switch r.Method {
	case http.MethodGet:
		post, err := h.services.GetPostByUserID(user.Id)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		result := models.MyPost{
			User:  user.Username,
			Posts: post,
		}
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
	case http.MethodPost:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
