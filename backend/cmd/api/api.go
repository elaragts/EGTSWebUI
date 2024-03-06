package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/app/handlers"
	"log"
	"net/http"
)

func Run(port string, distPath string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir(distPath))
	r.Handle("/*", http.StripPrefix("/", fs))
	r.Mount("/api", apiRoutes())
	r.Mount("/auth", authRoutes())
	http.ListenAndServe(port, r)
	log.Printf("Listening on 127.0.0.1%v", port)
}

func apiRoutes() chi.Router {
	r := chi.NewRouter()
	apiHandler := handlers.ApiHandler{}
	r.Get("/leaderboard", apiHandler.Leaderboard)
	return r
}

func authRoutes() chi.Router {
	r := chi.NewRouter()
	authHandler := handlers.AuthHandler{}
	r.Post("/signup", authHandler.Signup)
	return r
}
