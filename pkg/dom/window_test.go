package dom

import (
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - Window Creation

func TestNewWindow(t *testing.T) {
	window := GetWindow()
	if window == nil {
		t.Fatal("Expected non-nil window")
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Window Properties

func TestWindow_Document(t *testing.T) {
	window := GetWindow()
	if window == nil {
		t.Fatal("Expected non-nil window")
	}

	doc := window.Document()
	// In non-WASM mode, this might be nil
	// In WASM mode, it should return the document
	_ = doc // Just verify it doesn't panic
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Window EventTarget Interface

func TestWindow_AddEventListener(t *testing.T) {
	window := GetWindow()
	if window == nil {
		t.Fatal("Expected non-nil window")
	}

	// Test adding an event listener
	handler := func(e Event) {
		// Handler implementation
	}

	// This should not panic
	window.AddEventListener("test-event", handler)

	// Note: We can't actually trigger the event in tests easily,
	// but we verify the method exists and doesn't panic
}

func TestWindow_RemoveEventListener(t *testing.T) {
	window := GetWindow()
	if window == nil {
		t.Fatal("Expected non-nil window")
	}

	// This should not panic
	window.RemoveEventListener("test-event")
}
