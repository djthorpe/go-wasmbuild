//go:build !js

package dom

import (
	"fmt"
	"html"
	"io"
	"strings"

	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type comment struct {
	*node
}

/////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	startcomment = []byte("<!--")
	endcomment   = []byte("-->")
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *comment) String() string {
	var b strings.Builder
	b.WriteString("<DOMComment")
	fmt.Fprintf(&b, " data=%q length=%v", this.Data(), this.Length())
	b.WriteString(">")
	return b.String()
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *comment) Data() string {
	return this.cdata
}

func (this *comment) Length() int {
	return len(this.cdata)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *comment) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

// Child manipulation methods are no-ops for comment nodes (leaf nodes)
func (this *comment) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *comment) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	return nil
}

func (this *comment) RemoveChild(child dom.Node) {
}

func (this *comment) ReplaceChild(dom.Node, dom.Node) {
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *comment) v() *node {
	return this.node
}

func (this *comment) write(w io.Writer) (int, error) {
	s := 0
	if n, err := w.Write(startcomment); err != nil {
		return 0, err
	} else {
		s += n
	}
	if n, err := w.Write([]byte(html.EscapeString(this.cdata))); err != nil {
		return 0, err
	} else {
		s += n
	}
	if n, err := w.Write(endcomment); err != nil {
		return 0, err
	} else {
		s += n
	}
	return s, nil
}
