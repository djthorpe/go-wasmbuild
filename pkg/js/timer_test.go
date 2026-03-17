//go:build !(js && wasm)

package js_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

func TestSetTimeout_Fires(t *testing.T) {
	var fired atomic.Bool
	timer := js.SetTimeout(20*time.Millisecond, func() {
		fired.Store(true)
	})
	defer timer.Cancel()

	time.Sleep(60 * time.Millisecond)
	if !fired.Load() {
		t.Error("expected timeout callback to fire")
	}
}

func TestSetTimeout_Cancel(t *testing.T) {
	var fired atomic.Bool
	timer := js.SetTimeout(50*time.Millisecond, func() {
		fired.Store(true)
	})
	timer.Cancel()

	time.Sleep(100 * time.Millisecond)
	if fired.Load() {
		t.Error("expected timeout callback NOT to fire after Cancel")
	}
}

func TestSetTimeout_CancelTwice(t *testing.T) {
	timer := js.SetTimeout(50*time.Millisecond, func() {})
	// Should not panic
	timer.Cancel()
	timer.Cancel()
}

func TestSetInterval_FiresImmediately(t *testing.T) {
	var count atomic.Int32
	timer := js.SetInterval(1*time.Hour, func() {
		count.Add(1)
	})
	defer timer.Cancel()

	time.Sleep(20 * time.Millisecond)
	if count.Load() != 1 {
		t.Errorf("expected 1 immediate call, got %d", count.Load())
	}
}

func TestSetInterval_FiresRepeatedly(t *testing.T) {
	var count atomic.Int32
	timer := js.SetInterval(20*time.Millisecond, func() {
		count.Add(1)
	})
	defer timer.Cancel()

	time.Sleep(90 * time.Millisecond)
	n := count.Load()
	// Expect at least 3 ticks (1 immediate + ~3 interval ticks in 90ms with 20ms interval)
	if n < 3 {
		t.Errorf("expected at least 3 calls, got %d", n)
	}
}

func TestSetInterval_Cancel(t *testing.T) {
	var count atomic.Int32
	timer := js.SetInterval(20*time.Millisecond, func() {
		count.Add(1)
	})

	time.Sleep(50 * time.Millisecond)
	timer.Cancel()
	snapshot := count.Load()

	time.Sleep(60 * time.Millisecond)
	if count.Load() != snapshot {
		t.Errorf("expected no calls after Cancel, but count grew from %d to %d", snapshot, count.Load())
	}
}

func TestSetInterval_CancelTwice(t *testing.T) {
	timer := js.SetInterval(50*time.Millisecond, func() {})
	// Should not panic
	timer.Cancel()
	timer.Cancel()
}
