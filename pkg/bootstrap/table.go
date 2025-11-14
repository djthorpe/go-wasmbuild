package bootstrap

import (
	// Packages

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct {
	mvc.View
}

type tablerow struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewTable     = "mvc-bs-table"
	ViewTableRow  = "mvc-bs-tablerow"
	ViewTableFoot = "mvc-bs-tablefoot"
	ViewTableHead = "mvc-bs-tablehead"
)

func init() {
	mvc.RegisterView(ViewTable, newTableFromElement)
	mvc.RegisterView(ViewTableRow, newTableRowFromElement)
	mvc.RegisterView(ViewTableHead, newTableRowFromElement)
	mvc.RegisterView(ViewTableFoot, newTableRowFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

/*
func Table(args ...any) *table {
	// Return the table
	return mvc.NewViewEx(
		new(table), ViewTable, "TABLE",
		mvc.HTML("THEAD"), mvc.HTML("TBODY"), mvc.HTML("TFOOT"), nil,
		mvc.WithClass("table"), args,
	).(*table)
}
*/

func TableRow(args ...any) *tablerow {
	return tableRow(ViewTableRow, args...)
}

func tableRow(name string, args ...any) *tablerow {
	// Return the table row
	return mvc.NewView(
		new(tablerow), name, "TR", args,
	).(*tablerow)
}

func newTableFromElement(element Element) mvc.View {
	if element.TagName() != "TABLE" {
		return nil
	}
	return mvc.NewViewWithElement(new(table), element)
}

func newTableRowFromElement(element Element) mvc.View {
	if element.TagName() != "TR" {
		return nil
	}
	return mvc.NewViewWithElement(new(tablerow), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (table *table) SetView(view mvc.View) {
	table.View = view
}

func (tablerow *tablerow) SetView(view mvc.View) {
	tablerow.View = view
}

func (tablerow *tablerow) Append(cells ...any) mvc.View {
	// Write calls in td elements
	for _, cell := range cells {
		node := mvc.NodeFromAny(cell)
		switch tablerow.Name() {
		case ViewTableHead, ViewTableFoot:
			th := mvc.HTML("th", mvc.WithAttr("scope", "col"))
			th.AppendChild(node)
			tablerow.View.Append(th)
		default:
			td := mvc.HTML("td")
			td.AppendChild(node)
			tablerow.View.Append(td)
		}
	}
	return tablerow
}

func (table *table) Header(children ...any) mvc.View {
	// Create a header row
	header := tableRow(ViewTableHead, children...)
	table.View.Header(header)
	return table
}

func (table *table) Footer(children ...any) mvc.View {
	// Create a header row
	footer := tableRow(ViewTableFoot, children...)
	table.View.Footer(footer)
	return table
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithStripedRows() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return ErrInternalAppError.Withf("WithStripedRows: invalid view %q", o.Name())
		}
		if err := mvc.WithoutClass("table-striped-columns")(o); err != nil {
			return err
		}
		return mvc.WithClass("table-striped")(o)
	}
}

func WithStripedColumns() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return ErrInternalAppError.Withf("WithStripedColumns: invalid view %q", o.Name())
		}
		if err := mvc.WithoutClass("table-striped")(o); err != nil {
			return err
		}
		return mvc.WithClass("table-striped-columns")(o)
	}
}

func WithRowHover() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return ErrInternalAppError.Withf("WithRowHover: invalid view %q", o.Name())
		}
		return mvc.WithClass("table-hover")(o)
	}
}

func WithoutRowHover() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return ErrInternalAppError.Withf("WithRowHover: invalid view %q", o.Name())
		}
		return mvc.WithoutClass("table-hover")(o)
	}
}
