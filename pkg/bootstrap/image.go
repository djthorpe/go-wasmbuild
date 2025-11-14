package bootstrap

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type img struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewImage = "mvc-bs-img"
)

func init() {
	mvc.RegisterView(ViewImage, newImgFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Image(href string, args ...any) *img {
	// Return the img
	return mvc.NewView(
		new(img), ViewImage, "IMG",
		mvc.WithAttr("src", href), mvc.WithClass("img-fluid"), args,
	).(*img)
}

func newImgFromElement(element Element) mvc.View {
	if element.TagName() != "IMG" {
		return nil
	}
	return mvc.NewViewWithElement(new(img), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (img *img) SetView(view mvc.View) {
	img.View = view
}
