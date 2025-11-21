package mvc

import (
	"fmt"
	"os"
	"slices"

	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// A controller reacts to events from one or more views
type Controller interface {
	// Attach one or more views to this controller
	Attach(...View)

	// Detach one or more views from this controller
	Detach(...View)

	// Fire an action based on an event from a view
	EventListener(string, View)
}

type controller struct {
	self  Controller
	views []View
}

var _ Controller = (*controller)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewController(self Controller, view ...View) *controller {
	this := new(controller)
	this.self = self
	this.Attach(view...)
	return this
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CONTROLLER

func (c *controller) Attach(views ...View) {
	// Attach all views to the controller
	for _, view := range views {
		if view == nil || slices.Contains(c.views, view) {
			continue
		}
		c.views = append(c.views, view)

		// Add event listeners for this view
		for _, event := range events[view.Name()] {
			e := event // capture loop variable
			view.AddEventListener(e, func(evt dom.Event) {
				if view := ViewFromEvent(evt); view != nil {
					c.Self().EventListener(e, view)
				}
			})
		}
	}
}

func (c *controller) Detach(views ...View) {
	// Detach all views from the controller
	for _, view := range views {
		if view == nil {
			continue
		}
		for i, v := range c.views {
			if v == view {
				// Remove event listeners for this view
				for _, event := range events[view.Name()] {
					view.RemoveEventListener(event)
				}

				// Remove from slice
				c.views = append(c.views[:i], c.views[i+1:]...)
				break
			}
		}
	}
}

func (c *controller) EventListener(event string, view View) {
	fmt.Fprintf(os.Stderr, "Self %T", c.Self())
	fmt.Fprintf(os.Stderr, "Controller: EventListener not implemented for event %q from view %q\n", event, view.Name())
}

func (c *controller) Self() Controller {
	if c.self == nil {
		return c
	}
	return c.self
}
