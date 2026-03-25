package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// SideNavGroupItem returns a <cds-side-nav-menu-item> for a SideNavGroup.
func SideNavGroupItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu-item", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}
