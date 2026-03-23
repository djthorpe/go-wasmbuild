package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct{ base }

type tableHeader struct{ base }

type tableRow struct{ base }

var _ mvc.View = (*table)(nil)
var _ mvc.View = (*tableHeader)(nil)
var _ mvc.View = (*tableRow)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const templateTable = `
	<cds-table size="sm">
		<cds-table-head data-slot="header"></cds-table-head>
		<cds-table-body data-slot="body"></cds-table-body>
	</cds-table>
`

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTable, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(table), element, setView)
	})
	mvc.RegisterView(ViewTableHeader, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableHeader), element, setView)
	})
	mvc.RegisterView(ViewTableRow, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableRow), element, setView)
	})
}

// Table returns a minimal Carbon-styled data table.
func Table(args ...any) *table {
	return mvc.NewView(new(table), ViewTable, templateTable, setView, args).(*table)
}

// TableHeader returns a table header row.
func TableHeader(args ...any) *tableHeader {
	return mvc.NewView(new(tableHeader), ViewTableHeader, "cds-table-header-row", setView, args).(*tableHeader)
}

// TableRow returns a table body row.
func TableRow(args ...any) *tableRow {
	return mvc.NewView(new(tableRow), ViewTableRow, "cds-table-row", setView, args).(*tableRow)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (h *tableHeader) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("cds-table-header-cell", arg)
		}
	}
	return h.View.Content(args...)
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
