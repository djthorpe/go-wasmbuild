package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	mvc.View
}

type buttongroup struct {
	mvc.View
}

type buttontoolbar struct {
	mvc.View
}

var _ mvc.View = (*button)(nil)
var _ mvc.View = (*buttongroup)(nil)
var _ mvc.View = (*buttontoolbar)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// The template for buttons
	templateButton = `
		<button type="button" class="btn btn-primary text-nowrap">
			<slot></slot>
			<slot name="label"></slot>
		</button>
	`
)

func init() {
	mvc.RegisterView(ViewButton, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(button), element, setView)
	})
	mvc.RegisterView(ViewButtonGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(buttongroup), element, setView)
	})
	mvc.RegisterView(ViewButtonToolbar, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(buttontoolbar), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Button(args ...any) *button {
	return mvc.NewView(new(button), ViewButton, templateButton, setView, args).(*button)
}

func OutlineButton(args ...any) *button {
	return Button(mvc.WithoutClass("btn-primary"), mvc.WithClass(viewOutlineButtonClassPrefix, "btn-outline-primary"), args)
}

func CloseButton(args ...any) *button {
	return mvc.NewView(new(button), ViewButton, "BUTTON", setView, mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-close"), mvc.WithAriaLabel("close"), args).(*button)
}

func ButtonGroup(args ...any) mvc.View {
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", setView, mvc.WithAttr("role", "group"), mvc.WithClass("btn-group"), args)
}

/*
func ButtonToolbar(args ...any) mvc.View {
	return mvc.NewView(new(buttontoolbar), ViewButtonToolbar, "DIV", mvc.WithAttr("role", "toolbar"), mvc.WithClass("btn-toolbar"), args)
}

func VButtonGroup(args ...any) mvc.View {
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", mvc.WithAttr("role", "group"), mvc.WithClass("btn-group-vertical"), args)
}
*/

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (b *button) Label(children ...any) mvc.View {
	b.Root().ClassList().Add("position-relative")
	b.ReplaceSlot("label", mvc.HTML("SPAN", mvc.WithClass("position-absolute", "top-0", "start-100", "translate-middle", "badge", "rounded-pill", "bg-danger"), children))
	return b
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithSubmit() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewButton {
			return dom.ErrInternalAppError.With("WithSubmit: option only valid for button views")
		}
		return mvc.WithAttr("type", "submit")(o)
	}
}
