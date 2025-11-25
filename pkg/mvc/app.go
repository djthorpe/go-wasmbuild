// Package mvc provides a thin model-view-controller layer for building
// declarative WASM user interfaces using the go-wasmbuild DOM wrappers.
package mvc

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	impl "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// app is a simple app view
type app struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// ViewApp is the registered component name used for the root application
	// container within the MVC system.
	ViewApp = "mvc-app"
)

func init() {
	RegisterView(ViewApp, nil)
}

var (
	doc = impl.GetWindow().Document()
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// New creates an empty application root, attaches it to the document body and
// returns the view so callers can begin composing content.
func New(args ...any) *app {
	view := NewView(new(app), ViewApp, "DIV", func(self, child View) {
		self.(*app).View = child
	}, args...).Self().(*app)

	// Attach the view to the document body
	doc.Body().Prepend(view.Root())

	// Return the view
	return view
}

func (a *app) Run() {
	select {}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// elementFactory creates a new DOM element to be attached to a view.
func elementFactory(tagName string) dom.Element {
	return doc.CreateElement(tagName)
}

// textFactory creates a new DOM text node to be attached to a view.
func textFactory(text string) dom.Node {
	return doc.CreateTextNode(text)
}
