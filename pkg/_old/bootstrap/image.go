package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type image struct {
	component
}

// Ensure that image implements Component interface
var _ Component = (*image)(nil)

////////////////////////////////////////////////////////////////////////////////
// CONSTRUCTOR

// Image creates a Bootstrap image element with src and optional alt text
// The src parameter sets the image source URL, and options can be used to style the image
//
// Example:
//
//	Image("photo.jpg", WithAriaLabel("Profile photo"))
//	Image("logo.png", WithClass("img-fluid"))
//	Image("thumbnail.jpg", WithClass("img-thumbnail"))
func Image(src string, opt ...Opt) *image {
	c := newComponent(ImageComponent, dom.GetWindow().Document().CreateElement("IMG"))

	// Set src attribute
	if src != "" {
		c.root.SetAttribute("src", src)
	}

	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	return &image{
		component: *c,
	}
}
