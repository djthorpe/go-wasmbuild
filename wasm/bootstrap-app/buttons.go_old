package main

import (
	// Packages

	"fmt"

	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func Buttons() mvc.View {
	return bs.Container(bs.WithSize(bs.SizeFluid)).Append(
		// Heading with title
		bs.Heading(1, mvc.WithClass("mt-4", "mb-4")).Append("Bootstrap Button Examples"),

		// Button Variants
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Button Variants"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("Primary"),
			bs.Button(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Secondary"),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("Success"),
			bs.Button(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("Danger"),
			bs.Button(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("Warning"),
			bs.Button(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("Info"),
			bs.Button(bs.WithColor(bs.LIGHT), mvc.WithClass("me-2")).Append("Light"),
			bs.Button(bs.WithColor(bs.DARK), mvc.WithClass("me-2")).Append("Dark"),
		),

		// Outline Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Outline Buttons"),
		bs.Para(
			bs.OutlineButton(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("Primary"),
			bs.OutlineButton(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Secondary"),
			bs.OutlineButton(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("Success"),
			bs.OutlineButton(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("Danger"),
			bs.OutlineButton(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("Warning"),
			bs.OutlineButton(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("Info"),
			bs.OutlineButton(bs.WithColor(bs.LIGHT), mvc.WithClass("me-2")).Append("Light"),
			bs.OutlineButton(bs.WithColor(bs.DARK), mvc.WithClass("me-2")).Append("Dark"),
		),

		// Button Sizes
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Button Sizes"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), bs.WithSize(bs.SizeLarge), mvc.WithClass("me-2")).Append("Large button"),
			bs.Button(bs.WithColor(bs.SECONDARY), bs.WithSize(bs.SizeLarge), mvc.WithClass("me-2")).Append("Large button"),
		),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("Default button"),
			bs.Button(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Default button"),
		),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), bs.WithSize(bs.SizeSmall), mvc.WithClass("me-2")).Append("Small button"),
			bs.Button(bs.WithColor(bs.SECONDARY), bs.WithSize(bs.SizeSmall), mvc.WithClass("me-2")).Append("Small button"),
		),

		// Disabled Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Disabled Buttons"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2"), mvc.WithAttr("disabled", "")).Append("Primary button"),
			bs.Button(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2"), mvc.WithAttr("disabled", "")).Append("Button"),
			bs.OutlineButton(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2"), mvc.WithAttr("disabled", "")).Append("Primary button"),
			bs.OutlineButton(bs.WithColor(bs.SECONDARY), mvc.WithAttr("disabled", "")).Append("Button"),
		),

		// Buttons with Icons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Buttons with Icons"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append(
				bs.Icon("download"),
				" Download",
			),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(
				bs.Icon("check-circle"),
				" Confirm",
			),
			bs.Button(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append(
				bs.Icon("trash"),
				" Delete",
			),
			bs.Button(bs.WithColor(bs.INFO)).Append(
				bs.Icon("info-circle"),
				" Info",
			),
		),

		// Icon-only Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Icon-only Buttons"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append(bs.Icon("heart")),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(bs.Icon("star")),
			bs.Button(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append(bs.Icon("bell")),
			bs.Button(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append(bs.Icon("x-circle")),
			bs.OutlineButton(bs.WithColor(bs.SECONDARY)).Append(bs.Icon("gear")),
		),

		// Block Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Block Buttons"),
		createBlockButtons(),

		// Button Groups - Basic
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Button Groups"),
		bs.Para(
			createBasicButtonGroup().AddEventListener("click", func(node Node) {
				view := mvc.ViewFromNode(node).(mvc.ViewWithState)
				if view != nil && view.Name() == bs.ViewButton {
					view.Opts(bs.WithActive(!view.Active()))
				}

				// Get the group view
				group := mvc.ViewFromNode(node.ParentElement()).(mvc.ViewWithGroupState)
				if group != nil {
					fmt.Println("Active buttons in group:", group.Active())
				}
			}),
		),

		// Button Groups - Mixed Styles
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Mixed Styles"),
		bs.Para(
			createMixedButtonGroup(),
		),

		// Button Groups - Outlined
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Outlined Button Group"),
		bs.Para(
			createOutlinedButtonGroup(),
		),

		// Button Groups - Sizing
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Button Group Sizing"),
		createButtonGroupSizing(),

		// Button Toolbar
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Button Toolbar"),
		createButtonToolbar(),

		// Vertical Button Group
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Vertical Button Group"),
		bs.Para(
			createVerticalButtonGroup(),
		),

		// Toggle Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Toggle Buttons"),
		createToggleButtons(),

		// Button Group with Icons
		bs.Heading(3, mvc.WithClass("mt-3")).Append("Button Group with Icons"),
		bs.Para(
			createIconButtonGroup(),
		),
	)
}

// Helper function to create block buttons
func createBlockButtons() mvc.View {
	div := mvc.Div()
	div.Opts(mvc.WithClass("d-grid", "gap-2", "col-6", "mx-auto"))

	btn1 := bs.Button(bs.WithColor(bs.PRIMARY))
	btn1.Append("Block button 1")

	btn2 := bs.Button(bs.WithColor(bs.PRIMARY))
	btn2.Append("Block button 2")

	div.Append(btn1, btn2)
	return div
}

// Helper function to create a basic button group
func createBasicButtonGroup() mvc.View {
	group := bs.ButtonGroup(mvc.WithAttr("aria-label", "Basic example"))

	left := bs.Button(bs.WithColor(bs.PRIMARY))
	left.Append("Left")

	middle := bs.Button(bs.WithColor(bs.PRIMARY))
	middle.Append("Middle")

	right := bs.Button(bs.WithColor(bs.PRIMARY))
	right.Append("Right")

	group.Append(left, middle, right)
	return group
}

// Helper function to create a mixed styles button group
func createMixedButtonGroup() mvc.View {
	group := bs.ButtonGroup(mvc.WithAttr("aria-label", "Mixed styles example"))

	left := bs.Button(bs.WithColor(bs.DANGER))
	left.Append("Left")

	middle := bs.Button(bs.WithColor(bs.WARNING))
	middle.Append("Middle")

	right := bs.Button(bs.WithColor(bs.SUCCESS))
	right.Append("Right")

	group.Append(left, middle, right)
	return group
}

// Helper function to create an outlined button group
func createOutlinedButtonGroup() mvc.View {
	group := bs.ButtonGroup(mvc.WithAttr("aria-label", "Outlined example"))

	left := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	left.Append("Left")

	middle := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	middle.Append("Middle")

	right := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	right.Append("Right")

	group.Append(left, middle, right)
	return group
}

// Helper function to create button groups with different sizes
func createButtonGroupSizing() mvc.View {
	container := mvc.Div()

	// Large
	large := bs.ButtonGroup(mvc.WithClass("btn-group-lg"), mvc.WithAttr("aria-label", "Large button group"))
	large.Append(
		createButtonForGroup(bs.PRIMARY, "Left"),
		createButtonForGroup(bs.PRIMARY, "Middle"),
		createButtonForGroup(bs.PRIMARY, "Right"),
	)
	large.Opts(mvc.WithClass("me-2"))

	// Default
	normal := bs.ButtonGroup(mvc.WithAttr("aria-label", "Default button group"))
	normal.Append(
		createButtonForGroup(bs.PRIMARY, "Left"),
		createButtonForGroup(bs.PRIMARY, "Middle"),
		createButtonForGroup(bs.PRIMARY, "Right"),
	)
	normal.Opts(mvc.WithClass("me-2"))

	// Small
	small := bs.ButtonGroup(mvc.WithClass("btn-group-sm"), mvc.WithAttr("aria-label", "Small button group"))
	small.Append(
		createButtonForGroup(bs.PRIMARY, "Left"),
		createButtonForGroup(bs.PRIMARY, "Middle"),
		createButtonForGroup(bs.PRIMARY, "Right"),
	)

	para := bs.Para(large, normal, small)
	container.Append(para)
	return container
}

// Helper function to create button toolbar
func createButtonToolbar() mvc.View {
	toolbar := mvc.Div()
	toolbar.Opts(mvc.WithAttr("role", "toolbar"), mvc.WithAttr("aria-label", "Toolbar with button groups"))

	// First group
	group1 := bs.ButtonGroup(mvc.WithAttr("aria-label", "First group"))
	group1.Opts(mvc.WithClass("me-2"))
	for i := 1; i <= 4; i++ {
		btn := bs.Button(bs.WithColor(bs.PRIMARY))
		btn.Append(string(rune('0' + i)))
		group1.Append(btn)
	}

	// Second group
	group2 := bs.ButtonGroup(mvc.WithAttr("aria-label", "Second group"))
	group2.Opts(mvc.WithClass("me-2"))
	for i := 5; i <= 7; i++ {
		btn := bs.Button(bs.WithColor(bs.SECONDARY))
		btn.Append(string(rune('0' + i)))
		group2.Append(btn)
	}

	// Third group
	group3 := bs.ButtonGroup(mvc.WithAttr("aria-label", "Third group"))
	btn := bs.Button(bs.WithColor(bs.INFO))
	btn.Append("8")
	group3.Append(btn)

	toolbar.Append(group1, group2, group3)
	return toolbar
}

// Helper function to create vertical button group
func createVerticalButtonGroup() mvc.View {
	group := bs.ButtonGroup(mvc.WithClass("btn-group-vertical"), mvc.WithAttr("aria-label", "Vertical button group"))

	buttons := []string{"Button", "Button", "Button", "Button"}
	for _, text := range buttons {
		btn := bs.Button(bs.WithColor(bs.PRIMARY))
		btn.Append(text)
		group.Append(btn)
	}

	return group
}

// Helper function to create toggle buttons
func createToggleButtons() mvc.View {
	container := mvc.Div()

	para := bs.Para()
	para.Opts(mvc.WithClass("d-inline-flex", "gap-1"))

	// Normal toggle
	toggle1 := bs.Button(mvc.WithAttr("data-bs-toggle", "button"))
	toggle1.Append("Toggle button")

	// Active toggle
	toggle2 := bs.Button(mvc.WithClass("active"), mvc.WithAttr("data-bs-toggle", "button"), mvc.WithAttr("aria-pressed", "true"))
	toggle2.Append("Active toggle button")

	// Disabled toggle
	toggle3 := bs.Button(mvc.WithAttr("disabled", ""), mvc.WithAttr("data-bs-toggle", "button"))
	toggle3.Append("Disabled toggle button")

	para.Append(toggle1, toggle2, toggle3)
	container.Append(para)

	// Colored toggles
	para2 := bs.Para()
	para2.Opts(mvc.WithClass("d-inline-flex", "gap-1", "mt-2"))

	colorToggle1 := bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithAttr("data-bs-toggle", "button"))
	colorToggle1.Append("Toggle button")

	colorToggle2 := bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("active"), mvc.WithAttr("data-bs-toggle", "button"), mvc.WithAttr("aria-pressed", "true"))
	colorToggle2.Append("Active toggle button")

	colorToggle3 := bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithAttr("disabled", ""), mvc.WithAttr("data-bs-toggle", "button"))
	colorToggle3.Append("Disabled toggle button")

	para2.Append(colorToggle1, colorToggle2, colorToggle3)
	container.Append(para2)

	return container
}

// Helper function to create icon button group
func createIconButtonGroup() mvc.View {
	group := bs.ButtonGroup(mvc.WithAttr("aria-label", "Icon button group"))

	btn1 := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	btn1.Append(bs.Icon("align-left"))

	btn2 := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	btn2.Append(bs.Icon("align-center"))

	btn3 := bs.OutlineButton(bs.WithColor(bs.PRIMARY))
	btn3.Append(bs.Icon("align-right"))

	group.Append(btn1, btn2, btn3)
	return group
}

// Helper function to create a button for button groups
func createButtonForGroup(color bs.Color, text string) mvc.View {
	btn := bs.Button(bs.WithColor(color))
	btn.Append(text)
	return btn
}
