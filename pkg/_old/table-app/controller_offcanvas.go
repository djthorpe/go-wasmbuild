package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Offcanvas struct {
	view Component
}

// ShowHide interface
type ShowHide interface {
	Show()
	Hide()
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewOffcanvas() *Offcanvas {
	controller := new(Offcanvas)

	// Create the view
	view := bs.Offcanvas(
		bs.WithID("employeeDetails"),
		bs.WithPosition(bs.END),
	)

	// Set the header
	view.Header(
		bs.Heading(5).Append("Employee Details"),
		bs.CloseButton(bs.WithAttribute("data-bs-dismiss", "offcanvas")),
	)

	// Set the view in the controller
	controller.view = view

	// Return the controller
	return controller
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (controller *Offcanvas) View() Component {
	return controller.view
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (controller *Offcanvas) Show() {
	controller.view.(ShowHide).Show()
}
