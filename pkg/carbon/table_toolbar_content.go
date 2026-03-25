package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type tableToolbarContent struct{ base }

var _ mvc.View = (*tableToolbarContent)(nil)

func init() {
	mvc.RegisterView(ViewTableToolbarContent, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableToolbarContent), element, setView)
	})
}

// TableToolbarContent returns a <cds-table-toolbar-content> wrapper.
func TableToolbarContent(args ...any) *tableToolbarContent {
	return mvc.NewView(new(tableToolbarContent), ViewTableToolbarContent, "cds-table-toolbar-content", setView, args...).(*tableToolbarContent)
}
