package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type rule struct {
	component
}

// Ensure that rule implements Component interface
var _ Component = (*rule)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Rule creates a new horizontal rule (HR) element
func Rule(opt ...Opt) *rule {
	c := newComponent(RuleComponent, dom.GetWindow().Document().CreateElement("HR"))

	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &rule{
		component: *c,
	}
}

// VerticalRule creates a new vertical rule (DIV element with "vr" class)
func VerticalRule(opt ...Opt) *rule {
	c := newComponent(RuleComponent, dom.GetWindow().Document().CreateElement("DIV"))

	if err := c.applyTo(c.root, append([]Opt{WithClass("vr")}, opt...)...); err != nil {
		panic(err)
	}

	return &rule{
		component: *c,
	}
}
