//go:build js && wasm

package dom

import (
	// Package imports
	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	js.Value
	EventTarget
}

var _ Window = (*window)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func GetWindow() Window {
	return &window{
		Value:       js.Global(),
		EventTarget: js.NewEventTarget(js.Global()),
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (window *window) Document() Document {
	return newDocument(window.Value.Get("document"))
}
