package server

import (
	"encoding/json"

	"exa.ai.demo/exa"
	"exa.ai.demo/views"
)

func exaSearchRequest(form views.SearchForm) *exa.SearchRequest {
	form = form.WithDefaults()
	req := &exa.SearchRequest{
		Query:      form.Query,
		Type:       form.EffectiveSearchType(),
		Category:   form.Category,
		NumResults: int(form.NumResults),
	}
	if form.UsesDeepOutput() {
		req.OutputSchema = map[string]any{"type": "text"}
	}
	req.Contents = &exa.ContentsOptions{
		Highlights: &exa.HighlightsOptions{MaxCharacters: 4000},
		Extras:     &exa.ExtrasOptions{Links: 1},
	}
	return req
}

func marshalJSON(v any) string {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return `{"error":"Unable to format Exa response"}`
	}
	return string(bs)
}
