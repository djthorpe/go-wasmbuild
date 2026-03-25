package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type tableRow struct{ base }

var _ mvc.View = (*tableRow)(nil)

func init() {
	mvc.RegisterView(ViewTableRow, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableRow), element, setView)
	})
}

// TableRow returns a table body row.
func TableRow(args ...any) *tableRow {
	return mvc.NewView(new(tableRow), ViewTableRow, "cds-table-row", setView, args).(*tableRow)
}

func (r *tableRow) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("cds-table-cell", arg)
		}
	}
	return r.View.Content(args...)
}
