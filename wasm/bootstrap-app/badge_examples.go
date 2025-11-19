package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func BadgeExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Badge Examples"),
		bs.HRule(),
		bs.Heading(3, "Badges In Headings", mvc.WithClass("mt-5")), Example(Example_BadgeHeader_001),
	)
}

// Example_BadgeHeader_001 demonstrates badges in headings, with variations
// on color, badge style and borders
func Example_BadgeHeader_001() (mvc.View, string) {
	return bs.Container(
		bs.Heading(1, "Heading with default badge ", bs.Badge("1")),
		bs.Heading(2, "Green badge ", bs.Badge("2", bs.WithColor(bs.Success))),
		bs.Heading(3, "Heading with pill badge ", bs.PillBadge("3")),
		bs.Heading(4, "Heading with secondary badge ", bs.Badge("4", bs.WithColor(bs.Secondary))),
		bs.Heading(5, "Heading with badge with border ", bs.Badge("5", bs.WithBorder(bs.Black))),
		bs.Heading(6, "Heading with danger badge ", bs.Badge("6", bs.WithColor(bs.Danger))),
		bs.Heading(6, "Heading with outlined badge ", bs.Badge("6", bs.WithColor(bs.Light), bs.WithBorder(bs.Danger))),
	), sourcecode()
}
