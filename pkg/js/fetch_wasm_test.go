//go:build js && wasm

package js_test

import (
	"strings"
	"testing"

	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

func TestFetch_WASM_DataURL(t *testing.T) {
	// Test fetching a data URL - works without a server
	dataURL := "data:text/plain;base64,SGVsbG8gV29ybGQ="

	resp, err := js.Get(dataURL).Wait()
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

func TestFetch_WASM_JSON(t *testing.T) {
	// Test fetching JSON from a data URL
	// base64 encode: echo -n '{"name":"test","value":42}' | base64
	dataURL := "data:application/json;base64,eyJuYW1lIjoidGVzdCIsInZhbHVlIjo0Mn0="

	resp, err := js.Get(dataURL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	jsonValue, err := response.JSON().Wait()
	if err != nil {
		t.Fatalf("unexpected error parsing JSON: %v", err)
	}

	name := jsonValue.Get("name").String()
	if name != "test" {
		t.Errorf("expected name='test', got '%s'", name)
	}

	value := jsonValue.Get("value").Int()
	if value != 42 {
		t.Errorf("expected value=42, got %d", value)
	}
}

func TestFetch_WASM_Blob(t *testing.T) {
	dataURL := "data:text/plain,BlobTest"

	resp, err := js.Get(dataURL).Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	response := js.ResponseFrom(resp)
	blobValue, err := response.Blob().Wait()
	if err != nil {
		t.Fatalf("unexpected error getting blob: %v", err)
	}

	// Blob should have a size property
	size := blobValue.Get("size").Int()
	if size != 8 { // "BlobTest" = 8 bytes
		t.Errorf("expected blob size 8, got %d", size)
	}

	// Blob should have a type property
	blobType := blobValue.Get("type").String()
	if !strings.Contains(blobType, "text/plain") {
		t.Logf("blob type: %s", blobType)
	}
}
