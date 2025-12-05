package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func CardExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("card_examples.md"),
		ExampleCard("Single Card", Example_Cards_001),
		ExampleCard("Cards with Color", Example_Cards_002),
		ExampleCard("Header and Footer", Example_Cards_003),
		ExampleCard("Card Group", Example_Cards_004),
		ExampleCard("Image Card", Example_Cards_005),
	)
}

func Example_Cards_001() (mvc.View, string) {
	return bs.Card(
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
		quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	), sourcecode()
}

func Example_Cards_002() (mvc.View, string) {
	return bs.Row(
		bs.Col(
			bs.Card(
				bs.WithColor(bs.Primary), mvc.WithClass("m-1"),
				bs.Heading(5, "Primary Card"),
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			),
		),
		bs.Col(
			bs.Card(
				bs.WithColor(bs.Secondary), mvc.WithClass("m-1"),
				bs.Heading(5, bs.Icon("question-circle-fill"), " Info"),
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			),
		),
		bs.Col(
			bs.Card(
				bs.WithColor(bs.Success), mvc.WithClass("m-1"),
				bs.Heading(5, bs.Icon("check-circle"), " Success"),
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			),
		),
	), sourcecode()
}

func Example_Cards_003() (mvc.View, string) {
	return bs.Card(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	).Header(
		bs.Heading(5, "Are you sure?"),
	).Footer(
		bs.WithPosition(bs.Center),
		bs.Button("OK", mvc.WithClass("mx-1")), bs.Button("Cancel", mvc.WithClass("mx-1")),
	), sourcecode()
}

func Example_Cards_004() (mvc.View, string) {
	placeholder := `<svg role="img"><rect width="100%" height="100%" fill="#868e96"></rect><text x="50%" y="50%" text-anchor="middle" dominant-baseline="middle" fill="#dee2e6">Image</text></svg>`
	return bs.CardGroup(mvc.WithClass("m-3"),
		bs.Card(
			"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.",
		).Label(
			mvc.HTML(placeholder),
		),
		bs.Card(
			"This card has supporting text below as a natural lead-in to additional content.",
		).Label(
			mvc.HTML(placeholder),
		),
		bs.Card(
			"This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
		).Label(
			mvc.HTML(placeholder),
		),
	), sourcecode()
}

func Example_Cards_005() (mvc.View, string) {
	return bs.Card(
		"Some quick example text to build on the card title and make up the bulk of the card’s content.",
		mvc.WithStyle("width: 15rem;"),
	).Label(
		bs.Image("favicon.png"),
	), sourcecode()
}
