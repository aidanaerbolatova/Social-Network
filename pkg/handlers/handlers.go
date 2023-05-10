package handlers

import (
	"net/http"

	"Forum/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.HandlerFunc {
	router := http.NewServeMux()
	router.HandleFunc("/", h.AuthMiddleware(h.HomePage))
	router.HandleFunc("/signIn", h.SignIn)
	router.HandleFunc("/signUp", h.SignUp)
	router.HandleFunc("/logout", h.SignOut)

	router.HandleFunc("/signUp/google", h.googleSignUp)
	router.HandleFunc("/signUp/callback", h.callbackGoogleSignUp)
	router.HandleFunc("/signIn/google", h.googleSignIn)
	router.HandleFunc("/signIn/callback", h.callbackGoogleSignIn)

	router.HandleFunc("/signIn/github", h.githubSignIn)
	router.HandleFunc("/signIn/github/callback", h.signInCallbackGithub)
	router.HandleFunc("/signUp/github", h.githubSignUp)
	router.HandleFunc("/signUp/github/callback", h.signUpCallbackGithub)

	router.HandleFunc("/createPost", h.AuthMiddleware(h.CreatePost))
	router.HandleFunc("/categoryPost", h.CategoryPost)
	router.HandleFunc("/myposts", h.AuthMiddleware(h.MyPost))
	router.HandleFunc("/comment", h.AuthMiddleware(h.Comments))
	router.HandleFunc("/post", h.AuthMiddleware(h.Post))
	router.HandleFunc("/like", h.AuthMiddleware(h.LikePost))
	router.HandleFunc("/dislike", h.AuthMiddleware(h.DislikePost))
	router.HandleFunc("/delete", h.AuthMiddleware(h.DeletePost))
	router.HandleFunc("/deleteComment", h.AuthMiddleware(h.DeleteComment))
	router.HandleFunc("/edit", h.AuthMiddleware(h.EditPost))
	router.HandleFunc("/editComment", h.AuthMiddleware(h.EditComment))
	router.HandleFunc("/likedPost", h.AuthMiddleware(h.LikedPosts))
	router.HandleFunc("/dislikedPosts", h.AuthMiddleware(h.DislikedPosts))
	router.HandleFunc("/commentedPosts", h.AuthMiddleware(h.CommentedPosts))
	router.HandleFunc("/likeComment", h.AuthMiddleware(h.LikeComment))
	router.HandleFunc("/dislikeComment", h.AuthMiddleware(h.DislikeComment))
	router.HandleFunc("/following", h.AuthMiddleware(h.Following))
	router.HandleFunc("/follow", h.AuthMiddleware(h.Follow))
	router.HandleFunc("/unfollow", h.AuthMiddleware(h.UnFollowing))

	router.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))

	return h.RateLimitMiddleware(router)
}
