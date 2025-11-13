package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func CardExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Card Examples"),
		bs.HRule(),
		bs.Heading(3, "Single Card", mvc.WithClass("mt-5")), Example(Example_Cards_001),
		bs.Heading(3, "Cards With Color", mvc.WithClass("mt-5")), Example(Example_Cards_002),
		bs.Heading(3, "Cards With Header and Footer", mvc.WithClass("mt-5")), Example(Example_Cards_003),
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
	return bs.Grid(
		bs.Card(
			bs.Heading(5, "Primary Card"),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			bs.WithColor(bs.Primary),
			mvc.WithClass("m-1"),
		),
		bs.Card(
			bs.Heading(5, bs.Icon("question-circle-fill"), " Info"),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			bs.WithColor(bs.Secondary),
			mvc.WithClass("m-1"),
		),
		bs.Card(
			bs.Heading(5, bs.Icon("check-circle"), " Success"),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			bs.WithColor(bs.Success),
			mvc.WithClass("m-1"),
		),
	), sourcecode()
}

func Example_Cards_003() (mvc.View, string) {
	return bs.Grid(
		bs.Card(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		).Header(
			bs.Heading(5, "Are you sure?"),
		).Footer(
			bs.Button("OK", mvc.WithClass("mx-1")), bs.Button("Cancel", mvc.WithClass("mx-1")),
		),
	), sourcecode()
}
