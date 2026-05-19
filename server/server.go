package server

import (
	"log"
	"log/slog"
	"net/http"

	"exa.ai.demo/db"
	"exa.ai.demo/env"
	"exa.ai.demo/views"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router   chi.Router
	renderer *views.Renderer
	db       *db.Pool
	env      env.Env
}

func StartServer(env env.Env) {
	conn, err := db.Connect(env.DBPath)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer conn.Close()

	s := &Server{
		router:   chi.NewRouter(),
		renderer: views.NewRenderer(env),
		db:       conn,
		env:      env,
	}

	s.router.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	s.router.Get("/", s.handleHome)
	s.router.Get("/code", s.handleCode)
	s.router.Post("/search", s.handleSearch)

	addr := ":" + s.env.Port
	slog.Info("server starting", "addr", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatal(err)
	}
}
