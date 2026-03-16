package carbon

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

type tableRow struct {
	mvc.View
}

var _ mvc.View = (*table)(nil)
var _ mvc.View = (*tableRow)(nil)

// TableSize controls the row density of a Carbon data table.
type TableSize string

// TableSortDir controls the sort direction on a header cell.
type TableSortDir string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	TableXS TableSize = "xs" // extra compact
	TableSM TableSize = "sm" // 32px
	TableMD TableSize = "md" // 40px (default)
	TableLG TableSize = "lg" // 48px
	TableXL TableSize = "xl" // 64px
)

const (
	TableSortNone       TableSortDir = "none"
	TableSortAscending  TableSortDir = "ascending"
	TableSortDescending TableSortDir = "descending"
)

// Event name constants fired by the data-table web components.
const (
	// TableEventBeforeRowSelect fires on cds-table-row BEFORE its checked state
	// changes. event.detail.selected holds the NEW (upcoming) state.
	TableEventBeforeRowSelect = "cds-table-row-change-selection"

	// TableEventBeforeAllRowsSelect fires BEFORE the select-all state changes.
	TableEventBeforeAllRowsSelect = "cds-table-change-selection-all"

	// TableEventRowSelect fires on cds-table AFTER a row's selection changes.
	// event.detail.selectedRow is the toggled row; event.detail.selectedRows
	// holds all currently selected rows.
	TableEventRowSelect = "cds-table-row-selected"

	// TableEventAllRowsSelect fires on cds-table AFTER select-all changes.
	TableEventAllRowsSelect = "cds-table-row-all-selected"

	// TableEventBatchCancel fires when the Cancel button in cds-table-batch-actions
	// is clicked.
	TableEventBatchCancel = "cds-table-batch-actions-cancel-clicked"

	// TableEventSearch is fired by cds-table-toolbar-search on every keystroke.
	// cds-table handles this automatically to filter visible rows.
	TableEventSearch = "cds-search-input"

	// TableEventSorted is fired by cds-table after a column sort completes.
	TableEventSorted = "cds-table-sorted"
)

// templateTable provides the structural skeleton of a Carbon data table.
// data-slot="body" is the ContentSlot — TableRow children go here by default.
// data-slot="header" is populated via .Header().
const templateTable = `
	<cds-table>
		<cds-table-head><cds-table-header-row data-slot="header"></cds-table-header-row></cds-table-head>
		<cds-table-body data-slot="body"></cds-table-body>
	</cds-table>
`

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTable, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(table), element, func(self, child mvc.View) {
			self.(*table).View = child
		})
	})
	mvc.RegisterView(ViewTableRow, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableRow), element, func(self, child mvc.View) {
			self.(*tableRow).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// FACTORY FUNCTIONS

// Table returns a <cds-table> with pre-wired header and body slots.
// Pass TableRow children (they land in the body slot) and options as args.
// Call .Header() to set the header row.
//
//	cds.Table(
//	    cds.TableRow("Alice", "Engineer", "Active"),
//	    cds.TableRow("Bob",   "Designer", "Active"),
//	    cds.WithTableZebra(),
//	).Header("Name", "Role", "Status")
func Table(args ...any) *table {
	return mvc.NewView(new(table), ViewTable, templateTable, func(self, child mvc.View) {
		self.(*table).View = child
	}, args).(*table)
}

// TableRow returns a <cds-table-row>. Each string, dom.Element, or mvc.View
// arg is automatically wrapped in <cds-table-cell> via the Content() override.
//
//	cds.TableRow("Alice Martin", "Engineer", "Platform", "Active")
func TableRow(args ...any) *tableRow {
	return mvc.NewView(new(tableRow), ViewTableRow, "cds-table-row", func(self, child mvc.View) {
		self.(*tableRow).View = child
	}, args).(*tableRow)
}

// TableCell returns a plain <cds-table-cell> dom.Element. Use this when you
// need to apply attributes to an individual cell before passing it to TableRow.
func TableCell(args ...any) dom.Element {
	return mvc.HTML("cds-table-cell", args...)
}

// TableHeaderCell returns a <cds-table-header-cell> dom.Element. Useful when
// individual column options (e.g. WithTableHeaderSort) are needed; pass the
// result directly to Header().
//
//	t.Header(cds.TableHeaderCell("Name", cds.WithTableHeaderSort(cds.TableSortAscending)), "Role")
func TableHeaderCell(label string, args ...any) dom.Element {
	return mvc.HTML("cds-table-header-cell", append([]any{label}, args...)...)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Content wraps each string, dom.Element, or mvc.View arg in <cds-table-cell>
// before delegating to the embedded view. This fires both when TableRow() is
// called with args and when Content() is called explicitly later.
func (r *tableRow) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			args[i] = mvc.HTML("cds-table-cell", v)
		case dom.Element:
			if v.TagName() != "CDS-TABLE-CELL" {
				args[i] = mvc.HTML("cds-table-cell", v)
			}
		case mvc.View:
			if v.Root().TagName() != "CDS-TABLE-CELL" {
				args[i] = mvc.HTML("cds-table-cell", v)
			}
		}
	}
	return r.View.Content(args...)
}

// Header sets the header row. Each string arg is wrapped in
// <cds-table-header-cell>; a pre-built TableHeaderCell() element is inserted
// as-is so per-column options (e.g. sort direction) are preserved.
//
//	t.Header("Name", "Role", "Status")
//	t.Header(cds.TableHeaderCell("Name", cds.WithTableHeaderSort(cds.TableSortAscending)), "Role")
func (t *table) Header(args ...any) *table {
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			args[i] = mvc.HTML("cds-table-header-cell", v)
		case dom.Element:
			if v.TagName() != "CDS-TABLE-HEADER-CELL" {
				args[i] = mvc.HTML("cds-table-header-cell", v)
			}
		case mvc.View:
			if v.Root().TagName() != "CDS-TABLE-HEADER-CELL" {
				args[i] = mvc.HTML("cds-table-header-cell", v)
			}
		}
	}
	t.ReplaceSlot("header", mvc.HTML("cds-table-header-row", args...))
	return t
}

