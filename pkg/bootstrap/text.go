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

const (
	templateBlockquote = `
	<figure>
		<blockquote class="blockquote"><slot></slot></blockquote>
		<figcaption class="blockquote-footer"><slot name="label"></slot></figcaption>
	</figure>
	`
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
	return mvc.NewView(new(text), ViewText, "P", args).(*text)
}

func LeadPara(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", mvc.WithClass("lead"), args).(*text)
}

func Deleted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "DEL", args).(*text)
}

func Highlighted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "MARK", args).(*text)
}

func Smaller(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "SMALL", args).(*text)
}

func Strong(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "STRONG", args).(*text)
}

func Em(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "EM", args).(*text)
}

func Blockquote(args ...any) *text {
	return mvc.NewViewExEx(new(text), ViewText, templateBlockquote, args).(*text)
}

func Code(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "CODE", args).(*text)
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
