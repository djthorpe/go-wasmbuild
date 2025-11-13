//go:build !(js && wasm)

package js

import (
	"slices"
	"strings"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type tokenlist struct {
	values []string
}

var _ TokenList = (*tokenlist)(nil)

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTokenList(values ...string) *tokenlist {
	// Filter out empty strings to match DOMTokenList behavior
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			filtered = append(filtered, trimmed)
		}
	}
	return &tokenlist{
		values: filtered,
	}
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (tokenlist *tokenlist) Length() int {
	return len(tokenlist.values)
}

func (tokenlist *tokenlist) Value() string {
	return strings.Join(tokenlist.values, " ")
}

func (tokenlist *tokenlist) Values() []string {
	return tokenlist.values
}

func (tokenlist *tokenlist) Contains(value string) bool {
	value = strings.TrimSpace(value)
	// Empty strings are never in DOMTokenList
	if value == "" {
		return false
	}
	return slices.Contains(tokenlist.values, value)
}

func (tokenlist *tokenlist) Add(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		// Skip empty strings to match DOMTokenList behavior
		if value != "" && !slices.Contains(tokenlist.values, value) {
			tokenlist.values = append(tokenlist.values, value)
		}
	}
}

func (tokenlist *tokenlist) Remove(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		// Skip empty strings
		if value == "" {
			continue
		}
		for {
			// Remove all occurrences, not just the first one
			i := slices.Index(tokenlist.values, value)
			if i == -1 {
				break
			}
			tokenlist.values = append(tokenlist.values[:i], tokenlist.values[i+1:]...)
		}
	}
}

func (tokenlist *tokenlist) Toggle(value string, force ...bool) bool {
	// Panic if more than one force value
	if len(force) > 1 {
		panic("Toggle expects zero or one force parameters")
	}

	// Skip empty strings
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}

	if len(force) > 0 {
		if force[0] {
			tokenlist.Add(value)
			return true
		} else {
			tokenlist.Remove(value)
			return false
		}
	} else {
		if tokenlist.Contains(value) {
			tokenlist.Remove(value)
			return false
		} else {
			tokenlist.Add(value)
			return true
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (tokenlist *tokenlist) String() string {
	return strings.Join(tokenlist.Values(), " ")
}
