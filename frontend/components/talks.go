package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Talks renders index of talks I've given
type Talks struct {
	vecty.Core
}

// Render renders Talks component
func (tlks *Talks) Render() *vecty.HTML {
	return elem.Div(
		vecty.Markup(prop.ID("blog-container")),
		elem.Heading1(
			vecty.Markup(prop.Class("blog-header")),
			vecty.Text("Talks"),
		),
		elem.Paragraph(
			vecty.Markup(prop.Class("post-title")),
			elem.Anchor(
				vecty.Markup(
					prop.Class("post-title-text"),
					prop.Href("/public/bowery-golang.pdf"),
				),
				vecty.Text("Build Your Own Go CI Server"),
			),
			elem.Span(
				vecty.Markup(
					prop.Class("post-title-date"),
					vecty.UnsafeHTML(
						"2017, Aug 3 - <a href=\"https://www.meetup.com/Bowery-Go/events/241363507\">Bowery Golang</a>",
					),
				),
			),
		),
	)
}
