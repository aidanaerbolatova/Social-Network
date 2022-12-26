package handlers

import (
	"Forum"
	"errors"
	"html/template"
	"net/http"
	"strings"
)

var TemplateCategory = "templates/categories.html"

func (h *Handler) CategoryPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categoryPost" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplateCategory)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	var result []Forum.Post
	switch r.Method {
	case http.MethodGet:
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		tag, ok := r.Form["tag"]
		if !ok {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New("tags field not found"))
			return
		}
		tags := strings.Join(tag, " ")
		category, err := h.services.GetPostByTag(tags)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New("tags field not found"))
			return
		}
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		result = **category
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
