package bootstrap

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type offcanvas struct {
	mvc.View
	instance js.Value
}

var _ mvc.View = (*offcanvas)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateOffcanvas = `
		<div class="offcanvas offcanvas-start" tabindex="-1">
			<slot name="header"></slot>
			<slot></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewOffcanvas, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(offcanvas), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Offcanvas(id string, args ...any) *offcanvas {
	return mvc.NewView(new(offcanvas), ViewOffcanvas, templateOffcanvas, setView, mvc.WithAttr("id", id), mvc.WithAttr("aria-labelledby", id+"-label"), args).(*offcanvas)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (offcanvas *offcanvas) Header(children ...any) *offcanvas {
	offcanvas.View.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("offcanvas-header"), children))
	return offcanvas
}

func (offcanvas *offcanvas) Content(children ...any) mvc.View {
	offcanvas.View.ReplaceSlot("body", mvc.HTML("DIV", mvc.WithClass("offcanvas-body"), children))
	return offcanvas
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithOffcanvas returns an option which configures a button to open the offcanvas with the given ID
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

// Use WithOffcanvasScroll to enable <body> scrolling.
func WithOffcanvasScroll() mvc.Opt {
	return func(opts mvc.OptSet) error {
		if opts.Name() != ViewOffcanvas {
			return fmt.Errorf("WithOffcanvasScroll: invalid view type %q", opts.Name())
		}
		return mvc.WithAttr("data-bs-scroll", "true")(opts)
	}
}
