//go:build !(js && wasm)

package js

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Value is a wrapper around any value
type Value struct {
	t Proto
	v any
}

// Proto is the type representing the type
type Proto uint

const (
	UndefinedProto Proto = iota
	NullProto
	ArrayProto
	ObjectProto
	MapProto
	TextProto
	CommentProto
	WindowProto
	DocumentProto
	DocumentTypeProto
	ElementProto
	AttrProto
	NodeProto
	EventProto
	CustomEventProto
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - Proto

// Equal returns true if two Proto values are equal.
func (p Proto) Equal(other Proto) bool {
	return p == other
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewObject creates a new empty object.
func NewObject() Value {
	return Value{
		t: ObjectProto,
		v: map[string]any{},
	}
}

// NewArray creates a new  array with the given length.
func NewArray() Value {
	return Value{
		t: ArrayProto,
		v: make([]any, 0),
	}
}

// NewMap creates a new JavaScript Map.
func NewMap() Value {
	return Value{
		t: MapProto,
		v: make(map[any]any),
	}
}

// GetProto returns js.Undefined() for the non-wasm build.
func GetProto(path string) Value {
	return Undefined()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Returns the type of the given Value as a Proto.
func TypeOf(v Value) Proto {
	return v.t
}

// Global returns an undefined global object in the non-wasm build.
func Global() Value {
	return Undefined()
}

// Undefined returns the JavaScript undefined value.
func Undefined() Value {
	return Value{
		t: UndefinedProto,
		v: nil,
	}
}

// Null returns the JavaScript null value.
func Null() Value {
	return Value{
		t: NullProto,
		v: nil,
	}
}

// ValueOf wraps any Go value in a Value for use in non-WASM builds.
// Uses ObjectProto as the type since the native build has no JavaScript
// type system - all wrapped values are treated as generic objects.
// To retrieve the original Go value, use type assertion:
//
//	if str, ok := v.v.(string); ok { ... }
//
// In WASM builds, ValueOf converts Go values to their JavaScript equivalents.
func ValueOf(v any) Value {
	return Value{
		t: ObjectProto,
		v: v,
	}
}

///////////////////////////////////////////////////////////////////////////////
// VALUE METHODS (stubs for non-WASM builds)

// IsUndefined returns true if the value is undefined.
func (v Value) IsUndefined() bool {
	return v.t == UndefinedProto
}

// IsNull returns true if the value is null.
func (v Value) IsNull() bool {
	return v.t == NullProto
}

// Call is a stub that returns undefined in non-WASM builds.
func (v Value) Call(method string, args ...any) Value {
	return Undefined()
}

// New is a stub that returns undefined in non-WASM builds.
func (v Value) New(args ...any) Value {
	return Undefined()
}

// Get is a stub that returns undefined in non-WASM builds.
func (v Value) Get(key string) Value {
	return Undefined()
}

// Set is a stub that does nothing in non-WASM builds.
func (v Value) Set(key string, value any) {
}

// Bool returns the value as a bool (stub for non-WASM builds).
func (v Value) Bool() bool {
	if v.v == nil {
		return false
	}
	if b, ok := v.v.(bool); ok {
		return b
	}
	return false
}

// String returns the value as a string.
func (v Value) String() string {
	if v.v == nil {
		return ""
	}
	if s, ok := v.v.(string); ok {
		return s
	}
	return fmt.Sprint(v.v)
}
