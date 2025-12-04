package main

import (
	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func OffcanvasExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("offcanvas_examples.md"),
		bs.Container(
			bs.WithBorder(), mvc.WithClass("my-2"), bs.WithColor(bs.Light), mvc.WithClass("p-3"),
			bs.Heading(3, "Offcanvas Dialog"), Example(Example_Offcanvas_001),
		), bs.Container(
			bs.WithBorder(), mvc.WithClass("my-2"), bs.WithColor(bs.Light), mvc.WithClass("p-3"),
			bs.Heading(3, "Offcanvas Color and Theme"), Example(Example_Offcanvas_002),
		),
	)
}

func Example_Offcanvas_001() (mvc.View, string) {
	return bs.Container(
		bs.Offcanvas("dialog_001", bs.WithOffcanvasScroll(),
			bs.Markdown("Offcanvas can be dismissed by clicking the close button above, clicking outside the offcanvas, or pressing the ESC key.\n\n"+
				"The `bs.WithOffcanvasScroll()` option allows the body to be scrollable when the offcanvas is open.",
			),
		).Header(
			bs.Heading(5, "Offcanvas Dialog"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("dialog_001"), "Open Offcanvas"),
	), sourcecode()
}

func Example_Offcanvas_002() (mvc.View, string) {
	return bs.Container(
		bs.Offcanvas("offcanvas_002",
			bs.WithColor(bs.Secondary), bs.WithTheme("dark"), bs.WithPosition(bs.End),
			"This is the offcanvas content!",
		).Header(
			bs.Heading(5, "Offcanvas"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("offcanvas_002"), "Open Offcanvas"),
	), sourcecode()
}
