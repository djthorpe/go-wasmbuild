package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application displays examples of Bootstrap components
func main() {
	mvc.New().Content(
		bs.NavBar("main",
			bs.WithPosition(bs.Sticky|bs.Top), bs.WithTheme(bs.Dark), bs.WithSize(bs.Medium),
			bs.NavDropdown(
				bs.NavItem("#text", "Text"),
				bs.NavItem("#list", "Lists"),
				bs.NavItem("#badge", "Badges"),
				bs.NavItem("#icon", "Icons"),
			).Label("Typography"),
			bs.NavDropdown(
				bs.NavItem("#button", "Buttons"),
				bs.NavItem("#modal", "Modal"),
				bs.NavItem("#alert", "Alerts & Toasts"),
			).Label("Interactivity"),
			bs.NavDropdown(
				bs.NavItem("#input", "Input"),
			).Label("Forms & Controls"),
			bs.NavDropdown(
				bs.NavItem("#navbar", "Navbar"),
				bs.NavItem("#nav", "Accordion"),
				bs.NavItem("#nav", "Navigation"),
				bs.NavItem("#nav", "Pagination"),
			).Label("Navigation"),
			bs.NavDropdown(
				bs.NavItem("#border", "Borders"),
				bs.NavItem("#card", "Cards"),
			).Label("Decoration"),
			bs.NavDropdown(
				bs.NavItem("#table", "Tables"),
			).Label("Data"),
			bs.NavItem("https://github.com/djthorpe/go-wasmbuild", bs.Icon("github", mvc.WithClass("me-1")), "GitHub"),
		).Label(
			bs.Icon("bootstrap-fill"),
		),
		mvc.Router(mvc.WithClass("container-fluid", "my-2")).
			Page("#text", TextExamples()).
			Page("#border", BorderExamples()).
			Page("#badge", BadgeExamples()).
			Page("#list", ListExamples()).
			Page("#icon", IconExamples()).
			Page("#button", ButtonExamples()).
			Page("#card", CardExamples()).
			Page("#modal", ModalExamples()).
			Page("#input", InputExamples()).
			Page("#tooltips", TooltipExamples()).
			Page("#progress", ProgressExamples()).
			Page("#navbar", NavBarExamples()).
			Page("#nav", NavExamples()).
			Page("#table", TableExamples()).
			Page("#alert", AlertExamples()),
	)

	// Wait
	select {}
}
