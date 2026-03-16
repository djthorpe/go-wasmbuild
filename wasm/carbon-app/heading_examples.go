package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func HeadingExamples() mvc.View {
	return ExamplePage("Heading",
		cds.LeadPara("Carbon uses a productive heading scale for UI contexts. Each level maps to a fixed type token — size, weight, and line-height are all prescribed."),
		ExampleRow("Heading 1", Example_Heading_001),
		ExampleRow("Heading 2", Example_Heading_002),
		ExampleRow("Heading 3", Example_Heading_003),
		ExampleRow("Heading 4", Example_Heading_004),
		ExampleRow("Heading 5", Example_Heading_005),
		ExampleRow("Heading 6", Example_Heading_006),
	)
}

func Example_Heading_001() (mvc.View, string) {
	return cds.Heading(1, "42px, light"), sourcecode()
}

func Example_Heading_002() (mvc.View, string) {
	return cds.Heading(2, "32px"), sourcecode()
}

func Example_Heading_003() (mvc.View, string) {
	return cds.Heading(3, "28px"), sourcecode()
}

func Example_Heading_004() (mvc.View, string) {
	return cds.Heading(4, "20px"), sourcecode()
}

func Example_Heading_005() (mvc.View, string) {
	return cds.Heading(5, "16px, semibold"), sourcecode()
}

func Example_Heading_006() (mvc.View, string) {
	return cds.Heading(6, "14px, semibold"), sourcecode()
}
