package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
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
	ViewTable    = "mvc-bs-table"
	ViewTableRow = "mvc-bs-tablerow"
)

const (
	templateTable = `
		<table class="table">
			<thead><tr data-slot="header"></tr></thead>
			<tbody data-slot="body"></tbody>
			<tfoot><tr data-slot="footer"></tr></tfoot>
		</table>
	`
)

func init() {
	mvc.RegisterView(ViewTable, newTableFromElement)
	mvc.RegisterView(ViewTableRow, newTableRowFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Table(args ...any) *table {
	return mvc.NewViewExEx(new(table), ViewTable, templateTable, args...).(*table)
}

func Row(args ...any) *tablerow {
	return mvc.NewView(new(tablerow), ViewTableRow, "TR", args...).(*tablerow)
}

func newTableFromElement(element dom.Element) mvc.View {
	if element.TagName() != "TABLE" {
		return nil
	}
	return mvc.NewViewWithElement(new(table), element)
}

func newTableRowFromElement(element dom.Element) mvc.View {
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

func (table *table) Header(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TH", arg)
		default:
			args[i] = arg
		}
	}
	return table.View.ReplaceSlot("header", mvc.HTML("TR", args...))
}

func (table *table) Footer(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TH", arg)
		default:
			args[i] = arg
		}
	}
	return table.View.ReplaceSlot("footer", mvc.HTML("TR", args...))
}

func (table *table) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg := arg.(type) {
		case mvc.View:
			if _, ok := arg.(*tablerow); ok {
				args[i] = arg
			} else {
				panic("table.Content: invalid view type")
			}
		default:
			args[i] = arg
		}
	}
	return table.View.ReplaceSlot("body", mvc.HTML("TBODY", args...))
}

func (tablerow *tablerow) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TD", arg)
		default:
			args[i] = arg
		}
	}
	return tablerow.View.Content(args...)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

/*

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
*/
