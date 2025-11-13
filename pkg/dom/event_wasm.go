//go:build js && wasm

package dom

import (
	"fmt"

	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// EventTarget wraps a JavaScript EventTarget object
type eventtarget struct {
	js.Value
	listeners map[string][]js.Func
}

type event struct {
	js.Value
}

var _ EventTarget = (*eventtarget)(nil)
var _ Event = (*event)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewEventTarget(value js.Value) *eventtarget {
	return &eventtarget{
		Value:     value,
		listeners: make(map[string][]js.Func),
	}
}

func NewEvent(eventType string) *event {
	return &event{
		Value: js.EventProto.New(eventType),
	}
}

//////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (e *event) Type() string {
	return e.Get("type").String()
}

func (e *event) Target() any {
	target := e.Get("target")
	if target.IsUndefined() || target.IsNull() {
		return nil
	}
	switch {
	case js.TypeOf(target).Equal(js.ElementProto):
		return newElement(target)
	case js.TypeOf(target).Equal(js.WindowProto):
		return newWindow(target)
	default:
		panic(fmt.Sprintf("Unsupported event type %q", js.InstanceName(target)))
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
	jsCallback := js.NewFunc(func(this js.Value, args []js.Value) any {
		callback(&event{Value: args[0]})
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
