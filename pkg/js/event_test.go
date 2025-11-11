package js

import (
	"reflect"
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

func TestNewEventTargetDefaults(t *testing.T) {
	evt := NewEvent("custom")
	if evt == nil {
		t.Fatal("expected event instance")
	}
	target := evt.Target()
	if !reflect.DeepEqual(target, Undefined()) {
		t.Fatalf("expected undefined target, got %+v", target)
	}
}
