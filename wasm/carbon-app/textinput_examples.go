package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const textInputStyle = "max-width:28rem;padding-bottom:var(--cds-spacing-07,2rem);"

func InputExamples() mvc.View {
	return ExamplePage("Text input",
		cds.LeadPara(
			`Inputs start with the generic `, cds.Code("cds.Input()"), ` constructor. `,
			`Specialized variants such as `, cds.Code("cds.EmailInput()"), `, `, cds.Code("cds.TelInput()"), `, `,
			cds.Code("cds.SearchInput()"), `, `, cds.Code("cds.RangeInput()"), `, and `, cds.Code("cds.SecureInput()"),
			` map to the Carbon input family while keeping the same label, helper text, and placeholder style where the component supports it. `,
			`For longer multi-line content, use `, cds.Code("cds.TextArea()"), ` on the same page instead of treating it as a separate form primitive.`,
		),
		cds.Heading(4, "Single-line inputs"),
		ExampleRow("Basic", Example_Input_001, "A standard Carbon input with label, helper text, and placeholder."),
		ExampleRow("Typed inputs", Example_Input_002, "EmailInput, TelInput, and SearchInput cover common field semantics."),
		ExampleRow("Secure and range", Example_Input_003, "SecureInput adds the password affordance; RangeInput uses Carbon number input steppers on the right."),
		ExampleRow("Interactive", Example_Input_004, "Listen for input events and read the value from the input host element."),
		cds.Heading(4, "Multi-line input"),
		ExampleRow("Text area — Basic", Example_TextArea_001, "A standard text area with label, helper text, placeholder, and a larger editable region."),
		ExampleRow("Text area — Different sizes", Example_TextArea_002, "Use rows to give the field an appropriate initial height for the expected amount of content."),
		ExampleRow("Text area — Read-only and disabled", Example_TextArea_003, "Read-only preserves access to the content; disabled communicates that editing is not available."),
		ExampleRow("Text area — Interactive", Example_TextArea_004, "Listen for input events and read the value from the textarea host element."),
	)
}

func Example_Input_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textInputStyle),
		cds.Input(
			"Email address",
			"Use your work email so notifications go to the right inbox.",
			cds.WithInputPlaceholder("name@company.com"),
			cds.WithInputType(cds.InputEmail),
		),
	), sourcecode()
}

func Example_Input_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textInputStyle+"display:flex;flex-direction:column;gap:var(--cds-spacing-05,1rem);"),
		cds.EmailInput(
			"Contact email",
			"We will use this address for approvals and alerts.",
			cds.WithInputPlaceholder("ops@company.com"),
		),
		cds.TelInput(
			"Support phone",
			"Enter a direct number for urgent issues.",
			cds.WithInputPlaceholder("+1 415 555 0134"),
		),
		cds.SearchInput(
			"Search",
			"Search uses Carbon's dedicated search component with built-in affordances.",
			cds.WithInputPlaceholder("Search teams, repos, or projects"),
		),
	), sourcecode()
}

func Example_Input_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textInputStyle+"display:flex;flex-direction:column;gap:var(--cds-spacing-05,1rem);"),
		cds.SecureInput(
			"Admin password",
			"SecureInput uses Carbon's password field with the reveal toggle on the right.",
			cds.WithInputPlaceholder("Enter a strong password"),
		),
		cds.RangeInput(
			"Seats",
			"RangeInput is backed by Carbon number input, with increment/decrement controls on the right.",
			cds.WithInputValue("25"),
			cds.WithInputMin("1"),
			cds.WithInputMax("500"),
			cds.WithInputStep("5"),
		),
	), sourcecode()
}

func Example_Input_004() (mvc.View, string) {
	status := cds.HelperText(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03,0.5rem);display:block;"),
		"Start typing…",
	)
	input := cds.Input(
		"Project name",
		"This example reads the live value from the component when the input event fires.",
		cds.WithInputPlaceholder("Retail operations portal"),
	)
	input.AddEventListener("input", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			value := el.Value()
			if value == "" {
				status.Content("Start typing…")
			} else {
				status.Content("Current value: ", cds.Strong(value))
			}
		}
	})
	return cds.Section(
		mvc.WithAttr("style", textInputStyle),
		input,
		status,
	), sourcecode()
}
