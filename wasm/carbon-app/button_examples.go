package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	// btnPreviewStyle is applied to example sections that show standalone buttons.
	btnPreviewStyle = "display:flex;flex-wrap:wrap;align-items:flex-start;gap:var(--cds-spacing-03,0.5rem);padding:var(--cds-spacing-04,0.75rem);border:1px solid var(--cds-border-subtle-01,#e0e0e0);"
)

func ButtonExamples() mvc.View {
	return ExamplePage("Button",
		cds.LeadPara(
			`Buttons use the `, cds.Code("cds-button"), ` web component. `,
			`The default kind is `, cds.Code("primary"), `. `,
			`Pass `, cds.Code("cds.WithButtonKind()"), ` to change style, `,
			cds.Code("cds.WithButtonSize()"), ` to change size, and `,
			`pass `, cds.Code("cds.Icon()"), ` as a child to add an icon.`,
		),
		ExampleRow("Kinds", Example_Button_001, "Seven visual styles: primary (default), secondary, tertiary (outlined), ghost (text-only), and three danger variants."),
		ExampleRow("Sizes", Example_Button_002, "Five sizes from compact sm (32px) to expressive 2xl (80px)."),
		ExampleRow("With Icon", Example_Button_003, "Pass cds.Icon() as a child; the slot is set automatically."),
		ExampleRow("Interactive", Example_Button_004, "Click events bubble from cds-button to the parent container."),
		ExampleRow("Disabled", Example_Button_005, "Add cds.WithDisabled() to prevent interaction."),
		ExampleRow("Button Set", Example_Button_006, "ButtonSet groups buttons flush in a horizontal row."),
	)
}

func Example_Button_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button("Primary"),
		cds.Button("Secondary", cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button("Tertiary", cds.WithButtonKind(cds.ButtonTertiary)),
		cds.Button("Ghost", cds.WithButtonKind(cds.ButtonGhost)),
		cds.Button("Danger", cds.WithButtonKind(cds.ButtonDanger)),
		cds.Button("Danger ghost", cds.WithButtonKind(cds.ButtonDangerGhost)),
	), sourcecode()
}

func Example_Button_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button("Small", cds.WithButtonSize(cds.ButtonSM)),
		cds.Button("Medium"),
		cds.Button("Large", cds.WithButtonSize(cds.ButtonLG)),
		cds.Button("Extra large", cds.WithButtonSize(cds.ButtonXL)),
		cds.Button("2XL", cds.WithButtonSize(cds.ButtonXXL)),
	), sourcecode()
}

func Example_Button_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button(cds.Icon("add"), "Add item"),
		cds.Button(cds.Icon("trash-can"), "Delete", cds.WithButtonKind(cds.ButtonDanger)),
		cds.Button(cds.Icon("download"), "Download", cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button(cds.Icon("settings"), "Settings", cds.WithButtonKind(cds.ButtonGhost)),
	), sourcecode()
}

func Example_Button_004() (mvc.View, string) {
	response := cds.Para(mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);"), "Click a button")
	return cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button("Save"),
		cds.Button("Cancel", cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button("Reset", cds.WithButtonKind(cds.ButtonGhost)),
		response,
	).AddEventListener("click", func(e dom.Event) {
		btn := mvc.ViewFromEvent(e)
		if btn != nil && btn.Name() == cds.ViewButton {
			response.Content("Clicked: ", cds.Strong(btn.Root().TextContent()))
		}
	}), sourcecode()
}

func Example_Button_005() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button("Primary", cds.WithDisabled()),
		cds.Button("Secondary", cds.WithButtonKind(cds.ButtonSecondary), cds.WithDisabled()),
		cds.Button("Danger", cds.WithButtonKind(cds.ButtonDanger), cds.WithDisabled()),
	), sourcecode()
}

func Example_Button_006() (mvc.View, string) {
	response := cds.Para(mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);"), "Click a button")
	return cds.Section(
		cds.ButtonSet(
			cds.Button("Save"),
			cds.Button("Cancel", cds.WithButtonKind(cds.ButtonSecondary)),
			cds.Button("Delete", cds.WithButtonKind(cds.ButtonDanger)),
		),
		response,
	).AddEventListener("click", func(e dom.Event) {
		btn := mvc.ViewFromEvent(e)
		if btn != nil && btn.Name() == cds.ViewButton {
			response.Content("Clicked: ", cds.Strong(btn.Root().TextContent()))
		}
	}), sourcecode()
}