// Toolbar inserts a <cds-table-toolbar> as the first child of <cds-table> so
// that Carbon WC's firstUpdated() can auto-discover cds-table-batch-actions,
// cds-table-toolbar-content, and cds-table-toolbar-search via querySelector.
// Pass TableBatchActions() and/or TableToolbarContent() as args.
//
//	t := cds.Table(cds.WithTableSelectable(), rows...).
//	    Header("Name", "Role").
//	    Toolbar(
//	        cds.TableBatchActions(btnDelete),
//	        cds.TableToolbarContent(cds.TableToolbarSearch()),
//	    )
func (t *table) Toolbar(args ...any) *table {
	toolbar := mvc.HTML("cds-table-toolbar", args...)
	toolbar.SetAttribute("slot", "toolbar")
	t.Root().Prepend(toolbar)
	return t
}

///////////////////////////////////////////////////////////////////////////////
// CONVENIENCE

// DataTable builds a complete, ready-to-render table from string headers and
// a 2-D slice of cell values. Any additional options (WithTableSize, etc.)
// are applied to the table element.
//
//	cds.DataTable(
//	    []string{"Name", "Role", "Status"},
//	    [][]string{
//	        {"Alice", "Admin",   "Active"},
//	        {"Bob",   "Editor",  "Inactive"},
//	    },
//	    cds.WithTableSize(cds.TableLG),
//	)
func DataTable(headers []string, rows [][]string, opts ...any) *table {
	rowViews := make([]any, len(rows))
	for i, row := range rows {
		cells := make([]any, len(row))
		for j, v := range row {
			cells[j] = v
		}
		rowViews[i] = TableRow(cells...)
	}
	headerArgs := make([]any, len(headers))
	for i, h := range headers {
		headerArgs[i] = h
	}
	return Table(append(rowViews, opts...)...).Header(headerArgs...)
}

///////////////////////////////////////////////////////////////////////////////
// TOOLBAR HELPERS

// TableToolbar returns a <cds-table-toolbar> element.
// Pass it as the argument to Table.Toolbar() so it is placed inside
// <cds-table> where Carbon WC can auto-discover it.
func TableToolbar(args ...any) dom.Element {
	return mvc.HTML("cds-table-toolbar", args...)
}

// TableToolbarContent returns a <cds-table-toolbar-content> element for the
// right-hand side of the toolbar (search field, action buttons, etc.).
func TableToolbarContent(args ...any) dom.Element {
	return mvc.HTML("cds-table-toolbar-content", args...)
}

// TableToolbarSearch returns a <cds-table-toolbar-search> element.
// When placed inside a TableToolbar that is a sibling of a cds-table, the WC
// automatically filters visible rows on every keystroke via TableEventSearch.
func TableToolbarSearch(args ...any) dom.Element {
	return mvc.HTML("cds-table-toolbar-search", args...)
}

// TableBatchActions returns a <cds-table-batch-actions> element.
// Place it inside TableToolbar. When rows are selected the Carbon table WC
// activates this bar automatically and sets selectedRowsCount / totalRowsCount.
// Pass Button() children as batch operation actions.
func TableBatchActions(args ...any) dom.Element {
	return mvc.HTML("cds-table-batch-actions", args...)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTableSize sets the row density: xs, sm, md (default), lg, xl.
func WithTableSize(s TableSize) mvc.Opt {
	return mvc.WithAttr("size", string(s))
}

// WithTableSortable enables sortable column headers on the table.
func WithTableSortable() mvc.Opt {
	return mvc.WithAttr("is-sortable", "")
}

// WithTableZebra applies alternating row shading.
func WithTableZebra() mvc.Opt {
	return mvc.WithAttr("use-zebra-styles", "")
}

// WithTableHeaderSort sets the initial sort direction on a TableHeaderCell.
func WithTableHeaderSort(dir TableSortDir) mvc.Opt {
	return mvc.WithAttr("sort-direction", string(dir))
}

// WithTableSelectable enables checkbox row selection on the table.
func WithTableSelectable() mvc.Opt {
	return mvc.WithAttr("is-selectable", "")
}

// WithTableRadio switches row selection from checkboxes to radio buttons
// (single-select mode).
func WithTableRadio() mvc.Opt {
	return mvc.WithAttr("radio", "")
}

// WithTableRowSelected marks a row as initially selected.
func WithTableRowSelected() mvc.Opt {
	return mvc.WithAttr("selected", "")
}

// WithTableRowDisabled disables interaction on a row.
func WithTableRowDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithTableRowSelectionValue sets the value emitted in selection events for
// this row. Use unique values across rows.
func WithTableRowSelectionValue(v string) mvc.Opt {
	return mvc.WithAttr("selection-value", v)
}

// WithTableRowSelectionName sets the form field name shared by all row
// checkboxes or radios in the table.
func WithTableRowSelectionName(n string) mvc.Opt {
	return mvc.WithAttr("selection-name", n)
}

// WithTableRowSelectionLabel sets the accessible label for the row's
// selection checkbox or radio button.
func WithTableRowSelectionLabel(l string) mvc.Opt {
	return mvc.WithAttr("selection-label", l)
}

// WithToolbarSearchPersistent keeps the search field always expanded
// (non-collapsible).
func WithToolbarSearchPersistent() mvc.Opt {
	return mvc.WithAttr("persistent", "")
}
