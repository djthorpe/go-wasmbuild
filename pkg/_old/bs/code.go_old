package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type codeblock struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewCodeBlock = "mvc-bs-codeblock"
)

func init() {
	RegisterView(ViewCodeBlock, newCodeBlockFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func CodeBlock(opt ...Opt) *codeblock {
	opt = append([]Opt{WithClass("codeblock"), WithBorder(All), WithPadding(All, 2)}, opt...)
	view := &codeblock{NewView(ViewCodeBlock, "PRE", opt...)}

	// Create a body, append to root
	body := Code()
	view.Append(body)

	// Set the view's body to the code element
	view.Body(body.Root())

	return view
}

func newCodeBlockFromElement(element Element) View {
	if element.TagName() != "PRE" {
		return nil
	}
	return &codeblock{NewViewWithElement(element)}
}
