package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func CodeExamples() mvc.View {
	return ExamplePage("Code snippet",
		cds.LeadPara("Carbon's code snippet component comes in three forms: inline for embedding within prose, single-line for short commands, and multi-line for longer blocks. All variants include a copy-to-clipboard button."),
		ExampleRow("Inline", Example_Code_001),
		ExampleRow("Single-line", Example_Code_002),
		ExampleRow("Multi-line", Example_Code_003),
	)
}

func Example_Code_001() (mvc.View, string) {
	return cds.Para(
		"Use ", cds.CodeInline("cds.CodeInline(...)"), " to embed code within a sentence.",
	), sourcecode()
}

func Example_Code_002() (mvc.View, string) {
	return cds.CodeSingle("npm install @carbon/web-components"), sourcecode()
}

func Example_Code_003() (mvc.View, string) {
	return cds.CodeMulti(`{
  "dependencies": {
    "@carbon/web-components": "^2.0.0",
    "@carbon/styles": "^1.0.0"
  }
}`), sourcecode()
}
