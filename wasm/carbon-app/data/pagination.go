package data

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var paginationSizes = []carbon.Attr{"10", "20", "50"}

func PaginationView() []any {
	return []any{
		storybook.PageHeader("Pagination", "Pagination.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicPaginationStory(),
		),
	}
}

func basicPaginationStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := paginationSizes[1]
	enabled := true
	unknownPages := false

	pager := carbon.Pagination(
		mvc.WithAttr("page-sizes", "10,20,50"),
	)
	pager.SetCount(137)
	pager.SetLimit(20)
	pager.SetPage(3)
	pager.SetItemsPerPageText("Items per page")
	pager.SetPageSizeLabelText("Page size")
	pager.SetBackwardText("Previous page")
	pager.SetForwardText("Next page")

	status := carbon.Para()
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		carbon.Tile(
			carbon.WithFill(),
			mvc.WithStyle("padding:1rem"),
			carbon.Head(4, "Current pagination state"),
			status,
		),
		pager,
	)

	render := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		pager.SetEnabled(enabled)
		pager.SetPagesUnknown(unknownPages)
		switch currentSize {
		case "10":
			pager.SetLimit(10)
		case "50":
			pager.SetLimit(50)
		default:
			pager.SetLimit(20)
		}
		status.Content(fmt.Sprintf("Page %d, offset %d, limit %d, total %d", mvc.Page(pager), pager.Offset(), pager.Limit(), pager.Count()))
	}
	render()

	pager.AddEventListener(carbon.EventPaginationChanged, func(dom.Event) {
		render()
	})
	pager.AddEventListener(carbon.EventPaginationPageSize, func(dom.Event) {
		render()
	})

	return storybook.Story(
		"Pagination State",
		"Pagination exposes the shared pagination-state interface for page, offset, limit, and total count while still surfacing Carbon's page-change events. The preview keeps the component live so you can exercise the control directly.",
		canvas,
		pager,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Page size", currentSize, paginationSizes, func(size carbon.Attr) {
			currentSize = size
			render()
		}),
		storybook.CheckboxGroup("State", "Disabled", !enabled, func(disabled bool) {
			enabled = !disabled
			render()
		}),
		storybook.CheckboxGroup("Mode", "Pages unknown", unknownPages, func(value bool) {
			unknownPages = value
			render()
		}),
	)
}
