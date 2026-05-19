package server

import (
	"log/slog"
	"net/http"

	"exa.ai.demo/exa"
	"exa.ai.demo/views"
	"github.com/starfederation/datastar-go/datastar"
)

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	s.renderer.RenderSearch(w)
}

func (s *Server) handleCode(w http.ResponseWriter, r *http.Request) {
	var form views.SearchForm
	if err := datastar.ReadSignals(r, &form); err != nil {
		slog.Error("read signals", "err", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sse := datastar.NewSSE(w, r)
	views.PatchCodePanel(sse, form)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	var form views.SearchForm
	if err := datastar.ReadSignals(r, &form); err != nil {
		slog.Error("read signals", "err", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sse := datastar.NewSSE(w, r)
	views.PatchOutputLoading(sse, form)

	client := exa.NewClient(s.env.ExaAIAPIKey)
	resp, err := client.Search(r.Context(), exaSearchRequest(form))
	if err != nil {
		slog.Error("exa search", "err", err)
		views.PatchOutputJSON(sse, form, marshalJSON(map[string]string{"error": "Unable to search Exa. Check your API key and try again."}))
		return
	}

	views.PatchOutputResponse(sse, form, marshalJSON(resp), resp)
}
