package views

import (
	"bytes"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var TippyHead = Group([]Node{
	Link(Rel("stylesheet"), Href("https://unpkg.com/tippy.js@6/dist/tippy.css")),
	Raw(`<script src="https://unpkg.com/@popperjs/core@2"></script>`),
	Raw(`<script src="https://unpkg.com/tippy.js@6"></script>`),
})

var _ = styles.Style(`
	.tooltip-trigger {
		width: var(--size-5);
		height: var(--size-5);
		border: 0;
		border-radius: var(--radius-round);
		background: transparent;
		color: #6d7583;
		opacity: .42;
		font: 700 var(--font-size-0)/1 var(--font-sans);
	}
	.tooltip-trigger:hover, .tooltip-trigger:focus-visible { opacity: .9; }
	.label-with-tooltip { display: inline-flex; align-items: center; gap: var(--size-1); }
	.tippy-box[data-theme~='exa'] {
		background: #0f172a;
		color: white;
		font-size: var(--font-size-0);
		line-height: var(--font-lineheight-2);
	}
	.tippy-box[data-theme~='exa'] .tippy-arrow { color: #0f172a; }
	.tippy-box[data-theme~='exa'] a {
		color: white;
		text-decoration: underline dotted;
		text-underline-offset: 3px;
	}
`)

func Tooltip(ariaLabel string, content ...Node) Node {
	return Button(
		Type("button"),
		Class("tooltip-trigger"),
		Data("tooltip", renderTooltipContent(content...)),
		Data("init", "initTooltip(el)"),
		Attr("aria-label", ariaLabel),
		Text("ⓘ"),
	)
}

func LabelWithTooltip(label string, tooltip ...Node) Node {
	return Span(Class("label-with-tooltip"), Text(label), Tooltip(label+" help", tooltip...))
}

func renderTooltipContent(content ...Node) string {
	var b bytes.Buffer
	Group(content).Render(&b)
	return b.String()
}
