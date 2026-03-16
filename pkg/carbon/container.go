package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type section struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewSection, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(section), element, func(self, child mvc.View) {
			self.(*section).View = child
		})
	})
}

// Section returns a plain <div> mvc.View that can hold arbitrary children.
// Use it as a page-level container when composing router pages.
func Section(args ...any) mvc.View {
	return mvc.NewView(new(section), ViewSection, "DIV", func(self, child mvc.View) {
		self.(*section).View = child
	}, args)
}
