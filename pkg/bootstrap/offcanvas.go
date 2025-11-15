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

type offcanvas struct {
	mvc.View
}

var _ mvc.View = (*offcanvas)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewOffcanvas = "mvc-bs-offcanvas"
)

func init() {
	mvc.RegisterView(ViewOffcanvas, newOffcanvasFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

/*
func Offcanvas(id string, args ...any) *offcanvas {
	header := mvc.HTML("DIV", mvc.WithClass("offcanvas-header"))
	body := mvc.HTML("DIV", mvc.WithClass("offcanvas-body"))
	return mvc.NewViewEx(new(offcanvas), ViewOffcanvas, "DIV", header, body, nil, nil, mvc.WithAttr("id", id), mvc.WithClass("offcanvas", "offcanvas-start"), mvc.WithAttr("tabindex", "-1"), mvc.WithAttr("aria-labelledby", id+"-label"), args).(*offcanvas)
}
*/

func newOffcanvasFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newOffcanvasFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(offcanvas), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (offcanvas *offcanvas) SetView(view mvc.View) {
	offcanvas.View = view
}

func WithOffcanvas(id string) mvc.Opt {
	return func(opts mvc.OptSet) error {
		if opts.Name() != ViewButton {
			return fmt.Errorf("WithOffcanvas: invalid view type %q", opts.Name())
		}
		if err := mvc.WithAttr("data-bs-toggle", "offcanvas")(opts); err != nil {
			return err
		}
		if err := mvc.WithAttr("data-bs-target", "#"+id)(opts); err != nil {
			return err
		}
		return mvc.WithAttr("aria-controls", id)(opts)
	}
}
