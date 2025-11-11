//go:build !(js && wasm)

package js

import (
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
	value Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEventTarget() *eventtarget {
	return new(eventtarget)
}

func NewEvent(eventType string) *event {
	return &event{
		value: Value{
			t: EventProto,
			v: map[string]any{
				"type":   eventType,
				"target": Undefined(),
			},
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (e *event) Type() string {
	if e == nil {
		return ""
	}
	if data, ok := e.value.v.(map[string]any); ok {
		if v, ok := data["type"].(string); ok {
			return v
		}
	}
	return ""
}

func (e *event) Target() Value {
	if e == nil {
		return Undefined()
	}
	if data, ok := e.value.v.(map[string]any); ok {
		if v, ok := data["target"].(Value); ok {
			return v
		}
	}
	return Undefined()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (e *eventtarget) AddEventListener(eventType string, callback func(Event)) {
	// NO-OP
}

func (e *eventtarget) RemoveEventListener(eventType string) {
	// NO-OP
}
