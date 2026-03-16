package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func PaginationExamples() mvc.View {
	return ExamplePage("Pagination",
		cds.LeadPara(
			`Pagination uses the `, cds.Code("cds-pagination"), ` web component. `,
			`Pass `, cds.Code("cds.PaginationItem()"), ` children to populate the page-size dropdown. `,
			`Listen for `, cds.Code("PaginationEventPageChanged"), ` and `,
			cds.Code("PaginationEventPageSizeChanged"), ` to react to user navigation.`,
		),
		ExampleRow("Basic", Example_Pagination_001, "Fixed total with 10/20/50 page-size options."),
		ExampleRow("Sizes", Example_Pagination_002, "sm, md (default) and lg variants."),
		ExampleRow("Pages unknown", Example_Pagination_003, "Use WithPaginationPagesUnknown when the total item count is not available."),
		ExampleRow("With table", Example_Pagination_004, "Pagination bar below a data table; status line reflects the current page."),
	)
}

func Example_Pagination_001() (mvc.View, string) {
	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03);color:var(--cds-text-secondary,#525252);"),
		"Page 1 · 10 items per page",
	)
	p := cds.Pagination(
		cds.PaginationItem(10),
		cds.PaginationItem(20),
		cds.PaginationItem(50),
		cds.WithPaginationTotalItems(103),
		cds.WithPaginationPageSize(10),
	)
	p.Root().AddEventListener(cds.PaginationEventPageChanged, func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			status.Content("Page " + el.GetAttribute("page") + " · " + el.GetAttribute("page-size") + " items per page")
		}
	})
	p.Root().AddEventListener(cds.PaginationEventPageSizeChanged, func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			status.Content("Page " + el.GetAttribute("page") + " · " + el.GetAttribute("page-size") + " items per page")
		}
	})
	return cds.Section(p, status), sourcecode()
}

func Example_Pagination_002() (mvc.View, string) {
	return cds.Section(
		cds.Pagination(
			cds.PaginationItem(10),
			cds.PaginationItem(20),
			cds.WithPaginationTotalItems(50),
			cds.WithPaginationSize(cds.PaginationSM),
		),
		cds.Pagination(
			cds.PaginationItem(10),
			cds.PaginationItem(20),
			cds.WithPaginationTotalItems(50),
			cds.WithPaginationSize(cds.PaginationMD),
		),
		cds.Pagination(
			cds.PaginationItem(10),
			cds.PaginationItem(20),
			cds.WithPaginationTotalItems(50),
			cds.WithPaginationSize(cds.PaginationLG),
		),
	), sourcecode()
}

func Example_Pagination_003() (mvc.View, string) {
	return cds.Pagination(
		cds.PaginationItem(10),
		cds.PaginationItem(20),
		cds.WithPaginationPagesUnknown(),
		cds.WithPaginationPageSize(10),
	), sourcecode()
}

func Example_Pagination_004() (mvc.View, string) {
	const totalItems = 103
	const defaultPageSize = 10

	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03);color:var(--cds-text-secondary,#525252);"),
		"Showing items 1–10 of 103",
	)

	p := cds.Pagination(
		cds.PaginationItem(10),
		cds.PaginationItem(20),
		cds.PaginationItem(50),
		cds.WithPaginationTotalItems(totalItems),
		cds.WithPaginationPageSize(defaultPageSize),
	)
	p.Root().AddEventListener(cds.PaginationEventPageChanged, func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			status.Content("Page " + el.GetAttribute("page") + " · " + el.GetAttribute("page-size") + " items per page")
		}
	})
	p.Root().AddEventListener(cds.PaginationEventPageSizeChanged, func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			status.Content("Page " + el.GetAttribute("page") + " · " + el.GetAttribute("page-size") + " items per page")
		}
	})

	t := cds.DataTable(tableHeaders, tableRows, cds.WithTableSize(cds.TableLG))
	return cds.Section(t, p, status), sourcecode()
}
