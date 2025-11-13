//go:build js && wasm

package dom

import (
	// Package imports
	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type location struct {
	js.Value
}

var _ Location = (*location)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newLocation(value js.Value) Location {
	return &location{Value: value}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (l *location) Href() string {
	if l == nil {
		return ""
	}
	href := l.Value.Get("href")
	if href.IsUndefined() || href.IsNull() {
		return ""
	}
	return href.String()
}

func (l *location) Hash() string {
	if l == nil {
		return ""
	}
	hash := l.Value.Get("hash")
	if hash.IsUndefined() || hash.IsNull() {
		return ""
	}
	return hash.String()
}
