package auth

import (
	"net/http"

	"github.com/NhyiraAmofaSekyi/go-webserver/internal/middleware"
)

// NewRouter returns a new http.ServeMux with v1 routes configured
func NewRouter() *http.ServeMux {
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /signUp", SignUp)
	authRouter.HandleFunc("GET /check-username/{username}", CheckUsername)
	authRouter.HandleFunc("POST /farcasterSignUp", FarcasterSignUp)
	authRouter.HandleFunc("GET /SignOut", SignOut)
	authRouter.HandleFunc("GET /google/callback", Callback2)
	authRouter.HandleFunc("GET /login/callback", GoogleLogin)
	authRouter.HandleFunc("GET /dashboard", middleware.AuthMiddleware2(Dashboard))
	authRouter.HandleFunc("POST /emailAction", EmailAction)
	authRouter.HandleFunc("POST /emailVerification", EmailVerification)
	authRouter.HandleFunc("GET /check", Check)

	return authRouter
}
