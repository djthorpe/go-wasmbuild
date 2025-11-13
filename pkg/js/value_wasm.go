//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Value is an alias for js.Value for convenience.
type Value = js.Value

// Func is an alias for js.Func representing a JavaScript function.
type Func = js.Func

// Proto is an alias for js.Value representing a JavaScript prototype object.
type Proto = js.Value

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	// Constructors
	ArrayProto        Proto = js.Global().Get("Array")
	ObjectProto       Proto = js.Global().Get("Object")
	MapProto          Proto = js.Global().Get("Map")
	TextProto         Proto = js.Global().Get("Text")
	CommentProto      Proto = js.Global().Get("Comment")
	WindowProto       Proto = js.Global().Get("Window")
	DocumentProto     Proto = js.Global().Get("HTMLDocument")
	DocumentTypeProto Proto = js.Global().Get("DocumentType")
	ElementProto      Proto = js.Global().Get("HTMLElement")
	AttrProto         Proto = js.Global().Get("Attr")
	NodeProto         Proto = js.Global().Get("Node")
	EventProto        Proto = js.Global().Get("Event")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewObject creates a new empty JavaScript object.
func NewObject() js.Value {
	return ObjectProto.New()
}

// NewArray creates a new JavaScript array
func NewArray() js.Value {
	return ArrayProto.New()
}

// NewMap creates a new JavaScript Map.
func NewMap() js.Value {
	return MapProto.New()
}

// NewFunc creates a new function Value.
func NewFunc(fn func(this Value, args []Value) any) Func {
	return js.FuncOf(fn)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Returns the type of the given Value as a Proto.
func TypeOf(v Value) Proto {
	var proto = v
	if v.IsUndefined() || v.IsNull() {
		return js.Undefined()
	}

	for {
		proto = ObjectProto.Call("getPrototypeOf", proto)
		if proto.IsNull() || proto.IsUndefined() {
			panic(fmt.Sprint("Unknown constructor"))
		}

		// Check if this prototype matches any known types
		switch {
		case proto.Equal(EventProto.Get("prototype")):
			return EventProto
		case proto.Equal(DocumentProto.Get("prototype")):
			return DocumentProto
		case proto.Equal(ElementProto.Get("prototype")):
			return ElementProto
		case proto.Equal(TextProto.Get("prototype")):
			return TextProto
		case proto.Equal(CommentProto.Get("prototype")):
			return CommentProto
		case proto.Equal(WindowProto.Get("prototype")):
			return WindowProto
		case proto.Equal(DocumentTypeProto.Get("prototype")):
			return DocumentTypeProto
		case proto.Equal(AttrProto.Get("prototype")):
			return AttrProto
		case proto.Equal(NodeProto.Get("prototype")):
			return NodeProto
		case proto.Equal(MapProto.Get("prototype")):
			return MapProto
		case proto.Equal(ArrayProto.Get("prototype")):
			return ArrayProto
		case proto.Equal(ObjectProto.Get("prototype")):
			return ObjectProto
		}
	}
}

// Returns the instance name
func InstanceName(v Value) string {
	if v.IsUndefined() {
		return "undefined"
	} else if v.IsNull() {
		return "null"
	}
	return v.Get("constructor").Get("name").String()
}

// Global returns the JavaScript global object (window in browsers, global in Node.js).
func Global() Value {
	return js.Global()
}

// Undefined returns the JavaScript undefined value.
func Undefined() Value {
	return js.Undefined()
}

// Null returns the JavaScript null value.
func Null() Value {
	return js.Null()
}

// ToString converts a JavaScript Value to a Go string.
// This works for JavaScript strings and other types that can be converted to string.
func ToString(v Value) string {
	return v.String()
}
