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
		bs.Link("#list", mvc.WithClass("m-2")).Content("Lists"),
		bs.VRule(),
		bs.Link("#icon", mvc.WithClass("m-2")).Content("Icons"),
		bs.VRule(),
		bs.Link("#button", mvc.WithClass("m-2")).Content("Buttons"),
		bs.VRule(),
		bs.Link("#button-group", mvc.WithClass("m-2")).Content("Button Groups"),
		bs.VRule(),
		bs.Link("#card", mvc.WithClass("m-2")).Content("Cards"),
		bs.VRule(),
		bs.Link("#offcanvas", mvc.WithClass("m-2")).Content("Offcanvas"),
		bs.VRule(),
		bs.Link("#input", mvc.WithClass("m-2")).Content("Input"),

		mvc.Router(mvc.WithClass("container-fluid", "my-2")).Page(
			"#badge", BadgeExamples(),
		).Page(
			"#link", LinkExamples(),
		).Page(
			"#list", ListExamples(),
		).Page(
			"#button", ButtonExamples(),
		).Page(
			"#button-group", ButtonGroupExamples(),
		).Page(
			"#card", CardExamples(),
		).Page(
			"#icon", IconExamples(),
		).Page(
			"#offcanvas", OffcanvasExamples(),
		).Page(
			"#offcanvas", OffcanvasExamples(),
		).Page(
			"#input", InputExamples(),
		),
	)

	// Wait
	select {}
}

func InputExamples() mvc.View {
	rangevalue := mvc.CData("")
	return bs.Container().Content(
		bs.Heading(1).Content("Input Examples"),
		bs.HRule(),
		bs.Form("input").Content(
			bs.Card().Header(
				bs.Heading(4).Content("Enter your details"),
			).Content(
				bs.InputGroup(mvc.WithClass("my-2")).Content(
					bs.Input("username", bs.WithPlaceholder("Enter username"), bs.WithRequired(), bs.WithAutocomplete("user", "email")),
					"@",
					bs.Input("domain", bs.WithPlaceholder("Enter domain here"), bs.WithRequired(), bs.WithAutocomplete("domain")),
				),
				bs.Password("password", bs.WithPlaceholder("Enter password here"), mvc.WithClass("my-2"), bs.WithRequired(), bs.WithoutAutocomplete()),
				bs.Number(
					"number", bs.WithMinMax(-5, 5), bs.WithPlaceholder("Enter number here"), mvc.WithClass("my-2"), bs.WithRequired(), bs.WithoutAutocomplete(),
				).(mvc.ViewWithCaption).Caption(
					"Number of times",
				),
				bs.Textarea("description", bs.WithPlaceholder("Enter description here"), mvc.WithClass("my-2")),
				bs.InputGroup(mvc.WithClass("my-2")).Content(
					bs.Range("range", bs.WithMinMax(-5, 5)).AddEventListener("input", func(e Event) {
						r := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
						if r != nil {
							rangevalue.SetData(r.Value())
						}
					}),
					rangevalue,
				),
			).(mvc.ViewWithHeaderFooter).Footer(
				bs.Button(bs.WithColor(bs.Primary), bs.WithSubmit()).Content("Submit"),
			),
		),
	)
}

func ListExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("List Examples"),
		bs.HRule(),
		bs.Heading(3).Content("Unstyled List"),
		bs.UnstyledList().Content(
			"Item 1",
			"Item 2",
			"Item 3",
		),
		bs.HRule(),
		bs.Heading(3).Content("Ordered List"),
		bs.List().Content(
			"Item A",
			"Item B",
			"Item C",
		),
		bs.HRule(),
		bs.Heading(3).Content("Ordered Tree"),
		bs.List().Content(
			"Item A",
			"Item B",
			bs.List().Content(
				"Item B.1",
				"Item B.2",
			),
			"Item C",
		),
		bs.HRule(),
		bs.Heading(3).Content("Bulleted List"),
		bs.BulletList().Content(
			"Item A",
			"Item B",
			"Item C",
		),
		bs.HRule(),
		bs.Heading(3).Content("Link List"),
		bs.UnstyledList().Content(
			bs.Link("#link").Content("Default Link Color"),
			bs.Link("#link", bs.WithColor(bs.Secondary)).Content("Secondary Link Color"),
			bs.Link("#link", bs.WithColor(bs.Danger)).Content("Danger Link Color"),
		),
		bs.HRule(),
		bs.Heading(3).Content("List Group"),
		bs.ListGroup().Content(
			"Item 1",
			"Item 2",
			"Item 3",
		),
	)
}

