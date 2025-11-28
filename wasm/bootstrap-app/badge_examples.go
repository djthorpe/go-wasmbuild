package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func BadgeExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("badge_examples.md"),
		bs.HRule(),
		Example(Example_Badges_001),
	)
}

// Example_Badges_001 demonstrates badges in headings, with variations
// on color, badge style and borders
func Example_Badges_001() (mvc.View, string) {
	return bs.Container(
		bs.Heading(4, "Heading with default badge ", bs.Badge("1")),
		bs.Heading(4, "Green badge ", bs.Badge("2", bs.WithColor(bs.Success))),
		bs.Heading(4, "Heading with pill badge ", bs.PillBadge("3")),
		bs.Heading(4, "Heading with secondary badge ", bs.Badge("4", bs.WithColor(bs.Secondary))),
		bs.Heading(4, "Heading with badge with border ", bs.Badge("5", bs.WithBorder(bs.Black))),
		bs.Heading(4, "Heading with danger badge ", bs.Badge("6", bs.WithColor(bs.Danger))),
		bs.Heading(4, "Heading with outlined badge ", bs.Badge("6", bs.WithColor(bs.Light), bs.WithBorder(bs.Danger))),
	), sourcecode()
}
