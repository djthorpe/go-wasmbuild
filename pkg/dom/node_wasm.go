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

type node struct {
	js.Value
}

/*

type document struct {
	node
}

type doctype struct {
	node
}

type element struct {
	node
}

type text struct {
	node
}

type comment struct {
	node
}

type attr struct {
	node
}
*/

var _ Node = (*node)(nil)

//var _ Document = (*document)(nil)

/*
var _ DocumentType = (*doctype)(nil)
var _ Element = (*element)(nil)
var _ Text = (*text)(nil)
var _ Comment = (*comment)(nil)
var _ Attr = (*attr)(nil)
*/

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newNode(value js.Value) Node {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	// Create different kinds of nodes based on the Proto type
	node := &node{value}
	switch {
	//	case js.TypeOf(value).Equal(js.DocumentProto):
	//		return &document{node}
	default:
		return node
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

// Properties
func (node *node) ChildNodes() []Node {
	return nil
}
func (node *node) Contains(n Node) bool {
	return false
}
func (node *node) Equals(n Node) bool {
	return false
}
func (node *node) FirstChild() Node {
	return nil
}
func (node *node) HasChildNodes() bool {
	return false
}
func (node *node) IsConnected() bool {
	return false
}
func (node *node) LastChild() Node {
	return nil
}
func (node *node) NextSibling() Node {
	return nil
}
func (node *node) NodeName() string {
	return ""
}
func (node *node) NodeType() NodeType {
	return 0
}
func (node *node) OwnerDocument() Document {
	return nil
}
func (node *node) ParentElement() Element {
	return nil
}
func (node *node) ParentNode() Node {
	return nil
}
func (node *node) PreviousSibling() Node {
	return nil
}
func (node *node) TextContent() string {
	return ""
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Methods
func (node *node) AppendChild(Node) Node {
	return nil
}
func (node *node) CloneNode(bool) Node {
	return nil
}
func (node *node) InsertBefore(Node, Node) Node {
	return nil
}
func (node *node) RemoveChild(Node) {
}
