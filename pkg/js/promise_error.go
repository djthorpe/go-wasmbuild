package js

///////////////////////////////////////////////////////////////////////////////
// TYPES

// PromiseError wraps a Promise as an error, allowing promise chains to be flattened.
// When returned from a Then callback, the promise chain will wait for this promise
// to resolve or reject before continuing.
type PromiseError struct {
	promise *Promise
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewPromiseError creates a new PromiseError from a Promise.
func NewPromiseError(p *Promise) *PromiseError {
	return &PromiseError{promise: p}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Error implements the error interface.
func (e *PromiseError) Error() string {
	return "promise error"
}

// Promise returns the wrapped promise.
func (e *PromiseError) Promise() *Promise {
	return e.promise
}
