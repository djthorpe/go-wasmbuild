package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type para struct {
	component
}

// Ensure that para implements Component interface
var _ Component = (*para)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Para creates a new paragraph (P) element
func Para(opt ...Opt) *para {
	c := newComponent(ParaComponent, dom.GetWindow().Document().CreateElement("P"))

	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &para{
		component: *c,
	}
}
