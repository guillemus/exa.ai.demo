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
	if form.UsesSystemPrompt() {
		req.SystemPrompt = form.SystemPrompt
	}
	if form.UsesStreaming() {
		req.Stream = true
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
	bs, err := json.Marshal(v)
	if err != nil {
		return `{"error":"Unable to format Exa response"}`
	}
	return string(bs)
}

func marshalStreamChunks(chunks []exa.SearchStreamChunk) string {
	parts := make([]string, 0, len(chunks))
	for _, chunk := range chunks {
		bs, err := json.Marshal(chunk)
		if err != nil {
			return `{"error":"Unable to format Exa stream"}`
		}
		parts = append(parts, string(bs))
	}
	return strings.Join(parts, "\n")
}

func streamChunkContent(chunk exa.SearchStreamChunk) string {
	var b strings.Builder
	for _, choice := range chunk.Choices {
		b.WriteString(choice.Delta.Content)
	}
	return b.String()
}
