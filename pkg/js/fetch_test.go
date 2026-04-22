//go:build !(js && wasm)

package js_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

func TestFetch_Options(t *testing.T) {
	// Test that options configure the request correctly
	// This works in both native and WASM builds

	t.Run("WithMethod", func(t *testing.T) {
		// Options should not panic
		_ = js.WithMethod("POST")
		_ = js.WithMethod("GET")
		_ = js.WithMethod("PUT")
		_ = js.WithMethod("DELETE")
	})

	t.Run("WithHeader", func(t *testing.T) {
		_ = js.WithHeader("Content-Type", "application/json")
		_ = js.WithHeader("Authorization", "Bearer token")
	})

	t.Run("WithHeaders", func(t *testing.T) {
		_ = js.WithHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer token",
		})
	})

	t.Run("WithBody", func(t *testing.T) {
		_ = js.WithBody("test body")
		_ = js.WithBody([]byte("binary body"))
	})

	t.Run("WithJSON", func(t *testing.T) {
		_ = js.WithJSON(`{"key": "value"}`)
	})
}

func TestFetch_Get(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World"))
	}))
	defer server.Close()

	resp, err := js.Get(server.URL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	if !response.OK() {
		t.Errorf("expected OK response, got status %d", response.Status())
	}

	// Get text content
	textValue, err := response.Text().Wait()
	if err != nil {
		t.Fatalf("unexpected error getting text: %v", err)
	}

	text := textValue.String()
	if text != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", text)
	}
}

func TestFetch_ResponseProperties(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "test-value")
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := js.Get(server.URL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)

	if !response.OK() {
		t.Error("expected OK() to be true")
	}

	if response.Status() != 200 {
		t.Errorf("expected status 200, got %d", response.Status())
	}

	if response.StatusText() == "" {
		t.Error("expected StatusText to be non-empty")
	}

	headers := response.Headers()
	if headers == nil {
		t.Error("expected Headers to be defined")
	}

	if headers["X-Custom-Header"] != "test-value" {
		t.Errorf("expected custom header, got %v", headers)
	}
}

func TestFetch_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}))
	defer server.Close()

	resp, err := js.Get(server.URL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	if response.OK() {
		t.Error("expected non-OK response")
	}
	if response.Status() != http.StatusNotFound {
		t.Fatalf("expected 404 response, got %d", response.Status())
	}
}

func TestFetch_JSONErrorBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"verify JWT: token signature is invalid"}`))
	}))
	defer server.Close()

	resp, err := js.Get(server.URL).Then(func(value js.Value) (js.Value, error) {
		response := js.ResponseFrom(value)
		if response == nil {
			t.Fatal("expected response value")
		}
		if response.OK() {
			t.Fatal("expected non-OK response")
		}
		return value, nil
	}).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	if response.Status() != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", response.Status())
	}
	text, err := response.Text().Wait()
	if err != nil {
		t.Fatalf("unexpected body error: %v", err)
	}
	if got := text.String(); got != `{"message":"verify JWT: token signature is invalid"}` {
		t.Fatalf("unexpected body: %s", got)
	}
}

func TestFetch_WithOptions(t *testing.T) {
	var receivedMethod string
	var receivedHeader string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedMethod = r.Method
		receivedHeader = r.Header.Get("X-Custom")
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := js.Fetch(server.URL,
		js.WithMethod("POST"),
		js.WithHeader("X-Custom", "value"),
	).Wait()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	if !response.OK() {
		t.Errorf("expected OK response, got status %d", response.Status())
	}

	if receivedMethod != "POST" {
		t.Errorf("expected POST method, got %s", receivedMethod)
	}

	if receivedHeader != "value" {
		t.Errorf("expected X-Custom header 'value', got '%s'", receivedHeader)
	}
}

func TestFetch_Post(t *testing.T) {
	var receivedBody string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		receivedBody = string(buf[:n])
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	_, err := js.Post(server.URL, "test body").Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedBody != "test body" {
		t.Errorf("expected body 'test body', got '%s'", receivedBody)
	}
}

func TestFetch_Put(t *testing.T) {
	var receivedMethod string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedMethod = r.Method
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	_, err := js.Put(server.URL, "data").Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedMethod != "PUT" {
		t.Errorf("expected PUT method, got %s", receivedMethod)
	}
}

func TestFetch_Delete(t *testing.T) {
	var receivedMethod string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedMethod = r.Method
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	_, err := js.Delete(server.URL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedMethod != "DELETE" {
		t.Errorf("expected DELETE method, got %s", receivedMethod)
	}
}

func TestFetch_NetworkError(t *testing.T) {
	// Test that network errors are properly returned
	_, err := js.Get("http://localhost:99999/nonexistent").Wait()
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}
