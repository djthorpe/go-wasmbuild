package js

import (
	"fmt"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - Promise Creation

func TestNewPromise(t *testing.T) {
	value := NewObject()
	promise := NewPromise(value)

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}

	// Verify we can get the value back
	promiseValue := promise.Value()
	if !TypeOf(promiseValue).Equal(TypeOf(value)) {
		t.Errorf("Expected promise value type %v, got %v", TypeOf(value), TypeOf(promiseValue))
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - NewPromiseFromTask

func TestNewPromiseFromTask_Resolve(t *testing.T) {
	called := false
	var receivedValue any

	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		// Simulate some async work
		time.Sleep(10 * time.Millisecond)
		resolve("success")
	})

	promise.Then(func(v Value) error {
		called = true
		receivedValue = "resolved"
		return nil
	})

	// Wait for the async task to complete
	time.Sleep(50 * time.Millisecond)

	if !called {
		t.Error("Expected Then callback to be called after resolve")
	}

	if receivedValue != "resolved" {
		t.Errorf("Expected receivedValue to be 'resolved', got %v", receivedValue)
	}
}

func TestNewPromiseFromTask_Reject(t *testing.T) {
	called := false
	var receivedError error

	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		// Simulate some async work
		time.Sleep(10 * time.Millisecond)
		reject(fmt.Errorf("test error"))
	})

	promise.Catch(func(err error) {
		called = true
		receivedError = err
	})

	// Wait for the async task to complete
	time.Sleep(50 * time.Millisecond)

	if !called {
		t.Error("Expected Catch callback to be called after reject")
	}

	if receivedError == nil {
		t.Error("Expected receivedError to be non-nil")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Then

func TestPromise_Then(t *testing.T) {
	called := false

	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		time.Sleep(10 * time.Millisecond)
		resolve("test value")
	})

	promise.Then(func(v Value) error {
		called = true
		return nil
	})

	// Wait for async execution
	time.Sleep(50 * time.Millisecond)

	if !called {
		t.Error("Expected Then callback to be called")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Catch

func TestPromise_Catch(t *testing.T) {
	called := false

	promise := NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		time.Sleep(10 * time.Millisecond)
		reject(fmt.Errorf("test error"))
	})

	promise.Catch(func(err error) {
		called = true
	})

	// Wait for async execution
	time.Sleep(50 * time.Millisecond)

	if !called {
		t.Error("Expected Catch callback to be called")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Value Method

func TestPromise_Value(t *testing.T) {
	obj := NewObject()
	promise := NewPromise(obj)

	// Verify Value() returns the underlying value
	val := promise.Value()
	if !TypeOf(val).Equal(ObjectProto) {
		t.Errorf("Expected ObjectProto, got %v", TypeOf(val))
	}
}
