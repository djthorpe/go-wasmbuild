package mvc

import (
	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
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

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new application with a title
func New(title string) *app {
	doc := dom.GetWindow().Document()

	// TODO: Set document title

	// Create the application
	view := new(app)
	view.self = view
	view.name = ViewApp
	view.root = doc.Body()

	return view
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Event listener for the application attaches to the window, not the document
func (app *app) AddEventListener(event string, handler func(Node)) View {
	dom.GetWindow().AddEventListener(event, handler)
	return app
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Create a new DOM element to be attached to a view
func elementFactory(tagName string) Element {
	return dom.GetWindow().Document().CreateElement(tagName)
}

// Create a new DOM text node to be attached to a view
func textFactory(text string) Node {
	return dom.GetWindow().Document().CreateTextNode(text)
}
