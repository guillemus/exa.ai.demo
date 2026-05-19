package views

import (
	"io"
	"log/slog"

	"exa.ai.demo/env"

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

func (x *Renderer) RenderHome(w io.Writer) {
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
			Main(Data("signals", `{query: "Latest news on Nvidia", codeTab: "python", searchType: "auto"}`), PlaygroundPage()),
			If(x.env.Dev, DebugSignals()),
		},
	}))
}

func renderNode(w io.Writer, node Node) {
	if err := node.Render(w); err != nil {
		slog.Error("view render error", "err", err)
	}
}
