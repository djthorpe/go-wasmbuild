//go:build !(js && wasm)

package js

import (
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// EventTarget wraps a JavaScript EventTarget object
type eventtarget struct {
	// Empty for non-wasm builds
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEventTarget() *eventtarget {
	return new(eventtarget)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (e *eventtarget) AddEventListener(eventType string, callback func(Event)) {
	// NO-OP
}

func (e *eventtarget) RemoveEventListener(eventType string) {
	// NO-OP
}
