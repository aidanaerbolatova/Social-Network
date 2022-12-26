package handlers

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"Forum"
)

var TemplatePost = "templates/post.html"

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplatePost)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(Forum.User)
	result := Forum.GetComments{
		User:     user.Username,
		Post:     Forum.Post{},
		Comments: []Forum.Comments{},
	}
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		postId, ok := r.Form["postID"]
		if !ok {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New("tags field not found"))
			return
		}
		postID, err := strconv.Atoi(postId[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		post, err := h.services.GetPostByPostID(postID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, err)
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		comment, err := h.services.GetCommentByPost(postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		result.Post = post
		result.Comments = comment
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	}
}
