//go:build js

package dom

import (
	"fmt"
	"html"
	"io"
	"syscall/js"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type window struct {
	EventTarget
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// GetWindow returns a global window object
func GetWindow() dom.Window {
	value := jsutil.Global()
	return &window{
		jsutil.NewEventTarget(value),
	}
}

// GetWindowWithTitle returns a global window object with the specified title
func GetWindowWithTitle(title string) dom.Window {
	w := js.Global()
	w.Get("document").Set("title", title)
	return &window{w}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *window) String() string {
	str := "<DOMWindow"
	if doc := this.Document(); doc != nil {
		str += fmt.Sprint(" document=", doc)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *window) Document() dom.Document {
	return NewNode(this.Get("document")).(dom.Document)
}

// Write the HTML inside a node
func (this *window) Write(w io.Writer, node dom.Node) (int, error) {
	var s int
	if node == nil {
		return 0, dom.ErrBadParameter
	}

	jsVal := toJSValue(node)
	nodeType := dom.NodeType(jsVal.Get("nodeType").Int())

	switch nodeType {
	case dom.DOCUMENT_NODE:
		// For document, write DOCTYPE if present, then document element
		if doctype := node.FirstChild(); doctype != nil && doctype.NodeType() == dom.DOCUMENT_TYPE_NODE {
			if n, err := w.Write([]byte("<!DOCTYPE " + doctype.NodeName() + ">")); err != nil {
				return 0, err
			} else {
				s += n
			}
		}
		// Write document element (usually <html>)
		if html := jsVal.Get("documentElement"); !html.IsUndefined() {
			htmlStr := html.Get("outerHTML").String()
			if n, err := w.Write([]byte(htmlStr)); err != nil {
				return 0, err
			} else {
				s += n
			}
		}

	case dom.ELEMENT_NODE:
		// For elements, use outerHTML
		html := jsVal.Get("outerHTML").String()
		if n, err := w.Write([]byte(html)); err != nil {
			return 0, err
		} else {
			s += n
		}

	case dom.TEXT_NODE:
		// For text nodes, write the text content
		text := jsVal.Get("data").String()
		if n, err := w.Write([]byte(text)); err != nil {
			return 0, err
		} else {
			s += n
		}

	case dom.COMMENT_NODE:
		// For comment nodes, write as HTML comment
		comment := jsVal.Get("data").String()
		if n, err := w.Write([]byte("<!--" + comment + "-->")); err != nil {
			return 0, err
		} else {
			s += n
		}

	case dom.ATTRIBUTE_NODE:
		// For attribute nodes, format as name="value" with HTML escaping
		if attr, ok := node.(dom.Attr); ok {
			name := attr.Name()
			value := attr.Value()
			// Escape the value for HTML and wrap in quotes
			escaped := html.EscapeString(value)
			formatted := name + "=\"" + escaped + "\""
			if n, err := w.Write([]byte(formatted)); err != nil {
				return 0, err
			} else {
				s += n
			}
		}

	default:
		// For other node types, try outerHTML or fall back to textContent
		if html := jsVal.Get("outerHTML"); !html.IsUndefined() && html.Type() == js.TypeString {
			htmlStr := html.String()
			if n, err := w.Write([]byte(htmlStr)); err != nil {
				return 0, err
			} else {
				s += n
			}
		} else if text := jsVal.Get("textContent"); !text.IsUndefined() {
			textStr := text.String()
			if n, err := w.Write([]byte(textStr)); err != nil {
				return 0, err
			} else {
				s += n
			}
		}
	}

	return s, nil
}

func (this *window) NewMutationObserver(callback func()) dom.MutationObserver {
	return newMutationObserver(callback)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS
