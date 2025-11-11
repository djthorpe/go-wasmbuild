package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application displays examples of Bootstrap components
func main() {
	mvc.New().Content(
		bs.Link("#badge", mvc.WithClass("m-2")).Content("Badges"),
		bs.VRule(),
		bs.Link("#link", mvc.WithClass("m-2")).Content("Links"),
		bs.VRule(),
		bs.Link("#button", mvc.WithClass("m-2")).Content("Buttons"),
		mvc.Router(mvc.WithClass("container-fluid", "my-2")).Page(
			"#badge", BadgeExamples(),
		).Page(
			"#link", LinkExamples(),
		).Page(
			"#button", ButtonExamples(),
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
		bs.Heading(3).Content("Disabled Buttons TODO"),
	)
}
