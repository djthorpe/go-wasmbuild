package main

import (

	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Toast struct {
	view Component
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewToast() *Toast {
	controller := new(Toast)

	// Create toast with positioning
	toast := bs.Toast(
		bs.WithToastPosition(bs.BOTTOM|bs.END),
		bs.WithColor(bs.DARK),
		bs.WithMargin(bs.MarginAll, 2),
	)
	controller.view = toast

	// Return the controller
	return controller
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Toast) View() Component {
	return this.view
}

func (controller *Toast) Show() {
	controller.view.(ShowHide).Show()
}
