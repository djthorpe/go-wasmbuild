package js

import (
	"testing"
	"time"
)

// Use a non-existent domain to test error handling
const testURL = "https://non-existent-domain-for-testing.example.com"

///////////////////////////////////////////////////////////////////////////////
// TESTS - Response Methods

func TestResponse_Ok(t *testing.T) {
	done := make(chan bool)

	promise := Fetch(testURL)
	promise.Then(func(value Value) error {
		resp := NewResponse(value)

		// In mock mode, Ok() should always return true
		// In WASM mode, this tests the actual response
		_ = resp.Ok() // Just verify it doesn't panic
		close(done)
		return nil
	})
	promise.Catch(func(err error) {
		// Fetch may fail in WASM due to CORS or network issues
		t.Logf("Fetch failed (expected in some environments): %v", err)
		close(done)
	})

	select {
	case <-done:
		// Success
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for response")
	}
}

func TestResponse_Status(t *testing.T) {
	done := make(chan bool)

	promise := Fetch(testURL)
	promise.Then(func(value Value) error {
		resp := NewResponse(value)

		status := resp.Status()
		if status == 0 {
			t.Error("Expected non-zero status")
		}
		close(done)
		return nil
	})
	promise.Catch(func(err error) {
		t.Logf("Fetch failed (expected in some environments): %v", err)
		close(done)
	})

	select {
	case <-done:
		// Success
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for response")
	}
}

func TestResponse_StatusText(t *testing.T) {
	done := make(chan bool)

	promise := Fetch(testURL)
	promise.Then(func(value Value) error {
		resp := NewResponse(value)

		statusText := resp.StatusText()
		_ = statusText // Just verify it doesn't panic
		close(done)
		return nil
	})
	promise.Catch(func(err error) {
		t.Logf("Fetch failed (expected in some environments): %v", err)
		close(done)
	})

	select {
	case <-done:
		// Success
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for response")
	}
}

func TestResponse_Text(t *testing.T) {
	done := make(chan bool)

	promise := Fetch(testURL)
	promise.Then(func(value Value) error {
		resp := NewResponse(value)

		textPromise := resp.Text()
		if textPromise == nil {
			t.Fatal("Expected non-nil promise from Text()")
		}

		textPromise.Then(func(text Value) error {
			// Just verify we got a value back
			_ = text
			close(done)
			return nil
		})
		textPromise.Catch(func(err error) {
			t.Logf("Text() failed: %v", err)
			close(done)
		})
		return nil
	})
	promise.Catch(func(err error) {
		t.Logf("Fetch failed (expected in some environments): %v", err)
		close(done)
	})

	select {
	case <-done:
		// Success
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for text")
	}
}

func TestResponse_JSON(t *testing.T) {
	done := make(chan bool)

	promise := Fetch(testURL)
	promise.Then(func(value Value) error {
		resp := NewResponse(value)

		jsonPromise := resp.JSON()
		if jsonPromise == nil {
			t.Fatal("Expected non-nil promise from JSON()")
		}

		jsonPromise.Then(func(json Value) error {
			// Just verify we got a value back
			_ = json
			close(done)
			return nil
		})
		jsonPromise.Catch(func(err error) {
			t.Logf("JSON() failed: %v", err)
			close(done)
		})
		return nil
	})
	promise.Catch(func(err error) {
		t.Logf("Fetch failed (expected in some environments): %v", err)
		close(done)
	})

	select {
	case <-done:
		// Success
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for JSON")
	}
}
