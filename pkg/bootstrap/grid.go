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
	return mvc.NewView(new(grid), ViewGrid, "DIV", mvc.WithClass("row"), args).(*grid)
}

func newGridFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newGridFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(grid), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (grid *grid) SetView(view mvc.View) {
	grid.View = view
}

func (grid *grid) Append(children ...any) mvc.View {
	// Wrap all children in divs with class "col"
	for _, child := range children {
		col := mvc.HTML("DIV", mvc.WithClass("col"))
		col.AppendChild(mvc.NodeFromAny(child))
		grid.View.Append(col)
	}
	return grid
}
