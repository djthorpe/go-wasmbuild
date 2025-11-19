package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type grid struct {
	mvc.View
}

var _ mvc.View = (*grid)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewGrid = "mvc-bs-grid"
)

func init() {
	mvc.RegisterView(ViewGrid, newGridFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Grid(args ...any) *grid {
	g := new(grid)
	g.View = mvc.NewView(g, ViewGrid, "DIV", mvc.WithClass("row"), args)
	return g
}

func newGridFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newGridFromElement: invalid tag name %q", tagName))
	}
	g := new(grid)
	g.View = mvc.NewViewWithElement(g, element)
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (grid *grid) Self() mvc.View {
	return grid
}
