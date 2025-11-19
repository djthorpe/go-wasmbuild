package bootstrap

import (

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type icon struct {
	BootstrapView
}

var _ mvc.View = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewIcon = "mvc-bs-icon"
)

func init() {
	mvc.RegisterView(ViewIcon, newIconFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Icon(name string, args ...any) *icon {
	i := new(icon)
	i.BootstrapView.View = mvc.NewView(i, ViewIcon, "I", mvc.WithClass("bi-"+name), args)
	return i
}

func newIconFromElement(element Element) mvc.View {
	if element.TagName() != "I" {
		return nil
	}
	i := new(icon)
	i.BootstrapView.View = mvc.NewViewWithElement(i, element)
	return i
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (icon *icon) Append(children ...any) mvc.View {
	panic("Append: not supported for icon")
}

func (icon *icon) Self() mvc.View {
	return icon
}
