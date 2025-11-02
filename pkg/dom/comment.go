//go:build !(js && wasm)

package dom

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type comment struct {
	node
}

var _ Comment = (*comment)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newComment(document Document, cdata string) Comment {
	node := newNode(document, nil, "#comment", COMMENT_NODE, cdata)
	return &comment{
		node: node,
	}
}

// getNode implements nodeImpl interface
func (c *comment) getNode() *node {
	return &c.node
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
	panic("AppendChild not supported on Comment nodes")
}

func (c *comment) CloneNode(deep bool) Node {
	return newComment(nil, c.cdata)
}

func (c *comment) InsertBefore(newNode Node, refNode Node) Node {
	panic("InsertBefore not supported on Comment nodes")
}

func (c *comment) RemoveChild(Node) {
	panic("RemoveChild not supported on Comment nodes")
}
