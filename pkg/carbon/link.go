package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type link struct {
	mvc.View
}

var _ mvc.View = (*link)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewLink, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(link), element, func(self, child mvc.View) {
			self.(*link).View = child
		})
	})
}

// Link returns an anchor element styled with the Carbon cds--link class.
func Link(href string, args ...any) mvc.View {
	return mvc.NewView(new(link), ViewLink, "A", func(self, child mvc.View) {
		self.(*link).View = child
	}, mvc.WithClass("cds--link"), mvc.WithAttr("href", href), args)
}

// InlineLink returns a Carbon link styled for use within a body of text.
func InlineLink(href string, args ...any) mvc.View {
	return mvc.NewView(new(link), ViewLink, "A", func(self, child mvc.View) {
		self.(*link).View = child
	}, mvc.WithClass("cds--link", "cds--link--inline"), mvc.WithAttr("href", href), args)
}
