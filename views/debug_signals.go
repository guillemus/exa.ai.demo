package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	.debug-signals {
		position: fixed;
		right: var(--size-3);
		top: var(--size-3);
		z-index: 1000;
		width: var(--size-14);
		border: var(--border-size-1) solid var(--gray-7);
		border-radius: var(--radius-2);
		background: white;
		color: var(--gray-12);
		font-family: var(--font-code);
		font-size: var(--font-size-0);
		line-height: var(--font-lineheight-3);
	}
	.debug-signals summary {
		padding: var(--size-2);
		cursor: pointer;
		font-weight: 700;
	}
	.debug-signals pre {
		max-height: 50vh;
		overflow: auto;
		margin: 0;
		padding: var(--size-2);
		border-top: var(--border-size-1) solid var(--gray-3);
		white-space: pre-wrap;
	}
`)

func DebugSignals() Node {
	return Details(Class("debug-signals"),
		Summary(Text("Show signals")),
		Pre(Data("json-signals", "")),
	)
}
