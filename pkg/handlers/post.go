package handlers

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"Forum/models"
)

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplatePost)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	result := models.GetComments{
		User:     user.Username,
		Post:     models.Post{},
		Comments: []models.Comments{},
		Images:   []template.URL{},
	}
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		postId, ok := r.Form["postID"]
		if !ok {
			h.HandleErrorPage(w, http.StatusBadRequest, err)
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
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
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
		img := strings.Split(string(post.Image), " ")
		for _, w := range img {
			result.Images = append(result.Images, template.URL(w))
		}
		result.Post = post
		result.Comments = comment
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
