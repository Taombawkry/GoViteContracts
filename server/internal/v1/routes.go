// v1/routes.go
package v1

import (
	"net/http"

	databaseCfg "github.com/NhyiraAmofaSekyi/go-webserver/internal/db"
	"github.com/NhyiraAmofaSekyi/go-webserver/internal/middleware"
	"github.com/NhyiraAmofaSekyi/go-webserver/internal/v1/auth"
	"github.com/NhyiraAmofaSekyi/go-webserver/internal/v1/users"
)

func NewRouter(dbCFG *databaseCfg.DBConfig) *http.ServeMux {
	v1Router := http.NewServeMux()
	authRouter := auth.NewRouter()
	userRouter := users.NewRouter()

	v1Router.HandleFunc("GET /healthz", HealthzHandler)
	v1Router.HandleFunc("GET /secure", middleware.AuthMiddleware(SecureHandler))
	v1Router.Handle("/auth/", http.StripPrefix("/auth", authRouter))
	v1Router.Handle("/users/", http.StripPrefix("/users", userRouter))

	//git push origin --force --all
	//git push origin --force --tags
	return v1Router
}
