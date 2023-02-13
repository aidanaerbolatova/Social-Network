package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Forum/models"
)

func (h *Handler) Comments(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ctxKeyUser).(models.User)
	switch r.Method {
	case http.MethodGet:
		h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
		return
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		postId, ok1 := r.Form["postId"]
		comment, ok2 := r.Form["comment"]
		if !ok1 || !ok2 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
			return
		}
		if len(strings.TrimSpace(comment[0])) == 0 {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		postID, err := strconv.Atoi(postId[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		if user.Id == 0 {
			http.Redirect(w, r, "/signIn", 301)
			return
		}
		comments := models.Comments{
			UserId:    user.Id,
			PostId:    postID,
			Comment:   comment[0],
			Author:    user.Username,
			CreatedAt: time.Now().Format("2 Jan 15:04:05"),
		}
		err = h.services.AddComment(comments)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		post, err := h.services.GetPostByPostID(postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		if user.Username != post.Author {
			err = h.services.CreateNotification(user.Username, post.Author, fmt.Sprintf(" commented your '%v' title", post.Title))
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
