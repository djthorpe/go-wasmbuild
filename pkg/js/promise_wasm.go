//go:build js && wasm

package js

import (
	"sync"
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Promise wraps a JavaScript Promise and provides chainable handlers.
// In WASM, promises are fully async and non-blocking to avoid deadlocks.
// A Promise can only be executed once (via Run, Wait, or Done). Subsequent
// execution attempts are ignored due to an internal sync.Once guard.
type Promise struct {
	// The executor function that produces the initial value/error
	tryfn func() (Value, error)

	// The underlying JS Promise value (when wrapping an existing JS promise)
	jsPromise Value

	// Handler functions
	thenfn    func(value Value) (Value, error)
	catchfn   func(err error) error
	finallyfn func()

	// Completion callback - the async equivalent of Wait()
	donefn func(value Value, err error)

	// Track js.Func references for cleanup
	funcs []Func

	// Ensure promise executes only once
	once sync.Once
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	promiseProto = js.Global().Get("Promise")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise with the given executor function.
// The executor is the async operation that produces a value or error.
func NewPromise(tryfn func() (Value, error)) *Promise {
	return &Promise{
		tryfn: tryfn,
		funcs: make([]Func, 0, 4),
	}
}

// FromJSPromise wraps an existing JavaScript Promise value.
// Use this when you receive a promise from JS (e.g., fetch()).
func FromJSPromise(jsPromise Value) *Promise {
	return &Promise{
		jsPromise: jsPromise,
		funcs:     make([]Func, 0, 4),
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Then sets the success handler which is called when the promise resolves successfully.
// The handler receives the resolved value and can transform it or return an error.
// Returns the promise for chaining.
func (p *Promise) Then(thenfn func(value Value) (Value, error)) *Promise {
	p.thenfn = thenfn
	return p
}

// Catch sets the error handler which is called when the promise rejects or
// when the Then handler returns an error. The handler can recover by returning nil,
// or propagate/transform the error.
// Returns the promise for chaining.
func (p *Promise) Catch(catchfn func(err error) error) *Promise {
	p.catchfn = catchfn
	return p
}

// Finally sets a handler that is always called after the promise completes,
// regardless of success or failure. Useful for cleanup operations.
// Returns the promise for chaining.
func (p *Promise) Finally(finallyfn func()) *Promise {
	p.finallyfn = finallyfn
	return p
}

// Done sets a completion callback that receives the final value and error,
// and immediately starts executing the promise (non-blocking).
// This is the async alternative to Wait() for WASM - use this to get results.
// Returns the promise for chaining.
func (p *Promise) Done(donefn func(value Value, err error)) *Promise {
	p.donefn = donefn
	p.Run()
	return p
}

// Run executes the promise asynchronously (non-blocking).
// All handlers are called via the JS event loop when the promise settles.
// This returns immediately - use Done() or Wait() to receive the result.
// Can only be called once per Promise instance.
func (p *Promise) Run() {
	p.once.Do(func() {
		p.execute()
	})
}

// Wait executes the promise and blocks until completion.
// Returns the final value and any error that occurred.
//
// WARNING: In WASM, Wait() will DEADLOCK when called from JavaScript event
// handlers (e.g., click handlers, fetch callbacks, setTimeout). The WASM
// runtime is single-threaded - blocking with Wait() prevents the JavaScript
// event loop from running, which means promise callbacks can never execute.
//
// Safe to use: during initialization before event loop starts, or in tests.
// For async operations in event handlers, use Done() instead:
//
//	promise.Done(func(value Value, err error) {
//	    // handle result here
//	})
//
// Can only be called once per Promise instance (subsequent calls return zero values).
func (p *Promise) Wait() (Value, error) {
	var wg sync.WaitGroup
	var result Value
	var resultErr error

	p.once.Do(func() {
		wg.Add(1)
		originalDone := p.donefn
		p.donefn = func(value Value, err error) {
			result = value
			resultErr = err
			if originalDone != nil {
				originalDone(value, err)
			}
			wg.Done()
		}
		p.execute()
	})
	wg.Wait()

	return result, resultErr
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// execute runs the promise chain asynchronously via JS event loop
func (p *Promise) execute() {
	var jsPromise Value

	// If we have a JS promise already, use it; otherwise create one from tryfn
	if !p.jsPromise.IsUndefined() && !p.jsPromise.IsNull() {
		jsPromise = p.jsPromise
	} else if p.tryfn != nil {
		// Create a new JS Promise from the Go executor function
		jsPromise = p.createJSPromise()
	} else {
		// No executor and no JS promise - resolve immediately with undefined
		jsPromise = promiseProto.Call("resolve", js.Undefined())
	}

	// Chain .then() if we have a handler
	if p.thenfn != nil {
		thenFn := js.FuncOf(func(this js.Value, args []js.Value) any {
			value := js.Undefined()
			if len(args) > 0 {
				value = args[0]
			}

			result, err := p.thenfn(value)
			if err != nil {
				// Return a rejected promise to trigger catch
				return promiseProto.Call("reject", err.Error())
			}
			return result
		})
		p.funcs = append(p.funcs, thenFn)
		jsPromise = jsPromise.Call("then", thenFn)
	}

	// Chain .catch() if we have a handler
	if p.catchfn != nil {
		catchFn := js.FuncOf(func(this js.Value, args []js.Value) any {
			var err error
			if len(args) > 0 {
				err = js.Error{Value: args[0]}
			}

			if recoveredErr := p.catchfn(err); recoveredErr != nil {
				// Propagate the error
				return promiseProto.Call("reject", recoveredErr.Error())
			}
			// Error was recovered
			return js.Undefined()
		})
		p.funcs = append(p.funcs, catchFn)
		jsPromise = jsPromise.Call("catch", catchFn)
	}

	// Chain .finally() - always runs
	if p.finallyfn != nil {
		finallyFn := js.FuncOf(func(this js.Value, args []js.Value) any {
			p.finallyfn()
			return js.Undefined()
		})
		p.funcs = append(p.funcs, finallyFn)
		jsPromise = jsPromise.Call("finally", finallyFn)
	}

	// Add completion handler for Done() and cleanup
	successFn := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer p.cleanup()
		if p.donefn != nil {
			value := js.Undefined()
			if len(args) > 0 {
				value = args[0]
			}
			p.donefn(value, nil)
		}
		return js.Undefined()
	})
	p.funcs = append(p.funcs, successFn)

	errorFn := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer p.cleanup()
		if p.donefn != nil {
			var err error
			if len(args) > 0 {
				err = js.Error{Value: args[0]}
			}
			p.donefn(js.Undefined(), err)
		}
		return js.Undefined()
	})
	p.funcs = append(p.funcs, errorFn)

	jsPromise.Call("then", successFn, errorFn)
}

// createJSPromise creates a JS Promise from the Go tryfn executor.
// Uses setTimeout(0) to yield to the JS event loop, preventing blocking.
func (p *Promise) createJSPromise() Value {
	var executorFn Func
	executorFn = js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]

		// Use setTimeout(0) to yield to JS event loop before running executor.
		// This ensures we don't block the main thread.
		setTimeout := js.Global().Get("setTimeout")
		var timeoutFn Func
		timeoutFn = js.FuncOf(func(this js.Value, _ []js.Value) any {
			defer timeoutFn.Release()

			value, err := p.tryfn()
			if err != nil {
				reject.Invoke(err.Error())
			} else {
				resolve.Invoke(value)
			}
			return js.Undefined()
		})
		setTimeout.Invoke(timeoutFn, 0)

		return js.Undefined()
	})
	p.funcs = append(p.funcs, executorFn)

	return promiseProto.New(executorFn)
}

// cleanup releases all js.Func references to prevent memory leaks
func (p *Promise) cleanup() {
	for _, fn := range p.funcs {
		fn.Release()
	}
	p.funcs = nil
}
