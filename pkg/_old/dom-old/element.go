//go:build !js

package dom

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	*node
	classlist *tokenlist
	attrs     map[string]dom.Attr
}

var _ dom.Element = (*element)(nil)
var _ dom.EventTarget = (*element)(nil)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *element) String() string {
	return this.OuterHTML()
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *element) InnerHTML() string {
	buf := new(bytes.Buffer)
	for child := this.FirstChild(); child != nil; child = child.NextSibling() {
		writeNode(buf, child)
	}
	return buf.String()
}

func (this *element) SetInnerHTML(html string) {
	// Clear existing children
	this.node.children = []dom.Node{}
}

func (this *element) OuterHTML() string {
	buf := new(bytes.Buffer)
	this.write(buf)
	return buf.String()
}

func (this *element) TagName() string {
	if name := this.NodeName(); strings.HasPrefix(name, "#") {
		return name
	} else {
		return strings.ToUpper(name)
	}
}

func (this *element) Attributes() []dom.Attr {
	result := make([]dom.Attr, 0, len(this.attrs))
	for _, attr := range this.attrs {
		result = append(result, attr)
	}
	return result
}

func (this *element) HasAttributes() bool {
	return len(this.attrs) > 0
}

func (this *element) Style() dom.Style {
	// Not implemented for non-WASM builds
	return nil
}

func (this *element) SetAttribute(name, value string) dom.Attr {
	attr := this.document.CreateAttribute(name)
	attr.SetValue(value)
	getNode(attr).parent = this
	this.attrs[name] = attr

	// Sync classlist when class attribute is set
	if name == "class" && value != "" {
		classes := strings.Fields(value)
		this.classlist = NewTokenList(classes...)
	}

	return attr
}

func (this *element) GetAttribute(name string) string {
	attr := this.attrs[name]
	if attr == nil {
		return ""
	}
	return attr.Value()
}

func (this *element) GetAttributeNode(name string) dom.Attr {
	return this.attrs[name]
}

func (this *element) HasAttribute(name string) bool {
	_, exists := this.attrs[name]
	return exists
}

func (this *element) RemoveAttribute(name string) {
	if attr, exists := this.attrs[name]; exists {
		getNode(attr).parent = nil
		delete(this.attrs, name)
	}
}

func (this *element) RemoveAttributeNode(attr dom.Attr) {
	if attr == nil {
		return
	}
	name := attr.Name()
	if existing, exists := this.attrs[name]; exists && existing == attr {
		getNode(attr).parent = nil
		delete(this.attrs, name)
	}
}

func (this *element) SetAttributeNode(attr dom.Attr) dom.Attr {
	if attr == nil {
		return nil
	}
	name := attr.Name()
	oldAttr := this.attrs[name]
	if oldAttr != nil {
		getNode(oldAttr).parent = nil
	}
	getNode(attr).parent = this
	this.attrs[name] = attr
	return oldAttr
}

func (this *element) GetAttributeNames() []string {
	names := make([]string, 0, len(this.attrs))
	for name := range this.attrs {
		names = append(names, name)
	}
	return names
}

func (this *element) ClassList() dom.TokenList {
	return this.classlist
}

func (this *element) GetElementsByClassName(className string) []dom.Element {
	var result []dom.Element
	this.getElementsByClassName(className, &result)
	return result
}

func (this *element) getElementsByClassName(className string, result *[]dom.Element) {
	// Check if this element has the class by directly checking the attribute
	// This is more efficient than ClassList().Contains() which parses the entire class list
	classAttr := this.GetAttribute("class")
	if classAttr != "" {
		// Split and check for exact match
		for _, cls := range strings.Fields(classAttr) {
			if cls == className {
				*result = append(*result, this)
				break
			}
		}
	}
	// Recursively check children
	for _, child := range this.node.children {
		if elem, ok := child.(dom.Element); ok {
			if e, ok := elem.(*element); ok {
				e.getElementsByClassName(className, result)
			}
		}
	}
}

func (this *element) GetElementsByTagName(tagName string) []dom.Element {
	var result []dom.Element
	// Normalize tagName to uppercase for case-insensitive comparison
	tagName = strings.ToUpper(tagName)
	this.getElementsByTagName(tagName, &result)
	return result
}

func (this *element) getElementsByTagName(tagName string, result *[]dom.Element) {
	// Recursively check children only (not this element itself)
	for _, child := range this.node.children {
		if elem, ok := child.(dom.Element); ok {
			// Check if this child matches (case-insensitive comparison via TagName which returns uppercase)
			if elem.TagName() == tagName {
				*result = append(*result, elem)
			}
			// Recursively check child's descendants
			if e, ok := elem.(*element); ok {
				e.getElementsByTagName(tagName, result)
			}
		}
	}
}

func (this *element) Remove() {
	if this.node.parent != nil {
		this.node.parent.RemoveChild(this)
	}
}

func (this *element) ReplaceWith(nodes ...dom.Node) {
	parent := this.ParentNode()
	if parent == nil {
		return
	}

	// Insert all new nodes before this element
	for _, node := range nodes {
		parent.InsertBefore(node, this)
	}

	// Remove this element
	parent.RemoveChild(this)
}

