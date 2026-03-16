package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const iconStyle = "display:inline-flex;align-items:center;gap:var(--cds-spacing-03,0.5rem);margin-right:var(--cds-spacing-05,1rem);margin-bottom:var(--cds-spacing-04,0.75rem);"

func IconExamples() mvc.View {
	return ExamplePage("Icon",
		cds.LeadPara(
			`Icons are rendered using `,
			cds.InlineLink("https://carbondesignsystem.com/elements/icons/library/", "Carbon Icons"),
			`. Create an icon by name with `, cds.Code("cds.Icon()"), `. The icon inherits `,
			`the current text color, so tinting is done with any CSS color or Carbon token.`,
		),
		ExampleRow("Common Icons", Example_Icon_001),
		ExampleRow("Sizes", Example_Icon_002),
		ExampleRow("Color", Example_Icon_003),
	)
}

func Example_Icon_001() (mvc.View, string) {
	names := []string{
		"add", "close", "checkmark", "warning", "information",
		"search", "settings", "home", "user", "notification",
		"download", "upload", "edit", "trash-can", "copy",
	}
	items := make([]any, 0, len(names))
	for _, name := range names {
		items = append(items,
			mvc.HTML("span",
				mvc.WithAttr("style", iconStyle),
				cds.Icon(name),
				mvc.HTML("span", mvc.WithClass("cds--caption-01"), name),
			),
		)
	}
	return cds.Section(
		mvc.WithAttr("style", "display:flex;flex-wrap:wrap;"),
		items,
	), sourcecode()
}

func Example_Icon_002() (mvc.View, string) {
	sizes := []string{"16", "20", "24", "32"}
	items := make([]any, 0, len(sizes))
	for _, sz := range sizes {
		items = append(items,
			mvc.HTML("span",
				mvc.WithAttr("style", iconStyle),
				cds.Icon("add", mvc.WithAttr("size", sz)),
				mvc.HTML("span", mvc.WithClass("cds--caption-01"), sz+"px"),
			),
		)
	}
	return cds.Section(
		mvc.WithAttr("style", "display:flex;align-items:flex-end;flex-wrap:wrap;"),
		items,
	), sourcecode()
}

func Example_Icon_003() (mvc.View, string) {
	type coloredIcon struct {
		name  string
		color string
		label string
	}
	icons := []coloredIcon{
		{"warning--filled", "var(--cds-support-warning,#f1c21b)", "Warning"},
		{"error--filled", "var(--cds-support-error,#da1e28)", "Error"},
		{"checkmark--filled", "var(--cds-support-success,#24a148)", "Success"},
		{"information--filled", "var(--cds-support-info,#0043ce)", "Info"},
		{"star--filled", "var(--cds-interactive,#0f62fe)", "Star"},
	}
	items := make([]any, 0, len(icons))
	for _, ic := range icons {
		items = append(items,
			mvc.HTML("span",
				mvc.WithAttr("style", iconStyle+"color:"+ic.color+";"),
				cds.Icon(ic.name, mvc.WithAttr("size", "24")),
				mvc.HTML("span", mvc.WithClass("cds--caption-01"), ic.label),
			),
		)
	}
	return cds.Section(
		mvc.WithAttr("style", "display:flex;flex-wrap:wrap;"),
		items,
	), sourcecode()
}
