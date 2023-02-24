package handlers

import (
	"Forum/models"
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
		var user models.User
		cookie, err := request.Cookie("session_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, models.User{})))
				return
			}
			h.HandleErrorPage(writer, http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest)))
			return
		}
		token, err := h.services.GetToken(cookie.Value)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, models.User{})))
				return
			}
			h.HandleErrorPage(writer, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
			return
		}
		if token.ExpiresAT.Before(time.Now()) {
			if err := h.services.DeleteToken(cookie.Value); err != nil {
				return
			}
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, models.User{})))
			return
		}
		user, err = h.services.GetUserByToken(token.AuthToken)
		if err != nil {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, models.User{})))
			return
		}
		next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), ctxKeyUser, user)))
	}
}
