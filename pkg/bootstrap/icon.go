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
	return mvc.NewView(new(icon), ViewIcon, "I", mvc.WithClass("bi-"+name), args)
}

func newIconFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "I" {
		panic(fmt.Sprintf("newIconFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(icon), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (icon *icon) Append(children ...any) mvc.View {
	panic("Append: not supported for icon")
}

func (icon *icon) SetView(view mvc.View) {
	icon.View = view
}
