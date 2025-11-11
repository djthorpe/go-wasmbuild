package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type icon struct {
	component
}

// Ensure that icon implements Component interface
var _ Component = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Icon creates a new Bootstrap Icon element using the icon font.
// The iconName parameter should be the name of the icon without the "bi-" prefix
// (e.g., "heart-fill", "alarm", "bootstrap").
// See https://icons.getbootstrap.com/ for available icons.
//
// Example:
//
//	Icon("heart-fill", WithColor(DANGER), WithClass("fs-3"))
func Icon(iconName string, opt ...Opt) *icon {
	c := newComponent(IconComponent, dom.GetWindow().Document().CreateElement("I"))

	// Add the base Bootstrap Icons class and the specific icon class, then apply user options
	if err := c.applyTo(c.root, append([]Opt{WithClass("bi"), WithClass("bi-" + iconName)}, opt...)...); err != nil {
		panic(err)
	}

	// Return the icon component
	return &icon{
		component: *c,
	}
}
