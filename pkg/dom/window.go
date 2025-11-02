//go:build !(js && wasm)

package dom

import (
	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	EventTarget
	document Document
}

var _ Window = (*window)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	_window = &window{
		EventTarget: js.NewEventTarget(),
		document:    newDocument(nil),
	}
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func NewWindow() Window {
	return _window
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *window) Document() Document {
	return this.document
}
