package mvc

import (
	"fmt"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	counter     int
	counterLock sync.Mutex
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithCounter adds an id attribute with an incrementing counter to a view
func WithCounter(label string) Opt {
	return func(o OptSet) error {
		counterLock.Lock()
		defer counterLock.Unlock()
		counter++
		if label == "" {
			label = "id"
		}
		return WithAttr("id", fmt.Sprintf("%s-%d", label, counter))(o)
	}
}

// WithAriaLabel adds an aria-label attribute to a view
func WithAriaLabel(label string) Opt {
	return func(o OptSet) error {
		return WithAttr("aria-label", label)(o)
	}
}
