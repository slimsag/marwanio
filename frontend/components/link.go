package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Link renders a footer link with a dispatcher for onClick
type Link struct {
	vecty.Core
	Name string
	Link string
}

// Render renders Link
func (fl *Link) Render() *vecty.HTML {
	if fl.Link != "" {
		return elem.Anchor(
			vecty.Markup(
				prop.Class("footer-link"),
				prop.Href(fl.Link),
			),
			vecty.Text(fl.Name),
		)
	}

	return elem.Anchor(
		vecty.Markup(prop.Class("link")),
		vecty.Text(fl.Name),
	)
}
