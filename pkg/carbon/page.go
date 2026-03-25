package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Page returns a plain <div> view with no cds--content padding, suitable
// for use as a per-page wrapper inside the main content Section.
func Page(args ...any) *container {
	return mvc.NewView(new(container), ViewSection, "DIV", setView, args).(*container)
}
