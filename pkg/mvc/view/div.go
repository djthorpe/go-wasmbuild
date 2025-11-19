package view

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// div is a simple div view
type div struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewDiv = "mvc-div"
)

func init() {
	mvc.RegisterView(ViewDiv, newDivFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Div(opts ...mvc.Opt) mvc.View {
	d := new(div)
	args := make([]any, len(opts))
	for i, opt := range opts {
		args[i] = opt
	}
	d.View = mvc.NewView(d, ViewDiv, "DIV", args...)
	return d
}

func newDivFromElement(element dom.Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	d := new(div)
	d.View = mvc.NewViewWithElement(d, element)
	return d
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (d *div) Self() mvc.View {
	return d
}
