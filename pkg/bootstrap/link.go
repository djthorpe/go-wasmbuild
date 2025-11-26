package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
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

func init() {
	mvc.RegisterView(ViewLink, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(link), element, func(self, child mvc.View) {
			self.(*link).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Link(href string, args ...any) mvc.View {
	return mvc.NewView(new(link), ViewLink, "A", func(self, child mvc.View) {
		self.(*link).View = child
	}, mvc.WithAttr("href", href), args)
}

func IconLink(href string, args ...any) mvc.View {
	return mvc.NewView(new(link), ViewLink, "A", func(self, child mvc.View) {
		self.(*link).View = child
	}, mvc.WithClass("icon-link"), mvc.WithAttr("href", href), args)
}
