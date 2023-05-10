package handlers

import (
	"Forum/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(TemplateFollow)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(models.User)
	switch r.Method {
	case http.MethodGet:
		myFollowers, err := h.services.MyFollowers(user.Id)
		if err != nil {
			fmt.Println(err)
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		var followers []models.User
		for _, follower := range myFollowers {
			followerID, err := strconv.Atoi(follower)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
			user, err := h.services.GetUserByUserId(followerID)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
			followers = append(followers, user)
		}
		myFollowing, err := h.services.Following(user.Id)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		var following []models.User
		for _, myfollow := range myFollowing {
			followID, err := strconv.Atoi(myfollow)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
			user, err := h.services.GetUserByUserId(followID)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
			following = append(following, user)
		}
		result := models.AllFollow{
			User:      user,
			Following: following,
			Followers: followers,
		}
		if err := temp.Execute(w, result); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	default:
		h.HandleErrorPage(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
}
