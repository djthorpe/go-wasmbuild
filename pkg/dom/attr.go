//go:build !(js && wasm)

package dom

import (
	"bytes"
	"html"
	"io"
	"strconv"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type attr struct {
	node
}

var _ Attr = (*attr)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func newAttr(document Document, parent Element, name, value string) Attr {
	node := newNode(document, parent, name, ELEMENT_NODE, value)
	return &attr{
		node: node,
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (attr *attr) String() string {
	var b bytes.Buffer
	if _, err := attr.Write(&b); err != nil {
		return err.Error()
	} else {
		return b.String()
	}
}

func (attr *attr) Write(w io.Writer) (int, error) {
	var s int
	if n, err := w.Write([]byte(attr.name)); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte("=")); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(strconv.Quote(html.EscapeString(attr.cdata)))); err != nil {
		return s, err
	} else {
		s += n
	}
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// ATTR

func (attr *attr) OwnerElement() Element {
	return attr.ParentElement()
}

func (attr *attr) Name() string {
	return attr.name

}

func (attr *attr) Value() string {
	return attr.cdata
}

func (attr *attr) SetValue(value string) {
	attr.cdata = value
}
