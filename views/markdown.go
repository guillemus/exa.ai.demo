package views

import (
	"bytes"
	"log/slog"

	"github.com/yuin/goldmark"
)

func RenderMarkdown(src string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(src), &buf); err != nil {
		slog.Error("markdown render error", "err", err)
		return ""
	}
	return buf.String()
}
