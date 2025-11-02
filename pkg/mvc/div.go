package mvc

import (
	// Namespace imports
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// div is a simple div view
type div struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewDiv = "mvc-div"
)

func init() {
	RegisterView(ViewDiv, newDivFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Div(opts ...Opt) View {
	return NewView(new(div), ViewDiv, "DIV", opts...)
}

func newDivFromElement(element dom.Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return NewViewWithElement(new(div), element)
}
