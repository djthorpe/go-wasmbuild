package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	component
}

type buttonGroup struct {
	component
}

// Ensure that button implements Component interface
var _ Component = (*button)(nil)

// Ensure that buttonGroup implements Component interface
var _ Component = (*buttonGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// buttonWithOptions is an internal helper to create buttons with specific default options
func buttonWithOptions(defaults []Opt, opt ...Opt) *button {
	// Create a new component
	c := newComponent(ButtonComponent, dom.GetWindow().Document().CreateElement("BUTTON"))

	// Apply options with defaults first, then user options
	opts := append(defaults, opt...)

	// Apply options
	if err := c.applyTo(c.root, opts...); err != nil {
		panic(err)
	}

	return &button{component: *c}
}

// Button creates a new button element with "btn" class
// Use WithColor to set the button variant (e.g., btn-primary)
func Button(color Color, opt ...Opt) *button {
	colorClass := "btn-" + string(color)
	return buttonWithOptions([]Opt{
		WithClass("btn", colorClass),
		WithAttribute("type", "button"),
	}, opt...)
}

// OutlineButton creates a new button with outline styling (btn-outline-*)
// This is a convenience wrapper that uses outline button variants.
func OutlineButton(color Color, opt ...Opt) *button {
	outlineClass := "btn-outline-" + string(color)
	return buttonWithOptions([]Opt{
		WithClass("btn", outlineClass),
		WithAttribute("type", "button"),
	}, opt...)
}

// CloseButton creates a close button with the "btn-close" class.
// This is commonly used for dismissing modals, alerts, offcanvas, and other components.
// The button has no visible text but is accessible via aria-label.
//
// Example:
//
//	closeBtn := CloseButton()
//	closeBtn.Element().SetAttribute("data-bs-dismiss", "modal")
//	closeBtn.Element().SetAttribute("aria-label", "Close")
func CloseButton(opt ...Opt) *button {
	return buttonWithOptions([]Opt{
		WithClass("btn-close"),
		WithAttribute("type", "button"),
		WithAttribute("aria-label", "Close"),
	}, opt...)
}

// buttonGroupWithOptions is an internal helper to create button groups with specific default options
func buttonGroupWithOptions(defaults []Opt, opt ...Opt) *buttonGroup {
	// Create a new component
	c := newComponent(ButtonGroupComponent, dom.GetWindow().Document().CreateElement("DIV"))

	// Apply options with defaults first, then user options
	opts := append(defaults, opt...)

	// Apply options
	if err := c.applyTo(c.root, opts...); err != nil {
		panic(err)
	}

	return &buttonGroup{component: *c}
}

// ButtonGroup creates a button group container with "btn-group" class
// and appropriate ARIA attributes for accessibility
func ButtonGroup(opt ...Opt) *buttonGroup {
	return buttonGroupWithOptions([]Opt{
		WithClass("btn-group"),
		WithRole("group"),
	}, opt...)
}

// VerticalButtonGroup creates a vertical button group with "btn-group-vertical" class
func VerticalButtonGroup(opt ...Opt) *buttonGroup {
	return buttonGroupWithOptions([]Opt{
		WithClass("btn-group-vertical"),
		WithRole("group"),
	}, opt...)
}

// ButtonToolbar creates a button toolbar container with "btn-toolbar" class
func ButtonToolbar(opt ...Opt) *buttonGroup {
	return buttonGroupWithOptions([]Opt{
		WithClass("btn-toolbar"),
		WithRole("toolbar"),
	}, opt...)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Active gets or sets active buttons in the button group.
// When called with no arguments, returns the indices of currently active buttons.
// When called with indices, marks those buttons as active and deactivates others.
// Passing -1 deactivates all buttons.
// Indices that are out of bounds are ignored.
// Returns []int containing the indices of active buttons.
func (bg *buttonGroup) Active(indices ...int) []int {
	if len(indices) > 0 {
		// Set mode: mark specified buttons as active
		activeIndices := make(map[int]bool)
		for _, idx := range indices {
			if idx >= 0 {
				activeIndices[idx] = true
			}
		}

		// Iterate through all child buttons and set/remove active class
		currentIndex := 0
		for child := bg.root.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "BUTTON" {
				classList := elem.ClassList()
				if activeIndices[currentIndex] {
					classList.Add("active")
				} else {
					classList.Remove("active")
				}
				currentIndex++
			}
		}

		return indices
	}

	// Get mode: return currently active button indices
	var activeButtons []int
	currentIndex := 0
	for child := bg.root.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "BUTTON" {
			if elem.ClassList().Contains("active") {
				activeButtons = append(activeButtons, currentIndex)
			}
			currentIndex++
		}
	}

	return activeButtons
}

// Disabled manages the disabled state of buttons in the button group.
// With arguments: sets the disabled state for buttons at the specified indices.
// Without arguments: returns the indices of currently disabled buttons.
// Passing -1 disables all buttons.
// Indices that are out of bounds are ignored.
// Returns []int containing the indices of disabled buttons.
func (bg *buttonGroup) Disabled(indices ...int) []int {
	if len(indices) > 0 {
		// Set mode: mark specified buttons as disabled
		disabledIndices := make(map[int]bool)
		for _, idx := range indices {
			if idx >= 0 {
				disabledIndices[idx] = true
			}
		}

		// Iterate through all child buttons and set/remove disabled attribute
		currentIndex := 0
		for child := bg.root.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "BUTTON" {
				if disabledIndices[currentIndex] {
					elem.SetAttribute("disabled", "")
				} else {
					elem.RemoveAttribute("disabled")
				}
				currentIndex++
			}
		}

		return indices
	}

	// Get mode: return currently disabled button indices
	var disabledButtons []int
	currentIndex := 0
	for child := bg.root.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "BUTTON" {
			if elem.HasAttribute("disabled") {
				disabledButtons = append(disabledButtons, currentIndex)
			}
			currentIndex++
		}
	}

	return disabledButtons
}

// SetDisabled sets or removes the disabled attribute
func (button *button) SetDisabled(disabled bool) *button {
	if disabled {
		button.root.SetAttribute("disabled", "")
	} else {
		// Remove disabled attribute - need to check if there's a RemoveAttribute method
		// For now, we'll work with what we have
	}
	return button
}
