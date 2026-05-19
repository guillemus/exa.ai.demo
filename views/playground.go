package views

import (
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
	.slider-card {
		--slider-position: 66.666%;
		border: 1px solid var(--line);
		border-radius: var(--radius-3);
		padding: 28px 32px 24px;
	}
	.search-type-slider { user-select: none; touch-action: none; }
	.slider-track {
		position: relative;
		height: 28px;
		margin: 0 0 18px;
		cursor: pointer;
	}
	.slider-line, .slider-fill {
		position: absolute;
		left: 0;
		top: 10px;
		height: 8px;
		border-radius: var(--radius-round);
	}
	.slider-line { right: 0; background: #cfd2d8; }
	.slider-fill { width: var(--slider-position); background: linear-gradient(90deg, #9dc0ff, #356bf3); }
	.dot {
		position: absolute;
		top: 10px;
		width: 8px;
		height: 8px;
		border-radius: var(--radius-round);
		background: white;
		transform: translate(-50%, 0);
		pointer-events: none;
	}
	.dot-1 { left: 0; } .dot-2 { left: 33.333%; } .dot-3 { left: 66.666%; } .dot-4 { left: 100%; }
	.slider-thumb {
		position: absolute;
		left: var(--slider-position);
		top: 2px;
		width: 24px;
		height: 24px;
		border: 2px solid #5c84ee;
		border-radius: var(--radius-round);
		background: #cddcff;
		box-shadow: 0 2px 7px rgba(23,71,238,.35);
		transform: translateX(-50%);
		pointer-events: none;
	}
	.slider-labels { display: grid; grid-template-columns: repeat(4, 1fr); }
	.slider-option { display: flex; flex-direction: column; gap: 2px; border: 0; background: transparent; padding: 0; text-align: left; color: var(--text); }
	.slider-option:nth-child(4) { text-align: right; align-items: flex-end; }
	.slider-option span { color: #6d7583; font-size: 14px; }
	.slider-option.is-active strong { color: #0f172a; }
	.slider-option.is-active span { color: #4b5563; }
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
	.select-input { text-align: left; display: flex; align-items: center; justify-content: space-between; }
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

func PlaygroundPage() Node {
	return Div(Class("playground-shell"),
		Div(Class("playground-form"),
			HeaderBar(),
			QueryCard(),
			SearchTypeCard(),
			SimpleFields(),
			ContentsSection(),
			FiltersSection(),
		),
		CodePanel(),
		ExaBotBubble(),
	)
}

func HeaderBar() Node {
	return Div(Class("page-header"),
		Div(H1(Text("Search"))),
		Div(Class("header-actions"),
			Button(Type("button"), Class("ghost-button"), Text("▣ Docs")),
			Button(Type("button"), Class("ghost-button"), Text("◉ ✺ ◒ Copy for AI")),
		),
	)
}

func QueryCard() Node {
	return Section(Class("query-block"),
		Label(Class("field-label"), Text("Query")),
		Div(Class("query-card"),
			Textarea(Class("query-input"), Rows("2"), Attr("data-bind:query", ""), Text("Latest news on Nvidia")),
			Div(Class("query-footer"),
				Span(),
				Button(Type("button"), Class("primary-button"), Text("Search ↵")),
			),
		),
	)
}

func SearchTypeCard() Node {
	return Section(Class("section"),
		H2(Class("section-title"), Text("Search Type ⓘ")),
		Div(Class("slider-card search-type-slider"), Attr("data-search-type-slider", ""), Attr("data-init", "initSearchTypeSlider(el)"),
			Input(Type("hidden"), Attr("data-bind:search-type", ""), Value("auto"), Attr("data-search-type-value", "")),
			Div(Class("slider-track"), Attr("data-search-type-track", ""),
				Span(Class("slider-line")),
				Span(Class("slider-fill")),
				Span(Class("dot dot-1")),
				Span(Class("dot dot-2")),
				Span(Class("dot dot-3")),
				Span(Class("dot dot-4")),
				Span(Class("slider-thumb")),
			),
			Div(Class("slider-labels"),
				SliderOption("instant", "Instant", "200ms"),
				SliderOption("fast", "Fast", "450ms"),
				SliderOption("auto", "Auto", "1s (recommended)"),
				SliderOption("deep", "Deep", "4s-18s"),
			),
		),
	)
}

func SliderOption(value string, title string, subtitle string) Node {
	return Button(
		Type("button"),
		Class("slider-option"),
		Attr("data-search-type-option", value),
		Attr("data-on:click", "$searchType = '"+value+"'"),
		Attr("data-class:is-active", "$searchType == '"+value+"'"),
		Strong(Text(title)),
		Span(Text(subtitle)),
	)
}

func SimpleFields() Node {
	return Div(Class("field-stack"),
		FieldRow("Number of results", "Max: 100. Contact us for more results.", Input(Type("text"), Value("10"), Class("text-input"))),
		FieldRow("Result category", "", Button(Type("button"), Class("select-input"), Text("—⌄"))),
	)
}

func ContentsSection() Node {
	return Section(Class("section contents-section"),
		H2(Class("section-heading"), Text("Contents")),
		ToggleRow("Structured outputs", "Return structured outputs in addition to search results.", false),
		ToggleRow("Highlights", "Token efficient page excerpts", true),
		NestedFields(
			FieldRow("Max characters", "", Input(Type("text"), Placeholder("Default: 4000"), Class("text-input"))),
			FieldRow("Guiding query ⓘ", "", Input(Type("text"), Placeholder("e.g. key takeaways"), Class("text-input"))),
		),
		ToggleRow("Full webpage text", "", false),
		NestedFieldsDisabled(
			FieldRow("Max characters", "", Input(Type("text"), Value("20,000"), Disabled(), Class("text-input"))),
			FieldRow("Main content only ⓘ", "", Toggle(true)),
		),
		Div(Class("subsection"),
			Div(Class("copy"), Strong(Text("Livecrawl")), P(Class("muted"), Text("Manage content freshness"))),
			NestedFields(
				FieldRow("Max age ⓘ", "", UnitInput("Default: cache only", "hr")),
				FieldRow("Livecrawl timeout ⓘ", "", UnitInput("Max: 30000", "ms")),
			),
		),
		Button(Type("button"), Class("advanced-button"), Text("› Advanced Options")),
	)
}

func FiltersSection() Node {
	return Section(Class("section filters-section"),
		H2(Class("section-heading"), Text("Filters")),
		FieldRow("Include domains", "", Input(Type("text"), Placeholder("e.g. exa.ai, docs.exa.ai/reference"), Class("text-input"))),
		FieldRow("Exclude domains", "", Input(Type("text"), Placeholder("e.g. reddit.com, twitter.com"), Class("text-input"))),
		FieldRow("Published date range", "", Button(Type("button"), Class("select-input"), Text("Select date range  ◷"))),
		FieldRow("User location", "Select a country to localize search results", Button(Type("button"), Class("select-input"), Text("Select a country...⌄"))),
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

func ToggleRow(title string, description string, checked bool) Node {
	return Div(Class("toggle-row"),
		Div(Class("copy"), Strong(Text(title)), If(description != "", P(Class("muted"), Text(description)))),
		Toggle(checked),
	)
}

func Toggle(checked bool) Node {
	classes := "toggle"
	if checked {
		classes += " is-on"
	}
	return Button(Type("button"), Class(classes), Span())
}

func NestedFields(nodes ...Node) Node {
	return Div(Class("nested-fields"), Group(nodes))
}

func NestedFieldsDisabled(nodes ...Node) Node {
	return Div(Class("nested-fields disabled-nest"), Group(nodes))
}

func UnitInput(placeholder string, unit string) Node {
	return Div(Class("unit-input"), Input(Type("text"), Placeholder(placeholder), Class("text-input")), Span(Text(unit)))
}
