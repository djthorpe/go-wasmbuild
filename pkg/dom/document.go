//go:build !(js && wasm)

package dom

import (
	"bytes"
	"io"

	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type document struct {
	EventTarget
	node
}

var _ Document = (*document)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newDocument(parent Node) Document {
	node := newNode(nil, parent, "#document", DOCUMENT_NODE, "")
	return &document{
		EventTarget: js.NewEventTarget(),
		node:        node,
	}
}

func newHTMLDocument(parent Node) Document {
	document := newDocument(parent)

	// Create document structure
	// TODO: Add DOCTYPE
	html := newElement(document, "html")
	head := newElement(document, "head")
	body := newElement(document, "body")

	// Append children
	document.AppendChild(html)
	html.AppendChild(head)
	html.AppendChild(body)

	// Return the document
	return document
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

func (document *document) Head() Element {
	html := document.getChildNodesOfType(ELEMENT_NODE, func(n Node) bool {
		return n.(Element).TagName() == "HTML"
	})
	if len(html) != 1 {
		return nil
	}
	head := html[0].(*element).getChildNodesOfType(ELEMENT_NODE, func(n Node) bool {
		return n.(Element).TagName() == "HEAD"
	})
	if len(head) != 1 {
		return nil
	}
	return head[0].(Element)
}

func (document *document) Body() Element {
	html := document.getChildNodesOfType(ELEMENT_NODE, func(n Node) bool {
		return n.(Element).TagName() == "HTML"
	})
	if len(html) != 1 {
		return nil
	}
	head := html[0].(*element).getChildNodesOfType(ELEMENT_NODE, func(n Node) bool {
		return n.(Element).TagName() == "BODY"
	})
	if len(head) != 1 {
		return nil
	}
	return head[0].(Element)
}

func (document *document) Title() string {
	head := document.Head()
	if head == nil {
		return ""
	}
	title := document.getChildNodesOfType(ELEMENT_NODE, func(n Node) bool {
		return n.(Element).TagName() == "TITLE"
	})
	if len(title) != 1 {
		return ""
	}
	return title[0].(Element).TextContent()
}

func (document *document) Doctype() DocumentType {
	// TODO
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (document *document) CreateElement(name string) Element {
	return newElement(document, name)
}

func (document *document) CreateAttribute(name string) Attr {
	return newAttr(document, nil, name, "")
}

func (document *document) CreateComment(cdata string) Comment {
	return newComment(document, cdata)
}

func (document *document) CreateTextNode(cdata string) Text {
	return newTextNode(document, cdata)
}
