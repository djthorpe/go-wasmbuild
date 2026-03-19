package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type numberInput struct{ base }

var _ mvc.View = (*numberInput)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewNumberInput, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(numberInput), element, setView)
	}, EventInput, EventChange, EventInvalid, EventFocus, EventNoFocus)
}

// NumberInput returns a <cds-number-input> web component.
func NumberInput(args ...any) *numberInput {
	return mvc.NewView(new(numberInput), ViewNumberInput, "cds-number-input", setView, args).(*numberInput)
}

// AddEventListener registers an event handler on the number input.
// Carbon emits cds-number-input for user-initiated value changes, which is
// normalized here to both EventInput and EventChange.
func (n *numberInput) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	n.View.AddEventListener(numberInputEvent(event), handler)
	return n
}

// RemoveEventListener removes an event handler from the number input.
func (n *numberInput) RemoveEventListener(event string) mvc.View {
	n.View.RemoveEventListener(numberInputEvent(event))
	return n
}

// Label returns the number input label attribute when explicitly set.
func (n *numberInput) Label() string {
	return n.Root().GetAttribute("label")
}

// SetLabel sets the number input label attribute.
func (n *numberInput) SetLabel(label string) *numberInput {
	n.Root().SetAttribute("label", label)
	return n
}

// HelperText returns the helper text attribute when explicitly set.
func (n *numberInput) HelperText() string {
	return n.Root().GetAttribute("helper-text")
}

// SetHelperText sets or clears the helper text attribute.
func (n *numberInput) SetHelperText(text string) *numberInput {
	if text == "" {
		n.Root().RemoveAttribute("helper-text")
	} else {
		n.Root().SetAttribute("helper-text", text)
	}
	return n
}

// InvalidText returns the invalid text attribute when explicitly set.
func (n *numberInput) InvalidText() string {
	return n.Root().GetAttribute("invalid-text")
}

// SetInvalidText sets or clears the invalid text attribute.
func (n *numberInput) SetInvalidText(text string) *numberInput {
	if text == "" {
		n.Root().RemoveAttribute("invalid-text")
	} else {
		n.Root().SetAttribute("invalid-text", text)
	}
	return n
}

// Required reports whether the number input is marked required.
func (n *numberInput) Required() bool {
	return n.Root().HasAttribute("required")
}

// SetRequired marks the number input as required or optional.
func (n *numberInput) SetRequired(required bool) *numberInput {
	if required {
		n.Root().SetAttribute("required", "")
	} else {
		n.Root().RemoveAttribute("required")
	}
	return n
}

// Value returns the current number input value.
func (n *numberInput) Value() string {
	if value := n.Root().Value(); value != "" {
		return value
	}
	return n.Root().GetAttribute("value")
}

// SetValue sets the number input value.
func (n *numberInput) SetValue(value string) *numberInput {
	n.Root().SetValue(value)
	n.Root().SetAttribute("value", value)
	return n
}

// Min returns the minimum value attribute.
func (n *numberInput) Min() string {
	return n.Root().GetAttribute("min")
}

// SetMin sets the minimum value attribute.
func (n *numberInput) SetMin(value string) *numberInput {
	if value == "" {
		n.Root().RemoveAttribute("min")
	} else {
		n.Root().SetAttribute("min", value)
	}
	return n
}

// Max returns the maximum value attribute.
func (n *numberInput) Max() string {
	return n.Root().GetAttribute("max")
}

// SetMax sets the maximum value attribute.
func (n *numberInput) SetMax(value string) *numberInput {
	if value == "" {
		n.Root().RemoveAttribute("max")
	} else {
		n.Root().SetAttribute("max", value)
	}
	return n
}

// Step returns the step attribute.
func (n *numberInput) Step() string {
	return n.Root().GetAttribute("step")
}

// SetStep sets the step attribute.
func (n *numberInput) SetStep(value string) *numberInput {
	if value == "" {
		n.Root().RemoveAttribute("step")
	} else {
		n.Root().SetAttribute("step", value)
	}
	return n
}

// AllowEmpty reports whether an empty string is considered valid.
func (n *numberInput) AllowEmpty() bool {
	return n.Root().HasAttribute("allow-empty")
}

// SetAllowEmpty allows or disallows an empty string value.
func (n *numberInput) SetAllowEmpty(allow bool) *numberInput {
	if allow {
		n.Root().SetAttribute("allow-empty", "")
	} else {
		n.Root().RemoveAttribute("allow-empty")
	}
	return n
}

// HideSteppers reports whether the up/down steppers are hidden.
func (n *numberInput) HideSteppers() bool {
	return n.Root().HasAttribute("hide-steppers")
}

// SetHideSteppers shows or hides the up/down steppers.
func (n *numberInput) SetHideSteppers(hidden bool) *numberInput {
	if hidden {
		n.Root().SetAttribute("hide-steppers", "")
	} else {
		n.Root().RemoveAttribute("hide-steppers")
	}
	return n
}

// CheckValidity evaluates the number input constraints and updates the
// component invalid state.
func (n *numberInput) CheckValidity() bool {
	return checkNumberInputValidity(n)
}

// SetCustomValidity updates the number input invalid state directly. Pass an
// empty string to clear the invalid UI.
func (n *numberInput) SetCustomValidity(message string) *numberInput {
	setNumberInputCustomValidity(n, message)
	return n
}

func numberInputEvent(event string) string {
	switch event {
	case EventInput, EventChange:
		return numberInputChangeEvent
	default:
		return event
	}
}
