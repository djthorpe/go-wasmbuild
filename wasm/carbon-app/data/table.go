package data

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

type tableRecord struct {
	name   string
	role   string
	status string
}

var tableRecords = []tableRecord{
	{"Ada Lovelace", "Maintainer", "Ready"},
	{"Grace Hopper", "Reviewer", "In progress"},
	{"Margaret Hamilton", "Approver", "Ready"},
	{"Katherine Johnson", "Analyst", "Blocked"},
}

func TableView() []any {
	return []any{
		storybook.PageHeader("Table", "Table.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicTableStory(),
		),
	}
}

func basicTableStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	search := carbon.TableToolbarSearch().
		SetLabel("Filter rows").
		SetPlaceholder("Search people or status")
	table := carbon.Table()
	status := carbon.Compact()

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		carbon.TableToolbar(
			search,
			carbon.Button(carbon.With(carbon.KindSecondary), "Export"),
			carbon.Button(carbon.With(carbon.KindGhost), "Refresh"),
		),
		table,
		status,
	)

	render := func(filter string) {
		canvas.Apply(carbon.With(currentTheme)...)
		filter = strings.TrimSpace(strings.ToLower(filter))

		rows := make([]any, 0, len(tableRecords)+1)
		rows = append(rows, carbon.TableHeader("Name", "Role", "Status"))
		count := 0
		for _, record := range tableRecords {
			blob := strings.ToLower(record.name + " " + record.role + " " + record.status)
			if filter != "" && !strings.Contains(blob, filter) {
				continue
			}
			rows = append(rows, carbon.TableRow(record.name, record.role, record.status))
			count++
		}
		table.Content(rows...)
		if count == 1 {
			status.Content("Showing 1 matching row.")
		} else {
			status.Content(fmt.Sprintf("Showing %d matching rows.", count))
		}
	}
	render("")

	search.AddEventListener(carbon.EventInput, func(dom.Event) {
		render(search.Value())
	})
	search.AddEventListener(carbon.EventChange, func(dom.Event) {
		render(search.Value())
	})

	return storybook.Story(
		"Filterable Table",
		"The table wrapper keeps the header and body structure explicit. The toolbar story demonstrates how TableToolbarSearch stays in the direct slot while action buttons are grouped automatically into TableToolbarContent.",
		canvas,
		search,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render(search.Value())
		}),
	)
}