func (this *element) InsertAdjacentElement(position string, element dom.Element) dom.Element {
	if element == nil {
		return nil
	}

	switch strings.ToLower(position) {
	case "beforebegin":
		// Insert before this element
		if this.node.parent != nil {
			this.node.parent.InsertBefore(element, this)
			return element
		}
	case "afterbegin":
		// Insert as first child
		if len(this.node.children) > 0 {
			this.InsertBefore(element, this.node.children[0])
		} else {
			this.AppendChild(element)
		}
		return element
	case "beforeend":
		// Insert as last child
		this.AppendChild(element)
		return element
	case "afterend":
		// Insert after this element
		if this.node.parent != nil {
			nextSibling := this.NextSibling()
			if nextSibling != nil {
				this.node.parent.InsertBefore(element, nextSibling)
			} else {
				this.node.parent.AppendChild(element)
			}
			return element
		}
	}
	return nil
}

func (this *element) ID() string {
	return this.GetAttribute("id")
}

func (this *element) SetID(id string) {
	this.SetAttribute("id", id)
}

func (this *element) ClassName() string {
	return this.GetAttribute("class")
}

func (this *element) SetClassName(className string) {
	this.SetAttribute("class", className)
	// Invalidate cached classlist
	this.classlist = nil
}

func (this *element) Children() []dom.Element {
	var result []dom.Element
	for _, child := range this.node.children {
		if elem, ok := child.(dom.Element); ok {
			result = append(result, elem)
		}
	}
	return result
}

func (this *element) ChildElementCount() int {
	count := 0
	for _, child := range this.node.children {
		if _, ok := child.(dom.Element); ok {
			count++
		}
	}
	return count
}

func (this *element) FirstElementChild() dom.Element {
	for _, child := range this.node.children {
		if elem, ok := child.(dom.Element); ok {
			return elem
		}
	}
	return nil
}

func (this *element) LastElementChild() dom.Element {
	for i := len(this.node.children) - 1; i >= 0; i-- {
		if elem, ok := this.node.children[i].(dom.Element); ok {
			return elem
		}
	}
	return nil
}

func (this *element) NextElementSibling() dom.Element {
	if this.node.parent == nil {
		return nil
	}

	// Get parent as *node to access children
	parent := getNode(this.node.parent)
	if parent == nil {
		return nil
	}

	found := false
	for _, child := range parent.children {
		if found {
			if elem, ok := child.(dom.Element); ok {
				return elem
			}
		}
		if child == this {
			found = true
		}
	}
	return nil
}

func (this *element) PreviousElementSibling() dom.Element {
	if this.node.parent == nil {
		return nil
	}

	// Get parent as *node to access children
	parent := getNode(this.node.parent)
	if parent == nil {
		return nil
	}

	var prevElement dom.Element
	for _, child := range parent.children {
		if child == this {
			return prevElement
		}
		if elem, ok := child.(dom.Element); ok {
			prevElement = elem
		}
	}
	return nil
}

func (e *element) AddEventListener(eventType string, callback func(dom.Node)) dom.Element {
	// Event listeners are not supported in non-WASM builds
	// This is a no-op since there's no event loop outside the browser
	return e
}

// RemoveEventListener removes all event listeners of the specified type and releases their resources
func (e *element) RemoveEventListener(eventType string) dom.Element {
	// Event listeners are not supported in non-WASM builds
	return e
}

func (e *element) Blur() {
	// Not supported in non-WASM builds
}

func (e *element) Focus() {
	// Not supported in non-WASM builds
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *element) v() *node {
	return this.node
}

func (this *element) write(w io.Writer) (int, error) {
	s := 0

	// Sync classlist to class attribute before writing
	if this.classlist != nil && this.classlist.Length() > 0 {
		this.SetAttribute("class", this.classlist.Value())
	}

	// Write opening tag with attributes (lowercase to match browser behavior)
	tagName := strings.ToLower(this.node.name)
	tag := "<" + tagName

	// Add attributes in sorted order for consistent output
	if len(this.attrs) > 0 {
		// Get sorted attribute names
		names := make([]string, 0, len(this.attrs))
		for name := range this.attrs {
			names = append(names, name)
		}
		sort.Strings(names)

		// Write attributes in sorted order
		for _, name := range names {
			attr := this.attrs[name]
			tag += fmt.Sprintf(" %s=%q", attr.Name(), attr.Value())
		}
	}

	tag += ">"
	if n, err := w.Write([]byte(tag)); err != nil {
		return 0, err
	} else {
		s += n
	}

	// Write children
	for _, child := range this.node.children {
		if n, err := writeNode(w, child); err != nil {
			return 0, err
		} else {
			s += n
		}
	}

	// Write closing tag (lowercase to match browser behavior)
	if n, err := w.Write([]byte("</" + tagName + ">")); err != nil {
		return 0, err
	} else {
		s += n
	}

	return s, nil
}
