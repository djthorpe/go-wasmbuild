package bootstrap

import (
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type icon struct {
	mvc.View
}

var _ mvc.View = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewIcon, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(icon), element)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Icon(name string, args ...any) mvc.View {
	return mvc.NewView(new(icon), ViewIcon, "I", func(self, child mvc.View) {
		self.(*icon).View = child
	}, mvc.WithClass("bi-"+name), args)
}
