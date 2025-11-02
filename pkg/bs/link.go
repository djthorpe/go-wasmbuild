package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type link struct {
	View
}

var _ View = (*link)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewLink = "mvc-bs-link"
)

func init() {
	RegisterView(ViewLink, newLinkFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Link(href string, opt ...Opt) View {
	opt = append([]Opt{WithAttr("href", href)}, opt...)
	return NewView(new(link), ViewLink, "A", opt...)
}

func newLinkFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "A" {
		panic(fmt.Sprintf("newLinkFromElement: invalid tag name %q", tagName))
	}
	return NewViewWithElement(new(link), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (link *link) SetView(view View) {
	link.View = view
}
