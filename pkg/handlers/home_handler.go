package handlers

import (
	"Forum/models"
	"errors"
	"html/template"
	"net/http"
)

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
		return
	}
	if r.Method != http.MethodGet {
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	post, err := h.services.GetPost()
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	notifications, err := h.services.GetNotification(user.Username)

	result := models.MyPost{
		User:         user.Username,
		Posts:        post,
		Notification: notifications,
	}
	tmpl, err := template.ParseFiles(TemplateHome)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	if err := tmpl.Execute(w, result); err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
}
