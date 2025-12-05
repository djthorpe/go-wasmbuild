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
// PUBLIC METHODS

// Counter returns an attribute value function which generates unique IDs
func Counter(label string) string {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	if label == "" {
		label = "id"
	}
	return fmt.Sprintf("%s-%d", label, counter)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithAriaLabel adds an aria-label attribute to a view
func WithAriaLabel(label string) Opt {
	return func(o OptSet) error {
		return WithAttr("aria-label", label)(o)
	}
}
