package bs

import (
	"fmt"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// image are elements that represent image views
type image struct {
	View
}

var _ View = (*image)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewImage = "mvc-bs-image"
)

func init() {
	RegisterView(ViewImage, newImageFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Image(src string, opt ...Opt) *image {
	if src == "" {
		panic("Image: src cannot be empty")
	}
	return &image{NewView(ViewImage, "IMG", append(opt, WithClass("image-fluid"), WithAttr("src", src))...)}
}

func RoundedImage(src string, opt ...Opt) *image {
	return Image(src, append(opt, WithClass("rounded"))...)
}

func ThumbnailImage(src string, opt ...Opt) *image {
	return Image(src, append(opt, WithClass("img-thumbnail"))...)
}

func newImageFromElement(element Element) View {
	tagName := element.TagName()
	if tagName != "IMG" {
		panic(fmt.Sprintf("newImageFromElement: invalid tag name %q", tagName))
	}
	return &image{NewViewWithElement(element)}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (image *image) Append(children ...any) View {
	panic("Append: not supported for image")
}

func (image *image) Content(children ...any) View {
	panic("Content: not supported for image")
}
