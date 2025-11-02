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

type comment struct {
	js.Value
}

var _ Comment = (*comment)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newComment(value js.Value) Comment {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &comment{
		Value: value,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (c *comment) Data() string {
	return c.Value.Get("data").String()
}

func (c *comment) Length() int {
	return c.Value.Get("length").Int()
}

///////////////////////////////////////////////////////////////////////////////
// NODE INTERFACE METHODS

func (c *comment) ChildNodes() []Node {
	// Comment nodes cannot have children
	return nil
}

func (c *comment) Contains(n Node) bool {
	// Comment nodes cannot contain other nodes
	return false
}

func (c *comment) Equals(n Node) bool {
	if n == nil {
		return false
	}
	if nodeValue, ok := n.(interface{ Value() js.Value }); ok {
		return c.Value.Equal(nodeValue.Value())
	}
	return false
}

func (c *comment) FirstChild() Node {
	// Comment nodes cannot have children
	return nil
}

func (c *comment) HasChildNodes() bool {
	// Comment nodes cannot have children
	return false
}

func (c *comment) IsConnected() bool {
	return c.Value.Get("isConnected").Bool()
}

func (c *comment) LastChild() Node {
	// Comment nodes cannot have children
	return nil
}

func (c *comment) NextSibling() Node {
	nextSibling := c.Value.Get("nextSibling")
	return newNode(nextSibling)
}

func (c *comment) NodeName() string {
	return c.Value.Get("nodeName").String()
}

func (c *comment) NodeType() NodeType {
	return NodeType(c.Value.Get("nodeType").Int())
}

func (c *comment) OwnerDocument() Document {
	ownerDocument := c.Value.Get("ownerDocument")
	return newDocument(ownerDocument)
}

func (c *comment) ParentElement() Element {
	parentElement := c.Value.Get("parentElement")
	if parentElement.IsNull() || parentElement.IsUndefined() {
		return nil
	}
	return newElement(parentElement)
}

func (c *comment) ParentNode() Node {
	parentNode := c.Value.Get("parentNode")
	return newNode(parentNode)
}

func (c *comment) PreviousSibling() Node {
	previousSibling := c.Value.Get("previousSibling")
	return newNode(previousSibling)
}

func (c *comment) TextContent() string {
	textContent := c.Value.Get("textContent")
	if textContent.IsNull() || textContent.IsUndefined() {
		return ""
	}
	return textContent.String()
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (c *comment) AppendChild(Node) Node {
	// Comment nodes cannot have children - this will throw in JavaScript
	panic("AppendChild not supported on Comment nodes")
}

func (c *comment) CloneNode(deep bool) Node {
	result := c.Value.Call("cloneNode", deep)
	return newNode(result)
}

func (c *comment) InsertBefore(newNode Node, refNode Node) Node {
	// Comment nodes cannot have children - this will throw in JavaScript
	panic("InsertBefore not supported on Comment nodes")
}

func (c *comment) RemoveChild(Node) {
	// Comment nodes cannot have children - this will throw in JavaScript
	panic("RemoveChild not supported on Comment nodes")
}
