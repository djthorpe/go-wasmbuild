package bootstrap

import (
	// Packages

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type badge struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewBadge = "mvc-bs-badge"
)

func init() {
	mvc.RegisterView(ViewBadge, newBadgeFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Badge(args ...any) *badge {
	// Return the badge
	return mvc.NewView(
		new(badge), ViewBadge, "SPAN",
		mvc.WithClass("badge", "position-relative"), WithColor(Primary), args,
	).(*badge)
}

func PillBadge(args ...any) *badge {
	return Badge(args, mvc.WithClass("rounded-pill"))
}

func newBadgeFromElement(element Element) mvc.View {
	if element.TagName() != "SPAN" {
		return nil
	}
	return mvc.NewViewWithElement(new(badge), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (badge *badge) SetView(view mvc.View) {
	badge.View = view
}
