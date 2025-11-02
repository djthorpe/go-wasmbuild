//go:build js && wasm

package js

import (
	"errors"
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Promise wraps a JavaScript Promise object.
type Promise struct {
	value js.Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise from a JavaScript value.
func NewPromise(value js.Value) *Promise {
	return &Promise{value: value}
}

// NewPromiseFromTask creates a new Promise that executes a task in the background.
// The task function receives resolve and reject callbacks to fulfill or reject the promise.
// The task runs immediately but returns a Promise that can be chained with Then/Catch.
func NewPromiseFromTask(task func(resolve func(any), reject func(error))) *Promise {
	promiseConstructor := js.Global().Get("Promise")

	executor := js.FuncOf(func(this js.Value, args []js.Value) any {
		resolveJS := args[0]
		rejectJS := args[1]

		// Create Go-friendly resolve/reject functions
		resolve := func(value any) {
			resolveJS.Invoke(value)
		}
		reject := func(err error) {
			rejectJS.Invoke(err.Error())
		}

		// Execute the task
		go task(resolve, reject)

		return nil
	})

	jsPromise := promiseConstructor.New(executor)
	return NewPromise(jsPromise)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Then registers a callback to be called when the promise is fulfilled.
// The callback receives the resolved value and can return an error.
// If the callback returns an error, it will reject the promise chain.
// If the callback returns a PromiseError, the chain will flatten and wait for that promise.
// Returns a new promise for chaining.
func (p *Promise) Then(onFulfilled func(value Value) error) *Promise {
	promiseConstructor := js.Global().Get("Promise")

	newPromise := p.value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		var value js.Value
		if len(args) > 0 {
			value = args[0]
		} else {
			value = js.Undefined()
		}

		// Call the Go callback
		if err := onFulfilled(value); err != nil {
			// Check if it's a PromiseError (flattening)
			if promiseErr, ok := err.(*PromiseError); ok {
				// Return the wrapped promise's JavaScript value
				// JavaScript will automatically flatten the chain
				return promiseErr.Promise().Value()
			}
			// Return a rejected promise for regular errors
			return promiseConstructor.Call("reject", err.Error())
		}

		// Return the value (or undefined if callback succeeded with no return)
		return value
	}))

	return NewPromise(newPromise)
}

// Catch registers a callback to be called when the promise is rejected.
// The callback receives the rejection reason as a Go error.
// Returns the promise for chaining.
func (p *Promise) Catch(onRejected func(err error)) *Promise {
	callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		var errMsg string
		if len(args) > 0 {
			value := args[0]
			// Try to extract error message from JavaScript error object
			if value.Type() == js.TypeObject {
				if msg := value.Get("message"); !msg.IsUndefined() && !msg.IsNull() {
					errMsg = msg.String()
				} else {
					errMsg = value.String()
				}
			} else {
				errMsg = value.String()
			}
		} else {
			errMsg = "undefined error"
		}

		onRejected(errors.New(errMsg))
		return nil
	})

	p.value.Call("catch", callback)
	return p
}

// Value returns the underlying JavaScript value.
func (p *Promise) Value() Value {
	return p.value
}
