package main

import (
	"fmt"
	"strconv"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Pagination struct {
	view          Component
	offset, limit int
	count         int
}

// ActiveDisabled interface
type ActiveDisabled interface {
	Component
	Active(...int) []int
	Disabled(...int) []int
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewPagination(model *Model, toast *Toast, limit int) *Pagination {
	controller := new(Pagination)

	// Create the view
	pagination := bs.Pagination()

	// Add a listener for click events
	pagination.AddEventListener("click", func(node Node) {
		component := node.Component()
		if component == nil {
			return
		}
		if component.Name() != string(bs.PaginationItemComponent) {
			return
		}

		// We get the page number from the text of the component
		if page, err := strconv.ParseInt(component.Element().TextContent(), 10, 32); err == nil {
			if controller.Page() != int(page) {
				// Set the page
				controller.SetPage(int(page))

				// Update the pagination view
				controller.Update(model)
			}
		}
	})

	// Set controller options
	controller.view = pagination
	controller.offset = 0
	controller.limit = limit

	// Return the controller
	return controller
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Pagination) View() Component {
	return this.view
}

func (this *Pagination) Offset() int {
	if this.offset < 0 {
		return 0
	}
	return this.offset
}

func (this *Pagination) Limit() int {
	if this.limit <= 0 {
		return 10
	} else {
		return this.limit
	}
}

func (this *Pagination) Page() int {
	page := this.Offset() / this.Limit()
	return page + 1
}

func (this *Pagination) SetPage(page int) {
	if page < 1 {
		this.offset = 0
	} else {
		this.offset = (page - 1) * this.Limit()
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (controller *Pagination) Update(model *Model) {
	// The number of pages is the count of the rows in the model divided by the limit
	numPages := model.Count() / controller.Limit()
	if model.Count()%controller.Limit() != 0 {
		numPages++
	}
	if numPages != controller.count {
		controller.count = numPages

		// Clear existing page numbers
		controller.view.Empty()

		// Append pages
		for i := 1; i <= numPages; i++ {
			controller.view.Append(bs.PaginationItem(fmt.Sprint(i)))
		}
	}

	// Set active page
	controller.view.(ActiveDisabled).Active(controller.Page() - 1)
}
