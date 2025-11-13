//go:build js && wasm

package dom

import (
	// Package imports
	"bytes"
	"html"
	"io"

	js "github.com/djthorpe/go-wasmbuild/pkg/js"

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

func newComment(value js.Value) Comment {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &comment{newNode(value)}
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
	data := html.EscapeString(c.Value.Get("data").String())
	var written int
	if n, err := w.Write([]byte("<!--")); err != nil {
		return written, err
	} else {
		written += n
	}
	if n, err := w.Write([]byte(data)); err != nil {
		return written, err
	} else {
		written += n
	}
	if n, err := w.Write([]byte("-->")); err != nil {
		return written, err
	} else {
		written += n
	}
	return written, nil
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (c *comment) Data() string {
	return c.Value.Get("data").String()
}

func (c *comment) Length() int {
	return c.Value.Get("length").Int()
}
