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
		), bs.Container(
			bs.WithBorder(), mvc.WithClass("my-2"), bs.WithColor(bs.Light), mvc.WithClass("p-3"),
			bs.Heading(3, "Sticky Offcanvas"), Example(Example_Offcanvas_003),
		), bs.Container(
			bs.WithBorder(), mvc.WithClass("my-2"), bs.WithColor(bs.Light), mvc.WithClass("p-3"),
			bs.Heading(3, "Responsive Offcanvas"), Example(Example_Offcanvas_004),
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
			bs.Markdown("This offcanvas has a different color, theme and position."),
		).Header(
			bs.Heading(5, "Offcanvas Themes and Colors"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("offcanvas_002"), "Open Offcanvas"),
	), sourcecode()
}

func Example_Offcanvas_003() (mvc.View, string) {
	return bs.Container(
		bs.StickyOffcanvas("offcanvas_003",
			bs.WithTheme("dark"), bs.WithPosition(bs.Bottom), bs.WithOffcanvasScroll(),
			"Sticky offcanvas dialogs will not close when clicking outside of it, or pressing the ESC key. ",
			"Press the Close button above to close this offcanvas.",
		).Header(
			bs.Heading(5, "Offcanvas"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("offcanvas_003"), "Open Offcanvas"),
	), sourcecode()
}

func Example_Offcanvas_004() (mvc.View, string) {
	return bs.Container(
		bs.Offcanvas("offcanvas_004",
			bs.WithPosition(bs.Top), bs.WithSize(bs.Medium),
			bs.Markdown(`
				Responsive offcanvas classes hide content outside the viewport from a specified breakpoint and down. 
				Above that breakpoint, the contents within will behave as usual.

				In order to try this, resize the browser window to be less than "md" size (768px width), and this
				content will be hidden until you click the "Open Responsive Offcanvas" button.
			`),
		).Header(
			bs.Heading(5, "Responsive Offcanvas"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("offcanvas_004"), mvc.WithClass("d-md-none"), "Open Responsive Offcanvas"),
	), sourcecode()
}
