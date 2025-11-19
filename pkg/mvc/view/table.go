package view

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	reflect "github.com/djthorpe/go-wasmbuild/pkg/mvc/reflect"

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
	proto *reflect.Proto
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

// Create a Table with a prototype row
func Table(proto any, opts ...mvc.Opt) TableView {
	t := new(table)
	args := []any{mvc.HTML("thead"), mvc.HTML("tbody"), mvc.HTML("tfoot")}
	for _, o := range opts {
		args = append(args, o)
	}
	t.TableView = mvc.NewView(t, ViewTable, "table", args...).(TableView)
	if proto != nil {
		if proto := reflect.NewProto(proto); proto == nil {
			panic("Table: proto must be a struct")
		} else {
			t.proto = proto
		}
	}
	return t
}

// Create a TableRowEx (children are td or th)
func TableRowEx(name string, children ...any) mvc.View {
	self := new(tablerow)
	self.View = mvc.NewView(self, name, "tr")
	return self.Append(children...)
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
	t := new(table)
	t.TableView = mvc.NewViewWithElement(t, element).(TableView)
	return t
}

// Create a TableRow from an existing element
func newTableRowFromElement(element Element) mvc.View {
	if element.TagName() != "TR" {
		return nil
	}
	t := new(tablerow)
	t.View = mvc.NewViewWithElement(t, element)
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
			if table.proto != nil {
				// TODO: Ensure the child is the same type
				// TODO: Append this data as a new row based on the prototype
			} else {
				panic(fmt.Sprintf("table.Append: invalid child type %T", child))
			}
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
