package views

import (
	"encoding/json"
	"fmt"
	htmlescape "html"
	"strconv"
	"strings"

	"exa.ai.demo/exa"

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
	.code-tabs, .language-tabs, .output-tabs {
		display: flex; align-items: center; border-bottom: 1px solid #2a2a2a;
	}
	.code-tabs { height: 56px; gap: 22px; padding: 0 24px; }
	.language-tabs, .output-tabs { height: 36px; gap: 28px; padding: 0 24px; background: #111; }
	.code-tab, .language-tab, .output-tab { border: 0; background: transparent; color: #9a9a9a; font-size: 15px; font-weight: 600; }
	.code-tab { height: 40px; padding: 0 12px; border: 1px solid transparent; }
	.language-tab, .output-tab { align-self: stretch; padding: 0; border-bottom: 2px solid transparent; display: inline-flex; align-items: center; gap: 10px; }
	.code-tab.active { color: #84e8ff; background: #153840; border-color: #285b65; border-radius: var(--radius-2); }
	.language-tab.active, .output-tab.active { color: white; border-bottom-color: #75a7ff; }
	.tab-icon { color: #8f8f8f; font-size: 15px; }
	.language-tab.active .tab-icon { color: white; }
	.code-example { position: relative; padding: 24px; }
	.copy-code-button { position: absolute; top: 96px; right: 24px; z-index: 1; width: 32px; height: 32px; display: grid; place-items: center; border: 0; border-radius: var(--radius-2); background: #090909; color: #d8d8d8; font-size: 18px; }
	.copy-code-button:hover { background: #161616; color: white; }
	.install-line { margin: 0 0 24px; padding: 13px 16px; border-radius: var(--radius-2); background: var(--bg-code-soft); color: #f1f1d5; font-size: 14px; }
	.code-block, .highlighted-code .chroma { margin: 20px 24px; padding: 18px 24px; font-size: 14px; line-height: 2.05; color: #d7d7d7; background: transparent !important; overflow: auto; }
	.output-json { margin: 20px 24px; padding: 10px; overflow: auto; color: #d7d7d7; font-family: var(--font-code); font-size: 14px; line-height: 1.45; scrollbar-color: #3b3b3b #111; scrollbar-width: thin; }
	.output-json::-webkit-scrollbar { height: 10px; }
	.output-json::-webkit-scrollbar-track { background: #111; border-radius: var(--radius-round); }
	.output-json::-webkit-scrollbar-thumb { background: #3b3b3b; border-radius: var(--radius-round); border: 2px solid #111; }
	.output-json::-webkit-scrollbar-thumb:hover { background: #505050; }
	.output-json-line { white-space: pre; min-height: 1.45em; }
	.json-string { color: #a5d6ff; }
	.json-number { color: #79c0ff; }
	.json-keyword { color: #ff7b72; }
	.json-punct { color: #d2a8ff; }
	.code-example .highlighted-code .chroma { margin: 0; }
	.output-loading { margin: 28px 24px; color: #d7d7d7; font-size: 15px; }
	.visual-output { padding: 28px 24px 48px; color: #e7e7e7; font-family: var(--font-text); }
	.visual-output h3 { margin: 0 0 18px; font-size: 24px; color: white; }
	.visual-output h4 { margin: 30px 0 12px; font-size: 16px; color: #d8d8d8; }
	.output-content { color: #d6d6d6; line-height: 1.6; }
	.result-card { border: 1px solid #3a3a3a; border-radius: var(--radius-3); background: #222; margin: 16px 0; overflow: hidden; }
	.result-main { padding: 18px 22px 22px; }
	.result-title { font-size: 20px; color: white; margin-bottom: 12px; }
	.result-meta { color: #9a9a9a; margin-bottom: 8px; display: flex; gap: 16px; }
	.result-url { color: #e2e2e2; word-break: break-all; }
	.entity-toggle { border-top: 1px solid #3a3a3a; padding: 14px 22px; color: #e2e2e2; }
	.entity-table { margin: 10px 0 18px; border: 1px solid #3a3a3a; border-radius: var(--radius-2); overflow: hidden; }
	.entity-heading { padding: 16px; color: #aaa; letter-spacing: .08em; text-transform: uppercase; border-bottom: 1px solid #333; }
	.entity-row { display: grid; grid-template-columns: 180px 1fr; border-top: 1px solid #333; }
	.entity-row:first-child { border-top: 0; }
	.entity-cell { padding: 12px 16px; color: #ddd; }
	.entity-cell:first-child { color: #999; }
	.output-empty { min-height: calc(100vh - 92px); display: grid; place-items: center; color: #d6d6d6; font-family: var(--font-text); }
	.output-empty-inner { transform: translateY(80px); text-align: center; }
	.output-empty-icon { font-size: 42px; line-height: 1; margin-bottom: 18px; color: #e5e5e5; }
	.output-empty-text { font-size: 18px; }
	.highlighted-code pre { margin: 0; background: transparent !important; }
	.highlighted-code code { font-family: var(--font-code); background: transparent !important; }
	.highlighted-code .lnt, .highlighted-code .ln { color: #8b949e; padding-right: 16px; user-select: none; }
	@media (max-width: 1100px) {
		.code-panel { min-height: 560px; height: auto; }
	}
`)

type CodePanelData struct {
	Form       SearchForm
	OutputJSON string
	Response   *exa.SearchResponse
	Loading    bool
}

func CodePanel() Node {
	return Aside(Class("code-panel"), CodePanelContent(CodePanelData{Form: SearchForm{}}))
}

func CodePanelContent(data CodePanelData) Node {
	return Div(ID("code-panel-content"),
		Div(Class("code-tabs"),
			PanelTabButton("code", "▣ Code"),
			PanelTabButton("output", "◇ Output"),
		),
		Div(Data("show", "$panelTab == 'code'"),
			Nav(Class("language-tabs"), Attr("aria-label", "Code examples"),
				CodeTabButton("python", "♣", "Python"),
				CodeTabButton("javascript", "⬡", "Javascript"),
				CodeTabButton("curl", ">_", "curl"),
			),
			CodeExample("python", "pip install exa-py", PythonSearchCode(data.Form), HighlightCode("python", PythonSearchCode(data.Form))),
			CodeExample("javascript", "npm install exa-js", JavaScriptSearchCode(data.Form), HighlightCode("javascript", JavaScriptSearchCode(data.Form))),
			CodeExample("curl", "", CurlSearchCode(data.Form), HighlightCode("bash", CurlSearchCode(data.Form))),
		),
		Div(Data("show", "$panelTab == 'output'"), Attr("style", "display: none"),
			Nav(Class("output-tabs"), Attr("aria-label", "Search output"),
				OutputTabButton("json", "JSON"),
				OutputTabButton("visual", "Visual"),
			),
			OutputExample(data),
		),
	)
}

func PanelTabButton(tab, label string) Node {
	return Button(
		Type("button"),
		Class("code-tab"),
		Data("on:click", "$panelTab = '"+tab+"'"),
		Data("class:active", "$panelTab == '"+tab+"'"),
		Text(label),
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

func CodeExample(tab, install, code, highlighted string) Node {
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
	children = append(children,
		Button(Type("button"), Class("copy-code-button"), Data("on:click", "copyToClipboard("+strconv.Quote(code)+")"), Attr("aria-label", "Copy code"), Text("⧉")),
		Div(Class("highlighted-code"), Raw(highlighted)),
	)
	return Div(children...)
}

func OutputTabButton(tab, label string) Node {
	return Button(
		Type("button"),
		Class("output-tab"),
		Data("on:click", "$outputTab = '"+tab+"'"),
		Data("class:active", "$outputTab == '"+tab+"'"),
		Text(label),
	)
}

func OutputExample(data CodePanelData) Node {
	return Div(
		Div(Data("show", "$outputTab == 'json'"),
			If(data.Loading, Div(Class("output-loading"), Text("Searching Exa…"))),
			Iff(!data.Loading && data.OutputJSON != "", func() Node { return OutputJSON(data.OutputJSON) }),
			If(!data.Loading && data.OutputJSON == "", OutputEmptyState()),
		),
		Div(Data("show", "$outputTab == 'visual'"), Attr("style", "display: none"),
			If(data.Loading, Div(Class("output-loading"), Text("Searching Exa…"))),
			Iff(!data.Loading && data.Response != nil, func() Node { return VisualOutput(data.Response) }),
			If(!data.Loading && data.Response == nil, OutputEmptyState()),
		),
	)
}

func OutputJSON(output string) Node {
	lines := strings.Split(output, "\n")
	return Div(Class("output-json"), Group(Map(lines, func(line string) Node {
		return Div(Class("output-json-line"), Raw(highlightJSONLine(line)))
	})))
}

func highlightJSONLine(line string) string {
	var b strings.Builder
	for i := 0; i < len(line); {
		c := line[i]
		if c == '"' {
			end := i + 1
			for end < len(line) {
				if line[end] == '\\' {
					end += 2
					continue
				}
				if line[end] == '"' {
					end++
					break
				}
				end++
			}
			b.WriteString(`<span class="json-string">`)
			b.WriteString(htmlescape.EscapeString(line[i:end]))
			b.WriteString(`</span>`)
			i = end
			continue
		}
		if strings.ContainsRune("{}[]:,", rune(c)) {
			b.WriteString(`<span class="json-punct">`)
			b.WriteByte(c)
			b.WriteString(`</span>`)
			i++
			continue
		}
		if c == '-' || c >= '0' && c <= '9' {
			end := i + 1
			for end < len(line) && strings.ContainsRune("0123456789.eE+-", rune(line[end])) {
				end++
			}
			b.WriteString(`<span class="json-number">`)
			b.WriteString(htmlescape.EscapeString(line[i:end]))
			b.WriteString(`</span>`)
			i = end
			continue
		}
		matched := false
		for _, keyword := range []string{"true", "false", "null"} {
			if strings.HasPrefix(line[i:], keyword) {
				b.WriteString(`<span class="json-keyword">`)
				b.WriteString(keyword)
				b.WriteString(`</span>`)
				i += len(keyword)
				matched = true
				break
			}
		}
		if matched {
			continue
		}
		b.WriteString(htmlescape.EscapeString(line[i : i+1]))
		i++
	}
	return b.String()
}

func OutputEmptyState() Node {
	return Div(Class("output-empty"),
		Div(Class("output-empty-inner"),
			Div(Class("output-empty-icon"), Text("✧")),
			Div(Class("output-empty-text"), Text("Click Run to see a response")),
		),
	)
}

func VisualOutput(resp *exa.SearchResponse) Node {
	return Div(Class("visual-output"),
		OutputContent(resp),
		StructuredOutput(resp),
		Iff(len(resp.Results) > 0, func() Node {
			return Div(H3(Text(fmt.Sprintf("Results (%d)", len(resp.Results)))), Group(Map(resp.Results, ResultCard)))
		}),
	)
}

func OutputContent(resp *exa.SearchResponse) Node {
	return Div(
		H4(Text("Output Content")),
		Div(Class("output-content"), Text(outputContentText(resp))),
	)
}

func StructuredOutput(resp *exa.SearchResponse) Node {
	bs, err := json.MarshalIndent(structuredOutput(resp), "", "  ")
	if err != nil {
		return nil
	}
	return Div(
		H4(Text("Structured Output")),
		Div(Class("highlighted-code"), Raw(HighlightCode("json", string(bs)))),
	)
}

func ResultCard(result exa.Result) Node {
	return Div(Class("result-card"),
		Div(Class("result-main"),
			Div(Class("result-title"), Text(result.Title)),
			Div(Class("result-meta"), Span(Text("By Exa")), Iff(result.PublishedDate != nil, func() Node { return Span(Text(*result.PublishedDate)) })),
			Div(Class("result-url"), Text(result.URL)),
		),
		Iff(result.Extras != nil && len(result.Extras.Entities) > 0, func() Node { return EntityDetails(result.Extras.Entities) }),
	)
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
