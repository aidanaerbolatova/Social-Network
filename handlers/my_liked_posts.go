package handlers

import (
	"Forum"
	"database/sql"
	"errors"
	"html/template"
	"net/http"
)

var TemplateLikedPosts = "templates/likedPosts.html"

func (h *Handler) LikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedPost" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplateLikedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(Forum.User)
	result := Forum.MyPost{
		User:  user.Username,
		Posts: []Forum.Post{},
	}
	switch r.Method {
	case http.MethodGet:
		posts, err := h.services.LikedPosts(user.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		result.Posts = posts
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	}
}