func IconExamples() mvc.View {
	withButtonClasses := func(opts ...mvc.Opt) []mvc.Opt {
		base := []mvc.Opt{mvc.WithClass("d-flex", "align-items-center", "gap-2", "justify-content-center", "my-1", "px-3")}
		return append(base, opts...)
	}
	return bs.Container().Content(
		bs.Heading(1).Content("Icon Examples"),
		bs.HRule(),
		bs.Heading(3).Content("Common glyphs"),
		bs.Grid().Append(
			bs.Icon("alarm", mvc.WithClass("fs-1")),
			bs.Icon("award", mvc.WithClass("fs-1")),
			bs.Icon("bell", mvc.WithClass("fs-1")),
			bs.Icon("chat-dots", mvc.WithClass("fs-1")),
			bs.Icon("cloud", mvc.WithClass("fs-1")),
		),
		bs.HRule(),
		bs.Heading(3).Content("Colors"),
		bs.Grid().Append(
			bs.Icon("heart-fill", bs.WithColor(bs.Danger), mvc.WithClass("fs-1")),
			bs.Icon("sun-fill", bs.WithColor(bs.Warning), mvc.WithClass("fs-1")),
			bs.Icon("moon-stars", bs.WithColor(bs.Primary), mvc.WithClass("fs-1")),
			bs.Icon("toggle2-on", bs.WithColor(bs.Success), mvc.WithClass("fs-1")),
			bs.Icon("wifi", bs.WithColor(bs.Info), mvc.WithClass("fs-1")),
		),
		bs.HRule(),
		bs.Heading(3).Content("Buttons with icons"),
		bs.Grid().Append(
			bs.Button(withButtonClasses(bs.WithColor(bs.Primary))...).Content("Share", bs.Icon("share")),
			bs.OutlineButton(withButtonClasses(bs.WithColor(bs.Success))...).Content("Download", bs.Icon("cloud-arrow-down")),
			bs.Button(withButtonClasses(bs.WithColor(bs.Success))...).Content("Next", bs.Icon("arrow-right")),
			bs.Button(withButtonClasses(bs.WithColor(bs.Danger))...).Content(bs.Icon("trash"), "Delete"),
			bs.OutlineButton(withButtonClasses(bs.WithColor(bs.Dark))...).Content("Refresh", bs.Icon("arrow-repeat")),
		),
	)
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
		bs.UnstyledList().Content(
			bs.Link("#link").Content("Default Link Color"),
			bs.Link("#link", bs.WithColor(bs.Secondary)).Content("Secondary Link Color"),
			bs.Link("#link", bs.WithColor(bs.Danger)).Content("Danger Link Color"),
		),
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

func CardExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(1).Content("Cards"),
		bs.HRule(),
		bs.Heading(3).Content("Single Cards"),
		bs.Grid().Append(
			bs.Card().Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Primary)).Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Secondary)).Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Info)).Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Success)).Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Warning)).Content(
				"This is some card content!",
			),
			bs.Card(bs.WithColor(bs.Danger)).Content(
				"This is some card content!",
			),
		),
		bs.HRule(),
		bs.Heading(3).Content("Card Group"),
		bs.CardGroup().Content(
			bs.Card().Header("Header").Footer(
				"Card Footer",
			).Content(
				"This is some card content which is longer to show that the footers line up in the card groups",
			),
			bs.Card().Header("Header").Footer(
				"Card Footer",
			).Content("This is some card content!"),
		),
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
		bs.Offcanvas("end", bs.WithPosition(bs.End), bs.WithTheme(bs.Dark)).Header(
			mvc.HTML("H4", mvc.WithInnerText("This is the offcanvas title")),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		).Content(
			"This is the offcanvas content!",
		),
		bs.Offcanvas("top", bs.WithPosition(bs.Top)).Content("This is the offcanvas content!"),
		bs.Offcanvas("bottom", bs.WithPosition(bs.Bottom)).Content("This is the offcanvas content!"),
		bs.ButtonGroup().Content(
			bs.Button(bs.WithOffcanvas("start")).Content("Start"),
			bs.Button(bs.WithOffcanvas("end"), bs.WithColor(bs.Dark)).Content("End"),
			bs.Button(bs.WithOffcanvas("top")).Content("Top"),
			bs.Button(bs.WithOffcanvas("bottom")).Content("Bottom"),
		),
	)
}
