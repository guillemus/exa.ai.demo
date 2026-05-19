package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var _ = styles.Style(`
	.debug-signals {
		position: fixed;
		right: 12px;
		top: 12px;
		z-index: 1000;
		width: 320px;
		border: 1px solid var(--gray-7);
		border-radius: var(--radius-2);
		background: white;
		color: var(--gray-12);
		font-family: var(--font-code);
		font-size: 12px;
		line-height: 1.45;
	}
	.debug-signals summary {
		padding: 8px 10px;
		cursor: pointer;
		font-weight: 700;
	}
	.debug-signals pre {
		max-height: 420px;
		overflow: auto;
		margin: 0;
		padding: 10px;
		border-top: 1px solid var(--gray-3);
		white-space: pre-wrap;
	}
`)

func DebugSignals() Node {
	return Details(Class("debug-signals"),
		Summary(Text("Show signals")),
		Pre(Data("json-signals", "")),
	)
}
