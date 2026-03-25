package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Section returns a <section class="cds--content"> layout container.
func Section(args ...any) *container {
	return mvc.NewView(new(container), ViewSection, "SECTION", setView, mvc.WithClass("cds--content"), args).(*container)
}
