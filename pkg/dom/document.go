//go:build !(js && wasm)

package dom

import (
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

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (document *document) Head() Element {
	// TODO
	return nil
}

func (document *document) Body() Element {
	// TODO
	return nil
}

func (document *document) Title() string {
	// TODO
	return ""
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
	return nil
	// return newAttr(document, nil, name, ATTRIBUTE_NODE, "")
}

func (document *document) CreateComment(cdata string) Comment {
	return newComment(document, cdata)
}

func (document *document) CreateTextNode(cdata string) Text {
	return newTextNode(document, cdata)
}
