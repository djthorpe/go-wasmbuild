package extra

import (
	// Packages
	"fmt"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination_controller struct {
	mvc.Controller

	// Pagination state
	offset, limit, count uint
}

var _ mvc.Controller = (*pagination_controller)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func PaginationController(view mvc.View) *pagination_controller {
	// Check view is a Pagination view
	if view == nil || view.Name() != bs.ViewPagination {
		panic("Invalid view for PaginationController")
	}

	// Create controller and return it
	return mvc.NewController(new(pagination_controller), view).Self().(*pagination_controller)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CONTROLLER

func (c *pagination_controller) Attach(views ...mvc.View) {
	panic("Attach not implemented for PaginationController")
}

func (c *pagination_controller) Detach(views ...mvc.View) {
	panic("Detach not implemented for PaginationController")
}

func (c *pagination_controller) EventListener(event string, view mvc.View) {
	fmt.Println("PaginationController: Event", event, "from view", view.Name())
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Set the current offset, limit and count and update the pagination state
func (c *pagination_controller) Set(offset, limit, count uint) {
	c.offset = offset
	c.limit = limit
	c.count = count
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return the number of pages. Can return 0 if the number of pages
// cannot be determined
func (c *pagination_controller) Num() uint {
	// If we don't have a limit (page size) then return 0
	if c.limit == 0 {
		return 0
	}
	// If we have a count, return number of pages based on limit
	if c.count > 0 {
		if c.count%c.limit == 0 {
			return c.count / c.limit
		} else {
			return (c.count / c.limit) + 1
		}
	}
	// If we don't have a count, then return offset/limit + 1
	return (c.offset / c.limit) + 1
}

// Return the current page number. Can return 0 if the current page
// cannot be determined
func (c *pagination_controller) Cur() uint {
	// If we don't have a limit (page size) then return 0
	if c.limit == 0 {
		return 0
	}
	// Return current page based on offset/limit
	return (c.offset / c.limit) + 1
}
