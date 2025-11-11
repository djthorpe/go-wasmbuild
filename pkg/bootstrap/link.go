package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type link struct {
	mvc.View
}

var _ mvc.View = (*link)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewLink = "mvc-bs-link"
)

func init() {
	mvc.RegisterView(ViewLink, newLinkFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Link(href string, opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(link), ViewLink, "A", append([]mvc.Opt{mvc.WithAttr("href", href)}, opt...)...)
}

func newLinkFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "A" {
		panic(fmt.Sprintf("newLinkFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(link), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (link *link) SetView(view mvc.View) {
	link.View = view
}
