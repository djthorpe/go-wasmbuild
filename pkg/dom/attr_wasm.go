//go:build js && wasm

package dom

import (
	"bytes"
	"html"
	"io"
	"strconv"

	// Packages

	js "github.com/djthorpe/go-wasmbuild/pkg/js"

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

func newAttr(value js.Value) Attr {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	return &attr{newNode(value)}
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
	if n, err := w.Write([]byte(attr.Name())); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte("=")); err != nil {
		return s, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(strconv.Quote(html.EscapeString(attr.Value())))); err != nil {
		return s, err
	} else {
		s += n
	}
	return s, nil
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *attr) Name() string {
	return this.Get("name").String()
}

func (this *attr) Value() string {
	return this.Get("value").String()
}

func (this *attr) SetValue(value string) {
	this.Set("value", value)
}

func (this *attr) OwnerElement() Element {
	owner := this.Get("ownerElement")
	if owner.IsNull() || owner.IsUndefined() {
		return nil
	}
	return newElement(owner)
}
