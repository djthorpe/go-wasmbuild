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
	mvc.ViewWithCaption
	caption Element
}

type buttongroup struct {
	mvc.View
}

type buttontoolbar struct {
	mvc.View
}

var _ mvc.ViewWithCaption = (*button)(nil)
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
)

func init() {
	mvc.RegisterView(ViewButton, newButtonFromElement)
	mvc.RegisterView(ViewButtonGroup, newButtonGroupFromElement)
	mvc.RegisterView(ViewButtonToolbar, newButtonToolbarFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Button(args ...any) *button {
	caption := mvc.Placeholder()
	view := mvc.NewViewEx(
		new(button), ViewButton, "BUTTON",
		nil, nil, nil, caption,
		mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-primary"), args,
	).(*button)
	view.caption = caption
	return view
}

func OutlineButton(args ...any) *button {
	return mvc.NewView(new(button), ViewButton, "BUTTON", mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-outline", "btn-outline-primary"), args).(*button)
}

func CloseButton(args ...any) *button {
	return mvc.NewView(new(button), ViewButton, "BUTTON", mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-close"), mvc.WithAriaLabel("close"), args).(*button)
}

func ButtonToolbar(args ...any) mvc.View {
	return mvc.NewView(new(buttontoolbar), ViewButtonToolbar, "DIV", mvc.WithAttr("role", "toolbar"), mvc.WithClass("btn-toolbar"), args)
}

func ButtonGroup(args ...any) mvc.View {
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", mvc.WithAttr("role", "group"), mvc.WithClass("btn-group"), args)
}

func VButtonGroup(args ...any) mvc.View {
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", mvc.WithAttr("role", "group"), mvc.WithClass("btn-group-vertical"), args)
}

func newButtonFromElement(element Element) mvc.View {
	if element.TagName() != "BUTTON" {
		return nil
	}
	return mvc.NewViewWithElement(new(button), element)
}

func newButtonGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(buttongroup), element)
}

func newButtonToolbarFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(buttontoolbar), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (b *button) SetView(view mvc.View) {
	b.ViewWithCaption = view.(mvc.ViewWithCaption)
}

func (b *buttongroup) SetView(view mvc.View) {
	b.View = view
}

func (b *buttontoolbar) SetView(view mvc.View) {
	b.View = view
}

func (b *button) Caption(children ...any) mvc.ViewWithCaption {
	// Replace whatever is in the caption with the indicator badge
	// TODO: We actualy have to replace the caption element, which is currently a placeholder
	// and set the attribute so we can re-create the view from the element
	b.Root().ClassList().Add("position-relative")
	b.caption.ReplaceWith(mvc.HTML("SPAN", mvc.WithClass("position-absolute", "top-0", "start-100", "translate-middle", "badge", "rounded-pill", "bg-danger")))
	return b
}

func (b *button) Append(children ...any) mvc.View {
	// Close buttons cannot have children
	if b.Root().ClassList().Contains("btn-close") {
		panic("Append: not supported for close button")
	}
	// Call superclass
	return b.ViewWithCaption.Append(children...)
}

func (b *buttongroup) Append(children ...any) mvc.View {
	// TODO: Button groups can only include buttons
	// Call superclass
	return b.View.Append(children...)
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

/*
// Return true if button is disabled
func (b *button) Disabled() bool {
	return b.Root().HasAttribute("disabled")
}

// Return true if button is active
func (b *button) Active() bool {
	return b.Root().ClassList().Contains("active")
}

// Return elements which are active in the button group
func (b *buttongroup) Active() []Element {
	var elements []Element

	// Find active elements
	child := b.Root().FirstElementChild()
	for child != nil {
		if child.ClassList().Contains("active") {
			elements = append(elements, child)
		}
		child = child.NextElementSibling()
	}
	return elements
}

// Return elements which are disabled in the button group
func (b *buttongroup) Disabled() []Element {
	var elements []Element

	// Find disabled elements
	child := b.Root().FirstElementChild()
	for child != nil {
		if child.HasAttribute("disabled") {
			elements = append(elements, child)
		}
		child = child.NextElementSibling()
	}
	return elements
}
*/
