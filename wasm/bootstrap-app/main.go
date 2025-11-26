package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	bsextra "github.com/djthorpe/go-wasmbuild/pkg/bootstrap/extra"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application displays examples of Bootstrap components
func main() {
	controller := bsextra.NavbarController(navbar())

	// Run the application
	mvc.New(
		controller.Views()[0],
		router(),
	).Run()
}

func router() mvc.View {
	return mvc.Router(mvc.WithClass("container-fluid", "my-2")).
		Page("#text", TextExamples()).
		Page("#lists", ListExamples()).
		Page("#badges", BadgeExamples()).
		Page("#icons", IconExamples())

}

func navbar() mvc.View {
	return bs.NavBar("main",
		bs.WithPosition(bs.Sticky|bs.Top), bs.WithTheme(bs.Dark), bs.WithSize(bs.Medium),
		bs.NavDropdown(
			bs.NavItem("#text", "Text"),
			bs.NavItem("#lists", "Lists"),
			bs.NavItem("#badges", "Badges"),
			bs.NavItem("#icons", "Icons"),
		).Label("Typography"),
		bs.NavDropdown(
			bs.NavItem("#button", "Buttons"),
			bs.NavItem("#modal", "Modal"),
			bs.NavItem("#alert", "Alerts"),
			bs.NavDivider(),
			bs.NavItem("#toast", "Toasts"),
		).Label("Interactivity"),
		bs.NavDropdown(
			bs.NavItem("#input", "Input"),
		).Label("Forms & Controls"),
		bs.NavDropdown(
			bs.NavItem("#navbar", "Navbar"),
			bs.NavItem("#nav", "Accordion"),
			bs.NavItem("#nav", "Navigation"),
			bs.NavItem("#pagination", "Pagination"),
		).Label("Navigation"),
		bs.NavDropdown(
			bs.NavItem("#border", "Borders"),
			bs.NavItem("#card", "Cards"),
		).Label("Decoration"),
		bs.NavDropdown(
			bs.NavItem("#table", "Tables"),
		).Label("Data"),
		bs.NavItem("https://github.com/djthorpe/go-wasmbuild", bs.Icon("github"), mvc.WithClass("ms-auto")),
	).Label(
		bs.Icon("bootstrap-fill"),
	)
}
