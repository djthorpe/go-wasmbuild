package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type icon struct {
	mvc.View
}

var _ mvc.View = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewIcon, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(icon), element, func(self, child mvc.View) {
			self.(*icon).View = child
		})
	})
}

// Icon renders a Carbon icon using the <cds-icon> custom element defined in
// components.js. The name must be a valid @carbon/icons icon name
// (e.g. "add", "close", "warning--filled", "arrow--right").
//
// Default size is 16 px. Pass mvc.WithAttr("size", "24") to change it.
// Color is inherited from the CSS `color` property (fill: currentColor).
func Icon(name string, args ...any) mvc.View {
	return mvc.NewView(new(icon), ViewIcon, "cds-icon", func(self, child mvc.View) {
		self.(*icon).View = child
	}, mvc.WithAttr("name", name), args)
}
