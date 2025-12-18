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
	templateTable = `
		<table class="table">
			<thead><tr data-slot="header"></tr></thead>
			<tbody data-slot="body"></tbody>
			<tfoot><tr data-slot="footer"></tr></tfoot>
		</table>
	`
)

func init() {
	mvc.RegisterView(ViewTable, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(table), element, setView)
	})
	mvc.RegisterView(ViewTableRow, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tablerow), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Table(args ...any) *table {
	return mvc.NewView(new(table), ViewTable, templateTable, setView, args).(*table)
}

func TableRow(args ...any) *tablerow {
	return mvc.NewView(new(tablerow), ViewTableRow, "TR", setView, args).(*tablerow)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

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

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithStripedRows() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return dom.ErrInternalAppError.Withf("WithStripedRows: invalid view %q", o.Name())
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
			return dom.ErrInternalAppError.Withf("WithStripedColumns: invalid view %q", o.Name())
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
			return dom.ErrInternalAppError.Withf("WithRowHover: invalid view %q", o.Name())
		}
		return mvc.WithClass("table-hover")(o)
	}
}

func WithoutRowHover() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewTable {
			return dom.ErrInternalAppError.Withf("WithRowHover: invalid view %q", o.Name())
		}
		return mvc.WithoutClass("table-hover")(o)
	}
}
