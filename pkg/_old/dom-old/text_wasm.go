//go:build js

package dom

import (
	"fmt"
	"syscall/js"
)

/////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	*node
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *text) String() string {
	str := "<DOMText"
	str += fmt.Sprintf(" data=%q length=%v", this.Data(), this.Length())
	return str + ">"
}

/////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *text) Data() string {
	return this.Get("data").String()
}

func (this *text) Length() int {
	return this.Get("length").Int()
}

/////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *text) v() js.Value {
	return this.node.v()
}
