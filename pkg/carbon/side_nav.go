package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// SideNav returns a <cds-side-nav> shell panel.
func SideNav(args ...any) *navgroup {
	n := mvc.NewView(new(navgroup), ViewNav, templateShellSideNav, setView, args).(*navgroup)
	n.items = navItems(args...)
	return n
}
