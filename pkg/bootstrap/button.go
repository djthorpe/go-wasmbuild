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
	mvc.View
}

type buttongroup struct {
	mvc.View
}

var _ mvc.View = (*button)(nil)
var _ mvc.View = (*buttongroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewButton      = "mvc-bs-button"
	ViewButtonGroup = "mvc-bs-buttongroup"

	// The prefix class for outline buttons
	viewOutlineButtonClassPrefix = "btn-outline"
)

func init() {
	mvc.RegisterView(ViewButton, newButtonFromElement)
	mvc.RegisterView(ViewButtonGroup, newButtonGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Button(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(button), ViewButton, "BUTTON", append([]mvc.Opt{mvc.WithAttr("type", "button"), mvc.WithClass("btn"), mvc.WithClass("btn-primary")}, opt...)...)
}

func OutlineButton(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(button), ViewButton, "BUTTON", append([]mvc.Opt{mvc.WithAttr("type", "button"), mvc.WithClass("btn", "btn-outline"), mvc.WithClass("btn-outline-primary")}, opt...)...)
}

func CloseButton(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(button), ViewButton, "BUTTON", append([]mvc.Opt{mvc.WithAttr("type", "button"), mvc.WithClass("btn-close"), mvc.WithAriaLabel("close")}, opt...)...)
}

func ButtonGroup(opt ...mvc.Opt) mvc.View {
	opts := append([]mvc.Opt{mvc.WithAttr("role", "group"), mvc.WithClass("btn-group")}, opt...)
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", opts...)
}

func VButtonGroup(opt ...mvc.Opt) mvc.View {
	opts := append([]mvc.Opt{mvc.WithAttr("role", "group"), mvc.WithClass("btn-group-vertical")}, opt...)
	return mvc.NewView(new(buttongroup), ViewButtonGroup, "DIV", opts...)
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

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (b *button) SetView(view mvc.View) {
	b.View = view
}

func (b *buttongroup) SetView(view mvc.View) {
	b.View = view
}

func (b *button) Append(children ...any) mvc.View {
	// Close buttons cannot have children
	if b.Root().ClassList().Contains("btn-close") {
		panic("Append: not supported for close button")
	}
	// Call superclass
	return b.View.Append(children...)
}

func (b *buttongroup) Append(children ...any) mvc.View {
	// Button groups can only include buttons
	// Call superclass
	return b.View.Append(children...)
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
