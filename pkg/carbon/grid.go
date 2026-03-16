package carbon

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
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewGrid, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(grid), element, func(self, child mvc.View) {
			self.(*grid).View = child
		})
	})
}

// Grid returns a Carbon 16-column CSS grid container (cds--css-grid).
// Columns placed directly inside span implicitly; use Col* helpers to set span.
func Grid(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--css-grid"), args).(*grid)
}

// GridFullWidth returns a full-width CSS grid that stretches edge-to-edge.
func GridFullWidth(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--css-grid", "cds--css-grid--full-width"), args).(*grid)
}

// GridNarrow returns a narrow-gutter CSS grid variant.
func GridNarrow(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--css-grid", "cds--css-grid--narrow"), args).(*grid)
}

// GridCondensed returns a condensed-gutter CSS grid variant (1px gutters).
func GridCondensed(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--css-grid", "cds--css-grid--condensed"), args).(*grid)
}

// Col returns a column that spans 1 column (auto default).
func Col(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-1"), args).(*grid)
}

// Col2 spans 2 of 16 columns.
func Col2(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-2"), args).(*grid)
}

// Col4 spans 4 of 16 columns (one quarter).
func Col4(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-4"), args).(*grid)
}

// Col6 spans 6 of 16 columns.
func Col6(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-6"), args).(*grid)
}

// Col8 spans 8 of 16 columns (one half).
func Col8(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-8"), args).(*grid)
}

// Col10 spans 10 of 16 columns.
func Col10(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-10"), args).(*grid)
}

// Col12 spans 12 of 16 columns (three quarters).
func Col12(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-12"), args).(*grid)
}

// Col16 spans all 16 columns (full width).
func Col16(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", func(self, child mvc.View) {
		self.(*grid).View = child
	}, mvc.WithClass("cds--col-span-16"), args).(*grid)
}
