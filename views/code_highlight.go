package views

import (
	"bytes"
	htmlescape "html"
	"log/slog"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	chromastyles "github.com/alecthomas/chroma/v2/styles"
)

func HighlightCode(language string, src string) string {
	return highlightCode(language, src, true)
}

func HighlightCodeNoLines(language string, src string) string {
	return highlightCode(language, src, false)
}

func highlightCode(language string, src string, lineNumbers bool) string {
	lexer := lexers.Get(language)
	iterator, err := lexer.Tokenise(nil, src)
	if err != nil {
		slog.Error("code highlight error", "err", err)
		return `<pre class="code-block"><code>` + htmlescape.EscapeString(src) + `</code></pre>`
	}

	formatter := chromahtml.New(chromahtml.WithLineNumbers(lineNumbers))
	style := chromastyles.Get("github-dark")

	var buf bytes.Buffer
	if err := formatter.Format(&buf, style, iterator); err != nil {
		slog.Error("code format error", "err", err)
		return `<pre class="code-block"><code>` + htmlescape.EscapeString(src) + `</code></pre>`
	}

	return buf.String()
}
