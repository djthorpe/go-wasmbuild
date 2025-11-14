package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavBarExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Navbar Examples"), bs.HRule(),
		bs.Heading(3, "Primary", mvc.WithClass("mt-5")), Example(Example_Navbar_001),
		bs.Heading(3, "Dark", mvc.WithClass("mt-5")), Example(Example_Navbar_002),
	)
}

func Example_Navbar_001() (mvc.View, string) {
	return bs.Container(
		bs.NavBar(bs.WithTheme(bs.Dark), mvc.WithClass("m-3")).Label(
			bs.Icon("bootstrap-fill"), mvc.Text(" Bootstrap"),
		).Content(
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("file-text-fill"), " Text")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("border-style"), " Borders")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("badge-4k-fill"), " Badges")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("list-ul"), " Lists")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("emoji-smile-fill"), " Icons")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("toggle-on"), " Buttons")),
		),
	), sourcecode()
}

func Example_Navbar_002() (mvc.View, string) {
	return bs.Container(
		bs.NavBar(bs.WithColor(bs.Dark), bs.WithTheme(bs.Dark), mvc.WithClass("m-3")).Label(
			bs.Icon("bootstrap-fill"), mvc.Text(" Bootstrap"),
		).Content(
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("file-text-fill"), " Text")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("border-style"), " Borders")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("badge-4k-fill"), " Badges")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("list-ul"), " Lists")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("emoji-smile-fill"), " Icons")),
			bs.NavItem("#navbar", mvc.HTML("nobr", bs.Icon("toggle-on"), " Buttons")),
		),
	), sourcecode()
}
