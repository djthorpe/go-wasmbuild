package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	colStyleA = "background: var(--cds-layer-01, #f4f4f4); padding: var(--cds-spacing-04, 0.75rem); outline: 1px solid var(--cds-border-subtle-01, #e0e0e0);"
	colStyleB = "background: var(--cds-layer-02, #e0e0e0); padding: var(--cds-spacing-04, 0.75rem); outline: 1px solid var(--cds-border-subtle-01, #c6c6c6);"
)

func GridExamples() mvc.View {
	return ExamplePage("Grid",
		cds.LeadPara("Carbon uses a 16-column CSS grid. Columns are expressed as spans (e.g. col-span-4) and collapse naturally on smaller viewports. No row wrapper is needed."),
		cds.Grid(
			// 4 equal quarters (4+4+4+4 = 16)
			cds.Col4(mvc.WithAttr("style", colStyleA), cds.LabelText("4 col")),
			cds.Col4(mvc.WithAttr("style", colStyleB), cds.LabelText("4 col")),
			cds.Col4(mvc.WithAttr("style", colStyleA), cds.LabelText("4 col")),
			cds.Col4(mvc.WithAttr("style", colStyleB), cds.LabelText("4 col")),
			// two halves (8+8 = 16)
			cds.Col8(mvc.WithAttr("style", colStyleB), cds.LabelText("8 col")),
			cds.Col8(mvc.WithAttr("style", colStyleA), cds.LabelText("8 col")),
			// sidebar + content (4+12 = 16)
			cds.Col4(mvc.WithAttr("style", colStyleA), cds.LabelText("4 col sidebar")),
			cds.Col12(mvc.WithAttr("style", colStyleB), cds.LabelText("12 col content")),
			// full width
			cds.Col16(mvc.WithAttr("style", colStyleA), cds.LabelText("16 col full width")),
		),
	)
}
