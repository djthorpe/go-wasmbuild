package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct{ base }

var _ mvc.View = (*table)(nil)

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
}

// Table returns a minimal Carbon-styled data table.
func Table(args ...any) *table {
	return mvc.NewView(new(table), ViewTable, templateTable, setView, args).(*table)
}
