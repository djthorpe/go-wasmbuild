package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type structuredListRow struct{ base }

var _ mvc.View = (*structuredListRow)(nil)
var _ mvc.ActiveState = (*structuredListRow)(nil)
var _ mvc.ValueState = (*structuredListRow)(nil)

func init() {
	mvc.RegisterView(ViewStructuredListRow, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(structuredListRow), element, setView)
	})
}

// StructuredListRow returns a structured list body row.
func StructuredListRow(args ...any) *structuredListRow {
	return mvc.NewView(new(structuredListRow), ViewStructuredListRow, "cds-structured-list-row", setView, args).(*structuredListRow)
}

func (r *structuredListRow) Content(args ...any) mvc.View {
	for i, arg := range args {
		if !isStructuredListCell(arg, "CDS-STRUCTURED-LIST-CELL") {
			switch arg.(type) {
			case string, dom.Element, mvc.View:
				args[i] = StructuredListCell(arg)
			}
		}
	}
	return r.View.Content(args...)
}

func (r *structuredListRow) Active() bool {
	return boolProperty(r.Root(), "selected")
}

func (r *structuredListRow) SetActive(active bool) mvc.View {
	setBoolProperty(r.Root(), "selected", active)
	return r
}

func (r *structuredListRow) Value() string {
	return r.Root().GetAttribute("selection-value")
}

func (r *structuredListRow) SetValue(value string) mvc.View {
	if value == "" {
		r.Root().RemoveAttribute("selection-value")
	} else {
		r.Root().SetAttribute("selection-value", value)
	}
	return r
}
