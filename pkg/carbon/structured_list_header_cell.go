package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type structuredListHeaderCell struct{ base }

var _ mvc.View = (*structuredListHeaderCell)(nil)

func init() {
	mvc.RegisterView(ViewStructuredListTH, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(structuredListHeaderCell), element, setView)
	})
}

// StructuredListHeaderCell returns a structured list header cell.
func StructuredListHeaderCell(args ...any) *structuredListHeaderCell {
	return mvc.NewView(new(structuredListHeaderCell), ViewStructuredListTH, "cds-structured-list-header-cell", setView, args).(*structuredListHeaderCell)
}
