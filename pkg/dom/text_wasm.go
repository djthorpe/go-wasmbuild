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

type text struct {
	node
}

var _ Text = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newTextNode(value js.Value) Text {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &text{node{value}}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (t *text) Data() string {
	return t.node.Get("data").String()
}

func (t *text) SetData(cdata string) {
	t.node.Set("data", cdata)
}

func (t *text) Length() int {
	return t.node.Get("length").Int()
}
