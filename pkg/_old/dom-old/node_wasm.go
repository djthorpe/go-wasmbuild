//go:build js && wasm

package dom

import (
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type node struct {
	jsutil.Value
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewNode wraps a jsutil into a dom interface type
func NewNode(v jsutil.Value) dom.Node {
	proto := jsutil.TypeOf(v)
	if proto.IsNull() || proto.IsUndefined() {
		return nil
	}
	switch proto {
	case jsutil.DocumentProto:
		return &document{node: &node{v}}
	case jsutil.ElementProto:
		return &element{node: &node{v}}
	case jsutil.TextProto:
		return &text{node: &node{v}}
	case jsutil.CommentProto:
		return &comment{node: &node{v}}
	case jsutil.DocumentTypeProto:
		return &doctype{node: &node{v}}
	case jsutil.AttrProto:
		return &attr{node: &node{v}}
	case jsutil.NodeProto:
		return &node{v}
	default:
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// toValue returns the underlying Value from any node type
func toValue(n dom.Node) jsutil.Value {
	if n == nil {
		return jsutil.Undefined()
	}
	switch v := n.(type) {
	case *node:
		return v.Value
	case *element:
		return v.node.Value
	case *attr:
		return v.node.Value
	case *text:
		return v.node.Value
	case *comment:
		return v.node.Value
	case *doctype:
		return v.node.Value
	case *document:
		return v.node.Value
	}
	return jsutil.Undefined()
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *node) ChildNodes() []dom.Node {
	return fromNodeList(this.Get("childNodes"))
}

func (this *node) Contains(other dom.Node) bool {
	return this.Call("contains", toValue(other)).Bool()
}

func (this *node) Equals(other dom.Node) {
	this.Call("equals", toValue(other))
}

func (this *node) FirstChild() dom.Node {
	return NewNode(this.Get("firstChild"))
}

func (this *node) HasChildNodes() bool {
	return this.Call("hasChildNodes").Bool()
}

func (this *node) IsConnected() bool {
	return this.Get("isConnected").Bool()
}

func (this *node) LastChild() dom.Node {
	return NewNode(this.Get("lastChild"))
}

func (this *node) NextSibling() dom.Node {
	return NewNode(this.Get("nextSibling"))
}

func (this *node) NodeName() string {
	return this.Get("nodeName").String()
}

func (this *node) NodeType() dom.NodeType {
	return dom.NodeType(this.Get("nodeType").Int())
}

func (this *node) OwnerDocument() dom.Document {
	return NewNode(this.Get("ownerDocument")).(dom.Document)
}

func (this *node) ParentNode() dom.Node {
	return NewNode(this.Get("parentNode"))
}

func (this *node) ParentElement() dom.Element {
	return NewNode(this.Get("parentElement")).(dom.Element)
}

func (this *node) PreviousSibling() dom.Node {
	return NewNode(this.Get("previousSibling"))
}

func (this *node) TextContent() string {
	return this.Get("textContent").String()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *node) AppendChild(child dom.Node) dom.Node {
	this.Call("appendChild", toValue(child))
	return child
}

func (this *node) CloneNode(deep bool) dom.Node {
	return NewNode(this.Call("cloneNode", deep))
}

func (this *node) InsertBefore(child dom.Node, before dom.Node) dom.Node {
	if before == nil {
		return this.AppendChild(child)
	} else {
		this.Call("insertBefore", toValue(child), toValue(before))
		return child
	}
}

func (this *node) RemoveChild(child dom.Node) {
	this.Call("removeChild", toValue(child))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func fromNodeList(v js.Value) []dom.Node {
	var result []dom.Node
	for _, v := range nodeListToSlice(v) {
		result = append(result, NewNode(v))
	}
	return result
}

func nodeListToSlice(v js.Value) []js.Value {
	length := v.Get("length").Int()
	result := make([]js.Value, length)
	for i := 0; i < length; i++ {
		result[i] = v.Call("item", i)
	}
	return result
}
