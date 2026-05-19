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
		display: flex; align-items: center; border-bottom: 1px solid #2a2a2a;
	}
	.code-tabs { height: 56px; gap: 22px; padding: 0 24px; }
	.language-tabs { height: 36px; gap: 28px; padding: 0 24px; background: #111; }
	.code-tab, .language-tab { border: 0; background: transparent; color: #9a9a9a; font-size: 15px; font-weight: 600; }
	.code-tab { height: 40px; padding: 0 12px; border: 1px solid transparent; }
	.language-tab { align-self: stretch; padding: 0; border-bottom: 2px solid transparent; display: inline-flex; align-items: center; gap: 10px; }
	.code-tab.active { color: #84e8ff; background: #153840; border-color: #285b65; border-radius: var(--radius-2); }
	.language-tab.active { color: white; border-bottom-color: #75a7ff; }
	.tab-icon { color: #8f8f8f; font-size: 15px; }
	.language-tab.active .tab-icon { color: white; }
	.install-line { margin: 16px 24px 18px; padding: 13px 16px; border-radius: var(--radius-2); background: var(--bg-code-soft); color: #f1f1d5; font-size: 14px; }
	.code-block, .highlighted-code .chroma { margin: 20px 24px; font-size: 14px; line-height: 1.55; color: #d7d7d7; background: transparent; }
	.highlighted-code pre { margin: 0; }
	.highlighted-code code { font-family: var(--font-code); }
	.highlighted-code .ln { color: #8b949e; padding-right: 16px; }
	.highlighted-code .k, .highlighted-code .kn { color: #ff7b72; }
	.highlighted-code .s, .highlighted-code .s1, .highlighted-code .s2 { color: #a5d6ff; }
	.highlighted-code .mi, .highlighted-code .kc { color: #79c0ff; }
	.highlighted-code .nf, .highlighted-code .n { color: #d2d7de; }
	.highlighted-code .o, .highlighted-code .p { color: #c9d1d9; }
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
		Nav(Class("language-tabs"), Attr("aria-label", "Code examples"),
			CodeTabButton("python", "♣", "Python"),
			CodeTabButton("javascript", "⬡", "Javascript"),
			CodeTabButton("curl", ">_", "curl"),
		),
		CodeExample("python", "pip install exa-py", HighlightCode("python", `from exa_py import Exa

exa = Exa("47908a******************************")

result = exa.search(
    "Latest news on Nvidia",
    num_results = 10,
    type = "auto",
    contents = {
        "highlights": True
    }
)`)),
		CodeExample("javascript", "npm install exa-js", HighlightCode("javascript", `import Exa from "exa-js";

const exa = new Exa("47908a******************************");

const result = await exa.search("Latest news on Nvidia", {
  numResults: 10,
  type: "auto",
  contents: {
    highlights: true,
  },
});`)),
		CodeExample("curl", "", HighlightCode("bash", `curl https://api.exa.ai/search \
  --request POST \
  --header "Content-Type: application/json" \
  --header "x-api-key: 47908a******************************" \
  --data '{
    "query": "Latest news on Nvidia",
    "numResults": 10,
    "type": "auto",
    "contents": {
      "highlights": true
    }
  }'`)),
	)
}

func CodeTabButton(tab, icon, label string) Node {
	return Button(
		Type("button"),
		Class("language-tab"),
		Data("on:click", "$codeTab = '"+tab+"'"),
		Data("class:active", "$codeTab == '"+tab+"'"),
		Data("attr:aria-selected", "$codeTab == '"+tab+"'"),
		Span(Class("tab-icon"), Text(icon)),
		Span(Text(label)),
	)
}

func CodeExample(tab, install, highlighted string) Node {
	children := []Node{
		Class("code-example"),
		Data("show", "$codeTab == '"+tab+"'"),
	}
	if tab != "python" {
		children = append(children, Attr("style", "display: none"))
	}
	if install != "" {
		children = append(children, Div(Class("install-line"), Code(Text(install))))
	}
	children = append(children, Div(Class("highlighted-code"), Raw(highlighted)))
	return Div(children...)
}
