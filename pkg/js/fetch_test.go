package js

import (
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - Fetch Basic

func TestFetch(t *testing.T) {
	// Test that Fetch returns a promise
	promise := Fetch("https://example.com")

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}

	// Verify we can get the value
	_ = promise.Value()
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Fetch with Options

func TestFetch_WithMethod(t *testing.T) {
	// Create a simple option function
	withMethod := func(method string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Method = method
			return nil
		}
	}

	promise := Fetch("https://example.com", withMethod("POST"))

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}
}

func TestFetch_WithHeaders(t *testing.T) {
	// Create a headers option function
	withHeaders := func(headers map[string]string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Headers = headers
			return nil
		}
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	promise := Fetch("https://example.com", withHeaders(headers))

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}
}

func TestFetch_WithBody(t *testing.T) {
	// Create a body option function
	withBody := func(body string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Body = body
			return nil
		}
	}

	promise := Fetch("https://example.com", withBody(`{"key":"value"}`))

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Fetch with Multiple Options

func TestFetch_MultipleOptions(t *testing.T) {
	withMethod := func(method string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Method = method
			return nil
		}
	}

	withHeaders := func(headers map[string]string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Headers = headers
			return nil
		}
	}

	withBody := func(body string) FetchOpt {
		return func(opts *fetchopts) error {
			opts.Body = body
			return nil
		}
	}

	promise := Fetch(
		"https://example.com",
		withMethod("POST"),
		withHeaders(map[string]string{"Content-Type": "application/json"}),
		withBody(`{"data":"test"}`),
	)

	if promise == nil {
		t.Fatal("Expected non-nil promise")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Fetch Option Errors

func TestFetch_OptionError(t *testing.T) {
	// Create an option that returns an error
	errorOpt := func(opts *fetchopts) error {
		return &fetchError{msg: "option error"}
	}

	promise := Fetch("https://example.com", errorOpt)

	if promise == nil {
		t.Fatal("Expected non-nil promise even with error option")
	}

	// The error should be caught via Catch handler
	done := make(chan bool)
	promise.Catch(func(err error) {
		if err == nil {
			t.Error("Expected non-nil error in Catch handler")
		}
		if err.Error() != "option error" {
			t.Errorf("Expected 'option error', got '%s'", err.Error())
		}
		close(done)
	})

	// Wait for the promise to execute
	select {
	case <-done:
		// Success
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout waiting for error to be caught")
	}
}

///////////////////////////////////////////////////////////////////////////////
// HELPER - fetchError for testing

// fetchError is a simple error implementation for testing
type fetchError struct {
	msg string
}

func (e *fetchError) Error() string {
	return e.msg
}
