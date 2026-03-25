package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// OrderedList returns an ordered list (<ol>) view.
func OrderedList(args ...any) *list {
	l := mvc.NewView(new(list), ViewList, "OL", setView, args).(*list)
	l.syncPresentation()
	return l
}
