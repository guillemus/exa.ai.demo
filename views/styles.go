package views

import (
	"strings"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Styles struct {
	styleString strings.Builder
}

var styles Styles

func (x *Styles) Style(css string) uint8 {
	x.styleString.WriteString(css)
	return 0
}

func (x *Styles) Node() Node {
	return StyleEl(Raw(x.styleString.String()))
}
