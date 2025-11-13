//go:build !(js && wasm)

package js

import (
	"errors"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Promise wraps a JavaScript Promise object (mock for non-WASM builds).
type Promise struct {
	value        Value
	fulfilled    bool
	rejected     bool
	result       any
	thenHandler  func(value Value) any
	catchHandler func(reason Value) any
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise from a Value.
func NewPromise(value Value) *Promise {
	return &Promise{
		value: value,
	}
}

// NewPromiseFromTask creates a new Promise that executes a task in the background.
// The task function receives resolve and reject callbacks to fulfill or reject the promise.
// The task runs immediately but returns a Promise that can be chained with Then/Catch.
func NewPromiseFromTask(task func(resolve func(any), reject func(error))) *Promise {
	p := &Promise{
		value: Undefined(),
	}

	// Create resolve and reject callbacks
	resolve := func(value any) {
		p.fulfilled = true
		p.result = value
		if p.thenHandler != nil {
			resultValue := Value{t: ObjectProto, v: value}
			p.thenHandler(resultValue)
		}
	}

	reject := func(err error) {
		p.rejected = true
		p.result = err
		if p.catchHandler != nil {
			reasonValue := Value{t: ObjectProto, v: err}
			p.catchHandler(reasonValue)
		}
	}

	// Execute the task in background
	go task(resolve, reject)

	return p
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Then registers a callback to be called when the promise is fulfilled.
// The callback receives the resolved value and can return an error.
// If the callback returns an error, it will reject the promise chain.
// If the callback returns a PromiseError, the chain will flatten and wait for that promise.
// Returns a new promise for chaining.
func (p *Promise) Then(onFulfilled func(value Value) error) *Promise {
	newPromise := &Promise{
		value: Undefined(),
	}

	p.thenHandler = func(value Value) any {
		if err := onFulfilled(value); err != nil {
			// Check if it's a PromiseError (flattening)
			if promiseErr, ok := err.(*PromiseError); ok {
				// Chain the wrapped promise to the new promise
				wrappedPromise := promiseErr.Promise()
				wrappedPromise.thenHandler = func(v Value) any {
					newPromise.fulfilled = true
					newPromise.result = v.v
					if newPromise.thenHandler != nil {
						newPromise.thenHandler(v)
					}
					return nil
				}
				wrappedPromise.catchHandler = func(reason Value) any {
					newPromise.rejected = true
					newPromise.result = reason.v
					if newPromise.catchHandler != nil {
						newPromise.catchHandler(reason)
					}
					return nil
				}
				// If wrapped promise is already settled, handle it
				if wrappedPromise.fulfilled && wrappedPromise.thenHandler != nil {
					wrappedPromise.thenHandler(Value{t: ObjectProto, v: wrappedPromise.result})
				} else if wrappedPromise.rejected && wrappedPromise.catchHandler != nil {
					wrappedPromise.catchHandler(Value{t: ObjectProto, v: wrappedPromise.result})
				}
				return nil
			}
			// Reject the new promise for regular errors
			newPromise.rejected = true
			newPromise.result = err
			if newPromise.catchHandler != nil {
				reasonValue := Value{t: ObjectProto, v: err}
				newPromise.catchHandler(reasonValue)
			}
		} else {
			// Fulfill the new promise
			newPromise.fulfilled = true
			newPromise.result = value.v
			if newPromise.thenHandler != nil {
				newPromise.thenHandler(value)
			}
		}
		return nil
	}

	// If already fulfilled, call the handler immediately
	if p.fulfilled && p.thenHandler != nil {
		resultValue := Value{t: ObjectProto, v: p.result}
		p.thenHandler(resultValue)
	}

	return newPromise
}

// Catch registers a callback to be called when the promise is rejected.
// The callback receives the rejection reason as a Go error.
// Returns the promise for chaining.
func (p *Promise) Catch(onRejected func(err error)) *Promise {
	p.catchHandler = func(reason Value) any {
		// Convert Value to error for the callback
		if errVal, ok := reason.v.(error); ok {
			onRejected(errVal)
		} else {
			// Fallback: create a simple error
			onRejected(errors.New("unknown error"))
		}
		return nil
	}

	// If already rejected, call the handler immediately
	if p.rejected && p.catchHandler != nil {
		reasonValue := Value{t: ObjectProto, v: p.result}
		p.catchHandler(reasonValue)
	}

	return p
}

// Value returns the underlying Value.
func (p *Promise) Value() Value {
	return p.value
}
