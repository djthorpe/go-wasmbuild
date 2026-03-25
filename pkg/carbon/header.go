package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Header returns a <cds-header> UI shell header.
func Header(args ...any) *navgroup {
	opts, body, global := splitHeaderArgs(args...)
	n := mvc.NewView(new(navgroup), ViewNav, templateShellHeader, setView, opts).(*navgroup)
	n.items = navItems(body...)
	n.ReplaceSlotChildren("body", body...)
	if global != nil {
		n.ReplaceSlot("global", global)
	}
	bindNavItemClicks(n.items, func(item *navitem) {
		n.SetActive(item)
	})
	return n
}
