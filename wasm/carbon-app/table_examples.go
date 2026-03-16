package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TableExamples() mvc.View {
	return ExamplePage("Data table",
		cds.LeadPara(
			`Data tables use the `, cds.Code("cds-table"), ` web component family. `,
			`Use the `, cds.Code("cds.DataTable()"), ` convenience helper to build a table from string slices, `,
			`or compose `, cds.Code("cds.Table()"), ` / `, cds.Code("cds.TableRow()"), ` directly. `,
			`Call `, cds.Code(".Header()"), ` on the table to set the header row. `,
			`Add a toolbar above the table with `, cds.Code("cds.TableToolbar()"), `, `,
			cds.Code("cds.TableToolbarSearch()"), `, and `, cds.Code("cds.TableBatchActions()"), `.`,
		),
		ExampleRow("Basic", Example_Table_001, "A simple table built with the DataTable convenience helper."),
		ExampleRow("Sizes", Example_Table_002, "Five row densities: xs, sm, md (default), lg, xl."),
		ExampleRow("Zebra striping", Example_Table_003, "Alternating row shading with cds.WithTableZebra()."),
		ExampleRow("Sortable columns", Example_Table_004, "Enable column sorting; pass cds.TableHeaderCell() with a sort option to Header()."),
		ExampleRow("Interactive rows", Example_Table_005, "Attach a click listener to a TableRow; the name is captured in the closure."),
		ExampleRow("Row selection", Example_Table_007, "Checkbox selection with cds.WithTableSelectable(); listen for TableEventRowSelected."),
		ExampleRow("Toolbar search", Example_Table_008, "Built-in row filtering: TableToolbar + TableToolbarSearch as a sibling of the table."),
		ExampleRow("Batch actions", Example_Table_009, "Batch-actions bar activates automatically when rows are selected."),
	)
}

var (
	tableHeaders = []string{"Name", "Role", "Department", "Status"}
	tableRows    = [][]string{
		{"Alice Martin", "Engineer", "Platform", "Active"},
		{"Bob Chen", "Designer", "Product", "Active"},
		{"Carol Davis", "Manager", "Engineering", "On leave"},
		{"Dave Wilson", "Analyst", "Finance", "Active"},
		{"Eve Thompson", "Engineer", "Platform", "Inactive"},
	}
)

func Example_Table_001() (mvc.View, string) {
	return cds.DataTable(tableHeaders, tableRows), sourcecode()
}

func Example_Table_002() (mvc.View, string) {
	makeTable := func(size cds.TableSize, label string) mvc.View {
		return cds.Section(
			mvc.WithAttr("style", "margin-bottom:var(--cds-spacing-06,1.5rem);"),
			cds.Heading(5, label),
			cds.DataTable(tableHeaders, tableRows[:3], cds.WithTableSize(size)),
		)
	}
	return cds.Section(
		makeTable(cds.TableXS, "Extra small (xs)"),
		makeTable(cds.TableSM, "Small (sm)"),
		makeTable(cds.TableMD, "Medium (md — default)"),
		makeTable(cds.TableLG, "Large (lg)"),
		makeTable(cds.TableXL, "Extra large (xl)"),
	), sourcecode()
}

func Example_Table_003() (mvc.View, string) {
	return cds.DataTable(tableHeaders, tableRows, cds.WithTableZebra()), sourcecode()
}

func Example_Table_004() (mvc.View, string) {
	return cds.Table(
		cds.WithTableSortable(),
		cds.TableRow("Alice Martin", "Engineer", "Platform", "Active"),
		cds.TableRow("Bob Chen", "Designer", "Product", "Active"),
		cds.TableRow("Carol Davis", "Manager", "Engineering", "On leave"),
		cds.TableRow("Dave Wilson", "Analyst", "Finance", "Active"),
		cds.TableRow("Eve Thompson", "Engineer", "Platform", "Inactive"),
	).Header(
		cds.TableHeaderCell("Name", cds.WithTableHeaderSort(cds.TableSortAscending)),
		"Role",
		"Department",
		"Status",
	), sourcecode()
}

