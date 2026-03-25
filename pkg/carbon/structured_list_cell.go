package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type structuredListCell struct{ base }

var _ mvc.View = (*structuredListCell)(nil)

func init() {
	mvc.RegisterView(ViewStructuredListCell, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(structuredListCell), element, setView)
	})
}

// StructuredListCell returns a structured list body cell.
func StructuredListCell(args ...any) *structuredListCell {
	return mvc.NewView(new(structuredListCell), ViewStructuredListCell, "cds-structured-list-cell", setView, args).(*structuredListCell)
}
