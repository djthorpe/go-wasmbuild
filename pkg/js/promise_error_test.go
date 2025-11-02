package js

import (
	"fmt"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - PromiseError (Promise Flattening)

func TestPromiseError_Flattening(t *testing.T) {
	var result string

	// Create a promise that returns another promise via PromiseError
	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		time.Sleep(10 * time.Millisecond)
		resolve("first")
	})

	promise.Then(func(v Value) error {
		// Return a PromiseError that wraps an inner promise
		return NewPromiseError(NewPromiseFromTask(func(resolve func(any), reject func(error)) {
			time.Sleep(10 * time.Millisecond)
			resolve("second")
		}).Then(func(v Value) error {
			result = "inner promise resolved"
			return nil
		}))
	}).Catch(func(err error) {
		t.Errorf("Should not catch error: %v", err)
	})

	// Wait for promise chain to complete
	time.Sleep(100 * time.Millisecond)

	if result != "inner promise resolved" {
		t.Errorf("Expected inner promise to resolve, got: %s", result)
	}
}

func TestPromiseError_ErrorPropagation(t *testing.T) {
	called := false
	var caughtError error

	// Create a promise that returns a promise that rejects
	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		time.Sleep(10 * time.Millisecond)
		resolve("first")
	})

	promise.Then(func(v Value) error {
		// Return a PromiseError that wraps a promise that will reject
		return NewPromiseError(NewPromiseFromTask(func(resolve func(any), reject func(error)) {
			time.Sleep(10 * time.Millisecond)
			reject(fmt.Errorf("inner error"))
		}))
	}).Catch(func(err error) {
		called = true
		caughtError = err
	})

	// Wait for promise chain to complete
	time.Sleep(100 * time.Millisecond)

	if !called {
		t.Error("Expected catch to be called")
	}

	if caughtError == nil {
		t.Error("Expected error to be caught")
	}
}

func TestPromiseError_ErrorMethod(t *testing.T) {
	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		resolve("test")
	})

	promiseErr := NewPromiseError(promise)

	if promiseErr.Error() != "promise error" {
		t.Errorf("Expected 'promise error', got: %s", promiseErr.Error())
	}

	if promiseErr.Promise() != promise {
		t.Error("Expected Promise() to return the wrapped promise")
	}
}
