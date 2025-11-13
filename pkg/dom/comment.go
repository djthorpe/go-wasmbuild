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

type comment struct {
	node
}

var _ Comment = (*comment)(nil)

/////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	startcomment = []byte("<!--")
	endcomment   = []byte("-->")
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newComment(document Document, cdata string) Comment {
	node := newNode(document, nil, "#comment", COMMENT_NODE, cdata)
	return &comment{
		node: node,
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *comment) String() string {
	var b bytes.Buffer
	if _, err := c.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (c *comment) Write(w io.Writer) (int, error) {
	var s int
	if n, err := w.Write(startcomment); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(html.EscapeString(c.cdata))); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write(endcomment); err != nil {
		return s, err
	} else {
		s += n
	}
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (c *comment) Data() string {
	return c.cdata
}

func (c *comment) Length() int {
	return len(c.cdata)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (c *comment) AppendChild(Node) Node {
	panic("AppendChild[COMMENT_NODE] not supported")
}

func (c *comment) InsertBefore(newNode Node, refNode Node) Node {
	panic("InsertBefore[COMMENT_NODE] not supported")
}

func (c *comment) RemoveChild(Node) {
	panic("RemoveChild[COMMENT_NODE]  not supported")
}
