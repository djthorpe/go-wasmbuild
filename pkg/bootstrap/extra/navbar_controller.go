package extra

import (
	"fmt"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type navbar_controller struct {
	mvc.Controller
}

var _ mvc.Controller = (*navbar_controller)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavbarController(view mvc.View) *navbar_controller {
	if view == nil || view.Name() != bs.ViewNavBar {
		panic("Invalid view for NavbarController")
	}

	// Create controller and return it
	return mvc.NewController(new(navbar_controller), func(self, child mvc.Controller) {
		self.(*navbar_controller).Controller = child
	}, view).Self().(*navbar_controller)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CONTROLLER

func (c *navbar_controller) Attach(views ...mvc.View) {
	panic("Attach not implemented for NavbarController")
}

func (c *navbar_controller) Detach(views ...mvc.View) {
	panic("Detach not implemented for NavbarController")
}

func (c *navbar_controller) EventListener(event string, view mvc.View) {
	fmt.Println("NavbarController: Event", event, "from view", view)
}
