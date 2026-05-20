package server

import (
	"encoding/json"
	"strings"

	"exa.ai.demo/exa"
	"exa.ai.demo/views"
)

func exaSearchRequest(form views.SearchForm) *exa.SearchRequest {
	form = form.WithDefaults()
	req := &exa.SearchRequest{
		Query:              form.Query,
		Type:               form.EffectiveSearchType(),
		Category:           form.Category,
		NumResults:         int(form.NumResults),
		IncludeDomains:     splitList(form.IncludeDomains),
		ExcludeDomains:     splitList(form.ExcludeDomains),
		StartPublishedDate: form.StartPublishedDate,
		EndPublishedDate:   form.EndPublishedDate,
		UserLocation:       strings.ToUpper(form.UserLocation),
	}
	if form.UsesOutputSchema() {
		req.OutputSchema = map[string]any{"type": "text"}
	}
	req.Contents = &exa.ContentsOptions{
		LivecrawlTimeout: int(form.LivecrawlTimeout),
		MaxAgeHours:      int(form.MaxAgeHours),
		Extras:           &exa.ExtrasOptions{Links: 1},
	}
	if form.Highlights {
		req.Contents.Highlights = &exa.HighlightsOptions{MaxCharacters: int(form.HighlightMaxCharacters), Query: form.HighlightQuery}
	}
	if form.Text {
		req.Contents.Text = &exa.TextOptions{MaxCharacters: int(form.TextMaxCharacters)}
	}
	return req
}

func splitList(value string) []string {
	items := []string{}
	for part := range strings.SplitSeq(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			items = append(items, part)
		}
	}
	return items
}

func marshalJSON(v any) string {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return `{"error":"Unable to format Exa response"}`
	}
	return string(bs)
}