func Example_Table_005() (mvc.View, string) {
	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);color:var(--cds-text-secondary,#525252);"),
		"Click a row…",
	)
	rows := make([]any, len(tableRows))
	for i, r := range tableRows {
		name := r[0]
		row := cds.TableRow(r[0], r[1], r[2], r[3])
		row.AddEventListener("click", func(e dom.Event) {
			status.Content("Selected: ", cds.Strong(name))
		})
		rows[i] = row
	}
	t := cds.Table(rows...).Header("Name", "Role", "Department", "Status")
	return cds.Section(t, status), sourcecode()
}

func Example_Table_006() (mvc.View, string) {
	return cds.Table(
		cds.TableRow("Mercury", "0.39", "0", "Terrestrial"),
		cds.TableRow("Venus", "0.72", "0", "Terrestrial"),
		cds.TableRow("Earth", "1.00", "1", "Terrestrial"),
		cds.TableRow("Jupiter", "5.20", "95", "Gas giant"),
		cds.TableRow("Saturn", "9.58", "146", "Gas giant"),
	).Header(
		"Planet", "Distance (AU)", "Moons", "Type",
	), sourcecode()
}

func Example_Table_007() (mvc.View, string) {
	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);color:var(--cds-text-secondary,#525252);"),
		"Select rows using the checkboxes…",
	)
	t := cds.Table(
		cds.WithTableSelectable(),
		cds.TableRow(cds.WithTableRowSelectionValue("alice"), "Alice Martin", "Engineer", "Platform", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("bob"), "Bob Chen", "Designer", "Product", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("carol"), cds.WithTableRowSelected(), "Carol Davis", "Manager", "Engineering", "On leave"),
		cds.TableRow(cds.WithTableRowSelectionValue("dave"), "Dave Wilson", "Analyst", "Finance", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("eve"), cds.WithTableRowDisabled(), "Eve Thompson", "Engineer", "Platform", "Inactive"),
	).Header("Name", "Role", "Department", "Status")
	// TableEventRowSelect fires AFTER the row's selected attribute is updated.
	t.AddEventListener(cds.TableEventRowSelect, func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			val := el.GetAttribute("selection-value")
			if el.HasAttribute("selected") {
				status.Content("Selected: ", cds.Strong(val))
			} else {
				status.Content("Deselected: ", cds.Strong(val))
			}
		}
	})
	return cds.Section(t, status), sourcecode()
}

func Example_Table_008() (mvc.View, string) {
	t := cds.Table(
		cds.TableRow("Alice Martin", "Engineer", "Platform", "Active"),
		cds.TableRow("Bob Chen", "Designer", "Product", "Active"),
		cds.TableRow("Carol Davis", "Manager", "Engineering", "On leave"),
		cds.TableRow("Dave Wilson", "Analyst", "Finance", "Active"),
		cds.TableRow("Eve Thompson", "Engineer", "Platform", "Inactive"),
	).Header("Name", "Role", "Department", "Status")
	return cds.Section(
		cds.TableToolbar(
			cds.TableToolbarContent(
				cds.TableToolbarSearch(cds.WithToolbarSearchPersistent()),
			),
		),
		t,
	), sourcecode()
}

func Example_Table_009() (mvc.View, string) {
	btnDelete := cds.Button("Delete", cds.WithButtonKind(cds.ButtonGhost))
	btnDownload := cds.Button("Download", cds.WithButtonKind(cds.ButtonGhost))

	batchActions := cds.TableBatchActions(btnDelete, btnDownload)

	t := cds.Table(
		cds.WithTableSelectable(),
		cds.TableRow(cds.WithTableRowSelectionValue("alice"), "Alice Martin", "Engineer", "Platform", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("bob"), "Bob Chen", "Designer", "Product", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("carol"), "Carol Davis", "Manager", "Engineering", "On leave"),
		cds.TableRow(cds.WithTableRowSelectionValue("dave"), "Dave Wilson", "Analyst", "Finance", "Active"),
		cds.TableRow(cds.WithTableRowSelectionValue("eve"), "Eve Thompson", "Engineer", "Platform", "Inactive"),
	).Header("Name", "Role", "Department", "Status").
		Toolbar(
			batchActions,
			cds.TableToolbarContent(
				cds.TableToolbarSearch(cds.WithToolbarSearchPersistent()),
			),
		)

	return t, sourcecode()
}
