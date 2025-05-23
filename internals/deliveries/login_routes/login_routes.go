package login_routes

import (
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func AuthorizationRoutes(mux *http.ServeMux, authHandler *AuthHandler, tokenManager domain.TokenManager) {

	//login_routes
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		logLoginRegister(w, r, "/login_routes")
		authHandler.Login(w, r)
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		logLoginRegister(w, r, "/register")
		authHandler.Register(w, r)
	})
}

func logLoginRegister(w http.ResponseWriter, r *http.Request, route string) {
	log.Printf("Attempting to access %s route, Method: %s, URL: %s", route, r.Method, r.URL.Path)
}
