//go:build js && wasm

package bootstrap

import (
	"fmt"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type offcanvas struct {
	component
	header   Element
	body     Element
	instance js.Value // Bootstrap Offcanvas instance
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Offcanvas creates a new bootstrap offcanvas component.
// An offcanvas is a sidebar component that can be toggled via JavaScript to appear
// from the left (start), right (end), top, or bottom edge of the viewport.
//
// Use WithID() to set the ID (required for toggling).
// Use WithPosition() to set the position (START, END, TOP, BOTTOM).
// The offcanvas automatically creates a header and body section.
//
// Example:
//
//	Offcanvas(
//	    WithID("myMenu"),
//	    WithPosition(START),
//	    WithTheme(DARK),
//	).Header("Menu").Append("Content")
func Offcanvas(opt ...Opt) *offcanvas {
	// Create a new component
	c := newComponent(OffcanvasComponent, dom.GetWindow().Document().CreateElement("DIV"))

	// Apply options with offcanvas class and tabindex
	if err := c.applyTo(c.root, append(opt, WithClass("offcanvas"), WithAttribute("tabindex", "-1"))...); err != nil {
		panic(err)
	}

	// Handle position attribute - convert data-position to offcanvas-{position} class
	if position := c.root.GetAttribute("data-position"); position != "" {
		c.root.ClassList().Add("offcanvas-" + position)
		c.root.RemoveAttribute("data-position")
	}

	// Create header
	header := dom.GetWindow().Document().CreateElement("DIV")
	header.SetAttribute("class", "offcanvas-header")

	// Create body
	body := dom.GetWindow().Document().CreateElement("DIV")
	body.SetAttribute("class", "offcanvas-body")

	// Append header and body to root
	c.root.AppendChild(header)
	c.root.AppendChild(body)

	c.body = body

	// Create the offcanvas object
	o := &offcanvas{
		component: *c,
		header:    header,
		body:      body,
	}

	// Add event listeners to manage Bootstrap instance lifecycle
	c.root.AddEventListener("DOMNodeInserted", func(node Node) {
		fmt.Println("DOMNodeInserted fired")
		if o.instance.IsUndefined() {
			jsElement := c.root.(interface{ JSValue() js.Value }).JSValue()
			o.instance = js.Global().Get("bootstrap").Get("Offcanvas").New(jsElement)
		}
	})
	c.root.AddEventListener("DOMNodeRemoved", func(node Node) {
		fmt.Println("DOMNodeRemoved fired")
		if !o.instance.IsUndefined() {
			o.instance.Call("dispose")
			o.instance = js.Undefined()
		}
	})

	return o
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Header adds content to the offcanvas header.
// Accepts string, Component, or Element children.
// Returns *offcanvas to allow method chaining.
// Multiple calls append additional content to the header.
//
// Example:
//
//	offcanvas.Header(Icon("menu"), Heading("Navigation", 5))
func (o *offcanvas) Header(children ...any) *offcanvas {
	// Append children to header
	for _, child := range children {
		if component, ok := child.(Component); ok {
			o.header.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			o.header.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			o.header.AppendChild(node)
		}
	}

	return o
}

// Show makes the offcanvas visible by initializing Bootstrap's Offcanvas instance and calling show().
// This properly handles the backdrop, animations, and accessibility features.
func (o *offcanvas) Show() {
	// Create instance if not yet initialized (fallback if DOMNodeInserted hasn't fired)
	if o.instance.IsUndefined() {
		jsElement := o.root.(interface{ JSValue() js.Value }).JSValue()
		bootstrap := js.Global().Get("bootstrap")
		if bootstrap.IsUndefined() {
			return
		}
		o.instance = bootstrap.Get("Offcanvas").New(jsElement)
	}
	o.instance.Call("show")
}

// Hide makes the offcanvas hidden by using Bootstrap's Offcanvas instance.
func (o *offcanvas) Hide() {
	if !o.instance.IsUndefined() {
		o.instance.Call("hide")
	}
}

// Toggle toggles the visibility of the offcanvas using Bootstrap's Offcanvas instance.
func (o *offcanvas) Toggle() {
	// Create instance if not yet initialized (fallback if DOMNodeInserted hasn't fired)
	if o.instance.IsUndefined() {
		jsElement := o.root.(interface{ JSValue() js.Value }).JSValue()
		bootstrap := js.Global().Get("bootstrap")
		if bootstrap.IsUndefined() {
			return
		}
		o.instance = bootstrap.Get("Offcanvas").New(jsElement)
	}
	o.instance.Call("toggle")
}
