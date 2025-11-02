package bs

import (
	"slices"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type text struct {
	View
}

var _ View = (*text)(nil)

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
	RegisterView(ViewText, newTextFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Para(children ...any) View {
	return NewView(new(text), ViewText, "P").Append(children...)
}

func Deleted(children ...any) View {
	return NewView(new(text), ViewText, "DEL").Append(children...)
}

func Highlighted(children ...any) View {
	return NewView(new(text), ViewText, "MARK").Append(children...)
}

func Small(children ...any) View {
	return NewView(new(text), ViewText, "SMALL").Append(children...)
}

func Strong(children ...any) View {
	return NewView(new(text), ViewText, "STRONG").Append(children...)
}

func Em(children ...any) View {
	return NewView(new(text), ViewText, "EM").Append(children...)
}

func Blockquote(children ...any) View {
	return NewView(new(text), ViewText, "BLOCKQUOTE", WithClass("blockquote")).Append(children...)
}

func Code(children ...any) View {
	return NewView(new(text), ViewText, "CODE").Append(children...)
}

func newTextFromElement(element Element) View {
	if !slices.Contains(textTagNames, element.TagName()) {
		return nil
	}
	return NewViewWithElement(new(text), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (text *text) SetView(view View) {
	text.View = view
}
