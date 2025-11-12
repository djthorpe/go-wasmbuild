package bootstrap

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// containers are elements to wrap any content
type container struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewContainer = "mvc-bs-container"
)

func init() {
	mvc.RegisterView(ViewContainer, newContainerFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Container(args ...any) mvc.View {
	opts, content := gatherOpts(mvc.WithClass("container"), args)
	return mvc.NewView(new(container), ViewContainer, "DIV", opts...).Content(content...)
}

func FluidContainer(args ...any) mvc.View {
	opts, content := gatherOpts(mvc.WithClass("container-fluid"), args)
	return mvc.NewView(new(container), ViewContainer, "DIV", opts...).Content(content...)
}

func newContainerFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(container), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (container *container) SetView(view mvc.View) {
	container.View = view
}
