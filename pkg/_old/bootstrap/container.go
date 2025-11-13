package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type container struct {
	component
}

// Ensure that container implements Component interface
var _ Component = (*container)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new bootstrap container
func Container(opt ...Opt) *container {
	// Create a new component
	c := newComponent(ContainerComponent, dom.GetWindow().Document().CreateElement("DIV"))

	// Prepend default class before user options
	opts := append([]Opt{WithClass("container")}, opt...)

	// Apply options
	if err := c.applyTo(c.root, opts...); err != nil {
		panic(err)
	}

	return &container{component: *c}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (container *container) Element() Element {
	return container.root
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (container *container) Append(children ...any) Component {
	// Append Component, Element or string children to the root element
	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to root
		container.root.AppendChild(child.(Node))
	}

	// Return the container for chaining
	return container
}
