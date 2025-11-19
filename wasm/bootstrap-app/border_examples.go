package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func BorderExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Border Examples"),
		bs.HRule(),
		bs.Heading(3, "Border Utilities", mvc.WithClass("mt-5")), Example(Example_Borders_001),
	)
}

// Example_Borders_001 demonstrates borders around containers of
// different colours
func Example_Borders_001() (mvc.View, string) {
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
	), sourcecode()
}
