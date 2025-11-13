//go:build js && wasm

package js

import (
	"strings"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type tokenlist struct {
	classList js.Value
}

var _ dom.TokenList = (*tokenlist)(nil)

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new TokenList
func NewTokenList(values ...string) *tokenlist {
	// Create a temporary DOM element to get access to a real DOMTokenList
	document := js.Global().Get("document")
	self := &tokenlist{
		classList: document.Call("createElement", "div").Get("classList"),
	}

	// Add initial values if provided
	self.Add(values...)

	// Return the tokenlist
	return self
}

// Return a TokenList from a js.Value
func GetTokenList(value js.Value) *tokenlist {
	return &tokenlist{
		classList: value,
	}
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (tokenlist *tokenlist) Length() int {
	return tokenlist.classList.Get("length").Int()
}

func (tokenlist *tokenlist) Value() string {
	return tokenlist.classList.Get("value").String()
}

func (tokenlist *tokenlist) Values() []string {
	length := tokenlist.classList.Get("length").Int()
	values := make([]string, length)
	for i := 0; i < length; i++ {
		values[i] = tokenlist.classList.Call("item", i).String()
	}
	return values
}

func (tokenlist *tokenlist) Contains(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	return tokenlist.classList.Call("contains", value).Bool()
}

func (tokenlist *tokenlist) Add(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			tokenlist.classList.Call("add", value)
		}
	}
}

func (tokenlist *tokenlist) Remove(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			tokenlist.classList.Call("remove", value)
		}
	}
}

func (tokenlist *tokenlist) Toggle(value string, force ...bool) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	if len(force) > 0 {
		return tokenlist.classList.Call("toggle", value, force[0]).Bool()
	} else {
		return tokenlist.classList.Call("toggle", value).Bool()
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (tokenlist *tokenlist) String() string {
	return strings.Join(tokenlist.Values(), " ")
}
