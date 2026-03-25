package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type structuredListHeader struct{ base }

var _ mvc.View = (*structuredListHeader)(nil)

func init() {
	mvc.RegisterView(ViewStructuredListHead, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(structuredListHeader), element, setView)
	})
}

// StructuredListHeader returns a structured list header row.
func StructuredListHeader(args ...any) *structuredListHeader {
	return mvc.NewView(new(structuredListHeader), ViewStructuredListHead, "cds-structured-list-header-row", setView, args).(*structuredListHeader)
}

func (h *structuredListHeader) Content(args ...any) mvc.View {
	for i, arg := range args {
		if !isStructuredListCell(arg, "CDS-STRUCTURED-LIST-HEADER-CELL") {
			switch arg.(type) {
			case string, dom.Element, mvc.View:
				args[i] = StructuredListHeaderCell(arg)
			}
		}
	}
	return h.View.Content(args...)
}
