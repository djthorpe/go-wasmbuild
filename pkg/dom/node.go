//go:build !(js && wasm)

package dom

import (
	"slices"

	// Namespace imports

	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// nodeImpl is a private interface for types that wrap a node
type nodeImpl interface {
	getNode() *node
}

type node struct {
	document Document
	parent   Node
	name     string
	nodetype NodeType
	cdata    string
	children []Node
}

/*

type doctype struct {
	node
}

type element struct {
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

/*
var _ DocumentType = (*doctype)(nil)
var _ Element = (*element)(nil)
var _ Comment = (*comment)(nil)
var _ Attr = (*attr)(nil)
*/

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newNode(document Document, parent Node, name string, nodetype NodeType, cdata string) node {
	return node{document, parent, name, nodetype, cdata, nil}
}

// getNode implements nodeImpl interface
func (n *node) getNode() *node {
	return n
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

// Properties
func (node *node) ChildNodes() []Node {
	return node.children
}

func (node *node) Contains(n Node) bool {
	return slices.Contains(node.children, n)
}

func (n *node) Equals(other Node) bool {
	// Try to get the underlying node implementation
	if otherImpl, ok := other.(nodeImpl); ok {
		return n == otherImpl.getNode()
	}
	// Fallback to direct pointer comparison
	if other, ok := other.(*node); ok {
		return n == other
	}
	return false
}

func (node *node) FirstChild() Node {
	if len(node.children) > 0 {
		return node.children[0]
	}
	return nil
}

func (node *node) HasChildNodes() bool {
	return len(node.children) > 0
}

func (node *node) IsConnected() bool {
	return node.parent != nil
}

func (node *node) LastChild() Node {
	if len(node.children) > 0 {
		return node.children[len(node.children)-1]
	}
	return nil
}

func (node *node) NodeName() string {
	return node.name
}

func (node *node) NodeType() NodeType {
	return node.nodetype
}

func (node *node) OwnerDocument() Document {
	return node.document
}

func (node *node) ParentElement() Element {
	return nil
}

func (node *node) ParentNode() Node {
	return node.parent
}

func (node *node) NextSibling() Node {
	// Node should be connected
	if node.parent == nil {
		return nil
	}
	// Find next sibling
	nodes := node.parent.ChildNodes()
	for i, child := range nodes {
		if child.Equals(node) {
			if i+1 < len(nodes) {
				return nodes[i+1]
			} else {
				return nil
			}
		}
	}
	return nil
}

func (node *node) PreviousSibling() Node {
	// Node should be connected
	if node.parent == nil {
		return nil
	}
	// Find previous sibling
	nodes := node.parent.ChildNodes()

	for i, child := range nodes {
		if child.Equals(node) {
			if i-1 >= 0 {
				return nodes[i-1]
			} else {
				return nil
			}
		}
	}
	return nil
}

func (node *node) TextContent() string {
	if node.nodetype == TEXT_NODE || node.nodetype == COMMENT_NODE {
		return node.cdata
	}
	// For elements and documents, concatenate text from text nodes only (skip comments)
	var data string
	for _, child := range node.children {
		if child.NodeType() != COMMENT_NODE {
			data += child.TextContent()
		}
	}
	return data
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Methods
func (node *node) AppendChild(Node) Node {
	// TODO
	return nil
}

func (node *node) CloneNode(bool) Node {
	// TODO
	return nil
}

func (node *node) InsertBefore(Node, Node) Node {
	// TODO
	return nil
}

func (node *node) RemoveChild(Node) {
	// TODO
}
