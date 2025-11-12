package bootstrap

import (
	// Packages

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type navbar struct {
	mvc.ViewWithCaption
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavBar = "mvc-bs-navbar"
)

func init() {
	mvc.RegisterView(ViewNavBar, newNavBarFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(args ...any) *navbar {
	// Return the navbar
	return mvc.NewViewEx(
		new(navbar), ViewNavBar, "NAV",
		nil,
		mvc.HTML("DIV", mvc.WithClass("container-fluid")), // body
		nil,
		nil,
		mvc.WithClass("navbar", "bg-primary"), WithTheme(Dark), args,
	).(*navbar)
}

func newNavBarFromElement(element Element) mvc.View {
	if element.TagName() != "NAV" {
		return nil
	}
	return mvc.NewViewWithElement(new(navbar), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (navbar *navbar) SetView(view mvc.View) {
	navbar.ViewWithCaption = view.(mvc.ViewWithCaption)
}
