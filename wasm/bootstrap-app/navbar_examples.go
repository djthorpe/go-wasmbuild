package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavBarExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Navbar Examples"), bs.HRule(),
		bs.Heading(3, "With Label", mvc.WithClass("mt-5")), Example(Example_Navbar_001),
		bs.Heading(3, "Dropdowns", mvc.WithClass("mt-5")), Example(Example_Navbar_003),
		bs.Heading(3, "Event Listeners", mvc.WithClass("mt-5")), Example(Example_Navbar_004),
		bs.Heading(3, "Color and Theme", mvc.WithClass("mt-5")), Example(Example_Navbar_002),
		bs.Heading(3, "Position", mvc.WithClass("mt-5")), Example(Example_Navbar_005),
		bs.Heading(3, "Responsive", mvc.WithClass("mt-5")), Example(Example_Navbar_006),
		bs.Heading(3, "Search", mvc.WithClass("mt-5")), Example(Example_Navbar_007),
	)
}

func Example_Navbar_001() (mvc.View, string) {
	return bs.NavBar(
		"navbar-001", bs.WithTheme(bs.Dark), mvc.WithClass("my-3"),
	).Label(
		bs.Icon("bootstrap-fill"), mvc.Text(" Bootstrap"),
	).Content(
		bs.NavItem("#navbar", bs.Icon("file-text-fill"), " Text"),
		bs.NavItem("#navbar", bs.Icon("border-style"), " Borders"),
		bs.NavItem("#navbar", bs.Icon("badge-4k-fill"), " Badges"),
		bs.NavItem("#navbar", bs.Icon("list-ul"), " Lists"),
		bs.NavItem("#navbar", bs.Icon("emoji-smile-fill"), " Icons"),
		bs.NavItem("#navbar", bs.Icon("toggle-on"), " Buttons"),
	), sourcecode()
}

func Example_Navbar_002() (mvc.View, string) {
	return bs.NavBar(
		"navbar-002", bs.WithTheme(bs.Dark), bs.WithColor(bs.Dark), mvc.WithClass("my-3"),
		bs.NavDropdown(
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Primary)), " Primary", mvc.WithID("primary")),
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Secondary)), " Secondary", mvc.WithID("secondary")),
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Success)), " Success", mvc.WithID("success")),
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Warning)), " Warning", mvc.WithID("warning")),
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Danger)), " Danger", mvc.WithID("danger")),
			bs.NavItem("#navbar", bs.Icon("circle-fill", bs.WithColor(bs.Info)), " Info", mvc.WithID("info")),
		).Label("Colour"),
		bs.NavDropdown(
			bs.NavItem("#navbar", "Light", mvc.WithID("light")),
			bs.NavItem("#navbar", "Dark", mvc.WithID("dark")),
		).Label("Theme"),
	).AddEventListener("click", func(e dom.Event) {
		v := mvc.ViewFromEvent(e)
		if navbar := v.Parent().Parent(); navbar.Name() == bs.ViewNavBar {
			color := bs.Color(v.ID())
			switch color {
			case bs.Light, bs.Dark:
				navbar.Apply(bs.WithTheme(color))
			default:
				navbar.Apply(bs.WithColor(color))
			}
		}
	}), sourcecode()
}

func Example_Navbar_003() (mvc.View, string) {
	return bs.NavBar(
		"navbar-003", bs.WithTheme(bs.Dark), bs.WithColor(bs.Dark), mvc.WithClass("my-3"),
		bs.NavDropdown(
			bs.NavItem("#navbar", bs.Icon("file-text-fill"), " Open", bs.WithColor(bs.Success)),
			bs.NavItem("#navbar", bs.Icon("floppy"), " Save"),
			bs.NavDivider(),
			bs.NavItem("#navbar", bs.Icon("x-circle"), " Quit"),
		).Label("File"),
		bs.NavDropdown(
			bs.NavItem("#navbar", "Action 1"),
			bs.NavItem("#navbar", "Action 2"),
			bs.NavDivider(),
			bs.NavItem("#navbar", "Action 3"),
		).Label("Edit"),
	), sourcecode()
}

