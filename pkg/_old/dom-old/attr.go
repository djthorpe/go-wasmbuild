//go:build !js

package dom

import (
	"fmt"
	"html"
	"io"
	"strconv"
	"strings"

	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type attr struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *attr) String() string {
	var b strings.Builder
	b.WriteString("<DOMAttribute")
	if name := this.Name(); name != "" {
		fmt.Fprintf(&b, " %v=%q", name, this.Value())
	}
	b.WriteString(">")
	return b.String()
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *attr) Name() string {
	return this.name
}

func (this *attr) Value() string {
	return this.cdata
}

func (this *attr) SetValue(cdata string) {
	this.cdata = cdata
}

func (this *attr) OwnerElement() dom.Element {
	return this.ParentElement()
}

/////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (this *attr) CloneNode(bool) dom.Node {
	return NewNode(this.document, this.name, this.nodetype, this.cdata)
}

// Child manipulation methods are no-ops for attribute nodes (leaf nodes)
func (this *attr) AppendChild(child dom.Node) dom.Node {
	return nil
}

func (this *attr) InsertBefore(new dom.Node, ref dom.Node) dom.Node {
	return nil
}

func (this *attr) RemoveChild(child dom.Node) {
}

func (this *attr) ReplaceChild(dom.Node, dom.Node) {
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *attr) v() *node {
	return this.node
}

func (this *attr) write(w io.Writer) (int, error) {
	return w.Write([]byte(this.name + "=" + strconv.Quote(html.EscapeString(this.cdata))))
}
