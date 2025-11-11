//go:build !(js && wasm)

package dom

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type location struct {
	href string
}

var _ Location = (*location)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newLocation(href string) Location {
	return &location{href: href}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (l *location) Href() string {
	if l == nil {
		return ""
	}
	return l.href
}

func (l *location) Hash() string {
	if l == nil {
		return ""
	}
	for i := 0; i < len(l.href); i++ {
		if l.href[i] == '#' {
			return l.href[i:]
		}
	}
	return ""
}
