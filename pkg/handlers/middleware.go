package handlers

import (
	"Forum/models"
	"context"
	"database/sql"
	"errors"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"sync"
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

func (h *Handler) RateLimitMiddleware(next http.Handler) http.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var mutex sync.Mutex
	var clients = make(map[string]*client)

	go func() {
		for {
			time.Sleep(time.Minute)
			mutex.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mutex.Unlock()
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			h.HandleErrorPage(w, http.StatusInternalServerError, errors.New("error with split host and port"))
			return
		}
		mutex.Lock()
		if _, value := clients[ip]; !value {
			clients[ip] = &client{
				limiter: rate.NewLimiter(2, 6),
			}
		}
		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mutex.Unlock()
			h.HandleErrorPage(w, http.StatusTooManyRequests, errors.New("rate limit exceeded"))
			return
		}
		mutex.Unlock()
		next.ServeHTTP(w, r)
	}
}
