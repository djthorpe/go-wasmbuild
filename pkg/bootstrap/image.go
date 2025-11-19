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
	BootstrapView
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
	i := new(img)
	i.BootstrapView.View = mvc.NewView(
		i, ViewImage, "IMG",
		mvc.WithAttr("src", href), mvc.WithClass("img-fluid"), args,
	)
	return i
}

func newImgFromElement(element Element) mvc.View {
	if element.TagName() != "IMG" {
		return nil
	}
	i := new(img)
	i.BootstrapView.View = mvc.NewViewWithElement(i, element)
	return i
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (img *img) Self() mvc.View {
	return img
}
