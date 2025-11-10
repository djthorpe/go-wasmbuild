package views

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

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

type tablecell struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewTable     = "mvc-table"
	ViewTableRow  = "mvc-table-row"
	ViewTableCell = "mvc-table-cell"
)

func init() {
	mvc.RegisterView(ViewTable, newTableFromElement)
	mvc.RegisterView(ViewTableRow, newTableRowFromElement)
	mvc.RegisterView(ViewTableCell, newTableCellFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a Table
func Table(opts ...mvc.Opt) mvc.View {
	self := new(table)
	return mvc.NewView(self, ViewTable, "table", opts...)
}

// Create a TableRow (children are table cells)
func TableRow(children ...any) mvc.View {
	self := new(tablerow)
	return mvc.NewView(self, ViewTableRow, "tr").Append(children...)
}

// Create a TableCell
func TableCell(opts ...mvc.Opt) mvc.View {
	self := new(tablecell)
	return mvc.NewView(self, ViewTableCell, "td", opts...)
}

// Create a Table from an existing element
func newTableFromElement(element Element) mvc.View {
	if element.TagName() != "TABLE" {
		return nil
	}
	return mvc.NewViewWithElement(new(table), element)
}

// Create a TableRow from an existing element
func newTableRowFromElement(element Element) mvc.View {
	if element.TagName() != "TR" {
		return nil
	}
	return mvc.NewViewWithElement(new(tablerow), element)
}

// Create a TableCell from an existing element
func newTableCellFromElement(element Element) mvc.View {
	if element.TagName() != "TD" {
		return nil
	}
	return mvc.NewViewWithElement(new(tablecell), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (table *table) SetView(view mvc.View) {
	table.View = view
}

func (tablerow *tablerow) SetView(view mvc.View) {
	tablerow.View = view
}

func (tablecell *tablecell) SetView(view mvc.View) {
	tablecell.View = view
}

// Allow appending of mvc-table-row to mvc-table
func (table *table) Append(children ...any) mvc.View {
	for _, child := range children {
		switch child := child.(type) {
		case *tablerow:
			table.View.Append(child.Root())
		default:
			panic(fmt.Sprintf("table.Append: invalid child type %T", child))
		}
	}
	return table
}

// Wrap td elements in a table row
func (tablerow *tablerow) Append(children ...any) mvc.View {
	for _, child := range children {
		switch child := child.(type) {
		case *tablecell:
			tablerow.View.Append(child)
		case string:
			tablerow.View.Append(TableCell().Append(child))
		default:
			panic(fmt.Sprintf("tablerow.Append: invalid child type %T", child))
		}
	}
	return tablerow
}
