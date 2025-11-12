package bootstrap

import (
	// Package imports
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewText = "mvc-bs-text"
)

func init() {
	mvc.RegisterView(ViewText, newTextFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Para(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "P", opts...).Content(content...).(*text)
}

func LeadPara(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "P", append(opts, mvc.WithClass("lead"))...).Content(content...).(*text)
}

func newTextFromElement(element Element) mvc.View {
	if element.TagName() != "P" {
		return nil
	}
	return mvc.NewViewWithElement(new(codeblock), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (text *text) SetView(view mvc.View) {
	text.View = view
}
