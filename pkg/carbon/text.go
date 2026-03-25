package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct{ base }

var _ mvc.View = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, setView)
	})
}
