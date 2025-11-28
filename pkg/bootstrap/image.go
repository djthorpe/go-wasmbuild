package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type img struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewImage, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(img), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Image(href string, args ...any) *img {
	return mvc.NewView(new(img), ViewImage, "IMG", setView, mvc.WithAttr("src", href), mvc.WithClass("img-fluid"), args).(*img)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (img *img) Label(args ...any) mvc.View {
	// TODO
	return img
}