func Example_Navbar_004() (mvc.View, string) {
	var response = bs.Para("Click on the navbar items to see event listener output", mvc.WithClass("mt-2"))
	return bs.Container(
		bs.NavBar(
			"navbar-004",
			bs.WithTheme(bs.Dark), bs.WithColor(bs.Dark), mvc.WithClass("my-3"),
			bs.NavItem("#navbar", "Home", mvc.WithID("home")),
			bs.NavDropdown(
				bs.NavItem("#navbar", "Action 1", mvc.WithID("action-1")),
				bs.NavItem("#navbar", "Action 2", mvc.WithID("action-2")),
				bs.NavDivider(),
				bs.NavItem("#navbar", "Action 3", mvc.WithID("action-3")),
			).Label("Edit"),
		).AddEventListener("click", func(e dom.Event) {
			view := mvc.ViewFromEvent(e)
			if view != nil {
				response.Content("Navbar clicked: " + view.ID())
			}
		}),
		response,
	), sourcecode()
}

func Example_Navbar_005() (mvc.View, string) {
	return bs.NavBar(
		"navbar-005",
		bs.WithTheme(bs.Dark), bs.WithColor(bs.Dark),
		bs.NavDropdown(
			bs.NavItem("#navbar", "None", mvc.WithID("nav-position-none")),
			bs.NavItem("#navbar", "Top", mvc.WithID("nav-position-top")),
			bs.NavItem("#navbar", "Bottom", mvc.WithID("nav-position-bottom")),
			bs.NavItem("#navbar", "Sticky Top", mvc.WithID("nav-position-sticky-top")),
			bs.NavItem("#navbar", "Sticky Bottom", mvc.WithID("nav-position-sticky-bottom")),
		).Label("Position").AddEventListener("click", func(e dom.Event) {
			view := mvc.ViewFromEvent(e)
			navbar := view.Parent().Parent()
			switch view.ID() {
			case "nav-position-none":
				navbar.Apply(bs.WithPosition(bs.None))
			case "nav-position-top":
				navbar.Apply(bs.WithPosition(bs.Top))
			case "nav-position-bottom":
				navbar.Apply(bs.WithPosition(bs.Bottom))
			case "nav-position-sticky-top":
				navbar.Apply(bs.WithPosition(bs.Sticky | bs.Top))
			case "nav-position-sticky-bottom":
				navbar.Apply(bs.WithPosition(bs.Sticky | bs.Bottom))
			}
		})), sourcecode()
}

func Example_Navbar_006() (mvc.View, string) {
	return bs.NavBar(
		"navbar-006", bs.WithColor(bs.Light), bs.WithBorder(), bs.WithTheme(bs.Light), mvc.WithClass("my-3"), bs.WithSize(bs.Medium),
	).Label(
		mvc.Text("Collapse/expand the navbar"),
	).Content(
		bs.NavItem("#navbar", bs.Icon("file-text-fill"), " Text"),
		bs.NavItem("#navbar", bs.Icon("border-style"), " Borders"),
		bs.NavItem("#navbar", bs.Icon("badge-4k-fill"), " Badges"),
		bs.NavDropdown(
			bs.NavItem("#navbar", "Action 1", mvc.WithID("action-1"), bs.WithColor(bs.Primary)),
			bs.NavItem("#navbar", "Action 2", mvc.WithID("action-2")),
			bs.NavDivider(),
			bs.NavItem("#navbar", "Action 3", mvc.WithID("action-3")),
		).Label("Edit"),
	), sourcecode()
}

func Example_Navbar_007() (mvc.View, string) {
	// TODO: Appending a form tag needs to be done outside of the UL of NavItems
	return bs.NavBar(
		"navbar-007", bs.WithColor(bs.Light), bs.WithBorder(), bs.WithTheme(bs.Light), mvc.WithClass("my-3"), bs.WithSize(bs.Large),
		bs.NavDropdown(
			bs.NavItem("#navbar", "Action 1", mvc.WithID("action-1"), bs.WithColor(bs.Primary)),
			bs.NavItem("#navbar", "Action 2", mvc.WithID("action-2")),
			bs.NavDivider(),
			bs.NavItem("#navbar", "Action 3", mvc.WithID("action-3")),
		).Label("Edit"),
		bs.Form("search", bs.SearchInput("search", bs.WithPlaceholder("Search")), mvc.WithClass("ms-auto")),
	), sourcecode()
}
