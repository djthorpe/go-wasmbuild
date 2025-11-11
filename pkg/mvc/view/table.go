package view

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// INTERFACE

type TableView interface {
	mvc.ViewWithHeaderFooter
}

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct {
	TableView
}

type tablerow struct {
	mvc.View
}

var _ mvc.ViewWithHeaderFooter = (*table)(nil)
var _ mvc.View = (*tablerow)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewTable     = "mvc-table"
	ViewTableRow  = "mvc-tr"
	ViewTableHead = "mvc-thead"
	ViewTableFoot = "mvc-tfoot"
)

func init() {
	mvc.RegisterView(ViewTable, newTableFromElement)
	mvc.RegisterView(ViewTableRow, newTableRowFromElement)
	mvc.RegisterView(ViewTableHead, newTableRowFromElement)
	mvc.RegisterView(ViewTableFoot, newTableRowFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a Table
func Table(opts ...mvc.Opt) TableView {
	return mvc.NewViewEx(new(table), ViewTable, "table", mvc.HTML("thead"), mvc.HTML("tbody"), mvc.HTML("tfoot"), nil, opts...).(TableView)
}

// Create a TableRowEx (children are td or th)
func TableRowEx(name string, children ...any) mvc.View {
	self := new(tablerow)
	return mvc.NewView(self, name, "tr").Append(children...)
}

// Create a TableRow (children are td)
func TableRow(children ...any) mvc.View {
	return TableRowEx(ViewTableRow, children...)
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

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (table *table) SetView(view mvc.View) {
	table.TableView = view.(TableView)
}

func (tablerow *tablerow) SetView(view mvc.View) {
	tablerow.View = view
}

// Set header content
func (table *table) Header(children ...any) mvc.ViewWithHeaderFooter {
	table.TableView.Header(
		TableRowEx(ViewTableHead, children...),
	)
	return table
}

// Set footer content
func (table *table) Footer(children ...any) mvc.ViewWithHeaderFooter {
	table.TableView.Footer(
		TableRowEx(ViewTableFoot, children...),
	)
	return table
}

// Allow appending of mvc-table-row to mvc-table
func (table *table) Append(children ...any) mvc.View {
	for _, child := range children {
		switch child := child.(type) {
		case *tablerow:
			table.TableView.Append(child.Root())
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
		case string:
			switch tablerow.Name() {
			case ViewTableHead, ViewTableFoot:
				tablerow.View.Append(mvc.HTML("th", mvc.WithInnerText(child)))
			default:
				tablerow.View.Append(mvc.HTML("td", mvc.WithInnerText(child)))
			}
		default:
			panic(fmt.Sprintf("tablerow.Append: invalid child type %T", child))
		}
	}
	return tablerow
}
