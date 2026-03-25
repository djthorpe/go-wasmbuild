package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// ListItem returns a list item (<li>) view.
func ListItem(args ...any) *list {
	return mvc.NewView(new(list), ViewList, "LI", setView, args).(*list)
}
