package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct {
	BootstrapView
}

type tablerow struct {
	BootstrapView
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
	t := new(table)
	t.BootstrapView.View = mvc.NewViewExEx(t, ViewTable, templateTable, args...)
	return t
}

func Row(args ...any) *tablerow {
	t := new(tablerow)
	t.BootstrapView.View = mvc.NewView(t, ViewTableRow, "TR", args...)
	return t
}

func newTableFromElement(element dom.Element) mvc.View {
	if element.TagName() != "TABLE" {
		return nil
	}
	t := new(table)
	t.BootstrapView.View = mvc.NewViewWithElement(t, element)
	return t
}

func newTableRowFromElement(element dom.Element) mvc.View {
	if element.TagName() != "TR" {
		return nil
	}
	t := new(tablerow)
	t.BootstrapView.View = mvc.NewViewWithElement(t, element)
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (table *table) Self() mvc.View {
	return table
}

func (tablerow *tablerow) Self() mvc.View {
	return tablerow
}

func (table *table) Header(args ...any) *table {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TH", arg)
		default:
			args[i] = arg
		}
	}
	table.View.ReplaceSlot("header", mvc.HTML("TR", args...))
	return table
}

func (table *table) Footer(args ...any) *table {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TH", arg)
		default:
			args[i] = arg
		}
	}
	table.View.ReplaceSlot("footer", mvc.HTML("TR", args...))
	return table
}

func (table *table) Content(args ...any) *table {
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
	table.View.ReplaceSlot("body", mvc.HTML("TBODY", args...))
	return table
}

func (tablerow *tablerow) Content(args ...any) *tablerow {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, mvc.View:
			args[i] = mvc.HTML("TD", arg)
		default:
			args[i] = arg
		}
	}
	tablerow.ReplaceSlot("body", wrapChildren(args...))
	return tablerow
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
