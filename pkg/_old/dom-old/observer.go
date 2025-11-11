//go:build !js

package dom

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type mutationObserver struct {
	callback func()
}

// Ensure mutationObserver implements MutationObserver interface
var _ dom.MutationObserver = (*mutationObserver)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newMutationObserver(callback func()) *mutationObserver {
	return &mutationObserver{
		callback: callback,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (m *mutationObserver) Observe(target dom.Node, options map[string]interface{}) {
	// No-op for non-WASM builds
}

func (m *mutationObserver) Disconnect() {
	// No-op for non-WASM builds
}
