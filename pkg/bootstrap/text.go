package bootstrap

import (
	// Package imports
	"slices"

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

var (
	textTagNames = []string{
		"P",
		"DEL",
		"MARK",
		"SMALL",
		"STRONG",
		"EM",
		"BLOCKQUOTE",
		"CODE",
	}
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

func Deleted(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "DEL", opts...).Content(content...).(*text)
}

func Highlighted(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "MARK", opts...).Content(content...).(*text)
}

func Smaller(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "SMALL", opts...).Content(content...).(*text)
}

func Strong(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "STRONG", opts...).Content(content...).(*text)
}

func Em(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "EM", opts...).Content(content...).(*text)
}

func Blockquote(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "BLOCKQUOTE", append(opts, mvc.WithClass("blockquote"))...).Content(content...).(*text)
}

func Code(args ...any) *text {
	opts, content := gatherOpts(args)
	return mvc.NewView(new(text), ViewText, "CODE", opts...).Content(content...).(*text)
}

func newTextFromElement(element Element) mvc.View {
	if !slices.Contains(textTagNames, element.TagName()) {
		return nil
	}
	return mvc.NewViewWithElement(new(text), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (text *text) SetView(view mvc.View) {
	text.View = view
}
