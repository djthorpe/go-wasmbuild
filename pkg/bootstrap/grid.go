package bootstrap

import (

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type grid struct {
	mvc.View
}

var _ mvc.View = (*grid)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewGrid, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(grid), element, func(self, child mvc.View) {
			self.(*grid).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Row(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("row"), args).(*grid)
}

func Col(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col"), args).(*grid)
}

func Col2(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col"), args).(*grid)
}

func Col3(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col"), args).(*grid)
}

func Col4(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-4"), args).(*grid)
}

func Col5(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-5"), args).(*grid)
}

func Col6(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-6"), args).(*grid)
}

func Col8(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-8"), args).(*grid)
}

func Col9(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-9"), args).(*grid)
}

func Col10(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-10"), args).(*grid)
}

func Col12(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("col-12"), args).(*grid)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (grid *grid) SetView(view mvc.View) {
	grid.View = view
}
