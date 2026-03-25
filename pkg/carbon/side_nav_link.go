package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// SideNavLink returns a <cds-side-nav-link> top-level navigation entry.
func SideNavLink(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-link", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}
