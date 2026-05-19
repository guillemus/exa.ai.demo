package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	.code-panel {
		height: 100vh;
		overflow: auto;
		background: var(--bg-code-panel);
		color: #e9e9e9;
		position: relative;
		font-family: var(--font-code);
	}
	.code-tabs, .language-tabs {
		display: flex; gap: 20px; align-items: center; height: 56px; padding: 0 24px; border-bottom: 1px solid #2a2a2a;
	}
	.language-tabs { height: 48px; gap: 18px; }
	.code-tab, .language-tab { border: 0; background: transparent; color: #9a9a9a; font-size: 15px; }
	.code-tab.active { color: #dffaff; background: #153840; border: 1px solid #285b65; border-radius: var(--radius-2); padding: var(--size-1) var(--size-3); }
	.language-tab.active { color: white; font-weight: 700; border-bottom: 2px solid #75a7ff; padding-bottom: 13px; }
	.install-line { margin: var(--size-4) var(--size-5); padding: var(--size-3) var(--size-4); border-radius: var(--radius-2); background: var(--bg-code-soft); color: #f1f1d5; font-size: 14px; }
	.code-block, .highlighted-code .chroma { margin: 20px 24px; font-size: 14px; line-height: 1.55; color: #d7d7d7; background: transparent; }
	.highlighted-code pre { margin: 0; }
	.highlighted-code code { font-family: var(--font-code); }
	.highlighted-code .ln { color: #8b949e; padding-right: 16px; }
	.highlighted-code .k, .highlighted-code .kn { color: #ff7b72; }
	.highlighted-code .s, .highlighted-code .s1, .highlighted-code .s2 { color: #a5d6ff; }
	.highlighted-code .mi, .highlighted-code .kc { color: #79c0ff; }
	.highlighted-code .nf, .highlighted-code .n { color: #d2d7de; }
	.highlighted-code .o, .highlighted-code .p { color: #c9d1d9; }
	.exa-bot {
		position: fixed; right: 28px; bottom: 28px; width: 250px; height: 58px; border-radius: var(--radius-round); background: white;
		box-shadow: var(--shadow-6); display: flex; align-items: center; justify-content: space-between; padding: 0 var(--size-3) 0 var(--size-5); color: #666;
	}
	.exa-bot button { width: 38px; height: 38px; border-radius: var(--radius-round); border: 0; color: #999; background: var(--gray-2); font-size: 20px; }
	@media (max-width: 1100px) {
		.code-panel { min-height: 560px; height: auto; }
	}
`)

func CodePanel() Node {
	return Aside(Class("code-panel"),
		Div(Class("code-tabs"),
			Button(Type("button"), Class("code-tab active"), Text("▣ Code")),
			Button(Type("button"), Class("code-tab"), Text("◇ Output")),
		),
		Div(Class("language-tabs"),
			Button(Type("button"), Class("language-tab active"), Text("Python")),
			Button(Type("button"), Class("language-tab"), Text("Javascript")),
			Button(Type("button"), Class("language-tab"), Text("curl")),
		),
		Div(Class("install-line"), Code(Text("pip install exa-py"))),
		Div(Class("highlighted-code"), Raw(HighlightCode("python", `from exa_py import Exa

exa = Exa("47908a******************************")

result = exa.search(
    "Latest news on Nvidia",
    num_results = 10,
    type = "auto",
    contents = {
        "highlights": True
    }
)`))),
	)
}

func ExaBotBubble() Node {
	return Div(Class("exa-bot"), Span(Text("Ask ExaBot")), Button(Type("button"), Text("↑")))
}
