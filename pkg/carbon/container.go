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

// Section returns a <section class="cds--content"> layout container.
func Section(args ...any) *container {
	return mvc.NewView(new(container), ViewSection, "SECTION", setView, mvc.WithClass("cds--content"), args).(*container)
}

// Page returns a plain <div> view with no cds--content padding, suitable
// for use as a per-page wrapper inside the main content Section.
func Page(args ...any) *container {
	return mvc.NewView(new(container), ViewSection, "DIV", setView, args).(*container)
}
