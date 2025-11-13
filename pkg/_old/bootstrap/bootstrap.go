package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type app struct {
	Document
	component
}

// Ensure that app implements Application interface
var _ Application = (*app)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	AppComponent name = "wasmbuild-bootstrap-app"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new bootstrap application
func New() *app {
	// Get a document object
	doc := dom.GetWindow().Document()

	// Append a div to the body which will contain the application
	root := doc.CreateElement("DIV")
	root.SetAttribute("id", string(AppComponent))

	// Insert as first child of body (right after opening <body> tag)
	body := doc.Body()
	if firstChild := body.FirstChild(); firstChild != nil {
		body.InsertBefore(root, firstChild)
	} else {
		body.AppendChild(root)
	}

	// Return the document
	return &app{
		Document: doc,
		component: component{
			name: AppComponent,
			root: root,
		},
	}
}
