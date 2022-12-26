package handlers

import (
	"Forum/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", h.AuthMiddleware(h.HomePage))
	router.HandleFunc("/signIn", h.SignIn)
	router.HandleFunc("/signUp", h.SignUp)
	router.HandleFunc("/logout", h.SignOut)

	router.HandleFunc("/createPost", h.AuthMiddleware(h.CreatePost))
	router.HandleFunc("/categoryPost", h.CategoryPost)
	router.HandleFunc("/myposts", h.AuthMiddleware(h.MyPost))
	router.HandleFunc("/comment", h.AuthMiddleware(h.Comments))
	router.HandleFunc("/post", h.AuthMiddleware(h.Post))
	router.HandleFunc("/like", h.AuthMiddleware(h.LikePost))
	router.HandleFunc("/dislike", h.AuthMiddleware(h.DislikePost))
	router.HandleFunc("/likedPost", h.AuthMiddleware(h.LikedPosts))
	router.HandleFunc("/likeComment", h.AuthMiddleware(h.LikeComment))
	router.HandleFunc("/dislikeComment", h.AuthMiddleware(h.DislikeComment))

	router.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))

	return router
}
