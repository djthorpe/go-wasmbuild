package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func LinkExamples() mvc.View {
	return ExamplePage("Link",
		cds.LeadPara("Carbon provides two link variants: a standalone link for navigation actions, and an inline link that sits within the flow of body text."),
		ExampleRow("Link", Example_Link_001),
		ExampleRow("InlineLink", Example_Link_002),
	)
}

func Example_Link_001() (mvc.View, string) {
	return cds.Para(
		cds.Link("https://carbondesignsystem.com", "Carbon Design System"),
	), sourcecode()
}

func Example_Link_002() (mvc.View, string) {
	return cds.Para(
		"Visit the ",
		cds.InlineLink("https://carbondesignsystem.com", "Carbon Design System"),
		" for full documentation.",
	), sourcecode()
}
