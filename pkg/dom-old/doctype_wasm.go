//go:build js

package dom

import (
	"fmt"
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type doctype struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *doctype) String() string {
	str := "<DOMDocumentType"
	if name := this.Name(); name != "" {
		str += fmt.Sprintf(" name=%q", name)
	}
	if publicid := this.PublicId(); publicid != "" {
		str += fmt.Sprintf(" publicId=%q", publicid)
	}
	if systemid := this.SystemId(); systemid != "" {
		str += fmt.Sprintf(" systemId=%q", systemid)
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *doctype) Name() string {
	return this.Get("name").String()
}
func (this *doctype) PublicId() string {
	return this.Get("publicId").String()
}

func (this *doctype) SystemId() string {
	return this.Get("systemId").String()
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *doctype) v() js.Value {
	return this.node.v()
}
