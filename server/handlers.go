package server

import "net/http"

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	s.renderer.RenderHome(w)
}
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	s.renderer.RenderHome(w)
}
