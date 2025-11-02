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
	EventTarget
	document Document
}

var _ Window = (*window)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func NewWindow() Window {
	target := js.NewEventTarget(js.Global())
	document := newNode(target.Get("document")).(Document)
	return &window{target, document}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *window) Document() Document {
	return this.document
}
