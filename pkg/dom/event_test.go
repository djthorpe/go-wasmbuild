package dom

import (
	"testing"
)

func TestNewEventType(t *testing.T) {
	evt := NewEvent("wasmbuild")
	if evt == nil {
		t.Fatal("expected event instance")
	}
	if got := evt.Type(); got != "wasmbuild" {
		t.Fatalf("expected event type %q, got %q", "wasmbuild", got)
	}
}
