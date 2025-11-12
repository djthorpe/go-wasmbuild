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

func Badge(args ...any) mvc.View {
	opts, content := gatherOpts(mvc.WithClass("badge"), WithColor(Primary), args)
	return mvc.NewView(new(badge), ViewBadge, "SPAN", opts...).Content(content...)
}

func gatherOpts(args ...any) ([]mvc.Opt, []any) {
	var opts []mvc.Opt
	var content []any
	for _, arg := range args {
		switch v := arg.(type) {
		case []any:
			o, c := gatherOpts(v...)
			opts = append(opts, o...)
			content = append(content, c...)
		case mvc.Opt:
			opts = append(opts, v)
		default:
			content = append(content, v)
		}
	}
	return opts, content
}

func PillBadge(args ...any) mvc.View {
	return Badge(append(args, mvc.WithClass("rounded-pill"))...)
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
