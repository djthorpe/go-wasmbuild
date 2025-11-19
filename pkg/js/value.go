//go:build !(js && wasm)

package js

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
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - Proto

// Equal returns true if two Proto values are equal.
func (p Proto) Equal(other Proto) bool {
	return p == other
}

// IsUndefined returns true if the value is undefined
func (v Value) IsUndefined() bool {
	return v.t == UndefinedProto
}

// IsNull returns true if the value is null
func (v Value) IsNull() bool {
	return v.t == NullProto
}

// New returns a new value from the prototype
func (v Value) New(args ...any) Value {
	// In non-wasm, we can't really create new JS objects from prototypes easily
	// without a JS engine. For now, return Undefined or a mock.
	return Undefined()
}

// Call calls a method on the value
func (v Value) Call(method string, args ...any) Value {
	return Undefined()
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

// ToString converts a Value to a Go string.
func ToString(v Value) string {
	if v.v == nil {
		return ""
	}
	if s, ok := v.v.(string); ok {
		return s
	}
	return ""
}
