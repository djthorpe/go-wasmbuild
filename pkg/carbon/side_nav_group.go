package carbon

import (
	"strings"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// SideNavGroup returns a <cds-side-nav-menu> collapsible navigation group.
func SideNavGroup(title string, args ...any) *navitem {
	n := mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu", setView, append([]any{mvc.WithAttr("title", title), mvc.WithAttr("expanded", "")}, args...)...).(*navitem)
	if strings.TrimSpace(n.Root().Value()) == "" && !n.Root().HasAttribute("value") {
		n.Root().SetValue(title)
		n.Root().SetAttribute("value", title)
	}
	n.items = navItems(args...)
	return n
}
