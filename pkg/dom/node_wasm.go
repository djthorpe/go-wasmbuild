//go:build js && wasm

package dom

import (
	"fmt"
	"io"

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

var _ Node = (*node)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newNode(value js.Value) node {
	return node{value}
}

///////////////////////////////////////////////////////////////////////////////
// WRITER

func (node *node) Write(w io.Writer) (int, error) {
	return w.Write([]byte(node.Get("nodeName").String()))
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (node *node) ChildNodes() []Node {
	nodes := node.Value.Get("childNodes")
	length := nodes.Get("length").Int()
	result := make([]Node, 0, length)
	for i := 0; i < length; i++ {
		node := newNode(nodes.Index(i))
		result = append(result, &node)
	}
	return result
}

func (n *node) Contains(other Node) bool {
	if other == nil {
		return false
	}
	return n.Value.Call("contains", toValue(other)).Bool()
}

func (n *node) Equals(other Node) bool {
	if other == nil {
		return false
	}
	return n.Value.Equal(toValue(other))
}

func (n *node) FirstChild() Node {
	child := n.Value.Get("firstChild")
	if child.IsNull() || child.IsUndefined() {
		return nil
	}
	return &node{child}
}

func (n *node) HasChildNodes() bool {
	return n.Value.Call("hasChildNodes").Bool()
}

func (n *node) IsConnected() bool {
	return n.Value.Get("isConnected").Bool()
}

func (n *node) LastChild() Node {
	child := n.Value.Get("lastChild")
	if child.IsNull() || child.IsUndefined() {
		return nil
	}
	return &node{child}
}

func (n *node) NextSibling() Node {
	child := n.Value.Get("nextSibling")
	if child.IsNull() || child.IsUndefined() {
		return nil
	}
	return &node{child}
}

func (n *node) PreviousSibling() Node {
	child := n.Value.Get("previousSibling")
	if child.IsNull() || child.IsUndefined() {
		return nil
	}
	return &node{child}
}

func (n *node) NodeName() string {
	return n.Value.Get("nodeName").String()
}

func (n *node) NodeType() NodeType {
	return NodeType(n.Value.Get("nodeType").Int())
}

func (n *node) OwnerDocument() Document {
	return newDocument(n.Value.Get("ownerDocument"))
}

func (n *node) ParentElement() Element {
	node := n.Value.Get("parentElement")
	if node.IsNull() || node.IsUndefined() {
		return nil
	}
	return newElement(node)
}

func (n *node) ParentNode() Node {
	child := n.Value.Get("parentNode")
	if child.IsNull() || child.IsUndefined() {
		return nil
	}
	return &node{child}
}

func (n *node) TextContent() string {
	return n.Value.Get("textContent").String()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (n *node) AppendChild(child Node) Node {
	if child == nil {
		return nil
	}
	n.Value.Call("appendChild", toValue(child))
	return child
}

func (n *node) CloneNode(deep bool) Node {
	node := newNode(n.Value.Call("cloneNode", deep))
	return &node
}

func (n *node) InsertBefore(child Node, before Node) Node {
	if before == nil {
		return n.AppendChild(child)
	}
	n.Value.Call("insertBefore", toValue(child), toValue(before))
	return child
}

func (n *node) RemoveChild(child Node) {
	n.Value.Call("removeChild", toValue(child))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func toValue(n Node) js.Value {
	switch n := n.(type) {
	case nil:
		return js.Undefined()
	case *node:
		return n.Value
	case *text:
		return n.Value
	case *comment:
		return n.Value
	case *attr:
		return n.node.Value
	case *element:
		return n.node.Value
	case *document:
		return n.Value
	default:
		panic(fmt.Sprintf("toValue: invalid node type %T", n))
	}
}
