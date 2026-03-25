package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// HeaderNavItem returns a <cds-header-nav-item> link for the header menu bar.
func HeaderNavItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-header-nav-item", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}
