package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/app/handlers"
	myMiddleware "github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/middleware"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func Run(port string, distPath string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API and Auth routes
	r.Mount("/api", apiRoutes())
	r.Mount("/auth", authRoutes())
	r.Mount("/updater", updaterRoutes())
	//r.Get("/guide", func(w http.ResponseWriter, r *http.Request) {
	//	http.Redirect(w, r, "https://rentry.org/TaikoPublic", http.StatusFound)
	//})
	// Serve static files
	fileServer(r, "/", http.Dir(distPath), distPath)

	log.Printf("Listening on 127.0.0.1%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func apiRoutes() chi.Router {
	r := chi.NewRouter()
	apiHandler := handlers.ApiHandler{}
	r.Get("/leaderboard", apiHandler.Leaderboard)
	r.Get("/datatable", apiHandler.Datatable)
	r.Get("/user/{id}", apiHandler.GetUser)
	r.Get("/stats", apiHandler.Stats)

	// edit-options endpoints
	r.With(myMiddleware.RequireAuth).Get("/user/{id}/profile_options", apiHandler.GetProfileOptions)
	r.With(myMiddleware.RequireAuth).Get("/user/{id}/costume_options", apiHandler.GetCostumeOptions)
	r.With(myMiddleware.RequireAuth).Get("/user/{id}/song_options", apiHandler.GetSongOptions)
	r.With(myMiddleware.RequireAuth).Put("/user/{id}", apiHandler.UpdateUser)

	// settings endpoints
	r.With(myMiddleware.RequireAuth).Get("/user/{id}/access_codes", apiHandler.GetAccessCodes)
	r.With(myMiddleware.RequireAuth).Post("/user/{id}/access_codes", apiHandler.AddAccessCode)
	r.With(myMiddleware.RequireAuth).Delete("/user/{id}/access_codes", apiHandler.DeleteAccessCode)

	// fav-songs endpoints
	r.With(myMiddleware.RequireAuth).Get("/user/{id}/songs", apiHandler.GetFavouritedSongs)
	r.With(myMiddleware.RequireAuth).Put("/user/{id}/songs", apiHandler.AddFavouritedSong)
	r.With(myMiddleware.RequireAuth).Delete("/user/{id}/songs", apiHandler.DeleteFavouritedSong)

	return r
}

func authRoutes() chi.Router {
	r := chi.NewRouter()
	authHandler := handlers.AuthHandler{}
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.Post("/logout", authHandler.Logout)
	r.Get("/session", authHandler.Session)

	// settings endpoints
	r.With(myMiddleware.RequireAuth).Put("/user/{id}/username", authHandler.ChangeUsername)
	r.With(myMiddleware.RequireAuth).Put("/user/{id}/password", authHandler.ChangePassword)

	return r
}

func updaterRoutes() chi.Router {
	r := chi.NewRouter()
	updaterHandler := handlers.UpdaterHandler{}
	r.Get("/version", updaterHandler.GetUpdaterVersion)
	r.Post("/releases/{name}", updaterHandler.Releases)
	return r
}

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem, distPath string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		if _, err := root.Open(r.URL.Path); err != nil {
			// This ensures that we always serve index.html for any route not recognized
			// by the API, which lets Vue Router handle the routing
			http.ServeFile(w, r, filepath.Join(distPath, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	})
}
