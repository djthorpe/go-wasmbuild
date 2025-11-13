//go:build js

package dom

import (
	"fmt"
	"syscall/js"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type comment struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *comment) String() string {
	str := "<DOMComment"
	str += fmt.Sprintf(" data=%q length=%v", this.Data(), this.Length())
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *comment) Data() string {
	return this.Get("data").String()
}

func (this *comment) Length() int {
	return this.Get("length").Int()
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *comment) v() js.Value {
	return this.node.v()
}
