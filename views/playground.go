package views

import (
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	:root {
		--blue: #1747ee;
		--blue-soft: #8fb2ff;
		--line: var(--gray-3);
		--muted: var(--gray-6);
		--text: var(--gray-12);
		--bg-page: white;
		--bg-code-panel: var(--gray-12);
		--bg-code-soft: var(--gray-11);
		--font-app-size: 0.875rem;
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
		grid-template-columns: minmax(620px, 1.35fr) minmax(520px, 1fr);
		background: var(--bg-page);
	}
	.playground-form {
		height: 100vh;
		overflow: auto;
		padding: 50px 56px 44px;
		border-right: 1px solid var(--line);
	}
	@media (max-width: 1100px) {
		.playground-shell { grid-template-columns: 1fr; }
		.playground-form { height: auto; padding: 32px 22px; }
		.field-row { grid-template-columns: 1fr; gap: 8px; }
	}
	.page-header {
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
		gap: 24px;
		margin-bottom: 28px;
	}
	h1 { margin: 0; font-size: 28px; line-height: 1; font-weight: 500; letter-spacing: -0.04em; }
	.muted { color: var(--muted); margin: 0; }
	.header-actions { display: flex; gap: 8px; }
	.ghost-button {
		border: 1px solid var(--line);
		background: var(--bg-page);
		border-radius: var(--radius-2);
		padding: var(--size-1) var(--size-2);
		color: #6a6e76;
		font-size: 14px;
		text-decoration: none;
	}
	.query-block { margin-bottom: 36px; }
	.field-label { display: block; color: #656970; margin-bottom: 8px; }
	.query-card {
		border: 1px solid #dfe2e7;
		border-radius: var(--radius-3);
		box-shadow: 0 1px 2px rgba(17,24,39,.06), 0 6px 14px rgba(17,24,39,.08);
		padding: var(--size-3);
		min-height: 132px;
		display: flex;
		flex-direction: column;
		gap: var(--size-3);
	}
	.query-input {
		border: 0;
		outline: 0;
		resize: none;
		width: 100%;
		font-size: 15px;
		min-height: 56px;
	}
	.query-footer { display: flex; justify-content: space-between; align-items: center; }
	.primary-button {
		border: 0;
		color: white;
		border-radius: var(--radius-2);
		padding: var(--size-2) var(--size-3);
		font-size: 14px;
		background: linear-gradient(180deg, #002289, #1747ee);
		box-shadow: inset 0 -1.5px 2px rgba(130,170,255,.85), 0 5px 14px rgba(23,71,238,.28);
	}
	.section { margin: 28px 0; }
	.section-title, .section-heading { font-size: 18px; margin: 0 0 14px; font-weight: 600; }
	.field-stack { display: grid; gap: 16px; margin-bottom: 38px; }
	.field-row {
		display: grid;
		grid-template-columns: minmax(240px, 1fr) 300px;
		gap: 24px;
		align-items: center;
		min-height: 42px;
	}
	.field-copy label, .copy strong { display: block; font-size: 16px; font-weight: 500; }
	.field-copy small { display: block; color: var(--muted); margin-top: 3px; }
	.text-input, .select-input {
		width: 100%; height: 38px; border: 1px solid #dfe2e7; border-radius: var(--radius-2); background: var(--bg-page); padding: 0 var(--size-3); color: #15171b;
		font-size: 14px;
	}
	.text-input:disabled { color: #a1a5ad; background: #fcfcfd; }
	.prompt-input { height: 84px; padding: 10px var(--size-3); resize: vertical; }
	.select-input { text-align: left; display: flex; align-items: center; justify-content: space-between; }
	select.select-input { appearance: auto; cursor: pointer; }
	.contents-section { margin-top: 42px; }
	.toggle-row {
		display: grid;
		grid-template-columns: 1fr auto;
		align-items: center;
		gap: 20px;
		padding: 13px 0;
	}
	.toggle {
		width: 48px; height: 26px; border: 0; border-radius: var(--radius-round); background: #dedede; padding: 3px;
	}
	.toggle span { display: block; width: 20px; height: 20px; border-radius: 50%; background: white; box-shadow: 0 1px 2px rgba(0,0,0,.2); }
	.toggle.is-on { background: linear-gradient(180deg, #002289, #1747ee); }
	.toggle.is-on span { transform: translateX(22px); }
	.nested-fields {
		border-left: 2px solid var(--line);
		margin-left: 8px;
		padding: 4px 0 8px 22px;
		display: grid;
		gap: 10px;
	}
	.disabled-nest { opacity: .55; }
	.subsection { padding-top: 16px; }
	.unit-input { position: relative; }
	.unit-input span {
		position: absolute; right: 6px; top: 6px; height: 26px; min-width: 36px; display: grid; place-items: center;
		border: 1px solid var(--line); border-radius: var(--radius-2); color: #666; background: var(--bg-page);
	}
	.unit-input .text-input { padding-right: 54px; }
	.advanced-button { border: 0; background: transparent; color: #6d737c; padding: 18px 0; font-size: 15px; }
`)

const urlSignalsEffect = `
	$query;
	$panelTab;
	$codeTab;
	$outputTab;
	$searchType;
	$deepModel;
	$numResults;
	$category;
	$structuredOutputs;
	$streamResponse;
	$systemPromptEnabled;
	$systemPrompt;
	$highlights;
	$highlightMaxCharacters;
	$highlightQuery;
	$text;
	$textMaxCharacters;
	$maxAgeHours;
	$livecrawlTimeout;
	$includeDomains;
	$excludeDomains;
	$startPublishedDate;
	$endPublishedDate;
	$userLocation;
	syncSignalsToURL({
		query: $query,
		panelTab: $panelTab,
		codeTab: $codeTab,
		outputTab: $outputTab,
		searchType: $searchType,
		deepModel: $deepModel,
		numResults: $numResults,
		category: $category,
		structuredOutputs: $structuredOutputs,
		streamResponse: $streamResponse,
		systemPromptEnabled: $systemPromptEnabled,
		systemPrompt: $systemPrompt,
		highlights: $highlights,
		highlightMaxCharacters: $highlightMaxCharacters,
		highlightQuery: $highlightQuery,
		text: $text,
		textMaxCharacters: $textMaxCharacters,
		maxAgeHours: $maxAgeHours,
		livecrawlTimeout: $livecrawlTimeout,
		includeDomains: $includeDomains,
		excludeDomains: $excludeDomains,
		startPublishedDate: $startPublishedDate,
		endPublishedDate: $endPublishedDate,
		userLocation: $userLocation
	})
`

const codeRefreshEffect = `
	$query;
	$searchType;
	$deepModel;
	$numResults;
	$category;
	$structuredOutputs;
	$streamResponse;
	$systemPromptEnabled;
	$systemPrompt;
	$highlights;
	$highlightMaxCharacters;
	$highlightQuery;
	$text;
	$textMaxCharacters;
	$maxAgeHours;
	$livecrawlTimeout;
	$includeDomains;
	$excludeDomains;
	$startPublishedDate;
	$endPublishedDate;
	$userLocation;
	@get('/code')
`

func PlaygroundPage(state PageState) Node {
	return Div(Class("playground-shell"), Data("effect__debounce.150ms", urlSignalsEffect),
		Div(Class("playground-form"), Data("effect__debounce.150ms", codeRefreshEffect),
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
			A(Class("ghost-button"), Href("https://dashboard.exa.ai/playground/search"), Target("_blank"), Rel("noopener noreferrer"), Text("▣ Playground")),
			A(Class("ghost-button"), Href("https://exa.ai/docs/reference/search-api-guide"), Target("_blank"), Rel("noopener noreferrer"), Text("▣ Docs")),
			Button(Type("button"), Class("ghost-button"), Data("on:click", "copyToClipboard("+strconv.Quote(CopyForAIText)+")"), Text("◉ ✺ ◒ Copy for AI")),
		),
	)
}

func QueryCard(form SearchForm) Node {
	return Section(Class("query-block"),
		Label(Class("field-label"), Text("Query")),
		Div(Class("query-card"),
			Textarea(Class("query-input"), Rows("2"), Data("bind:query", ""), Text(form.Query)),
			Div(Class("query-footer"),
				Span(),
				Button(Type("button"), Class("primary-button"), Data("on:click", "@post('/search')"), Text("Search ↵")),
			),
		),
	)
}

func SearchTypeCard(form SearchForm) Node {
	return Section(Class("section"),
		H2(Class("section-title"), Text("Search Type ⓘ")),
		SearchTypeSlider(form),
	)
}

func SimpleFields(form SearchForm) Node {
	return Div(Class("field-stack"),
		FieldRow("Number of results", "Max: 100. Contact us for more results.", Input(
			Type("text"),
			Value(strconv.Itoa(int(form.NumResults))),
			Class("text-input"),
			Data("bind:num-results", ""),
		)),
		FieldRow("Result category", "", CategorySelect(form.Category)),
	)
}

func CategorySelect(category string) Node {
	return Select(Class("select-input"), Data("bind:category", ""),
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
		ToggleRow("Structured outputs", "Return structured outputs in addition to search results.", "structuredOutputs", form.StructuredOutputs),
		StructuredOutputFields(form),
		ToggleRow("Highlights", "Token efficient page excerpts", "highlights", form.Highlights),
		NestedFields("highlights", form.Highlights,
			FieldRow("Max characters", "", Input(
				Type("text"),
				Placeholder("Default: 4000"),
				Class("text-input"),
				Value(strconv.Itoa(int(form.HighlightMaxCharacters))),
				Data("bind:highlight-max-characters", ""),
				Data("attr:disabled", "!$highlights"),
				If(!form.Highlights, Disabled()),
			)),
			FieldRow("Guiding query ⓘ", "", Input(
				Type("text"),
				Placeholder("e.g. key takeaways"),
				Class("text-input"),
				Value(form.HighlightQuery),
				Data("bind:highlight-query", ""),
				Data("attr:disabled", "!$highlights"),
				If(!form.Highlights, Disabled()),
			)),
		),
		ToggleRow("Full webpage text", "", "text", form.Text),
		NestedFields("text", form.Text,
			FieldRow("Max characters", "", Input(
				Type("text"),
				Value(strconv.Itoa(int(form.TextMaxCharacters))),
				Data("bind:text-max-characters", ""),
				Data("attr:disabled", "!$text"),
				If(!form.Text, Disabled()),
				Class("text-input"),
			)),
		),
		Div(Class("subsection"),
			Div(Class("copy"), Strong(Text("Livecrawl")), P(Class("muted"), Text("Manage content freshness"))),
			NestedFields("", true,
				FieldRow("Max age ⓘ", "", UnitInput(
					"Default: cache only",
					"hr",
					"max-age-hours",
					signalIntValue(form.MaxAgeHours),
				)),
				FieldRow("Livecrawl timeout ⓘ", "", UnitInput(
					"Max: 30000",
					"ms",
					"livecrawl-timeout",
					strconv.Itoa(int(form.LivecrawlTimeout)),
				)),
			),
		),
		Button(Type("button"), Class("advanced-button"), Text("› Advanced Options")),
	)
}

func StructuredOutputFields(form SearchForm) Node {
	style := "display: none"
	if form.StructuredOutputs {
		style = ""
	}
	return Div(Class("nested-fields"), Data("show", "$structuredOutputs"), Attr("style", style),
		ToggleRow("Stream response", "Return OpenAI-compatible SSE chunks as they arrive.", "streamResponse", form.StreamResponse),
		ToggleRow("System prompt", "Instructions for synthesized output.", "systemPromptEnabled", form.SystemPromptEnabled),
		FieldRow("Prompt", "", Textarea(
			Class("text-input prompt-input"),
			Rows("3"),
			Data("bind:system-prompt", ""),
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
		)),
		FieldRow("Exclude domains", "", Input(
			Type("text"),
			Value(form.ExcludeDomains),
			Placeholder("e.g. reddit.com, twitter.com"),
			Class("text-input"),
			Data("bind:exclude-domains", ""),
		)),
		FieldRow("Published after", "", Input(
			Type("date"),
			Value(form.StartPublishedDate),
			Class("text-input"),
			Data("bind:start-published-date", ""),
		)),
		FieldRow("Published before", "", Input(
			Type("date"),
			Value(form.EndPublishedDate),
			Class("text-input"),
			Data("bind:end-published-date", ""),
		)),
		FieldRow("User location", "Two-letter ISO country code", Input(
			Type("text"),
			Value(form.UserLocation),
			Placeholder("US"),
			MaxLength("2"),
			Class("text-input"),
			Data("bind:user-location", ""),
		)),
	)
}

func FieldRow(label string, description string, control Node) Node {
	return Div(Class("field-row"),
		Div(Class("field-copy"),
			Label(Text(label)),
			If(description != "", Small(Text(description))),
		),
		Div(Class("field-control"), control),
	)
}

func ToggleRow(title string, description string, signal string, on bool) Node {
	return Div(Class("toggle-row"),
		Div(Class("copy"), Strong(Text(title)), If(description != "", P(Class("muted"), Text(description)))),
		Toggle(signal, on),
	)
}

func Toggle(signal string, on bool) Node {
	return Button(
		Type("button"),
		Class(toggleClass(on)),
		Data("on:click", "$"+signal+" = !$"+signal),
		Data("class:is-on", "$"+signal),
		Span(),
	)
}

func NestedFields(signal string, enabled bool, nodes ...Node) Node {
	class := "nested-fields"
	if signal != "" && !enabled {
		class += " disabled-nest"
	}
	attrs := []Node{Class(class)}
	if signal != "" {
		attrs = append(attrs, Data("class:disabled-nest", "!$"+signal))
	}
	attrs = append(attrs, Group(nodes))
	return Div(attrs...)
}

func toggleClass(on bool) string {
	if on {
		return "toggle is-on"
	}
	return "toggle"
}

func UnitInput(placeholder string, unit string, signal string, value string) Node {
	return Div(Class("unit-input"), Input(Type("text"), Value(value), Placeholder(placeholder), Class("text-input"), Data("bind:"+signal, "")), Span(Text(unit)))
}

func signalIntValue(value SignalInt) string {
	if value == 0 {
		return ""
	}
	return strconv.Itoa(int(value))
}
