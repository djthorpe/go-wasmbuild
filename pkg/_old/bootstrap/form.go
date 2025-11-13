package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type form struct {
	component
}

type input struct {
	component
}

type label struct {
	component
}

// Ensure that form implements Component interface
var _ Component = (*form)(nil)
var _ Component = (*input)(nil)
var _ Component = (*label)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Form creates a new bootstrap form element
func Form(opt ...Opt) *form {
	// Create a new component
	c := newComponent(FormComponent, dom.GetWindow().Document().CreateElement("FORM"))

	// Apply options
	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	// Return the component
	return &form{*c}
}

// Input creates a new bootstrap form input element
func Input(name string, opt ...Opt) *input {
	// Create a new component
	elem := dom.GetWindow().Document().CreateElement("INPUT")

	// Set default attributes
	elem.SetAttribute("type", "text")
	elem.SetAttribute("class", "form-control")
	elem.SetAttribute("name", name)

	c := newComponent(InputComponent, elem)

	// Apply options (these can override the defaults)
	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	// Return the component
	return &input{*c}
}

// NumberInput creates a new bootstrap form number input element
func NumberInput(name string, opt ...Opt) *input {
	// Append the type="number" attribute to the options
	opts := append([]Opt{WithAttribute("type", "number")}, opt...)
	return Input(name, opts...)
}

// Label creates a new bootstrap form label element
func Label(text string, opt ...Opt) *label {
	// Create a new component
	elem := dom.GetWindow().Document().CreateElement("LABEL")

	// Set default class
	elem.SetAttribute("class", "form-label")

	c := newComponent(LabelComponent, elem)

	// Apply options
	if err := c.applyTo(c.root, opt...); err != nil {
		panic(err)
	}

	// Add text content
	if text != "" {
		elem.AppendChild(dom.GetWindow().Document().CreateTextNode(text))
	}

	// Return the component
	return &label{*c}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Value gets or sets the value of the input element.
// When called with no arguments, it returns the current value.
// When called with one argument, it sets the value and returns it.
func (i *input) Value(value ...string) string {
	if len(value) > 0 {
		i.root.SetAttribute("value", value[0])
		return value[0]
	}
	return i.root.GetAttribute("value")
}
