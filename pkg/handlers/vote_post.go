package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"Forum/models"
)

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
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
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
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
		evaluates := models.Evaluate{
			PostId: postID,
			UserId: user.Id,
			Vote:   1,
		}
		reaction, err := h.services.CheckUserPost(user.Id, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		} // Check for evaluate from db
		if reaction.UserId != user.Id && reaction.PostId != postID {
			if err = h.services.CreateEvaluates(evaluates); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote != 1 {
			err = h.services.UpdateVote(user.Id, postID, 1)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote == 1 {
			if err = h.services.CheckVote(user.Id, postID, 1); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		count, err := h.services.EvaluateCount(postID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		err = h.services.UpdatePost(count.Like, count.Dislike, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		post, err := h.services.GetPostByPostID(postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		if post.Author != user.Username {
			err = h.services.CreateNotification(user.Username, post.Author, fmt.Sprintf(" liked your '%v' post", post.Title))
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

func (h *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
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
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
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
		evaluates := models.Evaluate{
			PostId: postID,
			UserId: user.Id,
			Vote:   -1,
		}
		reaction, err := h.services.CheckUserPost(user.Id, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		if reaction.UserId != user.Id && reaction.PostId != postID {
			if err = h.services.CreateEvaluates(evaluates); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote != -1 {
			err = h.services.UpdateVote(user.Id, postID, -1)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote == -1 {
			if err = h.services.CheckVote(user.Id, postID, -1); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}
		count, err := h.services.EvaluateCount(postID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		err = h.services.UpdatePost(count.Like, count.Dislike, postID)
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
			err = h.services.CreateNotification(user.Username, post.Author, fmt.Sprintf("  disliked your '%v' post", post.Title))
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
