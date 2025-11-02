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

type document struct {
	EventTarget
	js.Value
}

var _ Document = (*document)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newDocument(value js.Value) Document {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &document{
		EventTarget: js.NewEventTarget(value),
		Value:       value,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (d *document) Head() Element {
	return nil
}

func (d *document) Body() Element {
	return nil
}

func (d *document) Title() string {
	return ""
}

func (d *document) Doctype() DocumentType {
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// NODE INTERFACE METHODS

func (d *document) ChildNodes() []Node {
	childNodes := d.Value.Get("childNodes")
	length := childNodes.Get("length").Int()
	nodes := make([]Node, 0, length)
	for i := 0; i < length; i++ {
		node := childNodes.Index(i)
		if n := newNode(node); n != nil {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func (d *document) Contains(n Node) bool {
	if n == nil {
		return false
	}
	// Get the underlying js.Value from the node
	if nodeValue, ok := n.(interface{ Value() js.Value }); ok {
		return d.Value.Call("contains", nodeValue.Value()).Bool()
	}
	return false
}

func (d *document) Equals(n Node) bool {
	if n == nil {
		return false
	}
	if nodeValue, ok := n.(interface{ Value() js.Value }); ok {
		return d.Value.Equal(nodeValue.Value())
	}
	return false
}

func (d *document) FirstChild() Node {
	firstChild := d.Value.Get("firstChild")
	return newNode(firstChild)
}

func (d *document) HasChildNodes() bool {
	return d.Value.Call("hasChildNodes").Bool()
}

func (d *document) IsConnected() bool {
	return d.Value.Get("isConnected").Bool()
}

func (d *document) LastChild() Node {
	lastChild := d.Value.Get("lastChild")
	return newNode(lastChild)
}

func (d *document) NextSibling() Node {
	nextSibling := d.Value.Get("nextSibling")
	return newNode(nextSibling)
}

func (d *document) NodeName() string {
	return d.Value.Get("nodeName").String()
}

func (d *document) NodeType() NodeType {
	return NodeType(d.Value.Get("nodeType").Int())
}

func (d *document) OwnerDocument() Document {
	ownerDocument := d.Value.Get("ownerDocument")
	return newDocument(ownerDocument)
}

func (d *document) ParentElement() Element {
	parentElement := d.Value.Get("parentElement")
	if parentElement.IsNull() || parentElement.IsUndefined() {
		return nil
	}
	return newElement(parentElement)
}

func (d *document) ParentNode() Node {
	parentNode := d.Value.Get("parentNode")
	return newNode(parentNode)
}

func (d *document) PreviousSibling() Node {
	previousSibling := d.Value.Get("previousSibling")
	return newNode(previousSibling)
}

func (d *document) TextContent() string {
	textContent := d.Value.Get("textContent")
	if textContent.IsNull() || textContent.IsUndefined() {
		return ""
	}
	return textContent.String()
}

func (d *document) AppendChild(child Node) Node {
	if child == nil {
		return nil
	}
	if nodeValue, ok := child.(interface{ Value() js.Value }); ok {
		result := d.Value.Call("appendChild", nodeValue.Value())
		return newNode(result)
	}
	return nil
}

func (d *document) CloneNode(deep bool) Node {
	result := d.Value.Call("cloneNode", deep)
	return newNode(result)
}

func (d *document) InsertBefore(newChild, refChild Node) Node {
	if newChild == nil {
		return nil
	}
	var refValue js.Value
	if refChild != nil {
		if nodeValue, ok := refChild.(interface{ Value() js.Value }); ok {
			refValue = nodeValue.Value()
		}
	}
	if newValue, ok := newChild.(interface{ Value() js.Value }); ok {
		result := d.Value.Call("insertBefore", newValue.Value(), refValue)
		return newNode(result)
	}
	return nil
}

func (d *document) RemoveChild(child Node) {
	if child == nil {
		return
	}
	if nodeValue, ok := child.(interface{ Value() js.Value }); ok {
		d.Value.Call("removeChild", nodeValue.Value())
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (d *document) CreateElement(name string) Element {
	elem := d.Value.Call("createElement", name)
	return newElement(elem)
}

func (d *document) CreateAttribute(name string) Attr {
	attr := d.Value.Call("createAttribute", name)
	return newAttr(attr)
}

func (d *document) CreateComment(cdata string) Comment {
	comment := d.Value.Call("createComment", cdata)
	return newComment(comment)
}

func (d *document) CreateTextNode(cdata string) Text {
	text := d.Value.Call("createTextNode", cdata)
	return newTextNode(text)
}
