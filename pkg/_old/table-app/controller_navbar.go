package main

import (
	// Packages

	"fmt"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Nav struct {
	view  Component
	table Component
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewNav(table *Table) *Nav {
	controller := new(Nav)

	// Create the view
	controller.view = bs.NavBar(
		bs.WithClass("navbar-expand-lg"),
		bs.WithTheme(bs.DARK),
		bs.WithColor(bs.DANGER),
	).Header(
		bs.Icon("people-fill"),
		" Employee Data Example",
	)

	// Append the dropdowns
	controller.view.Append(
		controller.colorDropdown(),
		controller.optionsDropdown(),
	)

	// Set the table view
	controller.table = table.Table()

	// Return the controller
	return controller
}

func (controller *Nav) colorDropdown() Component {
	return bs.NavDropdown("Colour").Append(
		bs.NavItem("#", "Default").Apply(bs.WithID("default")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.PRIMARY), bs.WithClass("border")).Append(" "), " Primary").Apply(bs.WithID("primary")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.SECONDARY), bs.WithClass("border")).Append(" "), " Secondary").Apply(bs.WithID("secondary")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.SUCCESS), bs.WithClass("border")).Append(" "), " Success").Apply(bs.WithID("success")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.DANGER), bs.WithClass("border")).Append(" "), " Danger").Apply(bs.WithID("danger")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.WARNING), bs.WithClass("border")).Append(" "), " Warning").Apply(bs.WithID("warning")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.INFO), bs.WithClass("border")).Append(" "), " Info").Apply(bs.WithID("info")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.LIGHT), bs.WithClass("border")).Append(" "), " Light").Apply(bs.WithID("light")),
		bs.NavItem("#", bs.PillBadge(bs.WithBackground(bs.DARK), bs.WithClass("border")).Append(" "), " Dark").Apply(bs.WithID("dark")),
	).AddEventListener("click", func(target Node) {
		item := target.Component()
		if item == nil || item.Name() != string(bs.NavItemComponent) {
			return
		}
		if item.ID() == "default" {
			controller.table.Apply(bs.WithColor(bs.TRANSPARENT))
			return
		}
		controller.table.Apply(bs.WithColor(bs.Color(item.ID())))
	})
}

func (controller *Nav) optionsDropdown() Component {
	return bs.NavDropdown("Options").Append(
		// Border section
		bs.NavDropdownHeader("Border"),
		bs.NavItem("#", "Default").Apply(bs.WithID("border-default")),
		bs.NavItem("#", "Bordered").Apply(bs.WithID("border-bordered")),
		bs.NavItem("#", "Borderless").Apply(bs.WithID("border-borderless")),
		bs.NavDivider(false),

		// Size section
		bs.NavDropdownHeader("Size"),
		bs.NavItem("#", "Default").Apply(bs.WithID("size-default")),
		bs.NavItem("#", "Small").Apply(bs.WithID("size-small")),
		bs.NavDivider(false),

		// Group Divider section
		bs.NavDropdownHeader("Group Divider"),
		bs.NavItem("#", "Off").Apply(bs.WithID("divider-off")),
		bs.NavItem("#", "On").Apply(bs.WithID("divider-on")),
		bs.NavDivider(false),

		// Stripe section
		bs.NavDropdownHeader("Stripe"),
		bs.NavItem("#", "None").Apply(bs.WithID("stripe-none")),
		bs.NavItem("#", "Row").Apply(bs.WithID("stripe-row")),
		bs.NavItem("#", "Column").Apply(bs.WithID("stripe-column")),
		bs.NavDivider(false),

		// Hover section
		bs.NavDropdownHeader("Hover"),
		bs.NavItem("#", "Off").Apply(bs.WithID("hover-off")),
		bs.NavItem("#", "On").Apply(bs.WithID("hover-on")),
	).AddEventListener("click", func(target Node) {
		item := target.Component()
		if item == nil || item.Name() != string(bs.NavItemComponent) {
			return
		}

		// Apply the table options
		switch item.ID() {
		case "border-default":
			controller.table.Apply(bs.WithBordered()) // TODO: Remove both bordered and borderless
		case "border-bordered":
			controller.table.Apply(bs.WithBordered())
		case "border-borderless":
			controller.table.Apply(bs.WithBorderless())
		case "size-default":
			controller.table.Apply(bs.WithSize(bs.SizeDefault))
		case "size-small":
			controller.table.Apply(bs.WithSize(bs.SizeSmall))
		case "divider-off":
			controller.table.Apply(bs.WithoutGroupDivider())
		case "divider-on":
			controller.table.Apply(bs.WithGroupDivider())
		case "stripe-none":
			controller.table.Apply(bs.WithoutStriped())
		case "stripe-row":
			controller.table.Apply(bs.WithStripedRows())
		case "stripe-column":
			controller.table.Apply(bs.WithStripedColumns())
		case "hover-off":
			controller.table.Apply(bs.WithoutHover())
		case "hover-on":
			controller.table.Apply(bs.WithHover())
		}

		// Print out the current options
		fmt.Println(controller.table.Element().ClassList().Values())
	})
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Nav) View() Component {
	return this.view
}
