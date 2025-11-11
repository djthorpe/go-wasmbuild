//go:build !(js && wasm)

package dom

import (
	"io"

	// Namespace imports

	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	document Document
	parent   Node
	name     string
	nodetype NodeType
	cdata    string
	children []Node
}

var _ Node = (*node)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newNode(document Document, parent Node, name string, nodetype NodeType, cdata string) node {
	return node{document, parent, name, nodetype, cdata, nil}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

// Properties
func (node *node) ChildNodes() []Node {
	if len(node.children) == 0 {
		return []Node{}
	} else {
		return node.children
	}
}

func (node *node) Contains(n Node) bool {
	if n == nil {
		return false
	}
	if node.Equals(n) {
		return true
	}
	for _, child := range node.children {
		if child.Equals(n) {
			return true
		}
		if getnode(child).Contains(n) {
			return true
		}
	}
	return false
}

func (n *node) Equals(other Node) bool {
	return getnode(other) == n
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
	if node.nodetype == DOCUMENT_NODE {
		return true
	}
	if node.parent == nil {
		return false
	}
	return getnode(node.parent).IsConnected()
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
	if node.parent != nil && node.parent.NodeType() == ELEMENT_NODE {
		return node.parent.(Element)
	} else {
		return nil
	}
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
	if node.nodetype == TEXT_NODE || node.nodetype == COMMENT_NODE || node.nodetype == UNKNOWN_NODE {
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
// PUBLIC METHODS

// Append a child node and return it
func (node *node) AppendChild(child Node) Node {
	if child == nil {
		return nil
	}

	// Remove child from parent
	child_ := getnode(child)
	if child_ == nil {
		panic("AppendChild: does not implement *node")
	} else if child_.parent != nil {
		child_.parent.RemoveChild(child)
	}

	// Set new parent and append
	child_.parent = node
	node.children = append(node.children, child)

	// Return the child node
	return child
}

func (node *node) RemoveChild(child Node) {
	if child == nil || len(node.children) == 0 {
		return
	}

	// Iterate through children
	for i, c := range node.children {
		// Skip until child is found
		if c != child {
			continue
		}

		// Deattach child from parent
		child_ := getnode(child)
		if child_ == nil {
			panic("RemoveChild: does not implement *node")
		} else {
			child_.parent = nil
		}

		// Remove child from parent and return
		node.children = append(node.children[:i], node.children[i+1:]...)
		return
	}
}

func (node *node) RemoveAllChildren() {
	// Detach children from parent
	for _, c := range node.children {
		getnode(c).parent = nil
	}
	node.children = nil
}

func (node *node) InsertBefore(child Node, ref Node) Node {
	// Check parameters
	if child == nil {
		return nil
	}

	// 'child' is inserted at the end of parentNode's child nodes when 'ref' is nil
	if ref == nil {
		return node.AppendChild(child)
	}

	// insert node before ref
	child_ := getnode(child)
	if child_ == nil {
		panic("RemoveChild: does not implement *node")
	}
	for i := range node.children {
		if node.children[i] != ref {
			continue
		}
		// Detach from current parent
		if child_.parent != nil {
			child_.parent.RemoveChild(child)
		}
		// Attach in correct position
		child_.parent = node
		node.children = append(node.children[:i], append([]Node{child}, node.children[i:]...)...)
		return child
	}

	// ref not in children, return nil
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// WRITER

func (node *node) Write(w io.Writer) (int, error) {
	panic("Write: not implemented for node")
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return the underlying node from any DOM node object
func getnode(self Node) *node {
	if self == nil {
		return nil
	}
	switch n := self.(type) {
	case *document:
		return &n.node
	case *element:
		return &n.node
	case *attr:
		return &n.node
	case *text:
		return &n.node
	case *comment:
		return &n.node
	case *node:
		return n
	default:
		panic("getnode called with invalid implmentation")
	}
}

// Return nodes of a specific type
func (node *node) getChildNodesOfType(nodetype NodeType, filter func(Node) bool) []Node {
	result := make([]Node, 0, len(node.children))
	for _, child := range node.children {
		if child.NodeType() == nodetype {
			if filter == nil || filter(child) {
				result = append(result, child)
			}
		}
	}
	return result
}
