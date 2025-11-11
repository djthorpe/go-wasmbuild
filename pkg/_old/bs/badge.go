package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type badge struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewBadge = "mvc-bs-badge"
)

func init() {
	RegisterView(ViewBadge, newBadgeFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Badge(opt ...Opt) View {
	opt = append([]Opt{WithClass("badge"), WithColor(PRIMARY)}, opt...)
	return NewView(new(badge), ViewBadge, "SPAN", opt...)
}

func PillBadge(opt ...Opt) View {
	return Badge(append(opt, WithClass("rounded-pill"))...)
}

func newBadgeFromElement(element Element) View {
	if element.TagName() != "SPAN" {
		return nil
	}
	return NewViewWithElement(new(badge), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (badge *badge) SetView(view View) {
	badge.View = view
}
