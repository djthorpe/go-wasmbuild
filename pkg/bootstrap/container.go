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
	c := new(container)
	c.View = mvc.NewView(c, ViewContainer, "DIV", mvc.WithClass("container"), args)
	return c
}

func FluidContainer(args ...any) mvc.View {
	c := new(container)
	c.View = mvc.NewView(c, ViewContainer, "DIV", mvc.WithClass("container-fluid"), args)
	return c
}

func newContainerFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	c := new(container)
	c.View = mvc.NewViewWithElement(c, element)
	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (container *container) Self() mvc.View {
	return container
}
