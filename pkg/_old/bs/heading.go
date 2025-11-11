package bs

import (
	"fmt"
	"maps"
	"slices"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// heading represents a heading element, e.g. H1, H2, etc.
type heading struct {
	View
}

var _ View = (*heading)(nil)

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
	RegisterView(ViewHeading, newHeadingFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Heading(level int, opt ...Opt) View {
	if tagName, exists := headingLevels[level]; !exists {
		panic(fmt.Sprintf("Heading: invalid level %d", level))
	} else {
		return NewView(new(heading), ViewHeading, tagName, opt...)
	}
}

func newHeadingFromElement(element Element) View {
	tagName := element.TagName()
	if !slices.Contains(slices.Collect(maps.Values(headingLevels)), tagName) {
		panic(fmt.Sprintf("newHeadingFromElement: invalid tag name %q", tagName))
	}
	return NewViewWithElement(new(heading), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (h *heading) SetView(view View) {
	h.View = view
}
