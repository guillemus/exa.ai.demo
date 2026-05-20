package server

import (
	"log"
	"log/slog"
	"net/http"

	"exa.ai.demo/env"
	"exa.ai.demo/views"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router   chi.Router
	renderer *views.Renderer
	env      env.Env
}

func NewHandler(env env.Env) http.Handler {
	s := &Server{
		router:   chi.NewRouter(),
		renderer: views.NewRenderer(env),
		env:      env,
	}

	s.router.Get("/", s.handleHome)
	s.router.Get("/code", s.handleCode)
	s.router.Post("/search", s.handleSearch)
	if env.Dev {
		s.router.Handle("/*", http.FileServer(http.Dir("public")))
	}
	return s.router
}

func StartServer(env env.Env) {
	addr := ":" + env.Port
	slog.Info("server starting", "addr", addr)
	if err := http.ListenAndServe(addr, NewHandler(env)); err != nil {
		log.Fatal(err)
	}
}
