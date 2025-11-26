package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// containers are elements to wrap any content
type container struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewContainer, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(container), element, func(self, child mvc.View) {
			self.(*container).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Container(args ...any) mvc.View {
	return mvc.NewView(new(container), ViewContainer, "DIV", func(self, child mvc.View) {
		self.(*container).View = child
	}, mvc.WithClass("container"), args)
}

func FluidContainer(args ...any) mvc.View {
	return mvc.NewView(new(container), ViewContainer, "DIV", func(self, child mvc.View) {
		self.(*container).View = child
	}, mvc.WithClass("container-fluid"), args)
}
