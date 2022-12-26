package handlers

import (
	"Forum"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

const ctxKeyUser ctxKey = iota

type ctxKey int8

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var user Forum.User
		cookie, err := request.Cookie("session_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, Forum.User{})))
				return
			}
			h.HandleErrorPage(writer, http.StatusBadRequest, err)
			return
		}
		token, err := h.services.GetToken(cookie.Value)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, Forum.User{})))
				return
			}
			h.HandleErrorPage(writer, http.StatusInternalServerError, err)
			return
		}
		if token.ExpiresAT.Before(time.Now()) {
			if err := h.services.DeleteToken(cookie.Value); err != nil {
				return
			}
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, Forum.User{})))
			return
		}

		user, err = h.services.GetUserByToken(token.AuthToken)
		if err != nil {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, Forum.User{})))
			return
		}

		next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, user)))
	}
}
