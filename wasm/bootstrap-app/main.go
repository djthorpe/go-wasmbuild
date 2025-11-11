package main

import (
	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

// Application displays examples of Bootstrap components
func main() {
	mvc.New().Content(
		bs.Link("#badge", mvc.WithClass("m-2")).Content("Badges"),
		bs.VRule(),
		bs.Link("#link", mvc.WithClass("m-2")).Content("Links"),
		bs.VRule(),
		bs.Link("#button", mvc.WithClass("m-2")).Content("Buttons"),
		bs.VRule(),
		bs.Link("#button-group", mvc.WithClass("m-2")).Content("Button Groups"),
		bs.VRule(),
		bs.Link("#offcanvas", mvc.WithClass("m-2")).Content("Offcanvas"),

		mvc.Router(mvc.WithClass("container-fluid", "my-2")).Page(
			"#badge", BadgeExamples(),
		).Page(
			"#link", LinkExamples(),
		).Page(
			"#button", ButtonExamples(),
		).Page(
			"#button-group", ButtonGroupExamples(),
		).Page(
			"#offcanvas", OffcanvasExamples(),
		),
	)

	// Wait
	select {}
}

func BadgeExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.HRule(),
		bs.Heading(2).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.HRule(),
		bs.Heading(3).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.HRule(),
		bs.Heading(4).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.HRule(),
		bs.Heading(5).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.HRule(),
		bs.Heading(6).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
	)
}

func LinkExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("Link Examples"),
		bs.HRule(),
		bs.Link("#link").Content("Default Link Color"),
		bs.Link("#link", bs.WithColor(bs.Secondary)).Content("Secondary Link Color"),
		bs.Link("#link", bs.WithColor(bs.Danger)).Content("Danger Link Color"),
	)
}

func ButtonExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("Button Examples"),
		bs.HRule(),
		bs.Heading(3).Content("Standard Buttons"),
		bs.Button(mvc.WithClass("mx-1")).Content("Default Button"),
		bs.Button(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1")).Content("Secondary Button"),
		bs.Button(bs.WithColor(bs.Danger), mvc.WithClass("mx-1")).Content("Danger Button"),
		bs.HRule(),
		bs.Heading(3).Content("Outline Buttons"),
		bs.OutlineButton(mvc.WithClass("mx-1")).Content("Default Button"),
		bs.OutlineButton(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1")).Content("Secondary Button"),
		bs.OutlineButton(bs.WithColor(bs.Danger), mvc.WithClass("mx-1")).Content("Danger Button"),
		bs.HRule(),
		bs.Heading(3).Content("Close Buttons"),
		bs.CloseButton(mvc.WithClass("mx-1")),
		bs.HRule(),
		bs.Heading(3).Content("Large Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithSize(bs.Large)).Content("Default Button"),
		bs.HRule(),
		bs.Heading(3).Content("Small Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithSize(bs.Small)).Content("Default Button"),
		bs.HRule(),
		bs.Heading(3).Content("Disabled Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithDisabled(true)).Content("Default Button"),
		bs.Button(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1"), bs.WithDisabled(true)).Content("Secondary Button"),
		bs.Button(bs.WithColor(bs.Danger), mvc.WithClass("mx-1"), bs.WithDisabled(true)).Content("Danger Button"),
		bs.HRule(),
	)
}

func ButtonGroupExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("Button Groups"),
		bs.HRule(),
		bs.Heading(3).Content("Horizontal Button Group"),
		bs.ButtonGroup().Content(
			bs.Button().Content("Left"),
			bs.Button().Content("Middle"),
			bs.Button().Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2")).Content(
			bs.Button(bs.WithColor(bs.Danger)).Content("Left"),
			bs.Button(bs.WithColor(bs.Warning)).Content("Middle"),
			bs.Button(bs.WithColor(bs.Success)).Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2")).Content(
			bs.OutlineButton(bs.WithColor(bs.Danger)).Content("Left"),
			bs.OutlineButton(bs.WithColor(bs.Primary)).Content("Middle"),
			bs.OutlineButton(bs.WithColor(bs.Success)).Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2"), bs.WithSize(bs.Small)).Content(
			bs.Button(bs.WithColor(bs.Danger)).Content("Left"),
			bs.Button(bs.WithColor(bs.Warning)).Content("Middle"),
			bs.Button(bs.WithColor(bs.Success)).Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}), bs.ButtonGroup(mvc.WithClass("mx-2"), bs.WithSize(bs.Large)).Content(
			bs.Button(bs.WithColor(bs.Danger)).Content("Left"),
			bs.Button(bs.WithColor(bs.Warning)).Content("Middle"),
			bs.Button(bs.WithColor(bs.Success)).Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.HRule(),
		bs.Heading(3).Content("Vertical Button Group"),
		bs.VButtonGroup(mvc.WithClass("mx-2")).Content(
			bs.Button(bs.WithColor(bs.Danger)).Content("Top"),
			bs.Button(bs.WithColor(bs.Warning)).Content("Middle"),
			bs.Button(bs.WithColor(bs.Success)).Content("Bottom"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
	)
}

func OffcanvasExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("Offcanvas Example"),
		bs.HRule(),
		bs.Offcanvas("start", bs.WithPosition(bs.Start)).Header(
			mvc.HTML("H4", mvc.WithInnerText("This is the offcanvas title")),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		).Content(
			"This is the offcanvas content!",
		),
		bs.Offcanvas("end", bs.WithPosition(bs.End)).Content("This is the offcanvas content!"),
		bs.Offcanvas("top", bs.WithPosition(bs.Top)).Content("This is the offcanvas content!"),
		bs.Offcanvas("bottom", bs.WithPosition(bs.Bottom)).Content("This is the offcanvas content!"),
		bs.ButtonGroup().Content(
			bs.Button(bs.WithOffcanvas("start")).Content("Start"),
			bs.Button(bs.WithOffcanvas("end")).Content("End"),
			bs.Button(bs.WithOffcanvas("top")).Content("Top"),
			bs.Button(bs.WithOffcanvas("bottom")).Content("Bottom"),
		),
	)
}
