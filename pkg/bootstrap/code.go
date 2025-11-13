package bootstrap

import (
	// Package imports
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type codeblock struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewCodeBlock = "mvc-bs-codeblock"
)

func init() {
	mvc.RegisterView(ViewCodeBlock, newCodeBlockFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func CodeBlock(args ...any) *codeblock {
	return mvc.NewView(new(codeblock), ViewCodeBlock, "PRE", mvc.WithClass("codeblock"), args).(*codeblock)
}

func newCodeBlockFromElement(element Element) mvc.View {
	if element.TagName() != "PRE" {
		return nil
	}
	return mvc.NewViewWithElement(new(codeblock), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (codeblock *codeblock) SetView(view mvc.View) {
	codeblock.View = view
}
