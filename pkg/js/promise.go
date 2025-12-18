//go:build !(js && wasm)

package js

import "sync"

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Promise represents an asynchronous operation that can be chained with
// Then, Catch, and Finally handlers, similar to JavaScript Promises.
// A Promise can only be executed once (via Run, Wait, or Done). Subsequent
// execution attempts are ignored due to an internal sync.Once guard.
type Promise struct {
	tryfn     func() (Value, error)
	thenfn    func(value Value) (Value, error)
	catchfn   func(err error) error
	finallyfn func()
	donefn    func(value Value, err error)
	jsPromise Value // For API compatibility with WASM (unused in native)
	once      sync.Once
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromise creates a new Promise with the given executor function.
// The executor is the async operation that produces a value or error.
func NewPromise(tryfn func() (Value, error)) *Promise {
	return &Promise{
		tryfn: tryfn,
	}
}

// FromJSPromise wraps an existing JavaScript Promise value.
// In native Go, this creates a promise that resolves immediately with the value.
// This exists for API compatibility with the WASM build.
func FromJSPromise(jsPromise Value) *Promise {
	return &Promise{
		jsPromise: jsPromise,
		tryfn: func() (Value, error) {
			return jsPromise, nil
		},
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
// and immediately starts executing the promise asynchronously.
// In native Go, Wait() is preferred, but Done() provides API compatibility with WASM.
// Returns the promise for chaining.
func (p *Promise) Done(donefn func(value Value, err error)) *Promise {
	p.donefn = donefn
	p.Run()
	return p
}

// Run executes the promise asynchronously.
// Handlers are called in order: tryfn -> thenfn (on success) or catchfn (on error) -> finallyfn.
// Can only be called once per Promise instance.
func (p *Promise) Run() {
	p.once.Do(func() {
		go p.execute()
	})
}

// Wait executes the promise synchronously and blocks until completion.
// Returns the final value and any error that occurred.
// This is useful for testing or when you need the result immediately.
// Can only be called once per Promise instance (subsequent calls return zero values).
func (p *Promise) Wait() (Value, error) {
	var result Value
	var resultErr error
	var wg sync.WaitGroup

	p.once.Do(func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, resultErr = p.execute()
		}()
	})
	wg.Wait()

	return result, resultErr
}

// execute runs the promise chain and returns the final value and error
func (p *Promise) execute() (Value, error) {
	var err error
	value := Undefined() // Initialize to Undefined for explicit zero-value behavior

	// Always call finally and done at the end
	defer func() {
		if p.finallyfn != nil {
			p.finallyfn()
		}
		if p.donefn != nil {
			p.donefn(value, err)
		}
	}()

	// Execute the try function if set
	if p.tryfn != nil {
		value, err = p.tryfn()
	}

	// If successful and we have a then handler, call it
	if err == nil && p.thenfn != nil {
		value, err = p.thenfn(value)
	}

	// If there's an error and we have a catch handler, call it
	if err != nil && p.catchfn != nil {
		err = p.catchfn(err)
	}

	return value, err
}
