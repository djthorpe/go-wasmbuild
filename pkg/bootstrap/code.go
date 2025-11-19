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
	c := new(codeblock)
	c.View = mvc.NewView(c, ViewCodeBlock, "PRE", mvc.WithClass("codeblock"), args)
	return c
}

func newCodeBlockFromElement(element Element) mvc.View {
	if element.TagName() != "PRE" {
		return nil
	}
	c := new(codeblock)
	c.View = mvc.NewViewWithElement(c, element)
	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (codeblock *codeblock) Self() mvc.View {
	return codeblock
}
