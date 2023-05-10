package handlers

import (
	"Forum/models"
	"errors"
	"net/http"
	"strconv"
)

func (h *Handler) Following(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxKeyUser).(models.User)
	switch r.Method {
	case http.MethodGet:
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("tut"))
		return
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		authorId, ok1 := r.Form["userID"]
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
			return
		}
		authorID, err := strconv.Atoi(authorId[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		if user.Id == 0 {
			http.Redirect(w, r, "/signIn", 301)
			return
		}
		follow := models.Follow{
			UserId:   user.Id,
			AuthorId: authorID,
		}
		err = h.services.CheckFollow(follow)
		if user.Id != authorID && err != nil {
			err = h.services.CreateFollow(follow)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}

func (h *Handler) UnFollowing(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxKeyUser).(models.User)
	switch r.Method {
	case http.MethodGet:
		h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
		return
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		username, ok1 := r.Form["username"]
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
			return
		}
		userByUsername, err := h.services.GetUser(username[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		follow := models.Follow{
			UserId:   user.Id,
			AuthorId: userByUsername.Id,
		}
		err = h.services.DeleteFollow(follow)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
