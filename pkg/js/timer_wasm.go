//go:build js && wasm

package js

import (
	"syscall/js"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Timer wraps a browser setTimeout/setInterval handle.
type Timer struct {
	id       js.Value
	fn       Func
	kind     string // "timeout" or "interval"
	released bool   // guards against double-release of fn
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// SetTimeout calls fn once after the given delay using the browser's setTimeout.
func SetTimeout(delay time.Duration, fn func()) *Timer {
	t := &Timer{kind: "timeout"}
	t.fn = js.FuncOf(func(this js.Value, args []js.Value) any {
		fn()
		t.release()
		return nil
	})
	ms := delay.Milliseconds()
	t.id = js.Global().Call("setTimeout", t.fn, int(ms))
	return t
}

// SetInterval calls fn immediately and then repeatedly at the given interval
// using the browser's setInterval.
func SetInterval(interval time.Duration, fn func()) *Timer {
	t := &Timer{kind: "interval"}
	t.fn = js.FuncOf(func(this js.Value, args []js.Value) any {
		fn()
		return nil
	})
	ms := interval.Milliseconds()
	t.id = js.Global().Call("setInterval", t.fn, int(ms))
	// Fire immediately, matching the native behaviour
	fn()
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Cancel clears the browser timer and releases the associated JS function.
// Safe to call multiple times.
func (t *Timer) Cancel() {
	switch t.kind {
	case "timeout":
		js.Global().Call("clearTimeout", t.id)
	case "interval":
		js.Global().Call("clearInterval", t.id)
	}
	t.release()
}

// release releases the JS function exactly once.
func (t *Timer) release() {
	if !t.released {
		t.released = true
		t.fn.Release()
	}
}
