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
		bs.Link("#text", mvc.WithClass("m-2")).Content("Text"),
		bs.VRule(),
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
		bs.Link("#card", mvc.WithClass("m-2")).Content("Cards"),
		bs.VRule(),
		bs.Link("#modal", mvc.WithClass("m-2")).Content("Modal"),
		bs.VRule(),
		bs.Link("#input", mvc.WithClass("m-2")).Content("Input"),
		bs.VRule(),
		bs.Link("#tooltips", mvc.WithClass("m-2")).Content("Tooltips"),
		bs.VRule(),
		bs.Link("#progress", mvc.WithClass("m-2")).Content("Progress Bars"),
		bs.VRule(),
		bs.Link("#navbar", mvc.WithClass("m-2")).Content("Navbars"),
		bs.VRule(),
		bs.Link("#table", mvc.WithClass("m-2")).Content("Tables"),

		mvc.Router(mvc.WithClass("container-fluid", "my-2")).Page(
			"#text", TextExamples(),
		).Page(
			"#badge", BadgeExamples(),
		).Page(
			"#link", LinkExamples(),
		).Page(
			"#list", ListExamples(),
		).Page(
			"#button", ButtonExamples(),
		).Page(
			"#card", CardExamples(),
		).Page(
			"#icon", IconExamples(),
		).Page(
			"#modal", ModalExamples(),
		).Page(
			"#input", InputExamples(),
		).Page(
			"#tooltips", TooltipExamples(),
		).Page(
			"#progress", ProgressExamples(),
		).Page(
			"#navbar", NavBarExamples(),
		).Page(
			"#table", TableExamples(),
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
	withButtonClasses := func(opts ...any) []any {
		base := []any{mvc.WithClass("d-flex", "align-items-center", "gap-2", "justify-content-center", "my-1", "px-3")}
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
