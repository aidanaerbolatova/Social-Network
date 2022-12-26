package handlers

import (
	"Forum"
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/like" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplateMyPost, TemplateCategory, TemplateHome, TemplateLikedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(Forum.User)
	switch r.Method {
	case http.MethodGet:
		if err := temp.Execute(w, user); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		postId, ok1 := r.Form["postId"]
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, err)
			return
		}
		postID, err := strconv.Atoi(postId[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		if user.Id == 0 {
			http.Redirect(w, r, "/signIn", 301)
			return
		}
		evaluates := Forum.Evaluate{
			PostId: postID,
			UserId: user.Id,
			Vote:   1,
		}
		reaction, err := h.services.CheckUserPost(user.Id, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		} // Check for evaluate from db
		if reaction.UserId != user.Id && reaction.PostId != postID {
			if err = h.services.CreateEvaluates(evaluates); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote != 1 {
			err = h.services.UpdateVote(user.Id, postID, 1)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote == 1 {
			if err = h.services.CheckVote(user.Id, postID, 1); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		count, err := h.services.EvaluateCount(postID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, err)
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		err = h.services.UpdatePost(count.Like, count.Dislike, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	}
}

func (h *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislike" {
		h.HandleErrorPage(w, http.StatusNotFound, errors.New("error with url"))
		return
	}
	temp, err := template.ParseFiles(TemplateMyPost, TemplateCategory, TemplateHome, TemplateLikedPosts)
	if err != nil {
		h.HandleErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	user := r.Context().Value(ctxKeyUser).(Forum.User)
	switch r.Method {
	case http.MethodGet:
		if err := temp.Execute(w, user); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		postId, ok1 := r.Form["postId"]
		if !ok1 {
			h.HandleErrorPage(w, http.StatusBadRequest, err)
			return
		}
		postID, err := strconv.Atoi(postId[0])
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		if user.Id == 0 {
			http.Redirect(w, r, "/signIn", 301)
			return
		}
		evaluates := Forum.Evaluate{
			PostId: postID,
			UserId: user.Id,
			Vote:   -1,
		}
		reaction, err := h.services.CheckUserPost(user.Id, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		if reaction.UserId != user.Id && reaction.PostId != postID {
			if err = h.services.CreateEvaluates(evaluates); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote != -1 {
			err = h.services.UpdateVote(user.Id, postID, -1)
			if err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		if reaction.UserId == user.Id && reaction.PostId == postID && reaction.Vote == -1 {
			if err = h.services.CheckVote(user.Id, postID, -1); err != nil {
				h.HandleErrorPage(w, http.StatusInternalServerError, err)
				return
			}
		}
		count, err := h.services.EvaluateCount(postID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				h.HandleErrorPage(w, http.StatusNotFound, err)
				return
			}
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		err = h.services.UpdatePost(count.Like, count.Dislike, postID)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	}
}
