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
	b := new(badge)
	b.View = mvc.NewView(
		b, ViewBadge, "SPAN",
		mvc.WithClass("badge", "position-relative"), WithColor(Primary), args,
	)
	return b
}

func PillBadge(args ...any) *badge {
	return Badge(args, mvc.WithClass("rounded-pill"))
}

func newBadgeFromElement(element Element) mvc.View {
	if element.TagName() != "SPAN" {
		return nil
	}
	b := new(badge)
	b.View = mvc.NewViewWithElement(b, element)
	return b
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (badge *badge) Self() mvc.View {
	return badge
}
