package main

import (
	"strconv"

	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const textAreaStyle = "max-width:32rem;padding-bottom:var(--cds-spacing-07,2rem);"

func TextAreaExamples() mvc.View {
	return cds.Section(
		cds.LeadPara(
			`Text areas use the `, cds.Code("cds-textarea"), ` web component. `,
			`Use them when the user needs to enter longer multi-line content such as notes, descriptions, or comments. `,
			`As with other Carbon form controls, keep the label visible and use helper text for guidance rather than hiding instructions in the placeholder.`,
		),
		ExampleRow("Basic", Example_TextArea_001, "A standard text area with label, helper text, placeholder, and a larger editable region."),
		ExampleRow("Different sizes", Example_TextArea_002, "Use rows to give the field an appropriate initial height for the expected amount of content."),
		ExampleRow("Read-only and disabled", Example_TextArea_003, "Read-only preserves access to the content; disabled communicates that editing is not available."),
		ExampleRow("Interactive", Example_TextArea_004, "Listen for input events and read the value from the textarea host element."),
	)
}

func Example_TextArea_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textAreaStyle),
		cds.TextArea(
			"Project description",
			"Summarize the purpose, audience, and scope in a few sentences.",
			cds.WithTextAreaPlaceholder("Describe what this workspace is for and who it supports"),
			cds.WithTextAreaRows("5"),
		),
	), sourcecode()
}

func Example_TextArea_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textAreaStyle+"display:flex;flex-direction:column;gap:var(--cds-spacing-05,1rem);"),
		cds.TextArea(
			"Short note",
			"A compact text area works for concise updates.",
			cds.WithTextAreaPlaceholder("Add a short note"),
			cds.WithTextAreaRows("3"),
		),
		cds.TextArea(
			"Detailed summary",
			"A taller text area gives users enough room for longer structured content.",
			cds.WithTextAreaPlaceholder("Capture the full summary, constraints, and next steps"),
			cds.WithTextAreaRows("8"),
		),
	), sourcecode()
}

func Example_TextArea_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", textAreaStyle+"display:flex;flex-direction:column;gap:var(--cds-spacing-05,1rem);"),
		cds.TextArea(
			"Read-only notes",
			"Read-only fields can still be focused and copied.",
			cds.WithTextAreaValue("Incident review complete. Findings were shared with the operations team and the remediation plan was approved."),
			cds.WithTextAreaRows("4"),
			cds.WithTextAreaReadOnly(),
		),
		cds.TextArea(
			"Archived feedback",
			"Disabled fields indicate that editing is currently unavailable.",
			cds.WithTextAreaValue("This record is archived and can no longer be changed."),
			cds.WithTextAreaRows("3"),
			cds.WithTextAreaDisabled(),
		),
	), sourcecode()
}

func Example_TextArea_004() (mvc.View, string) {
	status := cds.HelperText(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-03,0.5rem);display:block;"),
		"Start typing…",
	)
	field := cds.TextArea(
		"Release notes",
		"This example reads the live value from the component when the input event fires.",
		cds.WithTextAreaPlaceholder("Summarize the notable changes in this release"),
		cds.WithTextAreaRows("5"),
	)
	field.AddEventListener("input", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			value := el.Value()
			if value == "" {
				status.Content("Start typing…")
			} else {
				status.Content("Current length: ", cds.Strong(strconv.Itoa(len(value))), " characters")
			}
		}
	})
	return cds.Section(
		mvc.WithAttr("style", textAreaStyle),
		field,
		status,
	), sourcecode()
}
