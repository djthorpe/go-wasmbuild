package mvc

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct {
	View
}

type tableHeader struct {
	View
}

type tableRow struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewTable       = "mvc-table"
	ViewTableHeader = "mvc-table-header"
	ViewTableRow    = "mvc-table-row"

	templateTable = `
		<table>
			<thead data-slot="header"></thead>
			<tbody data-slot="body"></tbody>
		</table>
	`
)

func init() {
	RegisterView(ViewTable, func(element dom.Element) View {
		return NewViewWithElement(new(table), element, func(self, child View) {
			self.(*table).View = child
		})
	})
	RegisterView(ViewTableHeader, func(element dom.Element) View {
		return NewViewWithElement(new(tableHeader), element, func(self, child View) {
			self.(*tableHeader).View = child
		})
	})
	RegisterView(ViewTableRow, func(element dom.Element) View {
		return NewViewWithElement(new(tableRow), element, func(self, child View) {
			self.(*tableRow).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Table creates a basic HTML table with header and body slots.
func Table(args ...any) *table {
	return NewView(new(table), ViewTable, templateTable, func(self, child View) {
		self.(*table).View = child
	}, args).(*table)
}

// TableHeader creates a table header row.
func TableHeader(args ...any) *tableHeader {
	return NewView(new(tableHeader), ViewTableHeader, "TR", func(self, child View) {
		self.(*tableHeader).View = child
	}, args).(*tableHeader)
}

// TableRow creates a table body row.
func TableRow(args ...any) *tableRow {
	return NewView(new(tableRow), ViewTableRow, "TR", func(self, child View) {
		self.(*tableRow).View = child
	}, args).(*tableRow)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Header replaces the table header slot.
func (t *table) Header(header *tableHeader) View {
	if header == nil {
		return t.ReplaceSlotChildren("header")
	}
	return t.ReplaceSlotChildren("header", header)
}

// Rows replaces the table body rows.
func (t *table) Rows(rows ...*tableRow) View {
	args := make([]any, 0, len(rows))
	for _, row := range rows {
		if row != nil {
			args = append(args, row)
		}
	}
	return t.ReplaceSlotChildren(ContentSlot, args...)
}

// Content replaces the header cells, wrapping simple content in TH elements.
func (h *tableHeader) Content(args ...any) View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, View:
			args[i] = HTML("TH", arg)
		}
	}
	return h.View.Content(args...)
}

// Content replaces the row cells, wrapping simple content in TD elements.
func (r *tableRow) Content(args ...any) View {
	for i, arg := range args {
		switch arg.(type) {
		case string, dom.Element, View:
			args[i] = HTML("TD", arg)
		}
	}
	return r.View.Content(args...)
}
