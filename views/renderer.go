package views

import (
	"context"
	"errors"
	"io"
	"log/slog"
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

func (x *Renderer) RenderSearch(w io.Writer) {
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
			Main(Data("signals", `{query: "Latest news on Nvidia", panelTab: "code", codeTab: "python", outputTab: "json", searchType: "auto", deepModel: "deep", numResults: 10, category: "company", structuredOutputs: false, streamResponse: false, systemPromptEnabled: false, systemPrompt: "", highlights: true, highlightMaxCharacters: 4000, highlightQuery: "", text: false, textMaxCharacters: 20000, maxAgeHours: "", livecrawlTimeout: 10000, includeDomains: "", excludeDomains: "", startPublishedDate: "", endPublishedDate: "", userLocation: ""}`), PlaygroundPage()),
			If(x.env.Dev, DebugSignals()),
		},
	}))
}

func renderNode(w io.Writer, node Node) {
	if err := node.Render(w); err != nil {
		slog.Error("view render error", "err", err)
	}
}

func PatchCodePanel(sse *datastar.ServerSentEventGenerator, form SearchForm) {
	ssePatch(sse, CodePanelContent(CodePanelData{Form: form}))
}

func PatchOutputLoading(sse *datastar.ServerSentEventGenerator, form SearchForm) {
	ssePatchSignals(sse, `{ "panelTab": "output" }`)
	ssePatch(sse, CodePanelContent(CodePanelData{Form: form, Loading: true}))
}

func PatchOutputJSON(sse *datastar.ServerSentEventGenerator, form SearchForm, output string) {
	ssePatch(sse, CodePanelContent(CodePanelData{Form: form, OutputJSON: output}))
}

func PatchOutputStream(sse *datastar.ServerSentEventGenerator, form SearchForm, output string, content string) {
	resp := &exa.SearchResponse{Output: &exa.DeepSearchOutput{Content: content, Grounding: []exa.GroundingInfo{}}}
	ssePatch(sse, CodePanelContent(CodePanelData{Form: form, OutputJSON: output, Response: resp}))
}

func PatchOutputResponse(sse *datastar.ServerSentEventGenerator, form SearchForm, output string, resp *exa.SearchResponse) {
	ssePatch(sse, CodePanelContent(CodePanelData{Form: form, OutputJSON: output, Response: resp}))
}

func ssePatchSignals(sse *datastar.ServerSentEventGenerator, signals string) {
	if err := sse.PatchSignals([]byte(signals)); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		slog.Error("sse.PatchSignals error", "err", err)
	}
}

func ssePatch(sse *datastar.ServerSentEventGenerator, node Node) {
	var sb strings.Builder
	if err := node.Render(&sb); err != nil {
		slog.Error("view render error", "err", err)
		return
	}

	if err := sse.PatchElements(sb.String(), datastar.WithSelectorID("code-panel-content"), datastar.WithModeOuter()); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		slog.Error("sse.PatchElements error", "err", err)
	}
}
