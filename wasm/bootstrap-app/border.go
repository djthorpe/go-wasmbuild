package main

import (
	// Packages
	"strings"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const exampleBorderCode = `
default := bs.Container(bs.WithBorder(), "Default Border")
primary := bs.Container(bs.WithBorder(bs.Primary), "Primary Border")
noborder := bs.Container(bs.WithoutBorder(), "No Border")
`

func BorderExamples() mvc.View {
	return bs.Container(
		bs.Heading(1).Content("Border Examples"),
		bs.HRule(),
		bs.Grid(
			BorderColors(),
			bs.CodeBlock(strings.TrimSpace(exampleBorderCode), bs.WithColor(bs.Light), mvc.WithClass("p-2"), bs.WithBorder(bs.Transparent)),
		),
	)
}

func BorderColors() mvc.View {
	classes := []mvc.Opt{
		mvc.WithAttr("style", "width: 7rem; height: 7rem;"), mvc.WithClass("m-3"), mvc.WithClass("shadow-lg"),
	}
	return bs.Grid(
		bs.Container(bs.WithBorder(), "Default Border", classes),
		bs.Container(bs.WithBorder(bs.Primary), "Primary Border", classes),
		bs.Container(bs.WithBorder(bs.Secondary), "Secondary Border", classes),
		bs.Container(bs.WithBorder(bs.Success), "Success Border", classes),
		bs.Container(bs.WithBorder(bs.Danger), "Danger Border", classes),
		bs.Container(bs.WithBorder(bs.Warning), "Warning Border", classes),
		bs.Container(bs.WithBorder(bs.Info), "Info Border", classes),
		bs.Container(bs.WithBorder(bs.Light), "Light Border", classes),
		bs.Container(bs.WithBorder(bs.Dark), "Dark Border", classes),
		bs.Container(bs.WithBorder(bs.Black), "Black Border", classes),
		bs.Container(bs.WithoutBorder(), "No Border", classes),
	)
}
