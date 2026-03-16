package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func FormsExamples() mvc.View {
	return cds.Section(
		cds.LeadPara(
			`Carbon forms are built around a few core rules: every field needs a clear label, `,
			`helper text should explain how to complete the field, validation should appear close to the control, `,
			`and the form should lead to one obvious primary action.`,
		),
		ExampleRow("Basic structure", Example_Forms_001, "How Carbon expects fields, helper text, and actions to be organized."),
		ExampleRow("Labels and helper text", Example_Forms_002, "The label tells users what the field is; helper text tells them how to complete it."),
		ExampleRow("Validation and errors", Example_Forms_003, "Errors should appear inline, next to the field that needs attention, with clear corrective guidance."),
		ExampleRow("Layout and actions", Example_Forms_004, "Keep forms linear, group related fields, and end with one dominant primary button."),
	)
}

func formNotes(title, intro string, bullets ...string) mvc.View {
	items := make([]any, 0, len(bullets))
	for _, bullet := range bullets {
		items = append(items, mvc.HTML("li", bullet))
	}
	children := []any{
		cds.Heading(5, title),
		cds.Para(intro),
	}
	if len(items) > 0 {
		children = append(children, mvc.HTML("ul", items...))
	}
	return cds.Section(children...)
}

func Example_Forms_001() (mvc.View, string) {
	return cds.Section(
		formNotes(
			"Carbon form anatomy",
			"A Carbon form usually flows top to bottom: fields first, supporting guidance near each field, then actions at the end.",
			"Use a visible label for every input; do not rely on placeholder text as the only label.",
			"Keep helper text directly below the field so the instruction stays attached to the control.",
			"Place optional and required cues in a consistent position across the entire form.",
			"Finish with one primary action and only add a secondary action when it is genuinely useful.",
		),
		cds.CodeMulti(`cds.Section(
  cds.LabelText("Email address"),
  cds.HelperText("Use your work email."),
	cds.Input(
		"Email address",
		"Use your work email.",
		cds.WithInputPlaceholder("name@company.com"),
	),
  cds.Button("Save", cds.WithButtonKind(cds.ButtonPrimary)),
)`),
	), sourcecode()
}

func Example_Forms_002() (mvc.View, string) {
	return formNotes(
		"Labels and helper text",
		"Carbon treats labels as mandatory structure, not optional decoration. Helper text should reduce ambiguity before the user makes a mistake.",
		"Write labels as nouns or short noun phrases: “Email address”, “Team name”, “Start date”.",
		"Use helper text for formatting hints, limits, or short contextual guidance.",
		"Avoid repeating the label inside helper text; add new information instead.",
		"If the field is optional, say so consistently rather than inventing special wording for each input.",
	), sourcecode()
}

func Example_Forms_003() (mvc.View, string) {
	return cds.Section(
		formNotes(
			"Validation and errors",
			"Validation should help the user recover quickly. Carbon patterns work best when the message says what is wrong and what the user should do next.",
			"Show validation near the relevant field instead of collecting all errors in a distant summary only.",
			"Use error text that is specific: “Enter a valid email address” is better than “Invalid input”.",
			"Trigger validation after interaction or on submit; avoid overwhelming the user with errors before they begin.",
			"Preserve the entered value whenever possible so the user can correct rather than retype.",
		),
		cds.CodeSingle(`Error text: Enter a valid email address in the format name@company.com`),
	), sourcecode()
}

func Example_Forms_004() (mvc.View, string) {
	return cds.Section(
		formNotes(
			"Layout and actions",
			"Carbon forms are easiest to complete when the layout follows the task order and only asks for information that is necessary at that moment.",
			"Group related inputs with headings or fieldsets so large forms feel chunked and scannable.",
			"Prefer a single-column layout unless the relationship between fields is obvious and the screen is wide enough.",
			"Use one primary action for submit, and keep cancel or secondary actions visually quieter.",
			"On submit, disable duplicate submissions or show loading feedback when the action takes noticeable time.",
		),
		cds.ButtonSet(
			cds.Button("Save changes"),
			cds.Button("Cancel", cds.WithButtonKind(cds.ButtonSecondary)),
		),
	), sourcecode()
}
