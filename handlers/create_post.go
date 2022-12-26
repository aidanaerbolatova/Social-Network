package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
	"time"

	"Forum"
)

var TemplateCreatePost = "templates/createPost.html"

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createPost" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplateCreatePost)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(Forum.User)
	switch r.Method {
	case http.MethodGet:
		if err := temp.Execute(w, user); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		title, ok1 := r.Form["title"]
		text, ok2 := r.Form["text"]
		tag, ok3 := r.Form["tag"]
		if !ok1 || !ok2 || !ok3 || len(text[0]) == 0 || len(title[0]) == 0 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New("data field not found"))
			return
		}
		tags := strings.Join(tag, " ")
		post := Forum.Post{
			UserId:     user.Id,
			Title:      title[0],
			Text:       text[0],
			Categories: tags,
			CreatedAt:  time.Now().Format("2 Jan 15:04:05"),
			Author:     user.Username,
		}
		err := h.services.CreatePosts(post)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
