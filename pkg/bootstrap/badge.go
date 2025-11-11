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

func Badge(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(badge), ViewBadge, "SPAN", append([]mvc.Opt{mvc.WithClass("badge"), WithColor(Primary)}, opt...)...)
}

func PillBadge(opt ...mvc.Opt) mvc.View {
	return Badge(append(opt, mvc.WithClass("rounded-pill"))...)
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
