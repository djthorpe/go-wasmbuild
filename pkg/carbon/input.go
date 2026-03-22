package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type input struct {
	base
	changeBaseline string
}

var _ mvc.View = (*input)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewInput, func(element dom.Element) mvc.View {
		i := mvc.NewViewWithElement(new(input), element, setView).(*input)
		initializeInput(i)
		return i
	}, EventInput, EventChange, EventInvalid, EventFocus, EventNoFocus)
}

// Input returns a <cds-text-input> web component.
func Input(args ...any) *input {
	i := mvc.NewView(new(input), ViewInput, "cds-text-input", setView, args).(*input)
	initializeInput(i)
	return i
}

// AddEventListener registers an event handler on the input.
// EventChange is bridged from focus loss because the Carbon host element does
// not emit a native change event itself.
func (i *input) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	if event == EventChange {
		ensureInputChangeBridge(i)
	}
	i.View.AddEventListener(event, handler)
	return i
}

// Label returns the input label attribute when explicitly set.
func (i *input) Label() string {
	return i.Root().GetAttribute("label")
}

// SetLabel sets the input label attribute.
func (i *input) SetLabel(label string) *input {
	i.Root().SetAttribute("label", label)
	return i
}

// HelperText returns the helper text attribute when explicitly set.
func (i *input) HelperText() string {
	return i.Root().GetAttribute("helper-text")
}

// SetHelperText sets or clears the helper text attribute.
func (i *input) SetHelperText(text string) *input {
	if text == "" {
		i.Root().RemoveAttribute("helper-text")
	} else {
		i.Root().SetAttribute("helper-text", text)
	}
	return i
}

// InvalidText returns the invalid text attribute when explicitly set.
func (i *input) InvalidText() string {
	return i.Root().GetAttribute("invalid-text")
}

// SetInvalidText sets or clears the invalid text attribute.
func (i *input) SetInvalidText(text string) *input {
	if text == "" {
		i.Root().RemoveAttribute("invalid-text")
	} else {
		i.Root().SetAttribute("invalid-text", text)
	}
	return i
}

// Required reports whether the input is marked required.
func (i *input) Required() bool {
	return i.Root().HasAttribute("required")
}

// SetRequired marks the input as required or optional.
func (i *input) SetRequired(required bool) *input {
	if required {
		i.Root().SetAttribute("required", "")
	} else {
		i.Root().RemoveAttribute("required")
	}
	return i
}

// Value returns the input value.
func (i *input) Value() string {
	if value := i.Root().Value(); value != "" {
		return value
	}
	return i.Root().GetAttribute("value")
}

// SetValue sets the input value.
func (i *input) SetValue(value string) *input {
	i.Root().SetValue(value)
	i.Root().SetAttribute("value", value)
	i.changeBaseline = value
	return i
}

// CheckValidity evaluates the input constraints and updates the component's
// invalid state.
func (i *input) CheckValidity() bool {
	return checkInputValidity(i)
}

// SetCustomValidity updates the input invalid state directly. Pass an empty
// string to clear the invalid UI.
func (i *input) SetCustomValidity(message string) *input {
	setInputCustomValidity(i, message)
	return i
}

func initializeInput(i *input) {
	if i == nil {
		return
	}
	i.changeBaseline = i.Value()
	ensureInputChangeBridge(i)
}
