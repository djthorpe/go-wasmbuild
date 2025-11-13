package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type badge struct {
	component
}

// Ensure that badge implements Component interface
var _ Component = (*badge)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Badge creates a new bootstrap badge (span element)
// Default badge uses "badge" class. Use WithBackground to set color variant.
func Badge(opt ...Opt) *badge {
	// Create a new component
	c := newComponent(BadgeComponent, dom.GetWindow().Document().CreateElement("SPAN"))

	// Apply options with badge class first
	if err := c.applyTo(c.root, append([]Opt{WithClass("badge")}, opt...)...); err != nil {
		panic(err)
	}

	// Return the component
	return &badge{*c}
}

// PillBadge creates a new bootstrap rounded-pill badge (span element)
func PillBadge(opt ...Opt) *badge {
	// Append the "rounded-pill" class to options
	opt = append(opt, WithClass("rounded-pill"))
	return Badge(opt...)
}
