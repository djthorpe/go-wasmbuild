//go:build js

package dom

import (
	"fmt"

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
	str := "<DOMAttribute"
	if name := this.Name(); name != "" {
		str += fmt.Sprintf(" %v=%q", name, this.Value())
	}
	return str + ">"
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

func (this *attr) OwnerElement() dom.Element {
	owner := this.Get("ownerElement")
	if owner.IsNull() || owner.IsUndefined() {
		return nil
	}
	return NewNode(owner).(dom.Element)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *attr) v() *node {
	return this.node
}
