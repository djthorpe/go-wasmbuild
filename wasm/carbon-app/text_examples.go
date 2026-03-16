package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TextExamples() mvc.View {
	return cds.Section(
		cds.LeadPara("Carbon's type scale covers body text, supporting copy, and inline emphasis. Each token is tuned for readability at its intended size and context."),
		ExampleRow("Lead Paragraph", Example_Text_001),
		ExampleRow("Paragraph", Example_Text_002),
		ExampleRow("Compact Paragraph", Example_Text_003),
		ExampleRow("Caption", Example_Text_004),
		ExampleRow("Helper Text", Example_Text_005),
		ExampleRow("Label", Example_Text_006),
		ExampleRow("Inline Styles", Example_Text_007),
	)
}

func Example_Text_001() (mvc.View, string) {
	return cds.LeadPara(lorem), sourcecode()
}

func Example_Text_002() (mvc.View, string) {
	return cds.Para(lorem), sourcecode()
}

func Example_Text_003() (mvc.View, string) {
	return cds.CompactPara(lorem), sourcecode()
}

func Example_Text_004() (mvc.View, string) {
	return cds.Caption("For image captions and secondary annotations."), sourcecode()
}

func Example_Text_005() (mvc.View, string) {
	return cds.HelperText("For form field hints and contextual guidance."), sourcecode()
}

func Example_Text_006() (mvc.View, string) {
	return cds.LabelText("For field labels and tight UI text."), sourcecode()
}

func Example_Text_007() (mvc.View, string) {
	return cds.Para(
		cds.Strong("strong"), " ",
		cds.Em("em"), " ",
		cds.Deleted("deleted"), " ",
		cds.Highlighted("highlighted"), " ",
		cds.Smaller("smaller"), " ",
		cds.Code("code"),
	), sourcecode()
}
