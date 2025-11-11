//go:build !(js && wasm)

package dom

import (
	"bytes"
	"errors"
	"io"
	"slices"
	"strings"

	// Packages

	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	html "golang.org/x/net/html"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type element struct {
	EventTarget
	node
	class TokenList
	attr  map[string]*attr
}

var _ Element = (*element)(nil)

/////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	startelementprefix = []byte("<")
	endelementprefix   = []byte("</")
	elementsuffix      = []byte(">")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newElement(document Document, name string) Element {
	node := newNode(document, nil, name, ELEMENT_NODE, "")
	return &element{
		EventTarget: NewEventTarget(),
		node:        node,
		class:       js.NewTokenList(),
		attr:        make(map[string]*attr, 10),
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (element *element) String() string {
	var b bytes.Buffer
	if _, err := element.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (element *element) Write(w io.Writer) (int, error) {
	var s int

	// Start Element
	if n, err := w.Write(startelementprefix); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(element.name)); err != nil {
		return s, err
	} else {
		s += n
	}

	// Attributes
	for _, attr := range element.attr {
		if n, err := w.Write([]byte(" ")); err != nil {
			return s, err
		} else {
			s += n
		}
		if n, err := attr.Write(w); err != nil {
			return s, err
		} else {
			s += n
		}
	}

	// Close Start Element
	if n, err := w.Write(elementsuffix); err != nil {
		return s, err
	} else {
		s += n
	}

	// Children
	for _, child := range element.children {
		if n, err := child.Write(w); err != nil {
			return s, err
		} else {
			s += n
		}
	}

	// End Element
	if n, err := w.Write(endelementprefix); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(element.name)); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write(elementsuffix); err != nil {
		return s, err
	} else {
		s += n
	}

	// Return success
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - PROPERTIES

// Return the tag name in uppercase
func (element *element) TagName() string {
	if name := element.node.name; strings.HasPrefix(name, "#") {
		return name
	} else {
		return strings.ToUpper(name)
	}
}

// Return the ID
func (element *element) ID() string {
	return element.GetAttribute("id")
}

// Set the ID
func (element *element) SetID(id string) {
	element.SetAttribute("id", id)
}

func (e *element) OuterHTML() string {
	var buf bytes.Buffer
	_, err := e.Write(&buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (e *element) InnerHTML() string {
	var buf bytes.Buffer
	for child := e.FirstChild(); child != nil; child = child.NextSibling() {
		if _, err := child.Write(&buf); err != nil {
			panic(err)
		}
	}
	return buf.String()
}

func (e *element) SetInnerHTML(value string) {
	reader := bytes.NewReader([]byte(value))
	tokenizer := html.NewTokenizer(reader)

	// Remove all the children
	e.RemoveAllChildren()

	// Short circuit if value is empty
	if value == "" {
		return
	}

	// Create a stack of elements
	stack := new(elements)
	stack.push(e)

	// Iterate through tokens
	for {
		// Check for error or EOF
		if tokenizer.Next() == html.ErrorToken {
			if errors.Is(tokenizer.Err(), io.EOF) {
				break
			}
			panic(tokenizer.Err())
		}

		// Create tree of tokens
		token := tokenizer.Token()
		switch token.Type {
		case html.CommentToken:
			stack.cur().AppendChild(e.document.CreateComment(token.Data))
		case html.TextToken:
			stack.cur().AppendChild(e.document.CreateTextNode(token.Data))
		case html.StartTagToken:
			child := e.document.CreateElement(token.Data)
			for _, attr := range token.Attr {
				if attr.Namespace == "" {
					child.SetAttribute(attr.Key, attr.Val)
				}
			}
			stack.cur().AppendChild(child)
			stack.push(child)
		case html.EndTagToken:
			// TODO: check the tag name for the popped token is the same
			stack.pop()
		default:
			token := tokenizer.Token()
			panic("Unsupported token " + token.String())
		}
	}

	if stack.cur() != e {
		panic("SetInnerHTML: invalid stack state")
	}
}

type elements struct {
	stack []Element
}

func (e *elements) cur() Element {
	if len(e.stack) == 0 {
		return nil
	}
	return e.stack[len(e.stack)-1]
}

func (e *elements) push(c Element) Element {
	e.stack = append(e.stack, c)
	return c
}

func (e *elements) pop() Element {
	c := e.cur()
	e.stack = e.stack[:len(e.stack)-1]
	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CLASSES

// Return the class attribute as a string
func (element *element) ClassName() string {
	return element.GetAttribute("class")
}

// Set the class as a atring
func (element *element) SetClassName(className string) {
	element.SetAttribute("class", className)
}

// Return a TokenList of the classes
func (element *element) ClassList() TokenList {
	return element.class
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - ATTRIBUTES

// Attributes
func (element *element) Attributes() []Attr {
	result := make([]Attr, 0, len(element.attr))
	for _, attr := range element.attr {
		result = append(result, attr)
	}
	return result
}

// Get the attribute
func (element *element) GetAttribute(name string) string {
	attr := element.GetAttributeNode(name)
	if attr == nil {
		return ""
	}
	return attr.Value()
}

// Get the attribute
func (element *element) GetAttributeNode(name string) Attr {
	if attr, exists := element.attr[name]; exists {
		return attr
	} else {
		return nil
	}
}

// Set the attribute
func (element *element) SetAttribute(name, value string) Attr {
	node := newAttr(element.document, element, name, value)
	element.SetAttributeNode(node)
	return node
}

// Set the attribute
func (element *element) SetAttributeNode(node Attr) Attr {
	if node == nil {
		return nil
	}

	var previous Attr

	// Detatch the previous node
	name := node.Name()
	if old := element.GetAttributeNode(name); old != nil {
		if existing, ok := old.(*attr); ok {
			existing.parent = nil
		}
		previous = old
	}

	// Set the parent of the node to the element
	node.(*attr).parent = element

	// Sync classlist when class attribute is set
	if name == "class" {
		element.class = js.NewTokenList(splitClassNames(node.Value())...)
	}

	// Set the attribute
	element.attr[name] = node.(*attr)

	// Return the previous attribute, if any
	return previous
}

// Remove an attribute
func (element *element) RemoveAttribute(name string) {
	if node := element.attr[name]; node != nil {
		element.RemoveAttributeNode(node)
	}
}

// Remove an attribute
func (element *element) RemoveAttributeNode(node Attr) {
	if node == nil {
		return
	}

	// Set the parent of the node to the element
	name := node.Name()
	if other := element.attr[name]; other == node {
		node.(*attr).parent = nil
		delete(element.attr, name)
	}
}

// Return an unsorted list of attribute names
func (element *element) GetAttributeNames() []string {
	result := make([]string, 0, len(element.attr))
	for name := range element.attr {
		result = append(result, name)
	}
	return result
}

// Return true if the element has a specific attribute
func (element *element) HasAttribute(name string) bool {
	_, exists := element.attr[name]
	return exists
}

// Return true if the element has any attribute
func (element *element) HasAttributes() bool {
	return len(element.attr) > 0
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DOM MANIPULATION

// Return the child elements
func (element *element) Children() []Element {
	var result []Element
	for _, child := range element.children {
		if elem, ok := child.(Element); ok {
			result = append(result, elem)
		}
	}
	return result
}

func (element *element) ChildElementCount() int {
	count := 0
	for _, child := range element.children {
		if _, ok := child.(Element); ok {
			count++
		}
	}
	return count
}

func (element *element) FirstElementChild() Element {
	for _, child := range element.children {
		if elem, ok := child.(Element); ok {
			return elem
		}
	}
	return nil
}

func (element *element) LastElementChild() Element {
	for i := len(element.children) - 1; i >= 0; i-- {
		if elem, ok := element.children[i].(Element); ok {
			return elem
		}
	}
	return nil
}

func (element *element) NextElementSibling() Element {
	if element.parent == nil {
		return nil
	}
	found := false
	for _, child := range element.parent.ChildNodes() {
		if found {
			if elem, ok := child.(Element); ok {
				return elem
			}
		}
		if child.Equals(element) {
			found = true
		}
	}
	return nil
}

func (element *element) PreviousElementSibling() Element {
	if element.parent == nil {
		return nil
	}
	var prev Element
	for _, child := range element.parent.ChildNodes() {
		if child.Equals(element) {
			return prev
		}
		if elem, ok := child.(Element); ok {
			prev = elem
		}
	}
	return nil
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
	var result []Element
	element.getElementsByClassName(className, &result)
	return result
}

func (element *element) GetElementsByTagName(tagName string) []Element {
	var result []Element
	element.getElementsByTagName(strings.ToUpper(tagName), &result)
	return result
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return class names from a list of class names as a string
func splitClassNames(value string) []string {
	return strings.Fields(value)
}

func (e *element) getElementsByClassName(className string, result *[]Element) {
	// Check if this element has the class by directly checking the attribute
	classAttr := e.GetAttribute("class")
	if classAttr != "" {
		if slices.Contains(splitClassNames(classAttr), className) {
			*result = append(*result, e)
		}
	}

	// Recursively check child elements
	for _, child := range e.Children() {
		if e, ok := child.(*element); ok {
			e.getElementsByClassName(className, result)
		}
	}
}

func (e *element) getElementsByTagName(tagName string, result *[]Element) {
	// Recursively check child elements
	for _, child := range e.Children() {
		if child.TagName() == tagName {
			*result = append(*result, child)
		}
		if e, ok := child.(*element); ok {
			e.getElementsByTagName(tagName, result)
		}
	}
}
