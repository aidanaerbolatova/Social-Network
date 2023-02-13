package handlers

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"

	"Forum/models"
)

var (
	TemplateLikedPosts     = "templates/html/likedPosts.html"
	TemplateDislikedPosts  = "templates/html/dislikedPosts.html"
	TemplateCommentedPosts = "templates/html/commented_posts.html"
)

func (h *Handler) LikedPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplateLikedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	result := models.MyPost{
		User:  user.Username,
		Posts: []models.Post{},
	}
	switch r.Method {
	case http.MethodGet:
		posts, err := h.services.LikedPosts(user.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		result.Posts = posts
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

}

func (h *Handler) DislikedPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplateDislikedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	result := models.MyPost{
		User:  user.Username,
		Posts: []models.Post{},
	}
	switch r.Method {
	case http.MethodGet:
		posts, err := h.services.DislikedPosts(user.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, err)
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		result.Posts = posts
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}

func (h *Handler) CommentedPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplateCommentedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	result := models.MyPost{
		User:  user.Username,
		Posts: []models.Post{},
	}
	switch r.Method {
	case http.MethodGet:
		posts, err := h.services.CommentedPosts(user.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		result.Posts = posts
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
