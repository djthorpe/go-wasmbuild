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
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "P", args)
	return t
}

func LeadPara(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "P", mvc.WithClass("lead"), args)
	return t
}

func Deleted(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "DEL", args)
	return t
}

func Highlighted(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "MARK", args)
	return t
}

func Smaller(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "SMALL", args)
	return t
}

func Strong(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "STRONG", args)
	return t
}

func Em(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "EM", args)
	return t
}

func Blockquote(args ...any) *text {
	t := new(text)
	t.View = mvc.NewViewExEx(t, ViewText, templateBlockquote, args)
	return t
}

func Code(args ...any) *text {
	t := new(text)
	t.View = mvc.NewView(t, ViewText, "CODE", args)
	return t
}

func newTextFromElement(element Element) mvc.View {
	if !slices.Contains(textTagNames, element.TagName()) {
		return nil
	}
	t := new(text)
	t.View = mvc.NewViewWithElement(t, element)
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (text *text) Self() mvc.View {
	return text
}
