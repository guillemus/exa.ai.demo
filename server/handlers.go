package server

import (
	"log/slog"
	"net/http"
	"strings"

	"exa.ai.demo/exa"
	"exa.ai.demo/views"
	"github.com/starfederation/datastar-go/datastar"
)

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	s.renderer.RenderSearch(w, r)
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
	req := exaSearchRequest(form)
	if form.UsesStreaming() {
		chunks := []exa.SearchStreamChunk{}
		resp := &exa.SearchResponse{}
		grounding := []exa.GroundingInfo{}
		var content strings.Builder
		err := client.StreamSearch(r.Context(), req, func(chunk exa.SearchStreamChunk) error {
			chunks = append(chunks, chunk)
			if chunk.StreamReset {
				content.Reset()
			}
			content.WriteString(streamChunkContent(chunk))
			if len(chunk.Results) > 0 {
				resp.Results = chunk.Results
			}
			if len(chunk.Grounding) > 0 {
				grounding = chunk.Grounding
			}
			if chunk.Output != nil {
				resp.Output = chunk.Output
			} else if content.Len() > 0 || len(grounding) > 0 {
				resp.Output = &exa.DeepSearchOutput{Content: content.String(), Grounding: grounding}
			}
			views.PatchOutputStream(sse, form, marshalStreamChunks(chunks), resp)
			return nil
		})
		if err != nil {
			slog.Error("exa stream search", "err", err)
			views.PatchOutputJSON(sse, form, marshalJSON(map[string]string{"error": "Unable to stream Exa response. Check your API key and try again."}))
		}
		return
	}

	resp, err := client.Search(r.Context(), req)
	if err != nil {
		slog.Error("exa search", "err", err)
		views.PatchOutputJSON(sse, form, marshalJSON(map[string]string{"error": "Unable to search Exa. Check your API key and try again."}))
		return
	}

	views.PatchOutputResponse(sse, form, marshalJSON(resp), resp)
}
