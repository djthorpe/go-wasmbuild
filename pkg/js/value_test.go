package js

import (
	"testing"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - NewObject

func TestNewObject(t *testing.T) {
	obj := NewObject()

	if !TypeOf(obj).Equal(ObjectProto) {
		t.Errorf("Expected ObjectProto, got %v", TypeOf(obj))
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - NewArray

func TestNewArray(t *testing.T) {
	arr := NewArray()

	if !TypeOf(arr).Equal(ArrayProto) {
		t.Errorf("Expected ArrayProto, got %v", TypeOf(arr))
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - NewMap

func TestNewMap(t *testing.T) {
	m := NewMap()

	if !TypeOf(m).Equal(MapProto) {
		t.Errorf("Expected MapProto, got %v", TypeOf(m))
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - TypeOf

func TestTypeOf(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		expected Proto
	}{
		{
			name:     "Object",
			value:    NewObject(),
			expected: ObjectProto,
		},
		{
			name:     "Array",
			value:    NewArray(),
			expected: ArrayProto,
		},
		{
			name:     "Map",
			value:    NewMap(),
			expected: MapProto,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TypeOf(tt.value)
			if !result.Equal(tt.expected) {
				t.Errorf("TypeOf() = %v, want %v", result, tt.expected)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Global

func TestGlobal(t *testing.T) {
	global := Global()

	// Just verify we can call Global() without panicking
	// The actual behavior differs between WASM and non-WASM builds
	_ = global
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Undefined

func TestUndefined(t *testing.T) {
	undef := Undefined()

	// Verify it's a valid value
	_ = undef
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Null

func TestNull(t *testing.T) {
	null := Null()

	// Verify it's a valid value
	_ = null
}
