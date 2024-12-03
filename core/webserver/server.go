package webserver

import (
	"fmt"
	"log"
	"net/http"
	"orydra/config"
	"orydra/handlers"
	"orydra/pkg/logger"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		httplog.RequestLogger(logger.HttpLogger),
		middleware.SetHeader("Cache-control", "no-store"),
		middleware.SetHeader("X-Robots-Tag", "noindex, nofollow, nosnippet, noarchive"),
	)
	// Assets
	staticPath, _ := filepath.Abs("static")
	fs := http.FileServer(http.Dir(staticPath))
	r.Route("/static", func(r chi.Router) {
		r.Handle("/*", http.StripPrefix("/static/", fs))
	})
	// Home route
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.HandleIndex)
		r.Get("/api/clients", handlers.HandleGetClients)
		r.Get("/api/client/{id}", handlers.HandleGetClientByID)
		r.Post("/api/client/update", handlers.HandleUpdateClient)
	})

	return r
}

func StartServer(r *chi.Mux) {
	envVars := config.SetEnv()
	fmt.Printf("Server started on 0.0.0.0:%d\n", envVars.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", envVars.PORT), r))
}
