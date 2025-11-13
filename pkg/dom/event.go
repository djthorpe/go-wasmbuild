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

// EventTarget wraps a JavaScript EventTarget object
type eventtarget struct {
	// Empty for non-wasm builds
}

type event struct {
	event  string
	target js.Value
}

var _ EventTarget = (*eventtarget)(nil)
var _ Event = (*event)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEventTarget() *eventtarget {
	return new(eventtarget)
}

func NewEvent(eventType string) *event {
	return &event{
		event:  eventType,
		target: js.Undefined(),
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (e *event) Type() string {
	if e == nil {
		return ""
	}
	return e.event
}

func (e *event) Target() any {
	if e == nil {
		return nil
	}
	return e.target
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (e *eventtarget) AddEventListener(eventType string, callback func(Event)) {
	// NO-OP
}

func (e *eventtarget) RemoveEventListener(eventType string) {
	// NO-OP
}
