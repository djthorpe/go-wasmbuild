//go:build js && wasm

package dom

import (
	"bytes"
	"io"

	// Packages

	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace import
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	node
	EventTarget
}

var _ Element = (*element)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newElement(value js.Value) Element {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &element{
		node:        newNode(value),
		EventTarget: NewEventTarget(value),
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (e *element) String() string {
	var b bytes.Buffer
	if _, err := e.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (e *element) Write(w io.Writer) (int, error) {
	return w.Write([]byte(e.Value.Get("outerHTML").String()))
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - PROPERTIES

// Return the tag name in uppercase
func (element *element) TagName() string {
	return element.Value.Get("tagName").String()
}

// Return the ID
func (element *element) ID() string {
	return element.Value.Get("id").String()
}

// Set the ID
func (element *element) SetID(id string) {
	element.Value.Set("id", id)
}

func (e *element) OuterHTML() string {
	return e.Value.Get("outerHTML").String()
}

func (e *element) InnerHTML() string {
	return e.Value.Get("innerHTML").String()
}

func (e *element) SetInnerHTML(value string) {
	e.Value.Set("innerHTML", value)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CLASSES

// Return the class attribute as a string
func (element *element) ClassName() string {
	return element.Value.Get("className").String()
}

// Set the class as a atring
func (element *element) SetClassName(className string) {
	element.Value.Set("className", className)
}

// Return a TokenList of the classes
func (element *element) ClassList() TokenList {
	return js.GetTokenList(element.Value.Get("classList"))
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - ATTRIBUTES

// Attributes
func (element *element) Attributes() []Attr {
	attrs := element.Value.Get("attributes")
	length := attrs.Get("length").Int()
	result := make([]Attr, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, newAttr(attrs.Call("item", i)))
	}
	return result
}

// Get the attribute
func (element *element) GetAttribute(name string) string {
	result := element.Value.Call("getAttribute", name)
	if result.IsNull() {
		return ""
	}
	return result.String()
}

// Get the attribute
func (element *element) GetAttributeNode(name string) Attr {
	attrNode := element.Value.Call("getAttributeNode", name)
	if attrNode.IsNull() {
		return nil
	}
	return newAttr(attrNode)
}

// Set the attribute
func (element *element) SetAttribute(name, value string) Attr {
	element.Value.Call("setAttribute", name, value)
	return element.GetAttributeNode(name)
}

// Set the attribute
func (element *element) SetAttributeNode(node Attr) Attr {
	if node == nil {
		return nil
	}
	result := element.Value.Call("setAttributeNode", toValue(node))
	if result.IsNull() {
		return nil
	}
	return newAttr(result)
}

// Remove an attribute
func (element *element) RemoveAttribute(name string) {
	element.Value.Call("removeAttribute", name)
}

// Remove an attribute
func (element *element) RemoveAttributeNode(node Attr) {
	if node == nil {
		return
	}

	// Wrap in a deferred recover to handle JS errors gracefully
	defer func() {
		if r := recover(); r != nil {
			// Error occurred (attribute wasn't attached to this element)
			// Silently ignore
		}
	}()

	element.Value.Call("removeAttributeNode", toValue(node))
}

// Return an unsorted list of attribute names
func (element *element) GetAttributeNames() []string {
	names := element.Value.Call("getAttributeNames")
	length := names.Get("length").Int()
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, names.Index(i).String())
	}
	return result
}

// Return true if the element has a specific attribute
func (element *element) HasAttribute(name string) bool {
	return element.Value.Call("hasAttribute", name).Bool()
}

// Return true if the element has any attribute
func (element *element) HasAttributes() bool {
	return element.Value.Call("hasAttributes").Bool()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DOM MANIPULATION

// Return the child elements
func (element *element) Children() []Element {
	children := element.Value.Get("children")
	length := children.Get("length").Int()
	result := make([]Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, newElement(children.Index(i)))
	}
	return result
}

func (element *element) ChildElementCount() int {
	return element.Value.Get("childElementCount").Int()
}

func (element *element) FirstElementChild() Element {
	child := element.Value.Get("firstElementChild")
	if child.IsNull() {
		return nil
	}
	return newElement(child)
}

func (element *element) LastElementChild() Element {
	child := element.Value.Get("lastElementChild")
	if child.IsNull() {
		return nil
	}
	return newElement(child)
}

func (element *element) NextElementSibling() Element {
	child := element.Value.Get("nextElementSibling")
	if child.IsNull() {
		return nil
	}
	return newElement(child)
}

func (element *element) PreviousElementSibling() Element {
	child := element.Value.Get("previousElementSibling")
	if child.IsNull() {
		return nil
	}
	return newElement(child)
}

func (element *element) ReplaceWith(nodes ...Node) {
	parent := element.ParentNode()
	if parent == nil {
		return
	}

	for _, child := range nodes {
		if child == nil {
			continue
		}
		parent.InsertBefore(child, element)
	}

	parent.RemoveChild(element)
}

func (element *element) Prepend(nodes ...Node) {
	if len(nodes) == 0 {
		return
	}

	first := element.FirstChild()
	for _, child := range nodes {
		if child == nil {
			continue
		}
		if first == nil {
			element.AppendChild(child)
		} else {
			element.InsertBefore(child, first)
		}
	}
}

func (element *element) Remove() {
	if parent := element.ParentNode(); parent != nil {
		parent.RemoveChild(element)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (element *element) GetElementsByClassName(className string) []Element {
	nodes := element.Value.Call("getElementsByClassName", className)
	length := nodes.Get("length").Int()
	result := make([]Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, newElement(nodes.Index(i)))
	}
	return result
}

func (element *element) GetElementsByTagName(tagName string) []Element {
	nodes := element.Value.Call("getElementsByTagName", tagName)
	length := nodes.Get("length").Int()
	result := make([]Element, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, newElement(nodes.Index(i)))
	}
	return result
}
