//go:build !(js && wasm)

package dom

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	EventTarget
	document Document
	location Location
}

var _ Window = (*window)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	_window = &window{
		EventTarget: NewEventTarget(),
		document:    newHTMLDocument(nil),
		location:    newLocation("about:blank"),
	}
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func GetWindow() Window {
	return _window
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (window *window) Document() Document {
	return window.document
}

func (window *window) Location() Location {
	return window.location
}
