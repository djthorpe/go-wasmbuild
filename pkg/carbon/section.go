package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type section struct{ base }

var _ mvc.View = (*section)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewSection, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(section), element, setView)
	})
}

// Section returns a <section class="cds--content"> layout container.
func Section(args ...any) *section {
	return mvc.NewView(new(section), ViewSection, "SECTION", setView, mvc.WithClass("cds--content"), args).(*section)
}
