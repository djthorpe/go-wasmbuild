//go:build !(js && wasm)

package js

import "time"

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Timer fires a callback once (after a delay) or repeatedly (on an interval).
// Mirrors the browser's setTimeout/setInterval API.
type Timer struct {
	fn     func()
	done   chan struct{}
	ticker *time.Ticker
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// SetTimeout calls fn once after the given delay.
func SetTimeout(delay time.Duration, fn func()) *Timer {
	t := &Timer{fn: fn, done: make(chan struct{})}
	go func() {
		select {
		case <-time.After(delay):
			t.fn()
		case <-t.done:
		}
	}()
	return t
}

// SetInterval calls fn immediately and then repeatedly at the given interval
// until Cancel is called.
func SetInterval(interval time.Duration, fn func()) *Timer {
	t := &Timer{
		fn:     fn,
		done:   make(chan struct{}),
		ticker: time.NewTicker(interval),
	}
	go func() {
		t.fn()
		for {
			select {
			case <-t.ticker.C:
				t.fn()
			case <-t.done:
				return
			}
		}
	}()
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Cancel stops the timer. Safe to call multiple times.
func (t *Timer) Cancel() {
	select {
	case <-t.done:
		// already cancelled
	default:
		close(t.done)
		if t.ticker != nil {
			t.ticker.Stop()
		}
	}
}
