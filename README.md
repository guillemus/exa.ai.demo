# Exa Search API Hypermedia Playground Demo

Quick and dirty demo of rebuilding the Exa Search API playground/search page with a hypermedia-first approach.

This is mostly proof-of-concept code, not production architecture. Main goal is to show how much interactivity you can get with server-rendered HTML, SSE patches, and barely any client-side JavaScript.

## Current caveat

The custom JavaScript is intentionally small, but brittle. It works for this demo, yet it is not the maintainable long-term shape.

A production version should move rich client-only behavior into Web Components and keep local client state inside those components.
