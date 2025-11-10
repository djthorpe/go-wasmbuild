//go:build js && wasm

package dom

import (
	// Package imports
	"bytes"
	"io"

	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type document struct {
	js.Value
	EventTarget
}

var _ Document = (*document)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newDocument(value js.Value) Document {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &document{
		Value:       value,
		EventTarget: js.NewEventTarget(value),
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (document *document) String() string {
	var b bytes.Buffer
	if _, err := document.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (document *document) Write(w io.Writer) (int, error) {
	var s int
	for _, child := range document.ChildNodes() {
		if n, err := child.Write(w); err != nil {
			return n, err
		} else {
			s += n
		}
	}
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (d *document) Head() Element {
	return newElement(d.Value.Get("head"))
}

func (d *document) Body() Element {
	return newElement(d.Value.Get("body"))
}

func (d *document) Title() string {
	return d.Value.Get("title").String()
}

func (d *document) Doctype() DocumentType {
	// TODO
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// NODE INTERFACE METHODS

func (d *document) ChildNodes() []Node {
	childNodes := d.Value.Get("childNodes")
	length := childNodes.Get("length").Int()
	nodes := make([]Node, 0, length)
	for i := 0; i < length; i++ {
		node := newNode(childNodes.Index(i))
		nodes = append(nodes, &node)
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
	node := newNode(d.Value.Get("firstChild"))
	return &node
}

func (d *document) HasChildNodes() bool {
	return d.Value.Call("hasChildNodes").Bool()
}

func (d *document) IsConnected() bool {
	return d.Value.Get("isConnected").Bool()
}

func (d *document) LastChild() Node {
	node := newNode(d.Value.Get("lastChild"))
	return &node
}

func (d *document) NextSibling() Node {
	node := newNode(d.Value.Get("nextSibling"))
	return &node
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
	node := newNode(d.Value.Get("parentNode"))
	return &node
}

func (d *document) PreviousSibling() Node {
	node := newNode(d.Value.Get("previousSibling"))
	return &node
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
		node := newNode(d.Value.Call("appendChild", nodeValue.Value()))
		return &node
	}
	return nil
}

func (d *document) CloneNode(deep bool) Node {
	node := newNode(d.Value.Call("cloneNode", deep))
	return &node
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
		node := newNode(d.Value.Call("insertBefore", newValue.Value(), refValue))
		return &node
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
	return newElement(d.Value.Call("createElement", name))
}

func (d *document) CreateAttribute(name string) Attr {
	return newAttr(d.Value.Call("createAttribute", name))
}

func (d *document) CreateComment(cdata string) Comment {
	return newComment(d.Value.Call("createComment", cdata))
}

func (d *document) CreateTextNode(cdata string) Text {
	return newTextNode(d.Value.Call("createTextNode", cdata))
}
