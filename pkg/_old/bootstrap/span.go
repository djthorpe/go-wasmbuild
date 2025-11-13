package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type span struct {
	component
}

// Ensure that span implements Component interface
var _ Component = (*span)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Span creates a new span element
func Span(opt ...Opt) *span {
	c := newComponent(SpanComponent, dom.GetWindow().Document().CreateElement("SPAN"))

	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &span{
		component: *c,
	}
}
