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

type text struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *text) String() string {
	var b strings.Builder
	b.WriteString("<DOMText")
	fmt.Fprintf(&b, " data=%q length=%v", this.Data(), this.Length())
	b.WriteString(">")
	return b.String()
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *text) Data() string {
	return this.cdata
}

func (this *text) Length() int {
	return len(this.cdata)
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *text) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

// Child manipulation methods are no-ops for text nodes (leaf nodes)
func (this *text) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *text) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	return nil
}

func (this *text) RemoveChild(child dom.Node) {
}

func (this *text) ReplaceChild(dom.Node, dom.Node) {
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *text) v() *node {
	return this.node
}

func (this *text) write(w io.Writer) (int, error) {
	return w.Write([]byte(html.EscapeString(this.cdata)))
}
