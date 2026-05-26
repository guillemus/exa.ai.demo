package views

import (
	"fmt"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	:root {
		--blue: #1747ee;
		--blue-soft: #8fb2ff;
		--line: #e5e5e5;
		--muted: var(--gray-6);
		--text: #171717;
		--bg-page: white;
		--bg-code-panel: var(--gray-12);
		--bg-code-soft: var(--gray-11);
		--font-app-size: 14px;
	}
	* { box-sizing: border-box; }
	body {
		margin: 0;
		font-family: var(--font-text);
		font-size: var(--font-app-size);
		color: var(--text);
		background: var(--bg-page);
	}
	button, input, textarea { font: inherit; }
	button { cursor: pointer; }
	.playground-shell {
		min-height: 100vh;
		display: grid;
		grid-template-columns: 854px minmax(var(--size-13), 1fr);
		background: var(--bg-page);
	}
	.playground-form {
		height: 100vh;
		overflow: auto;
		padding: 34px 48px 32px;
		border-right: var(--border-size-1) solid var(--line);
	}
	@media (max-width: 1100px) {
		.playground-shell { grid-template-columns: 1fr; }
		.playground-form { height: auto; padding: var(--size-7) var(--size-4); }
		.field-row { grid-template-columns: 1fr; gap: var(--size-2); }
	}
	.page-header {
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
		gap: var(--size-5);
		margin-bottom: var(--size-6);
	}
	h1 { margin: 0; font-size: 25px; line-height: 32px; font-weight: 500; letter-spacing: -0.02em; }
	.muted { color: var(--muted); margin: 0; }
	.header-actions { display: flex; gap: var(--size-2); }
	.ghost-button {
		display: inline-flex;
		align-items: center;
		min-height: 28px;
		border: var(--border-size-1) solid var(--line);
		background: var(--bg-page);
		border-radius: 6px;
		padding: 3px 8px;
		color: #6a6e76;
		font-size: 14px;
		line-height: 20px;
		text-decoration: none;
	}
	.copy-state-button { display: inline-grid; align-items: center; }
	.copy-state-button > span { grid-area: 1 / 1; transition: opacity .12s ease; }
	.copy-feedback { opacity: 0; visibility: hidden; }
	.copy-state-button.is-copied .copy-default { opacity: 0; visibility: hidden; }
	.copy-state-button.is-copied .copy-feedback { opacity: 1; visibility: visible; }
	.query-block { margin-bottom: 36px; }
	.field-label { display: block; color: #666; margin-bottom: 8px; font-size: 14px; }
	.query-card {
		border: 0;
		border-radius: 10px;
		box-shadow: rgba(0, 0, 0, 0.2) 0px 0px 0px 1px, rgba(0, 0, 0, 0.02) 0px 4px 3px 0px, rgba(0, 0, 0, 0.05) 0px 2px 3px 0px, rgba(0, 0, 0, 0.07) 0px 1px 2px 0px;
		padding: 10px;
		min-height: 100px;
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	.query-input {
		border: 0;
		outline: 0;
		resize: none;
		width: 100%;
		font-size: 15px;
		line-height: 20px;
		min-height: 40px;
		padding: 2px 3px;
	}
	.query-footer { display: flex; justify-content: space-between; align-items: center; }
	.primary-button {
		border: 0;
		color: white;
		border-radius: 8px;
		padding: 4px 10px;
		font-size: 16px;
		line-height: 24px;
		background: linear-gradient(180deg, #002289, #1747ee);
		box-shadow: rgb(99, 141, 255) 0px -1.5px 2px 0px inset, rgb(0, 67, 251) 0px 0px 10px 0px inset, rgb(0, 67, 251) 0px 0px 8px 0px inset;
	}
	.section { margin: 28px 0; }
	.section-title, .section-heading { font-size: 18px; line-height: 24px; margin: 0 0 12px; font-weight: 600; }
	.field-stack, .filters-section { display: grid; gap: var(--size-3); }
	.field-stack { margin-bottom: 36px; }
	.field-row {
		display: grid;
		grid-template-columns: minmax(var(--size-12), 1fr) 220px;
		gap: 32px;
		align-items: center;
		min-height: 46px;
	}
	.field-copy label, .copy strong { display: block; font-size: 16px; line-height: 20px; font-weight: 500; }
	.field-copy small { display: block; color: var(--muted); margin-top: var(--border-size-2); }
	.text-input, .select-input {
		width: 100%; height: 34px; border: 1px solid #e5e5e5; border-radius: 6px; background: var(--bg-page); padding: 0 10px; color: #171717;
		font-size: 14px;
		line-height: 20px;
	}
	.text-input:disabled { color: #a1a5ad; background: #fcfcfd; }
	.prompt-input { height: var(--size-10); padding: var(--size-2) var(--size-3); resize: vertical; }
	.select-input { text-align: left; display: flex; align-items: center; justify-content: space-between; }
	select.select-input { appearance: auto; cursor: pointer; }
	.contents-section { margin-top: 36px; }
	.toggle-row {
		display: grid;
		grid-template-columns: 1fr auto;
		align-items: center;
		gap: var(--size-4);
		padding: 10px 0;
	}
	.toggle {
		width: 42px; height: 22px; border: 0; border-radius: var(--radius-round); background-color: #dedede; padding: 2px; display: flex; align-items: center;
		transition: background-color .16s ease;
	}
	.toggle span { display: block; width: 18px; height: 18px; border-radius: 50%; background: white; box-shadow: var(--shadow-2); transform: translateX(0); transition: transform .16s ease; }
	.toggle.is-on { background-color: #1747ee; }
	.toggle.is-on span { transform: translateX(20px); }
	.nested-fields {
		border-left: var(--border-size-2) solid var(--line);
		margin-left: var(--size-2);
		padding: var(--size-1) 0 var(--size-2) var(--size-4);
		display: grid;
		gap: var(--size-2);
	}
	.disabled-nest { opacity: .55; }
	.subsection { padding-top: var(--size-3); }
	.unit-input { position: relative; }
	.unit-input span {
		position: absolute; right: var(--size-1); top: var(--size-1); height: var(--size-5); min-width: var(--size-7); display: grid; place-items: center;
		border: var(--border-size-1) solid var(--line); border-radius: var(--radius-2); color: #666; background: var(--bg-page);
	}
	.unit-input .text-input { padding-right: var(--size-9); }
	.advanced-button { border: 0; background: transparent; color: #6d737c; padding: var(--size-4) 0; font-size: var(--font-size-1); }
`)

func PlaygroundPage(state PageState) Node {
	return Div(
		Class("playground-shell"),
		Div(Class("playground-form"),
			HeaderBar(),
			QueryCard(state.Form),
			SearchTypeCard(state.Form),
			SimpleFields(state.Form),
			ContentsSection(state.Form),
			FiltersSection(state.Form),
		),
		CodePanel(state),
	)
}

func HeaderBar() Node {
	return Div(Class("page-header"),
		Div(H1(Text("Search"))),
		Div(Class("header-actions"),
			A(Class("ghost-button"), Href("https://exa.ai/docs/reference/search-api-guide"), Target("_blank"), Rel("noopener noreferrer"), Text("▣ Docs")),
			Button(
				Type("button"),
				Class("ghost-button copy-state-button"),
				Data("tooltip", "Copy context for using this endpoint with AI"),
				Data("init", "initTooltip(el)"),
				Data("on:click", "copyToClipboard("+strconv.Quote(CopyForAIText)+", el)"),
				Span(Class("copy-default"), Text("◉ ✺ ◒ Copy for AI")),
				Span(Class("copy-feedback"), Text("✓ Copied")),
			),
		),
	)
}

func effectSyncQueryParamAndGetCode(signal, param, defaultValue string) Node {
	value := "$" + signal
	return Group([]Node{
		Data("effect", fmt.Sprint(`
			if (`, value, ` !== `, defaultValue, `) {
				syncQueryParam('`, param, `', `, value, `)
			}
		`)),
		Data("effect__debounce.150ms", fmt.Sprint(`
			`, value, `;
			@get('/code')
		`)),
	})
}

func QueryCard(form SearchForm) Node {
	return Section(Class("query-block"),
		Label(Class("field-label"), Text("Query")),
		Div(Class("query-card"),
			Textarea(Class("query-input"), Rows("2"),
				Data("bind:query", ""),
				effectSyncQueryParamAndGetCode("query", "query", strconv.Quote("Latest news on Nvidia")),
				Text(form.Query),
			),
			Div(Class("query-footer"),
				Span(),
				Button(Type("button"), Class("primary-button"), Data("on:click", "@post('/search')"), Text("Search ↵")),
			),
		),
	)
}

func SearchTypeCard(form SearchForm) Node {
	return Section(Class("section"),
		H2(Class("section-title"), LabelWithTooltip("Search Type",
			A(Href("https://exa.ai/docs/reference/search-api-guide-for-coding-agents#search-types"), Target("_blank"), Rel("noreferrer"),
				Text("See docs"),
			),
			Text(" for more details"),
		)),

		SearchTypeSlider(form),
	)
}

func SimpleFields(form SearchForm) Node {
	return Div(Class("field-stack"),
		FieldRow("Number of results", "Max: 100. Contact us for more results.", Input(
			Type("number"),
			Value(strconv.Itoa(int(form.NumResults))),
			Class("text-input"),
			Data("bind:num-results", ""),
			effectSyncQueryParamAndGetCode("numResults", "numResults", "10"),
		)),
		FieldRow("Result category", "", CategorySelect(form.Category)),
	)
}

func CategorySelect(category string) Node {
	return Select(Class("select-input"), Data("bind:category", ""), effectSyncQueryParamAndGetCode("category", "category", strconv.Quote("company")),
		CategoryOption(category, "", "—"),
		CategoryOption(category, "company", "Company"),
		CategoryOption(category, "research paper", "Research Paper"),
		CategoryOption(category, "news article", "News Article"),
		CategoryOption(category, "github", "Github"),
		CategoryOption(category, "personal site", "Personal Site"),
		CategoryOption(category, "people", "People"),
		CategoryOption(category, "financial report", "Financial Report"),
	)
}

func CategoryOption(current string, value string, label string) Node {
	return Option(Value(value), If(current == value, Selected()), Text(label))
}

func ContentsSection(form SearchForm) Node {
	return Section(Class("section contents-section"),
		H2(Class("section-heading"), Text("Contents")),
		ToggleRow("Structured outputs", "Return structured outputs in addition to search results.", "structuredOutputs", form.StructuredOutputs, effectSyncQueryParamAndGetCode("structuredOutputs", "structuredOutputs", "false")),
		StructuredOutputFields(form),
		ToggleRow("Highlights", "Token efficient page excerpts", "highlights", form.Highlights, effectSyncQueryParamAndGetCode("highlights", "highlights", "true")),
		NestedFields("highlights", form.Highlights,
			FieldRow("Max characters", "", Input(
				Type("number"),
				Placeholder("Default: 4000"),
				Class("text-input"),
				Value(strconv.Itoa(int(form.HighlightMaxCharacters))),
				Data("bind:highlight-max-characters", ""),
				effectSyncQueryParamAndGetCode("highlightMaxCharacters", "highlightMaxCharacters", "4000"),
				Data("attr:disabled", "!$highlights"),
				If(!form.Highlights, Disabled()),
			)),
			FieldRowLabel(
				LabelWithTooltip("Guiding query", Text("Optional natural language description of what to have highlights focus on.")), "",
				Input(
					Type("text"),
					Placeholder("e.g. key takeaways"),
					Class("text-input"),
					Value(form.HighlightQuery),
					Data("bind:highlight-query", ""),
					effectSyncQueryParamAndGetCode("highlightQuery", "highlightQuery", "''"),
					Data("attr:disabled", "!$highlights"),
					If(!form.Highlights, Disabled()),
				),
			),
		),
		ToggleRow("Full webpage text", "", "text", form.Text, effectSyncQueryParamAndGetCode("text", "text", "false")),
		NestedFields("text", form.Text,
			FieldRow("Max characters", "", Input(
				Type("number"),
				Value(strconv.Itoa(int(form.TextMaxCharacters))),
				Data("bind:text-max-characters", ""),
				effectSyncQueryParamAndGetCode("textMaxCharacters", "textMaxCharacters", "20000"),
				Data("attr:disabled", "!$text"),
				If(!form.Text, Disabled()),
				Class("text-input"),
			)),
			ToggleRowLabel(LabelWithTooltip("Main content only", Text("Only return the main content of the page, excluding navbars, banners, footers, and similar page chrome.")), "", "textMainContentOnly", form.TextMainContentOnly, effectSyncQueryParamAndGetCode("textMainContentOnly", "textMainContentOnly", "true")),
		),
		Div(Class("subsection"),
			Div(Class("copy"), Strong(Text("Livecrawl")), P(Class("muted"), Text("Manage content freshness"))),
			NestedFields("", true,
				FieldRowLabel(LabelWithTooltip("Max age", Text("Max age of cached content before livecrawl. 0 = always livecrawl. -1 = never livecrawl (cache only).")), "", UnitInput(
					"Default: cache only",
					"hr",
					"max-age-hours",
					signalIntValue(form.MaxAgeHours),
					effectSyncQueryParamAndGetCode("maxAgeHours", "maxAgeHours", "''"),
				)),
				FieldRowLabel(LabelWithTooltip("Livecrawl timeout", Text("Maximum time to wait for live crawling before giving up.")), "", UnitInput(
					"Max: 30000",
					"ms",
					"livecrawl-timeout",
					strconv.Itoa(int(form.LivecrawlTimeout)),
					effectSyncQueryParamAndGetCode("livecrawlTimeout", "livecrawlTimeout", "10000"),
				)),
			),
		),
	)
}

func StructuredOutputFields(form SearchForm) Node {
	style := "display: none"
	if form.StructuredOutputs {
		style = ""
	}
	return Div(Class("nested-fields"), Data("show", "$structuredOutputs"), Attr("style", style),
		ToggleRow("Stream response", "Return OpenAI-compatible SSE chunks as they arrive.", "streamResponse", form.StreamResponse, effectSyncQueryParamAndGetCode("streamResponse", "streamResponse", "false")),
		ToggleRow("System prompt", "Instructions for synthesized output.", "systemPromptEnabled", form.SystemPromptEnabled, effectSyncQueryParamAndGetCode("systemPromptEnabled", "systemPromptEnabled", "false")),
		FieldRow("Prompt", "", Textarea(
			Class("text-input prompt-input"),
			Rows("3"),
			Data("bind:system-prompt", ""),
			effectSyncQueryParamAndGetCode("systemPrompt", "systemPrompt", "''"),
			Data("attr:disabled", "!$systemPromptEnabled"),
			If(!form.SystemPromptEnabled, Disabled()),
			Text(form.SystemPrompt),
		)),
	)
}

func FiltersSection(form SearchForm) Node {
	return Section(Class("section filters-section"),
		H2(Class("section-heading"), Text("Filters")),
		FieldRow("Include domains", "", Input(
			Type("text"),
			Value(form.IncludeDomains),
			Placeholder("e.g. exa.ai, docs.exa.ai/reference"),
			Class("text-input"),
			Data("bind:include-domains", ""),
			effectSyncQueryParamAndGetCode("includeDomains", "includeDomains", "''"),
		)),
		FieldRow("Exclude domains", "", Input(
			Type("text"),
			Value(form.ExcludeDomains),
			Placeholder("e.g. reddit.com, twitter.com"),
			Class("text-input"),
			Data("bind:exclude-domains", ""),
			effectSyncQueryParamAndGetCode("excludeDomains", "excludeDomains", "''"),
		)),
		FieldRow("Published after", "", Input(
			Type("date"),
			Value(form.StartPublishedDate),
			Class("text-input"),
			Data("bind:start-published-date", ""),
			effectSyncQueryParamAndGetCode("startPublishedDate", "startPublishedDate", "''"),
		)),
		FieldRow("Published before", "", Input(
			Type("date"),
			Value(form.EndPublishedDate),
			Class("text-input"),
			Data("bind:end-published-date", ""),
			effectSyncQueryParamAndGetCode("endPublishedDate", "endPublishedDate", "''"),
		)),
		FieldRow("User location", "Two-letter ISO country code", Input(
			Type("text"),
			Value(form.UserLocation),
			Placeholder("US"),
			MaxLength("2"),
			Class("text-input"),
			Data("bind:user-location", ""),
			effectSyncQueryParamAndGetCode("userLocation", "userLocation", "''"),
		)),
	)
}

func FieldRow(label string, description string, control Node) Node {
	return FieldRowLabel(Text(label), description, control)
}

func FieldRowLabel(label Node, description string, control Node) Node {
	return Div(Class("field-row"),
		Div(Class("field-copy"),
			Label(label),
			If(description != "", Small(Text(description))),
		),
		Div(Class("field-control"), control),
	)
}

func ToggleRow(title string, description string, signal string, on bool, children ...Node) Node {
	return ToggleRowLabel(Text(title), description, signal, on, children...)
}

func ToggleRowLabel(title Node, description string, signal string, on bool, children ...Node) Node {
	return Div(Class("toggle-row"),
		Div(Class("copy"), Strong(title), If(description != "", P(Class("muted"), Text(description)))),
		Toggle(signal, on, children...),
	)
}

func Toggle(signal string, on bool, children ...Node) Node {
	return Button(
		Type("button"),
		Class(toggleClass(on)),
		Data("on:click", "$"+signal+" = !$"+signal),
		Data("class:is-on", "$"+signal),
		Group(children),
		Span(),
	)
}

func NestedFields(signal string, enabled bool, nodes ...Node) Node {
	class := "nested-fields"
	if signal != "" && !enabled {
		class += " disabled-nest"
	}
	return Div(
		Class(class),
		If(signal != "", Data("class:disabled-nest", "!$"+signal)),
		Group(nodes),
	)
}

func toggleClass(on bool) string {
	if on {
		return "toggle is-on"
	}
	return "toggle"
}

func UnitInput(placeholder string, unit string, signal string, value string, children ...Node) Node {
	return Div(Class("unit-input"), Input(
		Type("number"),
		Value(value),
		Placeholder(placeholder),
		Class("text-input"),
		Data("bind:"+signal, ""),
		Group(children),
	), Span(Text(unit)))
}

func signalIntValue(value SignalInt) string {
	if value == 0 {
		return ""
	}
	return strconv.Itoa(int(value))
}
