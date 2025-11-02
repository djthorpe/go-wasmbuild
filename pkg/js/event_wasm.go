//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// EventTarget wraps a JavaScript EventTarget object
type eventtarget struct {
	Value
	listeners map[string][]js.Func
}

type event struct {
	Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEventTarget(value Value) *eventtarget {
	return &eventtarget{
		Value:     value,
		listeners: make(map[string][]js.Func),
	}
}

func NewEvent(value Value) *event {
	return &event{
		Value: value,
	}
}

//////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (e *eventtarget) AddEventListener(eventType string, callback func(Event)) {
	// Panic if e is nil
	if e.IsUndefined() || e.IsNull() {
		panic("EventTarget Value is nil")
	}

	// Initialize event listeners map if needed
	if e.listeners == nil {
		e.listeners = make(map[string][]js.Func)
	}

	// Create a JS function wrapper
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("TODO: Callback for event type:", eventType)
		return nil
	})

	// Store the callback to prevent garbage collection
	e.listeners[eventType] = append(e.listeners[eventType], jsCallback)

	// Add event listener
	e.Value.Call("addEventListener", eventType, jsCallback)
}

// RemoveEventListener removes all event listeners of the specified type and releases their resources
func (e *eventtarget) RemoveEventListener(eventType string) {
	// Panic if Value is nil
	if e.Value.IsUndefined() || e.Value.IsNull() {
		panic("EventTarget Value is nil")
	}

	// Get listeners for the event type, and remove them
	if e.listeners == nil {
		return
	} else if listeners, ok := e.listeners[eventType]; !ok || len(listeners) == 0 {
		return
	} else {
		// Remove each listener from the DOM and release the js.Func
		for _, jsCallback := range listeners {
			e.Value.Call("removeEventListener", eventType, jsCallback)
			jsCallback.Release()
		}

		// Remove from the map
		delete(e.listeners, eventType)
	}
}

func (e *event) Type() string {
	return e.Get("type").String()
}

func (e *event) Target() Value {
	target := e.Get("target")
	if target.IsUndefined() || target.IsNull() {
		return Undefined()
	}
	return target
}
