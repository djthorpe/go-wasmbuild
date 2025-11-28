package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type badge struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewBadge, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(badge), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Badge(args ...any) *badge {
	return mvc.NewView(new(badge), ViewBadge, "SPAN", setView, mvc.WithClass("badge", "position-relative"), WithColor(Primary), args).(*badge)
}

func PillBadge(args ...any) *badge {
	return Badge(args, mvc.WithClass("rounded-pill"))
}
