package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// text are elements that represent text views
type icon struct {
	mvc.View
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

func Icon(name string, args ...any) mvc.View {
	i := new(icon)
	i.View = mvc.NewView(i, ViewIcon, "I", mvc.WithClass("bi-"+name), args)
	return i
}

func newIconFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "I" {
		panic(fmt.Sprintf("newIconFromElement: invalid tag name %q", tagName))
	}
	i := new(icon)
	i.View = mvc.NewViewWithElement(i, element)
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
