package bootstrap

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type heading struct {
	component
	level int
}

// Ensure that heading implements Component interface
var _ Component = (*heading)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Heading creates a new heading element with the specified level (1-6)
func Heading(level int, opt ...Opt) *heading {
	if level < 1 || level > 6 {
		panic("heading level must be between 1 and 6")
	}

	c := newComponent(HeadingComponent, dom.GetWindow().Document().CreateElement(fmt.Sprintf("H%d", level)))
	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &heading{
		component: *c,
		level:     level,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

// Level returns the heading level (1-6)
func (h *heading) Level() int {
	return h.level
}
