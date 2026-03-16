package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const selectStyle = "max-width:28rem;padding-bottom:var(--cds-spacing-07,2rem);"

func SelectExamples() mvc.View {
	return ExamplePage("Select",
		cds.LeadPara(
			`Select uses the `, cds.Code("cds-select"), ` family. `,
			`Use them when the user should choose one item from a known set of options. `,
			`Keep the label explicit, provide helper text when the choice needs explanation, and use a placeholder when the user must make an active selection.`,
		),
		ExampleRow("Basic", Example_Select_001, "A standard select control with label, helper text, placeholder, and options."),
		ExampleRow("Grouped options", Example_Select_002, "Use option groups when the list is long or naturally clustered."),
		ExampleRow("Read-only and disabled", Example_Select_003, "Read-only preserves the current value; disabled communicates that the control is unavailable."),
		ExampleRow("Interactive", Example_Select_004, "Listen for cds-select-selected and read the selected value from the host element."),
	)
}

func Example_Select_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", selectStyle),
		cds.Select(
			"Environment",
			"Choose the deployment environment for this release.",
			cds.WithSelectPlaceholder("Select an environment"),
			cds.SelectItem("Development", "dev"),
			cds.SelectItem("Staging", "staging"),
			cds.SelectItem("Production", "prod"),
		),
	), sourcecode()
}

func Example_Select_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", selectStyle),
		cds.Select(
			"Region",
			"Options are grouped to make large lists easier to scan.",
			cds.WithSelectPlaceholder("Select a region"),
			cds.SelectGroup("North America",
				cds.SelectItem("United States", "us"),
				cds.SelectItem("Canada", "ca"),
			),
			cds.SelectGroup("Europe",
				cds.SelectItem("Germany", "de"),
				cds.SelectItem("United Kingdom", "uk"),
			),
		),
	), sourcecode()
}

func Example_Select_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", selectStyle+"display:flex;flex-direction:column;gap:var(--cds-spacing-05,1rem);"),
		cds.Select(
			"Current plan",
			"Read-only preserves the selected value without allowing a change.",
			cds.WithSelectValue("team"),
			cds.WithSelectReadOnly(),
			cds.SelectItem("Starter", "starter"),
			cds.SelectItem("Team", "team", cds.WithSelectItemSelected()),
			cds.SelectItem("Enterprise", "enterprise"),
		),
		cds.Select(
			"Archived workspace",
			"Disabled selectors indicate that no interaction is currently available.",
			cds.WithSelectValue("archived"),
			cds.WithSelectDisabled(),
			cds.SelectItem("Active", "active"),
			cds.SelectItem("Archived", "archived", cds.WithSelectItemSelected()),
		),
	), sourcecode()
}

func Example_Select_004() (mvc.View, string) {
	status := cds.HelperText(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03,0.5rem);display:block;"),
		"Choose an option…",
	)
	field := cds.Select(
		"Team",
		"This example reads the selected value when Carbon's select event fires.",
		cds.WithSelectPlaceholder("Select a team"),
		cds.SelectItem("Design", "design"),
		cds.SelectItem("Engineering", "engineering"),
		cds.SelectItem("Operations", "operations"),
	)
	field.AddEventListener("cds-select-selected", func(e dom.Event) {
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
		mvc.WithAttr("style", selectStyle),
		field,
		status,
	), sourcecode()
}
