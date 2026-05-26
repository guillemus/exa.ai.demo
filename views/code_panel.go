package views

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"exa.ai.demo/exa"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	.code-panel {
		height: 100vh;
		overflow-y: auto;
		overflow-x: hidden;
		background: var(--bg-code-panel);
		color: #e9e9e9;
		position: relative;
		font-family: var(--font-code);
		scrollbar-color: #3b3b3b #111;
		scrollbar-width: thin;
	}
	.code-panel * { min-width: 0; }
	.code-panel::-webkit-scrollbar, .code-panel *::-webkit-scrollbar { width: var(--size-2); height: var(--size-2); }
	.code-panel::-webkit-scrollbar-track, .code-panel *::-webkit-scrollbar-track { background: #111; border-radius: var(--radius-round); }
	.code-panel::-webkit-scrollbar-thumb, .code-panel *::-webkit-scrollbar-thumb { background: #3b3b3b; border-radius: var(--radius-round); border: var(--border-size-2) solid #111; }
	.code-panel::-webkit-scrollbar-thumb:hover, .code-panel *::-webkit-scrollbar-thumb:hover { background: #505050; }
	.code-tabs, .language-tabs, .output-tabs {
		display: flex; align-items: center; border-bottom: var(--border-size-1) solid #2a2a2a;
	}
	.code-tabs { height: 52px; gap: 24px; padding: 0 24px; }
	.language-tabs, .output-tabs { height: 36px; gap: 30px; padding: 0 24px; background: #111; }
	.code-tab, .language-tab, .output-tab { border: 0; background: transparent; color: #9a9a9a; font-size: 16px; line-height: 24px; font-weight: 600; }
	.code-tab { height: 36px; padding: 0 12px; border: var(--border-size-1) solid transparent; }
	.language-tab, .output-tab { align-self: stretch; padding: 0; border-bottom: var(--border-size-2) solid transparent; display: inline-flex; align-items: center; gap: var(--size-2); }
	.code-tab.active { color: #84e8ff; background: #153840; border-color: #285b65; border-radius: var(--radius-2); }
	.language-tab.active, .output-tab.active { color: white; border-bottom-color: #75a7ff; }
	.tab-icon { color: #8f8f8f; font-size: var(--font-size-1); }
	.language-tab.active .tab-icon { color: white; }
	.code-example { position: relative; padding: 24px; }
	.copy-code-button {
		position: absolute;
		top: calc(var(--size-9) + var(--size-7) + var(--size-2));
		right: var(--size-5);
		z-index: 1;
		width: 32px;
		height: 32px;
		display: grid;
		place-items: center;
		border: 0;
		border-radius: var(--radius-2);
		background: #090909;
		color: #d8d8d8;
		font-size: var(--font-size-2);
	}
	.copy-code-button:hover, .copy-code-button.is-copied { background: #161616; color: white; }
	.copy-code-button.copy-state-button { display: grid; }
	.install-line { margin: 0 0 24px; padding: 12px 16px; border-radius: 6px; background: #1d1d1d; color: #f1f1d5; font-size: 14px; line-height: 20px; }
	.code-block, .code-example .highlighted-code .chroma {
		margin: var(--size-4) var(--size-5);
		padding: var(--size-4) var(--size-5);
		font-size: 14px;
		line-height: 24px;
		color: #d7d7d7;
		background: transparent !important;
		overflow: auto;
	}
	.output-json {
		margin: var(--size-4) var(--size-5);
		padding: var(--size-2);
		overflow: auto;
		color: #d7d7d7;
		font-family: var(--font-code);
		font-size: var(--font-size-0);
		line-height: var(--font-lineheight-3);
		scrollbar-color: #3b3b3b #111;
		scrollbar-width: thin;
	}
	.output-json::-webkit-scrollbar { height: var(--size-2); }
	.output-json::-webkit-scrollbar-track { background: #111; border-radius: var(--radius-round); }
	.output-json::-webkit-scrollbar-thumb { background: #3b3b3b; border-radius: var(--radius-round); border: var(--border-size-2) solid #111; }
	.output-json::-webkit-scrollbar-thumb:hover { background: #505050; }
	.output-json .chroma, .output-json pre {
		margin: 0;
		background: transparent !important;
		overflow: visible !important;
		white-space: pre;
	}
	.code-example .highlighted-code .chroma { margin: 0; }
	.output-loading { margin: var(--size-6) var(--size-5); color: #d7d7d7; font-size: var(--font-size-1); }
	.visual-output { display: grid; gap: var(--size-7); padding: var(--size-7) var(--size-5) var(--size-8); color: #e7e7e7; font-family: var(--font-text); }
	.visual-section { display: grid; gap: var(--size-4); }
	.visual-heading { margin: 0; font-size: var(--font-size-4); line-height: var(--font-lineheight-1); color: white; }
	.output-content { color: #d6d6d6; font-size: var(--font-size-2); line-height: var(--font-lineheight-3); }
	.output-content > :first-child { margin-top: 0; }
	.output-content > :last-child { margin-bottom: 0; }
	.output-content p, .output-content ul, .output-content ol { margin: 0 0 var(--size-3); }
	.output-content ul, .output-content ol { padding-left: var(--size-5); }
	.output-content li { margin: var(--size-2) 0; }
	.output-content strong { color: #fff; }
	.output-content code { padding: var(--size-1) var(--size-2); border-radius: var(--radius-2); background: #111; font-family: var(--font-code); }
	.output-content a { color: #84e8ff; }
	.structured-output .highlighted-code { overflow: hidden; }
	.structured-output .highlighted-code .chroma { margin: 0; padding: 0; overflow: hidden; font-size: var(--font-size-1); line-height: var(--font-lineheight-3); }
	.structured-output .highlighted-code table { width: 100%; table-layout: fixed; }
	.structured-output .highlighted-code pre { white-space: pre-wrap !important; overflow-wrap: anywhere; }
	.result-list { display: grid; gap: var(--size-5); }
	.result-card { border: var(--border-size-1) solid #3a3a3a; border-radius: var(--radius-3); background: #222; margin: 0; overflow: hidden; }
	.result-main { padding: var(--size-5); }
	.result-title { font-size: var(--font-size-4); line-height: var(--font-lineheight-1); color: white; margin-bottom: var(--size-3); }
	.result-meta { color: #9a9a9a; margin-bottom: var(--size-2); display: flex; gap: var(--size-3); }
	.result-url { color: #e2e2e2; word-break: break-all; }
	.result-content-toggle { margin-top: var(--size-4); display: flex; flex-direction: column; }
	.result-section-label { margin: 0 0 var(--size-2); color: #9a9a9a; font-size: var(--font-size-0); font-weight: 700; letter-spacing: var(--font-letterspacing-3); text-transform: uppercase; }
	.result-expand-check { display: none; }
	.result-content-button { order: 2; align-self: center; margin-top: var(--size-3); padding: var(--size-2) var(--size-3); border: var(--border-size-1) solid #4a4a4a; border-radius: var(--radius-round); color: #f0f0d8; cursor: pointer; }
	.result-content-button .show-less { display: none; }
	.result-content-preview { order: 1; max-height: var(--size-13); overflow: hidden; color: #d6d6d6; line-height: var(--font-lineheight-3); position: relative; }
	.result-content-preview::after { content: ""; position: absolute; left: 0; right: 0; bottom: 0; height: var(--size-8); background: linear-gradient(transparent, #222); }
	.result-expand-check:checked ~ .result-content-preview { max-height: none; }
	.result-expand-check:checked ~ .result-content-preview::after { display: none; }
	.result-expand-check:checked ~ .result-content-button .show-more { display: none; }
	.result-expand-check:checked ~ .result-content-button .show-less { display: inline; }
	.entity-toggle { border-top: var(--border-size-1) solid #3a3a3a; padding: var(--size-3) var(--size-5); color: #e2e2e2; }
	.entity-table { margin: var(--size-2) 0 var(--size-4); border: var(--border-size-1) solid #3a3a3a; border-radius: var(--radius-2); overflow: hidden; }
	.entity-heading { padding: var(--size-3); color: #aaa; letter-spacing: var(--font-letterspacing-3); text-transform: uppercase; border-bottom: var(--border-size-1) solid #333; }
	.entity-row { display: grid; grid-template-columns: minmax(var(--size-12), 1fr) 2fr; border-top: var(--border-size-1) solid #333; }
	.entity-row:first-child { border-top: 0; }
	.entity-cell { padding: var(--size-3); color: #ddd; }
	.entity-cell:first-child { color: #999; }
	.output-empty { min-height: calc(100vh - var(--size-10)); display: grid; place-items: center; color: #d6d6d6; font-family: var(--font-text); }
	.output-empty-inner { transform: translateY(var(--size-9)); text-align: center; }
	.output-empty-icon { font-size: var(--font-size-6); line-height: var(--font-lineheight-0); margin-bottom: var(--size-4); color: #e5e5e5; }
	.output-empty-text { font-size: var(--font-size-2); }
	.highlighted-code pre { margin: 0; background: transparent !important; }
	.highlighted-code code { font-family: var(--font-code); background: transparent !important; }
	.highlighted-code .lnt, .highlighted-code .ln { color: #8b949e; padding-right: var(--size-3); user-select: none; }
	@media (max-width: 1100px) {
		.code-panel { min-height: 70vh; height: auto; }
	}
`)

type CodePanelData struct {
	Form       SearchForm
	PanelTab   string
	CodeTab    string
	OutputTab  string
	OutputJSON string
	Response   *exa.SearchResponse
	Loading    bool
}

func CodePanel(state PageState) Node {
	return Aside(Class("code-panel"),
		CodePanelContent(CodePanelData{
			Form:      state.Form,
			PanelTab:  state.PanelTab,
			CodeTab:   state.CodeTab,
			OutputTab: state.OutputTab,
		}),
	)
}

func CodePanelContent(data CodePanelData) Node {
	return Div(ID("code-panel-content"),
		Div(Class("code-tabs"),
			PanelTabButton("code", "▣ Code", data.PanelTab),
			PanelTabButton("output", "◇ Output", data.PanelTab),
		),
		Div(Data("show", "$panelTab == 'code'"), Attr("style", showStyle(data.PanelTab != "output")), CodeContent(data.Form, data.CodeTab)),
		Div(Data("show", "$panelTab == 'output'"), Attr("style", showStyle(data.PanelTab == "output")),
			Nav(Class("output-tabs"), Attr("aria-label", "Search output"),
				OutputTabButton("json", "JSON", data.OutputTab),
				OutputTabButton("visual", "Visual", data.OutputTab),
			),
			OutputExample(data),
		),
	)
}

func activeTabClass(base string, active bool) string {
	if active {
		return base + " active"
	}
	return base
}

func CodeContent(form SearchForm, codeTab string) Node {
	return Div(ID("code-panel-code"),
		Nav(Class("language-tabs"), Attr("aria-label", "Code examples"),
			CodeTabButton("python", "♣", "Python", codeTab),
			CodeTabButton("javascript", "⬡", "Javascript", codeTab),
			CodeTabButton("curl", ">_", "curl", codeTab),
		),
		CodeExample("python", "pip install exa-py", PythonSearchCode(form), HighlightCode("python", PythonSearchCode(form)), codeTab),
		CodeExample("javascript", "npm install exa-js", JavaScriptSearchCode(form), HighlightCode("javascript", JavaScriptSearchCode(form)), codeTab),
		CodeExample("curl", "", CurlSearchCode(form), HighlightCode("bash", CurlSearchCode(form)), codeTab),
	)
}

func PanelTabButton(tab, label string, current string) Node {
	return Button(
		Type("button"),
		Class(activeTabClass("code-tab", current == tab)),
		Data("on:click", "$panelTab = '"+tab+"'"),
		Data("class:active", "$panelTab == '"+tab+"'"),
		Text(label),
	)
}

func CodeTabButton(tab, icon, label string, current string) Node {
	return Button(
		Type("button"),
		Class(activeTabClass("language-tab", current == tab)),
		Data("on:click", "$codeTab = '"+tab+"'"),
		Data("class:active", "$codeTab == '"+tab+"'"),
		Data("attr:aria-selected", "$codeTab == '"+tab+"'"),
		Span(Class("tab-icon"), Text(icon)),
		Span(Text(label)),
	)
}

func CodeExample(tab, install, code, highlighted string, current string) Node {
	children := []Node{
		Class("code-example"),
		Data("show", "$codeTab == '"+tab+"'"),
	}
	if current != tab {
		children = append(children, Attr("style", "display: none"))
	}
	if install != "" {
		children = append(children, Div(Class("install-line"), Code(Text(install))))
	}
	children = append(children,
		Button(
			Type("button"),
			Class("copy-code-button copy-state-button"),
			Data("on:click", "copyToClipboard("+strconv.Quote(code)+", el)"),
			Attr("aria-label", "Copy code"),
			Span(Class("copy-default"), Text("⧉")),
			Span(Class("copy-feedback"), Text("✓")),
		),
		Div(Class("highlighted-code"), Raw(highlighted)),
	)
	return Div(children...)
}

func OutputTabButton(tab, label string, current string) Node {
	return Button(
		Type("button"),
		Class(activeTabClass("output-tab", current == tab)),
		Data("on:click", "$outputTab = '"+tab+"'"),
		Data("class:active", "$outputTab == '"+tab+"'"),
		Text(label),
	)
}

func OutputExample(data CodePanelData) Node {
	return Div(
		Div(Data("show", "$outputTab == 'json'"), Attr("style", showStyle(data.OutputTab != "visual")),
			If(data.Loading, Div(Class("output-loading"), Text("Searching Exa…"))),
			Iff(!data.Loading && data.OutputJSON != "", func() Node { return OutputJSON(data.OutputJSON) }),
			If(!data.Loading && data.OutputJSON == "", OutputEmptyState()),
		),
		Div(Data("show", "$outputTab == 'visual'"), Attr("style", showStyle(data.OutputTab == "visual")),
			If(data.Loading, Div(Class("output-loading"), Text("Searching Exa…"))),
			Iff(!data.Loading && data.Response != nil, func() Node { return VisualOutput(data.Response, data.Form) }),
			If(!data.Loading && data.Response == nil, OutputEmptyState()),
		),
	)
}

func showStyle(show bool) string {
	if show {
		return ""
	}
	return "display: none"
}

func OutputJSON(output string) Node {
	return Div(Class("output-json"), Raw(HighlightCodeNoLines("json", output)))
}

func OutputEmptyState() Node {
	return Div(Class("output-empty"),
		Div(Class("output-empty-inner"),
			Div(Class("output-empty-icon"), Text("✧")),
			Div(Class("output-empty-text"), Text("Click Run to see a response")),
		),
	)
}

func VisualOutput(resp *exa.SearchResponse, form SearchForm) Node {
	return Div(Class("visual-output"),
		Iff(form.StructuredOutputs, func() Node { return OutputContent(resp) }),
		Iff(form.StructuredOutputs, func() Node { return StructuredOutput(resp) }),
		Iff(len(resp.Results) > 0, func() Node {
			return Div(Class("visual-section"), H3(Class("visual-heading"), Text(fmt.Sprintf("Results (%d)", len(resp.Results)))), Div(Class("result-list"), Group(Map(resp.Results, ResultCard))))
		}),
	)
}

func OutputContent(resp *exa.SearchResponse) Node {
	return Div(Class("visual-section"),
		H4(Class("visual-heading"), Text("Output Content")),
		Div(Class("output-content"), Raw(RenderMarkdown(outputContentText(resp)))),
	)
}

func StructuredOutput(resp *exa.SearchResponse) Node {
	bs, err := json.MarshalIndent(structuredOutput(resp), "", "  ")
	if err != nil {
		return nil
	}
	return Div(Class("visual-section structured-output"),
		H4(Class("visual-heading"), Text("Structured Output")),
		Div(Class("highlighted-code"), Raw(HighlightCode("json", string(bs)))),
	)
}

func ResultCard(result exa.Result) Node {
	return Div(Class("result-card"),
		Div(Class("result-main"),
			Div(Class("result-title"), Text(result.Title)),
			Div(Class("result-meta"), Span(Text("By Exa")), Iff(result.PublishedDate != nil, func() Node { return Span(Text(*result.PublishedDate)) })),
			Div(Class("result-url"), Text(result.URL)),
			Iff(len(result.Highlights) > 0, func() Node {
				return ResultExpandable(resultToggleID("highlights", result), "Highlights", markdownList(result.Highlights))
			}),
			Iff(result.Text != "", func() Node { return ResultExpandable(resultToggleID("text", result), "Text", result.Text) }),
		),
		Iff(result.Extras != nil && len(result.Extras.Entities) > 0, func() Node { return EntityDetails(result.Extras.Entities) }),
	)
}

func ResultExpandable(id string, label string, content string) Node {
	return Div(Class("result-content-toggle"),
		Div(Class("result-section-label"), Text(label)),
		Input(Type("checkbox"), ID(id), Class("result-expand-check")),
		Div(Class("result-content-preview output-content"), Raw(RenderMarkdown(content))),
		Label(Attr("for", id), Class("result-content-button"), Span(Class("show-more"), Text("Show More")), Span(Class("show-less"), Text("Show Less"))),
	)
}

func markdownList(items []string) string {
	var b strings.Builder
	for _, item := range items {
		b.WriteString("- ")
		b.WriteString(item)
		b.WriteString("\n")
	}
	return b.String()
}

func resultToggleID(kind string, result exa.Result) string {
	base := result.ID
	if base == "" {
		base = result.URL
	}
	return "result-" + kind + "-" + safeID(base)
}

func safeID(value string) string {
	var b strings.Builder
	for _, r := range value {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
		} else {
			b.WriteByte('-')
		}
	}
	return strings.Trim(b.String(), "-")
}

func EntityDetails(entities []exa.Entity) Node {
	return Details(Class("entity-toggle"),
		Summary(Text(fmt.Sprintf("Show Entities (%d)", len(entities)))),
		Group(Map(entities, EntityTable)),
	)
}

func EntityTable(entity exa.Entity) Node {
	return Div(Class("entity-table"),
		Div(Class("entity-heading"), Text(entity.Type)),
		Iff(entity.Properties.Name != nil, func() Node { return EntityRow("Name", *entity.Properties.Name) }),
		Iff(entity.Properties.Headquarters != nil, func() Node { return EntityRow("Headquarters", headquartersText(entity.Properties.Headquarters)) }),
		Iff(entity.Properties.Workforce != nil && entity.Properties.Workforce.Total != nil, func() Node {
			return EntityRow("Employees", fmt.Sprintf("%d", *entity.Properties.Workforce.Total))
		}),
		Iff(entity.Properties.WebTraffic != nil && entity.Properties.WebTraffic.VisitsMonthly != nil, func() Node {
			return EntityRow("Monthly Visits", compactNumber(*entity.Properties.WebTraffic.VisitsMonthly))
		}),
	)
}

func EntityRow(label, value string) Node {
	return Div(Class("entity-row"), Div(Class("entity-cell"), Text(label)), Div(Class("entity-cell"), Text(value)))
}

func outputContentText(resp *exa.SearchResponse) string {
	if resp.Output != nil {
		text := contentText(resp.Output.Content)
		if !strings.Contains(text, "[") {
			text += citationMarkers(resp.Output.Grounding)
		}
		return text
	}
	if resp.Context != "" {
		return resp.Context
	}
	return resultContent(resp.Results)
}

func structuredOutput(resp *exa.SearchResponse) any {
	if resp.Output != nil {
		return resp.Output
	}
	return map[string]any{
		"content":   resultContent(resp.Results),
		"grounding": resultGrounding(resp.Results),
	}
}

func contentText(content any) string {
	s, ok := content.(string)
	if ok {
		return s
	}
	bs, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return ""
	}
	return string(bs)
}

func citationMarkers(grounding []exa.GroundingInfo) string {
	if len(grounding) == 0 || len(grounding[0].Citations) == 0 {
		return ""
	}
	var b strings.Builder
	for i := range grounding[0].Citations {
		fmt.Fprintf(&b, " [%d]", i+1)
	}
	return b.String()
}

func resultContent(results []exa.Result) string {
	parts := []string{}
	for i, result := range results {
		text := resultSnippet(result)
		if text != "" {
			parts = append(parts, fmt.Sprintf("%s [%d]", text, i+1))
		}
	}
	return strings.Join(parts, " ")
}

func resultSnippet(result exa.Result) string {
	if len(result.Highlights) > 0 {
		return strings.Join(result.Highlights, " ")
	}
	if result.Summary != "" {
		return result.Summary
	}
	if result.Text != "" {
		return result.Text
	}
	return result.Title
}

func resultGrounding(results []exa.Result) []map[string]any {
	citations := []map[string]string{}
	for _, result := range results {
		citations = append(citations, map[string]string{"url": result.URL, "title": result.Title})
	}
	return []map[string]any{{
		"field":      "content",
		"citations":  citations,
		"confidence": "high",
	}}
}

func compactNumber(n int) string {
	if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return fmt.Sprintf("%d", n)
}

func headquartersText(h *exa.HeadquartersInfo) string {
	parts := []string{}
	if h.City != nil {
		parts = append(parts, *h.City)
	}
	if h.Country != nil {
		parts = append(parts, *h.Country)
	}
	if len(parts) == 0 && h.Address != nil {
		parts = append(parts, *h.Address)
	}
	return strings.Join(parts, ", ")
}
