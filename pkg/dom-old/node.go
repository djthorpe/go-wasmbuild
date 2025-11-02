//go:build !(js && wasm)

package js

import (
	"fmt"
	"io"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"

	// Namespace imports
	. "github.com/djthorpe/go-dom"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	document Document
	parent   Node
	name     string
	nodetype NodeType
	children []Node
	cdata    string
}

var _ Node = (*node)(nil)

/////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewNode(doc dom.Document, name string, nodetype dom.NodeType, cdata string) dom.Node {
	node := &node{doc, nil, name, nodetype, nil, cdata}
	switch nodetype {
	case dom.DOCUMENT_NODE:
		return &document{node, nil, nil, nil, nil}
	case dom.DOCUMENT_TYPE_NODE:
		return &doctype{node, "", ""}
	case dom.ELEMENT_NODE:
		return &element{node, NewTokenList(), map[string]dom.Attr{}}
	case dom.TEXT_NODE:
		return &text{node}
	case dom.COMMENT_NODE:
		return &comment{node}
	case dom.ATTRIBUTE_NODE:
		return &attr{node}
	default:
		return node
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *node) String() string {
	var b strings.Builder
	b.WriteString("<DOMNode")
	if this.name != "" {
		fmt.Fprintf(&b, " name=%q", this.name)
	}
	if this.nodetype != dom.UNKNOWN_NODE {
		fmt.Fprint(&b, " type=", this.nodetype)
	}
	if this.parent != nil {
		b.WriteString(" parent=<DOMNode")
		if name := this.parent.NodeName(); name != "" {
			fmt.Fprintf(&b, " name=%q", name)
		}
		b.WriteString(">")
		return b.String()
	}
	for c := this.FirstChild(); c != nil; c = c.NextSibling() {
		fmt.Fprint(&b, " child=", c)
	}
	b.WriteString(">")
	return b.String()
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *node) Contains(child dom.Node) bool {
	for _, c := range this.children {
		if c == child {
			return true
		}
	}
	for _, c := range this.children {
		if c.Contains(child) {
			return true
		}
	}
	return false
}

func (this *node) Equals(other dom.Node) bool {
	return this == getNode(other)
}

func (this *node) ChildNodes() []dom.Node {
	result := make([]dom.Node, len(this.children))
	for i := range this.children {
		result[i] = this.children[i]
	}
	return result
}

func (this *node) FirstChild() dom.Node {
	if len(this.children) > 0 {
		return this.children[0]
	} else {
		return nil
	}
}

func (this *node) IsConnected() bool {
	return this.parent != nil
}

func (this *node) LastChild() dom.Node {
	last := len(this.children) - 1
	if last >= 0 {
		return this.children[last]
	} else {
		return nil
	}
}

func (this *node) NextSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return findNextChild(getNode(this.parent), this)
	}
}

func (this *node) NodeName() string {
	return this.name
}

func (this *node) NodeType() dom.NodeType {
	return this.nodetype
}

func (this *node) OwnerDocument() dom.Document {
	return this.document
}

func (this *node) ParentNode() dom.Node {
	return this.parent
}

func (this *node) ParentElement() dom.Element {
	if this.parent != nil && this.parent.NodeType() == dom.ELEMENT_NODE {
		return this.parent.(dom.Element)
	} else {
		return nil
	}
}

func (this *node) PreviousSibling() dom.Node {
	if this.parent == nil {
		return nil
	} else {
		return findPreviousChild(getNode(this.parent), this)
	}
}

func (this *node) TextContent() string {
	if this.nodetype == dom.TEXT_NODE {
		return this.cdata
	}
	var data string
	for _, child := range this.children {
		data += child.TextContent()
	}
	return data
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *node) AppendChild(child dom.Node) dom.Node {
	node := getNode(child)
	if node.parent != nil {
		node.parent.RemoveChild(child)
	}
	node.parent = this
	this.children = append(this.children, child)
	return child
}

