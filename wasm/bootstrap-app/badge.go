package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func BadgeExamples() mvc.View {
	return bs.Container(
		BadgeHeaders(),
		bs.HRule(),
		BadgeWithIndicator(),
	)
}

func BadgeHeaders() mvc.View {
	return bs.Grid().Content(
		bs.Container(mvc.WithClass("my-2")).Content(
			bs.Heading(1, "Example heading ", bs.Badge("New")),
			bs.Heading(2, "Example heading ", bs.Badge("New", bs.WithColor(bs.Success))),
			bs.Heading(3, "Example heading ", bs.PillBadge("New")),
			bs.Heading(4, "Example heading ", bs.Badge("New", bs.WithColor(bs.Secondary))),
			bs.Heading(5, "Example heading ", bs.Badge("New")),
			bs.Heading(6, "Example heading ", bs.Badge("New")),
		), bs.Container().Content(
			bs.Para(
				`Headings can include badges, which can be created with the bs.Badge and bs.PillBadge functions.
				The color of the badge can be changed with the bs.WithColor option.`,
			),
			bs.CodeBlock(bs.WithColor(bs.Light), mvc.WithClass("p-3"), mvc.WithClass("border", "border-dark-subtle")).Content(
				`bs.Heading(1, "Example heading ", bs.Badge("New")),
bs.Heading(2, "Example heading ", bs.Badge("New", bs.WithColor(bs.Success))),
bs.Heading(3, "Example heading ", bs.PillBadge("New")),
bs.Heading(4, "Example heading ", bs.Badge("New", bs.WithColor(bs.Secondary))),
bs.Heading(5, "Example heading ", bs.Badge("New")),
bs.Heading(6, "Example heading ", bs.Badge("New"))`,
			),
		),
	)
}

func BadgeWithIndicator() mvc.View {
	return bs.Grid().Content(
		bs.Container(mvc.WithClass("my-2")).Content(
			bs.Heading(1, "Example heading ", bs.Badge("New")),
			bs.Heading(2, "Example heading ", bs.Badge("New", bs.WithColor(bs.Success))),
			bs.Heading(3, "Example heading ", bs.PillBadge("New")),
			bs.Heading(4, "Example heading ", bs.Badge("New", bs.WithColor(bs.Secondary))),
			bs.Heading(5, "Example heading ", bs.Badge("New")),
			bs.Heading(6, "Example heading ", bs.Badge("New")),
		), bs.Container().Content(
			bs.Para("Badges may include indicators to highlight new content, and which can be changed with the `Caption` method"),
			bs.CodeBlock(bs.WithColor(bs.Light), mvc.WithClass("p-3"), mvc.WithClass("border", "border-dark-subtle")).Content(
				`bs.Badge("New").Caption("99+")`,
			),
		),
	)
}
