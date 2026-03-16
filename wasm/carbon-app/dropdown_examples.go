package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const dropdownStyle = "max-width:28rem;padding-bottom:var(--cds-spacing-07,2rem);"

func DropdownExamples() mvc.View {
	return cds.Section(
		cds.LeadPara(
			`Dropdown uses the `, cds.Code("cds-dropdown"), ` family. `,
			`Use it when you want Carbon's richer custom picker rather than the simpler select control. `,
			`It works well for curated option lists where the selected item should be reflected in the trigger button.`,
		),
		ExampleRow("Basic", Example_Dropdown_001, "A standard dropdown with label, helper text, placeholder, and options."),
		ExampleRow("Preselected", Example_Dropdown_002, "Set the current value up front to reflect an existing choice."),
		ExampleRow("Disabled", Example_Dropdown_003, "Disabled dropdowns communicate that the list is currently unavailable."),
		ExampleRow("Interactive", Example_Dropdown_004, "Listen for cds-dropdown-selected and read the selected value from the host element."),
	)
}

func Example_Dropdown_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", dropdownStyle),
		cds.Dropdown(
			"Workspace",
			"Choose the workspace that should receive this update.",
			cds.WithDropdownPlaceholder("Choose an option"),
			cds.DropdownItem("Design system", "design-system"),
			cds.DropdownItem("Platform", "platform"),
			cds.DropdownItem("Operations", "operations"),
		),
	), sourcecode()
}

func Example_Dropdown_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", dropdownStyle),
		cds.Dropdown(
			"Severity",
			"Use a preselected item when the user is editing an existing value.",
			cds.WithDropdownValue("high"),
			cds.WithDropdownPlaceholder("Choose a severity"),
			cds.DropdownItem("Low", "low"),
			cds.DropdownItem("Medium", "medium"),
			cds.DropdownItem("High", "high", cds.WithDropdownItemSelected()),
		),
	), sourcecode()
}

func Example_Dropdown_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", dropdownStyle),
		cds.Dropdown(
			"Archived project",
			"Disabled dropdowns keep the current context visible while blocking interaction.",
			cds.WithDropdownValue("locked"),
			cds.WithDropdownDisabled(),
			cds.DropdownItem("Active", "active"),
			cds.DropdownItem("Locked", "locked", cds.WithDropdownItemSelected()),
		),
	), sourcecode()
}

func Example_Dropdown_004() (mvc.View, string) {
	status := cds.HelperText(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03,0.5rem);display:block;"),
		"Choose an option…",
	)
	field := cds.Dropdown(
		"Team",
		"This example reads the selected value when Carbon's dropdown event fires.",
		cds.WithDropdownPlaceholder("Choose a team"),
		cds.DropdownItem("Design", "design"),
		cds.DropdownItem("Engineering", "engineering"),
		cds.DropdownItem("Operations", "operations"),
	)
	field.AddEventListener("cds-dropdown-selected", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			value := el.Value()
			if value == "" {
				status.Content("Choose an option…")
			} else {
				status.Content("Selected value: ", cds.Strong(value))
			}
		}
	})
	return cds.Section(
		mvc.WithAttr("style", dropdownStyle),
		field,
		status,
	), sourcecode()
}
