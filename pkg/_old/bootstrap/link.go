package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type link struct {
	component
}

// Ensure that link implements Component interface
var _ Component = (*link)(nil)

////////////////////////////////////////////////////////////////////////////////
// CONSTRUCTOR

// Link creates a Bootstrap link (anchor) element with optional href and content
// The href parameter sets the link destination, and options can be used to style the link
//
// Example:
//
//	Link("/home", WithColor(PRIMARY)).Append("Home Page")
//	Link("#section").Append("Jump to Section")
//	Link("https://example.com", WithClass("link-offset-2")).Append("External Link")
func Link(href string, opt ...Opt) *link {
	c := newComponent(LinkComponent, dom.GetWindow().Document().CreateElement("A"))

	// Set href attribute
	if href != "" {
		c.root.SetAttribute("href", href)
	}

	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &link{
		component: *c,
	}
}
