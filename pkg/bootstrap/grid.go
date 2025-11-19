package bootstrap

import (

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type grid struct {
	BootstrapView
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
	g.BootstrapView.View = mvc.NewView(g, ViewGrid, "DIV", mvc.WithClass("row"), args)
	return g
}

func newGridFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	g := new(grid)
	g.BootstrapView.View = mvc.NewViewWithElement(g, element)
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (grid *grid) Self() mvc.View {
	return grid
}
