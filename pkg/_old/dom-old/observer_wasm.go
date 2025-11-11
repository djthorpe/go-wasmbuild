//go:build js

package dom

import (
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type mutationObserver struct {
	observer js.Value
	callback js.Func
}

// Ensure mutationObserver implements MutationObserver interface
var _ dom.MutationObserver = (*mutationObserver)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newMutationObserver(callback func()) *mutationObserver {
	// Create the callback function
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback()
		return nil
	})

	// Create the MutationObserver
	observer := js.Global().Get("MutationObserver").New(jsCallback)

	return &mutationObserver{
		observer: observer,
		callback: jsCallback,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (m *mutationObserver) Observe(target dom.Node, options map[string]interface{}) {
	// Convert options map to JS object
	jsOptions := js.Global().Get("Object").New()
	for key, value := range options {
		// Convert Go slices to JavaScript arrays
		switch v := value.(type) {
		case []interface{}:
			jsArray := js.Global().Get("Array").New(len(v))
			for i, item := range v {
				jsArray.SetIndex(i, item)
			}
			jsOptions.Set(key, jsArray)
		case []string:
			jsArray := js.Global().Get("Array").New(len(v))
			for i, item := range v {
				jsArray.SetIndex(i, item)
			}
			jsOptions.Set(key, jsArray)
		default:
			jsOptions.Set(key, value)
		}
	}

	// Start observing
	m.observer.Call("observe", toJSValue(target), jsOptions)
}

func (m *mutationObserver) Disconnect() {
	m.observer.Call("disconnect")
	m.callback.Release()
}
