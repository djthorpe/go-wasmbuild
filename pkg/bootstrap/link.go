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
	BootstrapView
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

func Link(href string, args ...any) mvc.View {
	l := new(link)
	l.BootstrapView.View = mvc.NewView(l, ViewLink, "A", mvc.WithAttr("href", href), args)
	return l
}

func newLinkFromElement(element Element) mvc.View {
	if element.TagName() != "A" {
		panic(fmt.Sprintf("newLinkFromElement: invalid tag name %q", element.TagName()))
	}
	l := new(link)
	l.BootstrapView.View = mvc.NewViewWithElement(l, element)
	return l
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (link *link) Self() mvc.View {
	return link
}
