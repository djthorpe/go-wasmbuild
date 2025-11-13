package bootstrap

import (
	"fmt"
	"maps"
	"slices"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
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

const (
	ViewHeading = "mvc-bs-heading"
)

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
	mvc.RegisterView(ViewHeading, newHeadingFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Heading(level int, args ...any) mvc.View {
	tagName, exists := headingLevels[level]
	if !exists {
		panic(fmt.Sprintf("Heading: invalid level %d", level))
	}
	return mvc.NewView(new(heading), ViewHeading, tagName, args)
}

func newHeadingFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if !slices.Contains(slices.Collect(maps.Values(headingLevels)), tagName) {
		panic(fmt.Sprintf("newHeadingFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(heading), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (h *heading) SetView(view mvc.View) {
	h.View = view
}
