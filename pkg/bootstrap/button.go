package bootstrap

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	BootstrapView
}

type buttongroup struct {
	BootstrapView
}

type buttontoolbar struct {
	BootstrapView
}

var _ mvc.View = (*button)(nil)
var _ mvc.View = (*buttongroup)(nil)
var _ mvc.View = (*buttontoolbar)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewButton        = "mvc-bs-button"
	ViewButtonGroup   = "mvc-bs-buttongroup"
	ViewButtonToolbar = "mvc-bs-buttontoolbar"

	// The prefix class for outline buttons
	viewOutlineButtonClassPrefix = "btn-outline"

	// The template for buttons
	templateButton = `
		<button type="button" class="btn btn-primary text-nowrap"><slot></slot><slot name="label"></slot></button>
	`
)

func init() {
	mvc.RegisterView(ViewButton, newButtonFromElement)
	mvc.RegisterView(ViewButtonGroup, newButtonGroupFromElement)
	mvc.RegisterView(ViewButtonToolbar, newButtonToolbarFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Button(args ...any) *button {
	b := new(button)
	b.BootstrapView.View = mvc.NewViewExEx(b, ViewButton, templateButton, args)
	return b
}

func OutlineButton(args ...any) *button {
	b := new(button)
	b.View = mvc.NewView(b, ViewButton, "BUTTON", mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-outline", "btn-outline-primary"), args)
	return b
}

func CloseButton(args ...any) *button {
	b := new(button)
	b.View = mvc.NewView(b, ViewButton, "BUTTON", mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-close"), mvc.WithAriaLabel("close"), args)
	return b
}

func ButtonToolbar(args ...any) *buttontoolbar {
	b := new(buttontoolbar)
	allArgs := append([]any{mvc.WithAttr("role", "toolbar"), mvc.WithClass("btn-toolbar")}, args...)
	b.BootstrapView.View = mvc.NewView(b, ViewButtonToolbar, "div", allArgs...)
	return b
}

func ButtonGroup(args ...any) *buttongroup {
	b := new(buttongroup)
	b.View = mvc.NewView(b, ViewButtonGroup, "DIV", mvc.WithAttr("role", "group"), mvc.WithClass("btn-group"), args)
	return b
}

func VButtonGroup(args ...any) *buttongroup {
	b := new(buttongroup)
	b.View = mvc.NewView(b, ViewButtonGroup, "DIV", mvc.WithAttr("role", "group"), mvc.WithClass("btn-group-vertical"), args)
	return b
}

func newButtonFromElement(element Element) mvc.View {
	if element.TagName() != "BUTTON" {
		return nil
	}
	b := new(button)
	b.BootstrapView.View = mvc.NewViewWithElement(b, element)
	return b
}

func newButtonGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	b := new(buttongroup)
	b.BootstrapView.View = mvc.NewViewWithElement(b, element)
	return b
}

func newButtonToolbarFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	b := new(buttontoolbar)
	b.BootstrapView.View = mvc.NewViewWithElement(b, element)
	return b
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (b *button) Self() mvc.View {
	return b
}

func (b *buttongroup) Self() mvc.View {
	return b
}

func (b *buttontoolbar) Self() mvc.View {
	return b
}

func (b *button) Label(children ...any) *button {
	b.Root().ClassList().Add("position-relative")
	b.ReplaceSlot("label", mvc.HTML("SPAN", mvc.WithClass("position-absolute", "top-0", "start-100", "translate-middle", "badge", "rounded-pill", "bg-danger"), children))
	return b
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithSubmit() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewButton {
			return ErrInternalAppError.With("WithSubmit: option only valid for button views")
		}
		return mvc.WithAttr("type", "submit")(o)
	}
}
