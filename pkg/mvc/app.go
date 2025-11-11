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
	view
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
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

// Create a new application
func New() *app {
	// Create the application
	view := new(app)
	view.self = view
	view.name = ViewApp
	view.root = elementFactory("div")
	if view.root == nil {
		panic("document has no body element")
	}

	// Prepend the application div to the document body
	doc.Body().Prepend(view.root)

	// Return the view
	return view
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Create a new DOM element to be attached to a view
func elementFactory(tagName string) dom.Element {
	return doc.CreateElement(tagName)
}

// Create a new DOM text node to be attached to a view
func textFactory(text string) dom.Node {
	return doc.CreateTextNode(text)
}
