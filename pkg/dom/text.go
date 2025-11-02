//go:build !(js && wasm)

package dom

import (
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

func newTextNode(document Document, cdata string) Text {
	node := newNode(document, nil, "#text", TEXT_NODE, cdata)
	return &text{
		node: node,
	}
}

// getNode implements nodeImpl interface
func (t *text) getNode() *node {
	return &t.node
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (t *text) Data() string {
	return t.cdata
}

func (t *text) Length() int {
	return len(t.cdata)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (text *text) AppendChild(Node) Node {
	panic("AppendChild not supported on Text nodes")
}

func (text *text) CloneNode(deep bool) Node {
	return newTextNode(nil, text.cdata)
}

func (text *text) InsertBefore(newNode Node, refNode Node) Node {
	panic("InsertBefore not supported on Text nodes")
}

func (text *text) RemoveChild(Node) {
	panic("RemoveChild not supported on Text nodes")
}
