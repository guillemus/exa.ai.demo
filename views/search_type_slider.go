package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	.slider-card {
		--slider-position: 50%;
		border: var(--border-size-1) solid var(--line);
		border-radius: var(--radius-3);
		padding: var(--size-6) var(--size-7) var(--size-5);
	}
	.search-type-slider { user-select: none; touch-action: none; }
	.slider-track {
		position: relative;
		height: var(--size-6);
		margin: 0 0 var(--size-4);
		cursor: pointer;
	}
	.slider-line, .slider-fill {
		position: absolute;
		left: 0;
		top: var(--size-2);
		height: var(--size-2);
		border-radius: var(--radius-round);
	}
	.slider-line { right: 0; background: #cfd2d8; overflow: hidden; }
	.slider-fill { width: var(--slider-position); max-width: 100%; background: linear-gradient(90deg, #9dc0ff, #356bf3); }
	.dot {
		position: absolute;
		top: var(--size-2);
		width: var(--size-2);
		height: var(--size-2);
		border-radius: var(--radius-round);
		background: white;
		transform: translate(-50%, 0);
		pointer-events: none;
	}
	.dot-1 { left: 0; } .dot-2 { left: 25%; } .dot-3 { left: 50%; } .dot-4 { left: 100%; }
	.slider-thumb {
		position: absolute;
		left: clamp(var(--size-3), var(--slider-position), calc(100% - var(--size-3)));
		top: var(--border-size-2);
		width: var(--size-5);
		height: var(--size-5);
		border: var(--border-size-2) solid #5c84ee;
		border-radius: var(--radius-round);
		background: #cddcff;
		box-shadow: var(--shadow-2);
		transform: translateX(-50%);
		pointer-events: none;
	}
	.slider-labels { position: relative; height: var(--size-8); }
	.slider-option {
		position: absolute;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: var(--border-size-2);
		border: 0;
		background: transparent;
		padding: 0;
		text-align: center;
		color: var(--text);
		transform: translateX(-50%);
	}
	.slider-option:nth-child(1) { left: 0; align-items: flex-start; text-align: left; transform: none; }
	.slider-option:nth-child(2) { left: 25%; }
	.slider-option:nth-child(3) { left: 50%; }
	.slider-option:nth-child(4) { left: 100%; text-align: right; align-items: flex-end; transform: translateX(-100%); }
	.slider-option span { color: #6d7583; font-size: var(--font-size-0); }
	.slider-option.is-active strong { color: #0f172a; }
	.slider-option.is-active span { color: #4b5563; }
	.deep-model-row {
		display: grid;
		grid-template-columns: 1fr auto;
		align-items: center;
		gap: var(--size-3);
		margin: var(--size-5) calc(var(--size-7) * -1) calc(var(--size-5) * -1);
		padding: var(--size-4) var(--size-7);
		border-top: var(--border-size-1) solid var(--line);
	}
	.deep-model-buttons { display: flex; gap: var(--size-3); }
	.deep-model-button {
		border: var(--border-size-1) solid var(--line);
		border-radius: var(--radius-2);
		background: white;
		padding: var(--size-2) var(--size-4);
		color: #555;
		font-weight: 600;
	}
	.deep-model-button.is-active {
		border-color: var(--blue);
		background: #eef4ff;
		color: var(--blue);
	}
`)

func SearchTypeSlider(form SearchForm) Node {
	return Div(
		Class("slider-card search-type-slider"),
		Attr("style", "--slider-position: "+searchTypePosition(form.SearchType)),
		Data("search-type-slider", ""),
		Data("init", "initSearchTypeSlider(el)"),
		Data("effect", "setSearchTypeValue(el, $searchType, true)"),
		Input(Type("hidden"), Value(form.SearchType), Data("bind:search-type", ""), Data("search-type-value", "")),
		Div(Class("slider-track"), Data("search-type-track", ""),
			Span(Class("slider-line")),
			Span(Class("slider-fill")),
			Span(Class("dot dot-1")),
			Span(Class("dot dot-2")),
			Span(Class("dot dot-3")),
			Span(Class("dot dot-4")),
			Span(Class("slider-thumb")),
		),
		Div(Class("slider-labels"),
			SliderOption("instant", "Instant", "200ms", form.SearchType),
			SliderOption("fast", "Fast", "450ms", form.SearchType),
			SliderOption("auto", "Auto", "1s (recommended)", form.SearchType),
			SliderOption("deep", "Deep", "4s-18s", form.SearchType),
		),
		DeepModelControls(form),
	)
}

func searchTypePosition(searchType string) string {
	switch searchType {
	case "instant":
		return "0%"
	case "fast":
		return "25%"
	case "auto", "":
		return "50%"
	case "deep":
		return "100%"
	}
	return "50%"
}

func DeepModelControls(form SearchForm) Node {
	style := "display: none"
	if form.SearchType == searchTypeDeep {
		style = ""
	}
	return Div(Class("deep-model-row"), Data("show", "$searchType == 'deep'"), Attr("style", style),
		Strong(Text("Deep model")),
		Div(Class("deep-model-buttons"),
			DeepModelButton("deep-lite", form.DeepModel),
			DeepModelButton("deep", form.DeepModel),
			DeepModelButton("deep-reasoning", form.DeepModel),
		),
	)
}

func DeepModelButton(value string, current string) Node {
	return Button(
		Type("button"),
		Class(activeClass("deep-model-button", current == value)),
		Data("on:click", "$deepModel = '"+value+"'"),
		Data("class:is-active", "$deepModel == '"+value+"'"),
		Text(value),
	)
}

func SliderOption(value string, title string, subtitle string, current string) Node {
	return Button(
		Type("button"),
		Class(activeClass("slider-option", current == value)),
		Data("search-type-option", value),
		Data("on:click", "$searchType = '"+value+"'"),
		Data("class:is-active", "$searchType == '"+value+"'"),
		Strong(Text(title)),
		Span(Text(subtitle)),
	)
}

func activeClass(base string, active bool) string {
	if active {
		return base + " is-active"
	}
	return base
}
