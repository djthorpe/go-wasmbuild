package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type grid struct{ base }

var _ mvc.View = (*grid)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewGrid, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(grid), element, setView)
	})
}

// Grid returns a Carbon 16-column CSS grid container (cds--css-grid).
func Grid(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", setView, mvc.WithClass("cds--css-grid"), args).(*grid)
}

// GridFullWidth returns a full-width CSS grid that stretches edge-to-edge.
func GridFullWidth(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", setView, mvc.WithClass("cds--css-grid", "cds--css-grid--full-width"), args).(*grid)
}

// GridNarrow returns a narrow-gutter CSS grid variant.
func GridNarrow(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", setView, mvc.WithClass("cds--css-grid", "cds--css-grid--narrow"), args).(*grid)
}

// GridCondensed returns a condensed-gutter CSS grid variant (1px gutters).
func GridCondensed(args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", setView, mvc.WithClass("cds--css-grid", "cds--css-grid--condensed"), args).(*grid)
}

func gridColumn(span int, args ...any) *grid {
	return mvc.NewView(new(grid), ViewGrid, "DIV", setView,
		mvc.WithClass("cds--css-grid-column", fmt.Sprintf("cds--col-span-%d", span)), args).(*grid)
}

// Col spans 1 of 16 columns.
func Col(args ...any) *grid {
	return gridColumn(1, args...)
}

// Col2 spans 2 of 16 columns.
func Col2(args ...any) *grid {
	return gridColumn(2, args...)
}

// Col4 spans 4 of 16 columns (one quarter).
func Col4(args ...any) *grid {
	return gridColumn(4, args...)
}

// Col6 spans 6 of 16 columns.
func Col6(args ...any) *grid {
	return gridColumn(6, args...)
}

// Col8 spans 8 of 16 columns (one half).
func Col8(args ...any) *grid {
	return gridColumn(8, args...)
}

// Col10 spans 10 of 16 columns.
func Col10(args ...any) *grid {
	return gridColumn(10, args...)
}

// Col12 spans 12 of 16 columns (three quarters).
func Col12(args ...any) *grid {
	return gridColumn(12, args...)
}

// Col16 spans all 16 columns (full width).
func Col16(args ...any) *grid {
	return gridColumn(16, args...)
}

// ColSpan spans n of 16 columns, where n must be between 1 and 16.
func ColSpan(n int, args ...any) *grid {
	if n < 1 || n > 16 {
		panic(fmt.Sprintf("carbon.ColSpan: n must be 1–16, got %d", n))
	}
	return gridColumn(n, args...)
}
