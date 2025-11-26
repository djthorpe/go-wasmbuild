package bootstrap

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// heading represents a heading element, e.g. H1, H2, etc.
type heading struct {
	mvc.View
}

var _ mvc.View = (*heading)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	headingLevels = map[int]string{
		1: "H1",
		2: "H2",
		3: "H3",
		4: "H4",
		5: "H5",
		6: "H6",
	}
)

func init() {
	mvc.RegisterView(ViewHeading, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(heading), element, func(self, child mvc.View) {
			self.(*heading).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Heading(level int, args ...any) mvc.View {
	tagName, exists := headingLevels[level]
	if !exists {
		panic(fmt.Sprintf("Heading: invalid level %d", level))
	}
	return mvc.NewView(new(heading), ViewHeading, tagName, func(self, child mvc.View) {
		self.(*heading).View = child
	}, args)
}
