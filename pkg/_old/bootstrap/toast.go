//go:build js && wasm

package bootstrap

import (
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type toast struct {
	component
	header Element
	body   Element
}

var _ Component = (*toast)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Toast creates a new Bootstrap toast notification component
// Structure: <div class="toast"><div class="toast-body">...</div></div>
// Header is added only if Header() method is called
func Toast(opt ...Opt) *toast {
	doc := dom.GetWindow().Document()

	// Create main toast div
	div := doc.CreateElement("DIV")
	div.SetAttribute("role", "alert")
	div.SetAttribute("aria-live", "assertive")
	div.SetAttribute("aria-atomic", "true")

	// Create toast body
	body := doc.CreateElement("DIV")
	body.SetAttribute("class", "toast-body")
	div.AppendChild(body)

	// Create component and apply options with toast class
	c := newComponent(ToastComponent, div)
	c.body = body
	if err := c.applyTo(div, append(opt, WithClass("toast"))...); err != nil {
		panic(err)
	}

	return &toast{
		component: *c,
		header:    nil,
		body:      body,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Header adds content to the toast header section.
// Creates the header element if it doesn't exist yet.
// Each child argument is appended to the header.
// Accepts string, Component, or Element children.
// Returns *toast to allow method chaining.
func (t *toast) Header(children ...any) *toast {
	doc := dom.GetWindow().Document()

	// Create header if it doesn't exist
	if t.header == nil {
		t.header = doc.CreateElement("DIV")
		t.header.SetAttribute("class", "toast-header")
		// Insert header before body
		t.root.InsertBefore(t.header, t.body)
	}

	for _, child := range children {
		if component, ok := child.(Component); ok {
			t.header.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			t.header.AppendChild(doc.CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			t.header.AppendChild(node)
		}
	}

	return t
}

// Show displays the toast using Bootstrap's JavaScript API
func (t *toast) Show() {
	jsElement := t.root.(interface{ JSValue() js.Value }).JSValue()
	bootstrap := js.Global().Get("bootstrap")
	if bootstrap.IsUndefined() {
		return
	}

	toastClass := bootstrap.Get("Toast")
	if toastClass.IsUndefined() {
		return
	}

	// Call getOrCreateInstance and show
	instance := toastClass.Call("getOrCreateInstance", jsElement)
	if !instance.IsUndefined() {
		instance.Call("show")
	}
}

// Hide hides the toast using Bootstrap's JavaScript API
func (t *toast) Hide() {
	jsElement := t.root.(interface{ JSValue() js.Value }).JSValue()
	bootstrap := js.Global().Get("bootstrap")
	if bootstrap.IsUndefined() {
		return
	}

	toastClass := bootstrap.Get("Toast")
	if toastClass.IsUndefined() {
		return
	}

	instance := toastClass.Call("getInstance", jsElement)
	if !instance.IsUndefined() {
		instance.Call("hide")
	}
}

// Dispose removes the toast instance
func (t *toast) Dispose() {
	jsElement := t.root.(interface{ JSValue() js.Value }).JSValue()
	bootstrap := js.Global().Get("bootstrap")
	if bootstrap.IsUndefined() {
		return
	}

	toastClass := bootstrap.Get("Toast")
	if toastClass.IsUndefined() {
		return
	}

	instance := toastClass.Call("getInstance", jsElement)
	if !instance.IsUndefined() {
		instance.Call("dispose")
	}
}
