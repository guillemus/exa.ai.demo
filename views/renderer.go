package views

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"exa.ai.demo/env"
	"exa.ai.demo/exa"
	"github.com/starfederation/datastar-go/datastar"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type Renderer struct {
	env env.Env
}

func NewRenderer(env env.Env) *Renderer {
	return &Renderer{env: env}
}

func (x *Renderer) RenderSearch(w io.Writer, r *http.Request) {
	renderNode(w, HTML5(HTML5Props{
		Title:       "Search | Exa API",
		Description: "Exa Search API playground markup",
		Language:    "en",
		Head: []Node{
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Link(Rel("preconnect"), Href("https://api.exa.ai"), Attr("crossorigin", "anonymous")),
			Link(Rel("dns-prefetch"), Href("https://api.exa.ai")),
			Link(Rel("preload"), Href("/public/fonts/figtree-latin-var.woff2"), As("font"), Type("font/woff2"), Attr("crossorigin")),
			Link(Rel("preload"), Href("/public/fonts/jetbrains-mono-latin-var.woff2"), As("font"), Type("font/woff2"), Attr("crossorigin")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/open-props@1.7.23/open-props.min.css")),
			Link(Rel("stylesheet"), Href("/public/styles.css")),
			Raw(`<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@v1.0.1/bundles/datastar.js"></script>`),
			styles.Node(),
			JS,
		},
		Body: []Node{
			Iff(true, func() Node {
				state := initialPageState(r.URL.Query())
				return Main(Data("signals", state.Signals()), PlaygroundPage(state))
			}),
			If(x.env.Dev, DebugSignals()),
		},
	}))
}

type PageState struct {
	Form      SearchForm
	PanelTab  string
	CodeTab   string
	OutputTab string
}

func initialPageState(q url.Values) PageState {
	form := SearchForm{
		Query:                  queryString(q, "query", "Latest news on Nvidia"),
		CodeTab:                "python",
		OutputTab:              "json",
		SearchType:             queryString(q, "searchType", "auto"),
		DeepModel:              queryString(q, "deepModel", "deep"),
		NumResults:             SignalInt(queryInt(q, "numResults", 10)),
		Category:               queryString(q, "category", "company"),
		StructuredOutputs:      queryBool(q, "structuredOutputs", false),
		StreamResponse:         queryBool(q, "streamResponse", false),
		SystemPromptEnabled:    queryBool(q, "systemPromptEnabled", false),
		SystemPrompt:           queryString(q, "systemPrompt", ""),
		Highlights:             queryBool(q, "highlights", true),
		HighlightMaxCharacters: SignalInt(queryInt(q, "highlightMaxCharacters", 4000)),
		HighlightQuery:         queryString(q, "highlightQuery", ""),
		Text:                   queryBool(q, "text", false),
		TextMaxCharacters:      SignalInt(queryInt(q, "textMaxCharacters", 20000)),
		MaxAgeHours:            SignalInt(queryInt(q, "maxAgeHours", 0)),
		LivecrawlTimeout:       SignalInt(queryInt(q, "livecrawlTimeout", 10000)),
		IncludeDomains:         queryString(q, "includeDomains", ""),
		ExcludeDomains:         queryString(q, "excludeDomains", ""),
		StartPublishedDate:     queryString(q, "startPublishedDate", ""),
		EndPublishedDate:       queryString(q, "endPublishedDate", ""),
		UserLocation:           queryString(q, "userLocation", ""),
	}
	return PageState{
		Form:      form,
		PanelTab:  "code",
		CodeTab:   form.CodeTab,
		OutputTab: form.OutputTab,
	}
}

func (s PageState) Signals() string {
	signals := map[string]any{
		"query":                  s.Form.Query,
		"panelTab":               s.PanelTab,
		"codeTab":                s.Form.CodeTab,
		"outputTab":              s.Form.OutputTab,
		"searchType":             s.Form.SearchType,
		"deepModel":              s.Form.DeepModel,
		"numResults":             int(s.Form.NumResults),
		"category":               s.Form.Category,
		"structuredOutputs":      s.Form.StructuredOutputs,
		"streamResponse":         s.Form.StreamResponse,
		"systemPromptEnabled":    s.Form.SystemPromptEnabled,
		"systemPrompt":           s.Form.SystemPrompt,
		"highlights":             s.Form.Highlights,
		"highlightMaxCharacters": int(s.Form.HighlightMaxCharacters),
		"highlightQuery":         s.Form.HighlightQuery,
		"text":                   s.Form.Text,
		"textMaxCharacters":      int(s.Form.TextMaxCharacters),
		"maxAgeHours":            signalIntOrEmpty(s.Form.MaxAgeHours),
		"livecrawlTimeout":       int(s.Form.LivecrawlTimeout),
		"includeDomains":         s.Form.IncludeDomains,
		"excludeDomains":         s.Form.ExcludeDomains,
		"startPublishedDate":     s.Form.StartPublishedDate,
		"endPublishedDate":       s.Form.EndPublishedDate,
		"userLocation":           s.Form.UserLocation,
	}
	bs, err := json.Marshal(signals)
	if err != nil {
		return `{}`
	}
	return string(bs)
}

func queryString(q url.Values, key string, fallback string) string {
	if q.Has(key) {
		return q.Get(key)
	}
	return fallback
}

func queryBool(q url.Values, key string, fallback bool) bool {
	if q.Has(key) {
		return q.Get(key) == "true"
	}
	return fallback
}

func queryInt(q url.Values, key string, fallback int) int {
	if !q.Has(key) {
		return fallback
	}
	n, err := strconv.Atoi(q.Get(key))
	if err != nil {
		return fallback
	}
	return n
}

func signalIntOrEmpty(value SignalInt) any {
	if value == 0 {
		return ""
	}
	return int(value)
}

func renderNode(w io.Writer, node Node) {
	if err := node.Render(w); err != nil {
		slog.Error("view render error", "err", err)
	}
}

func PatchCodePanel(sse *datastar.ServerSentEventGenerator, form SearchForm) {
	form = form.WithDefaults()
	ssePatchID(sse, CodeContent(form, form.CodeTab), "code-panel-code")
}

func PatchOutputLoading(sse *datastar.ServerSentEventGenerator, form SearchForm) {
	form = form.WithDefaults()
	ssePatchSignals(sse, `{ "panelTab": "output" }`)
	ssePatchCodePanel(sse, CodePanelContent(CodePanelData{Form: form, PanelTab: "output", OutputTab: form.OutputTab, Loading: true}))
}

func PatchOutputJSON(sse *datastar.ServerSentEventGenerator, form SearchForm, output string) {
	form = form.WithDefaults()
	ssePatchCodePanel(sse, CodePanelContent(CodePanelData{Form: form, PanelTab: "output", OutputTab: form.OutputTab, OutputJSON: output}))
}

func PatchOutputStream(sse *datastar.ServerSentEventGenerator, form SearchForm, output string, content string) {
	form = form.WithDefaults()
	resp := &exa.SearchResponse{Output: &exa.DeepSearchOutput{Content: content, Grounding: []exa.GroundingInfo{}}}
	ssePatchCodePanel(sse, CodePanelContent(CodePanelData{Form: form, PanelTab: "output", OutputTab: form.OutputTab, OutputJSON: output, Response: resp}))
}

func PatchOutputResponse(sse *datastar.ServerSentEventGenerator, form SearchForm, output string, resp *exa.SearchResponse) {
	form = form.WithDefaults()
	ssePatchCodePanel(sse, CodePanelContent(CodePanelData{Form: form, PanelTab: "output", OutputTab: form.OutputTab, OutputJSON: output, Response: resp}))
}

func ssePatchSignals(sse *datastar.ServerSentEventGenerator, signals string) {
	if err := sse.PatchSignals([]byte(signals)); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		slog.Error("sse.PatchSignals error", "err", err)
	}
}

func ssePatchCodePanel(sse *datastar.ServerSentEventGenerator, node Node) {
	ssePatchID(sse, node, "code-panel-content")
}

func ssePatchID(sse *datastar.ServerSentEventGenerator, node Node, id string) {
	var sb strings.Builder
	if err := node.Render(&sb); err != nil {
		slog.Error("view render error", "err", err)
		return
	}

	if err := sse.PatchElements(sb.String(), datastar.WithSelectorID(id), datastar.WithModeOuter()); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		slog.Error("sse.PatchElements error", "err", err)
	}
}
