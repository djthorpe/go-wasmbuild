package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// figure are elements that represent figure views
type figure struct {
	View
}

// figurecaption are elements that represent figure caption, within the figure
type figurecaption struct {
	View
}

var _ View = (*figure)(nil)
var _ View = (*figurecaption)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewFigure        = "mvc-bs-figure"
	ViewFigureCaption = "mvc-bs-figurecaption"
)

func init() {
	RegisterView(ViewFigure, newFigureFromElement)
	RegisterView(ViewFigureCaption, newFigureCaptionFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Figure(opt ...Opt) *figure {
	return &figure{NewView(ViewFigure, "FIGURE", append(opt, WithClass("figure"))...)}
}

func newFigureFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "FIGURE" {
		panic(fmt.Sprintf("newFigureFromElement: invalid tag name %q", tagName))
	}
	return &figure{NewViewWithElement(element)}
}

func newFigureCaptionFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "FIGCAPTION" {
		panic(fmt.Sprintf("newFigureCaptionFromElement: invalid tag name %q", tagName))
	}
	return &figurecaption{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (figure *figure) Append(children ...any) View {
	// Remove any existing caption
	caption := figure.caption()
	if caption != nil {
		caption.Root().Remove()
	}

	// Append the content
	figure.View.Content(children...)

	// Append the caption
	if caption != nil {
		figure.View.Append(caption)
	}

	// Return the figure
	return figure
}

func (figure *figure) Content(children ...any) View {
	// Remove any existing caption
	caption := figure.caption()
	if caption != nil {
		caption.Root().Remove()
	}

	// Set the content
	figure.View.Content(children...)

	// Append the caption
	if caption != nil {
		figure.View.Append(caption)
	}

	// Return the figure
	return figure
}

func (figure *figure) Caption(children ...any) *figure {
	// Remove any existing caption first
	caption := figure.caption()
	if caption != nil {
		caption.Root().Remove()
	}

	// We append the caption to the figure at the bottom of the view
	figure.View.Append(NewView(ViewFigureCaption, "FIGCAPTION", WithClass("figure-caption")).Append(children...))

	// Return the figure
	return figure
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// Return the figurecaption element
func (figure *figure) caption() *figurecaption {
	element := figure.Root().LastElementChild()
	if element == nil || element.TagName() != "FIGCAPTION" {
		return nil
	}
	return newFigureCaptionFromElement(element).(*figurecaption)
}
