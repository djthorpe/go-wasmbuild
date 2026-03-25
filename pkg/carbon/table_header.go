package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type tableHeader struct{ base }

var _ mvc.View = (*tableHeader)(nil)

func init() {
	mvc.RegisterView(ViewTableHeader, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableHeader), element, setView)
	})
}

// TableHeader returns a table header row.
func TableHeader(args ...any) *tableHeader {
	return mvc.NewView(new(tableHeader), ViewTableHeader, "cds-table-header-row", setView, args).(*tableHeader)
}

func (h *tableHeader) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("cds-table-header-cell", arg)
		}
	}
	return h.View.Content(args...)
}
