//go:build js

package dom

import (
	"fmt"
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

func NewTokenList(values ...string) *tokenlist {
	// Create a temporary DOM element to get access to a real DOMTokenList
	document := js.Global().Get("document")
	element := document.Call("createElement", "div")
	classList := element.Get("classList")

	// Add initial values if provided, filtering out empty strings
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			classList.Call("add", trimmed)
		}
	}

	return &tokenlist{
		classList: classList,
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
	// Empty strings are never in DOMTokenList
	if value == "" {
		return false
	}
	return tokenlist.classList.Call("contains", value).Bool()
}

func (tokenlist *tokenlist) Add(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		// Skip empty strings to avoid DOMTokenList errors
		if value != "" {
			tokenlist.classList.Call("add", value)
		}
	}
}

func (tokenlist *tokenlist) Remove(values ...string) {
	for _, value := range values {
		value = strings.TrimSpace(value)
		// Skip empty strings to avoid DOMTokenList errors
		if value != "" {
			tokenlist.classList.Call("remove", value)
		}
	}
}

func (tokenlist *tokenlist) Toggle(value string, force ...bool) bool {
	value = strings.TrimSpace(value)
	// Skip empty strings to avoid DOMTokenList errors
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
	str := "<DOMTokenList"
	values := tokenlist.Values()
	if len(values) > 0 {
		str += fmt.Sprint(" ", strings.Join(values, ","))
	}
	return str + ">"
}
