package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	bsextra "github.com/djthorpe/go-wasmbuild/pkg/bootstrap/extra"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func PaginationExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Pagination Examples"), bs.HRule(),
		bs.Heading(3, "Pagination", mvc.WithClass("mt-5")), Example(Example_Pagination_001),
		bs.Heading(3, "Events", mvc.WithClass("mt-5")), Example(Example_Pagination_002),
		bs.Heading(3, "Icons & Theme", mvc.WithClass("mt-5")), Example(Example_Pagination_003),
		bs.Heading(3, "Controller", mvc.WithClass("mt-5")), Example(Example_Pagination_004),
	)
}

func Example_Pagination_001() (mvc.View, string) {
	return bs.Pagination(
		"Prev", "1", "2",
		bs.PaginationItem("...", bs.WithDisabled(true)),
		bs.PaginationItem("3", bs.WithActive(true)),
		"Next",
	), sourcecode()
}

func Example_Pagination_002() (mvc.View, string) {
	response := bs.Para("Click on a pagination item")
	return bs.Container(
		response,
		bs.Pagination(
			"Prev", "1", "2",
			bs.PaginationItem("...", bs.WithDisabled(true)),
			bs.PaginationItem("3", bs.WithActive(true)),
			"Next",
		).AddEventListener("click", func(e dom.Event) {
			view := mvc.ViewFromEvent(e)
			if view.Name() == bs.ViewPaginationItem {
				response.Content(bs.Code(view.Root().TextContent()), " clicked")
			} else {
				response.Content()
			}
		})), sourcecode()
}

func Example_Pagination_003() (mvc.View, string) {
	return bs.Pagination(
		bs.Icon("arrow-left"),
		"1", "2", "3",
		bs.PaginationItem("4", bs.WithActive(true)),
		"5",
		bs.Icon("arrow-right"),
		bs.WithTheme(bs.Dark),
	), sourcecode()
}

func Example_Pagination_004() (mvc.View, string) {
	// Create the pagination view
	view := bs.Pagination(
		bs.Icon("arrow-left"),
		"1", "2", "3",
		bs.PaginationItem("4", bs.WithActive(true)),
		"5",
		bs.Icon("arrow-right"),
		bs.WithTheme(bs.Dark),
	)

	// Create controller and attach the pagination view
	bsextra.PaginationController(view)

	// Return the view for display
	return view, sourcecode()
}
