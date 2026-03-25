package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type container struct{ base }

var _ mvc.View = (*container)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewSection, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(container), element, setView)
	})
}
