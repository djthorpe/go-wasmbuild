package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination struct {
	mvc.View
}

var _ mvc.View = (*pagination)(nil)

// PaginationSize controls the row density of the pagination bar.
type PaginationSize string

const (
	PaginationSM PaginationSize = "sm"
	PaginationMD PaginationSize = "md" // default
	PaginationLG PaginationSize = "lg"
)

// Event name constants fired by cds-pagination.
const (
	// PaginationEventPageChanged fires on cds-pagination after the current
	// page is changed. event.detail.{page, pageSize} hold the new values.
	PaginationEventPageChanged = "cds-pagination-changed-current"

	// PaginationEventPageSizeChanged fires after the items-per-page
	// select is changed. event.detail.{page, pageSize} hold the new values.
	PaginationEventPageSizeChanged = "cds-page-sizes-select-changed"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewPagination, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(pagination), element, func(self, child mvc.View) {
			self.(*pagination).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// FACTORY FUNCTIONS

// Pagination returns a <cds-pagination> element. Pass PaginationItem children
// to populate the page-size dropdown, and WithPagination* options to configure.
//
//	cds.Pagination(
//	    cds.PaginationItem(10),
//	    cds.PaginationItem(20),
//	    cds.PaginationItem(50),
//	    cds.WithPaginationTotalItems(150),
//	    cds.WithPaginationPageSize(10),
//	)
func Pagination(args ...any) *pagination {
	return mvc.NewView(new(pagination), ViewPagination, "cds-pagination", func(self, child mvc.View) {
		self.(*pagination).View = child
	}, args).(*pagination)
}

// PaginationItem returns a <cds-select-item> for the page-size dropdown.
// Pass the number of items per page as value.
//
//	cds.PaginationItem(10)
func PaginationItem(pageSize int) dom.Element {
	v := itoa(pageSize)
	el := mvc.HTML("cds-select-item", mvc.WithAttr("value", v), v)
	return el
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithPaginationTotalItems sets the total number of items being paginated.
func WithPaginationTotalItems(n int) mvc.Opt {
	return mvc.WithAttr("total-items", itoa(n))
}

// WithPaginationPage sets the current page (1-based).
func WithPaginationPage(p int) mvc.Opt {
	return mvc.WithAttr("page", itoa(p))
}

// WithPaginationPageSize sets the current page size (items per page).
func WithPaginationPageSize(n int) mvc.Opt {
	return mvc.WithAttr("page-size", itoa(n))
}

// WithPaginationSize sets the visual size of the pagination bar.
func WithPaginationSize(s PaginationSize) mvc.Opt {
	return mvc.WithAttr("size", string(s))
}

// WithPaginationDisabled disables all controls in the pagination bar.
func WithPaginationDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithPaginationPagesUnknown hides the page-number select; use when the total
// number of pages is not known ahead of time.
func WithPaginationPagesUnknown() mvc.Opt {
	return mvc.WithAttr("pages-unknown", "")
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := make([]byte, 0, 20)
	if n < 0 {
		buf = append(buf, '-')
		n = -n
	}
	digits := make([]byte, 0, 20)
	for n > 0 {
		digits = append(digits, byte('0'+n%10))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		buf = append(buf, digits[i])
	}
	return string(buf)
}
