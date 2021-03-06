package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/marwan-at-work/marwanio/frontend/js-wrappers/marked"
	"github.com/marwan-at-work/marwanio/frontend/stores/blogposts"
	"github.com/marwan-at-work/vecty-router"
)

// PostView represents a post
type PostView struct {
	vecty.Core
}

// Render returns every title
func (pv *PostView) Render() *vecty.HTML {
	// TODO: safely check with ok var
	id := router.GetNamedVar(pv)["id"]
	p, err := blogposts.GetByID(id)
	if err == blogposts.ErrNotFound {
		return pv.renderErr()
	}

	output := marked.Marked(p.Markdown)

	return elem.Div(
		vecty.Markup(
			prop.Class("blogpost-container"),
			vecty.UnsafeHTML(output),
		),
	)
}

func (pv *PostView) renderErr() *vecty.HTML {
	return elem.Div(
		vecty.Text("not found"),
	)
}
