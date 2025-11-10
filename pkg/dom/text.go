//go:build !(js && wasm)

package dom

import (
	"bytes"
	"html"
	"io"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	node
}

var _ Text = (*text)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newTextNode(document Document, cdata string) Text {
	node := newNode(document, nil, "#text", TEXT_NODE, cdata)
	return &text{
		node: node,
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (t *text) String() string {
	var b bytes.Buffer
	if _, err := t.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (t *text) Write(w io.Writer) (int, error) {
	return w.Write([]byte(html.EscapeString(t.cdata)))
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (t *text) Data() string {
	return t.cdata
}

func (t *text) Length() int {
	return len(t.cdata)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (t *text) AppendChild(Node) Node {
	panic("AppendChild[TEXT_NODE] not supported")
}

func (t *text) InsertBefore(newNode Node, refNode Node) Node {
	panic("InsertBefore[TEXT_NODE] not supported")
}

func (t *text) RemoveChild(Node) {
	panic("RemoveChild[TEXT_NODE]  not supported")
}
