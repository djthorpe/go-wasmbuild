package views

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
	return mvc.NewView(new(div), ViewDiv, "DIV", opts...)
}

func newDivFromElement(element dom.Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(div), element)
}