func (this *node) CloneNode(deep bool) dom.Node {
	clone := NewNode(this.document, this.name, this.nodetype, this.cdata)
	if deep {
		getNode(clone).children = make([]dom.Node, len(this.children))
		for i := range this.children {
			child := this.children[i].CloneNode(deep)
			getNode(child).parent = clone
		}
	}
	return clone
}

func (this *node) HasChildNodes() bool {
	return len(this.children) > 0
}

func (this *node) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	// Check parameters
	if new == nil {
		return nil
	}
	// 'new' is inserted at the end of parentNode's child nodes
	// when 'ref' is nil
	if ref == nil {
		return this.AppendChild(new)
	}
	// insert node before ref
	node := getNode(new)
	for i := range this.children {
		if this.children[i] != ref {
			continue
		}
		// Detach new from current parent
		if node.parent != nil {
			node.parent.RemoveChild(new)
		}
		// Attach new to this
		node.parent = this
		this.children = append(this.children[:i], append([]dom.Node{new}, this.children[i:]...)...)
		return new
	}
	// Ref not in children, return nil
	return nil
}

func (this *node) RemoveChild(child dom.Node) {
	for i, c := range this.children {
		if c != child {
			continue
		}
		// Deattach child from parent
		getNode(child).parent = nil
		// Remove child from parent
		this.children = append(this.children[:i], this.children[i+1:]...)
		return
	}
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *node) v() *node {
	return this
}

// getNode returns the internal *node from any node type
// This replaces the nodevalue interface v() method
func getNode(n dom.Node) *node {
	switch v := n.(type) {
	case *node:
		return v
	case *element:
		return v.node
	case *attr:
		return v.node
	case *text:
		return v.node
	case *comment:
		return v.node
	case *doctype:
		return v.node
	case *document:
		return v.node
	default:
		panic("getNode: unknown node type")
	}
}

// writeNode serializes any node type to HTML
// This replaces the nodevalue interface write() method
func writeNode(w io.Writer, n dom.Node) (int, error) {
	switch v := n.(type) {
	case *element:
		return v.write(w)
	case *attr:
		return v.write(w)
	case *text:
		return v.write(w)
	case *comment:
		return v.write(w)
	case *doctype:
		return v.write(w)
	case *node:
		return v.write(w)
	case *document:
		return v.write(w)
	default:
		panic("writeNode: unknown node type")
	}
}

// findNextChild finds the next sibling of child in parent's children
// This replaces the nodevalue interface nextChild() method
func findNextChild(parent *node, child dom.Node) dom.Node {
	if child == nil {
		return nil
	}
	for i, c := range parent.children {
		// Note: We use c.Equals(child) instead of c != child to properly compare
		// the underlying *node pointers. Equals() is safe here - it only calls
		// getNode() (a pure type switch) and does pointer comparison with ==.
		// No recursion occurs.
		if !c.Equals(child) {
			continue
		}
		if i < len(parent.children)-1 {
			return parent.children[i+1]
		} else {
			return nil
		}
	}
	return nil
}

// findPreviousChild finds the previous sibling of child in parent's children
// This replaces the nodevalue interface previousChild() method
func findPreviousChild(parent *node, child dom.Node) dom.Node {
	if child == nil {
		return nil
	}
	for i, c := range parent.children {
		// Note: We use c.Equals(child) instead of c != child to properly compare
		// the underlying *node pointers. Equals() is safe here - it only calls
		// getNode() (a pure type switch) and does pointer comparison with ==.
		// No recursion occurs.
		if !c.Equals(child) {
			continue
		}
		if i > 0 {
			return parent.children[i-1]
		} else {
			return nil
		}
	}
	return nil
}

func (this *node) write(w io.Writer) (int, error) {
	s := 0
	if n, err := w.Write([]byte("<" + this.name + ">")); err != nil {
		return 0, err
	} else {
		s += n
	}
	for _, child := range this.children {
		if n, err := writeNode(w, child); err != nil {
			return 0, err
		} else {
			s += n
		}
	}
	if n, err := w.Write([]byte("</" + this.name + ">")); err != nil {
		return 0, err
	} else {
		s += n
	}
	return s, nil
}
