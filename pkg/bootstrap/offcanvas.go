package bootstrap

import (
	"fmt"

	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type offcanvas struct {
	mvc.View
}

var _ mvc.ViewWithVisibility = (*offcanvas)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewOffcanvas = "mvc-bs-offcanvas"
)

const (
	templateOffcanvas = `
		<div class="offcanvas offcanvas-start" tabindex="-1">
			<slot name="header"></slot>
			<slot></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewOffcanvas, newOffcanvasFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Offcanvas(id string, args ...any) *offcanvas {
	o := new(offcanvas)
	o.View = mvc.NewViewExEx(o, ViewOffcanvas, templateOffcanvas, mvc.WithAttr("id", id), mvc.WithAttr("aria-labelledby", id+"-label"), args)
	return o
}

func newOffcanvasFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newOffcanvasFromElement: invalid tag name %q", tagName))
	}
	o := new(offcanvas)
	o.View = mvc.NewViewWithElement(o, element)
	return o
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (offcanvas *offcanvas) Self() mvc.View {
	return offcanvas
}

func (offcanvas *offcanvas) Header(children ...any) mvc.View {
	return offcanvas.View.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("offcanvas-header"), children))
}

func (offcanvas *offcanvas) Content(children ...any) mvc.View {
	return offcanvas.View.ReplaceSlot("", mvc.HTML("DIV", mvc.WithClass("offcanvas-body"), children))
}

func (offcanvas *offcanvas) Apply(opts ...mvc.Opt) mvc.View {
	offcanvas.View.Apply(opts...)
	// TODO: Clear inline positioning styles directly via the style object
	return offcanvas
}

func (offcanvas *offcanvas) Show() mvc.ViewWithVisibility {
	// TODO: Clear issues for non-wasm builds
	if j := js.GetProto("bootstrap.Offcanvas"); j.IsUndefined() {
		panic("bootstrap.Offcanvas is undefined")
	} else {
		j.New("#" + offcanvas.ID()).Call("show")
	}
	return offcanvas
}

func (offcanvas *offcanvas) Hide() mvc.ViewWithVisibility {
	// TODO: Clear issues for non-wasm builds
	if j := js.GetProto("bootstrap.Offcanvas"); j.IsUndefined() {
		panic("bootstrap.Offcanvas is undefined")
	} else {
		j.New("#" + offcanvas.ID()).Call("hide")
	}
	return offcanvas
}

// Returns true if the view is visible
func (offcanvas *offcanvas) Visible() bool {
	return true
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

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
