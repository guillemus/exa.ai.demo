package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
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
	.deep-model-row {
		display: grid;
		grid-template-columns: 1fr auto;
		align-items: center;
		gap: 16px;
		margin: 24px -32px -24px;
		padding: 20px 32px;
		border-top: 1px solid var(--line);
	}
	.deep-model-buttons { display: flex; gap: 16px; }
	.deep-model-button {
		border: 1px solid var(--line);
		border-radius: var(--radius-2);
		background: white;
		padding: 10px 20px;
		color: #555;
		font-weight: 600;
	}
	.deep-model-button.is-active {
		border-color: var(--blue);
		background: #eef4ff;
		color: var(--blue);
	}
`)

func SearchTypeSlider() Node {
	return Div(Class("slider-card search-type-slider"), Data("search-type-slider", ""), Data("init", "initSearchTypeSlider(el)"),
		Input(Type("hidden"), Data("bind:search-type", ""), Value("auto"), Data("search-type-value", "")),
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
			SliderOption("instant", "Instant", "200ms"),
			SliderOption("fast", "Fast", "450ms"),
			SliderOption("auto", "Auto", "1s (recommended)"),
			SliderOption("deep", "Deep", "4s-18s"),
		),
		DeepModelControls(),
	)
}

func DeepModelControls() Node {
	return Div(Class("deep-model-row"), Data("show", "$searchType == 'deep'"), Attr("style", "display: none"),
		Strong(Text("Deep model")),
		Div(Class("deep-model-buttons"),
			DeepModelButton("deep-lite"),
			DeepModelButton("deep"),
			DeepModelButton("deep-reasoning"),
		),
	)
}

func DeepModelButton(value string) Node {
	return Button(
		Type("button"),
		Class("deep-model-button"),
		Data("on:click", "$deepModel = '"+value+"'"),
		Data("class:is-active", "$deepModel == '"+value+"'"),
		Text(value),
	)
}

func SliderOption(value string, title string, subtitle string) Node {
	return Button(
		Type("button"),
		Class("slider-option"),
		Data("search-type-option", value),
		Data("on:click", "$searchType = '"+value+"'"),
		Data("class:is-active", "$searchType == '"+value+"'"),
		Strong(Text(title)),
		Span(Text(subtitle)),
	)
}
