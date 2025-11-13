package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TextExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Text Examples"), bs.HRule(),
		bs.Heading(3, "Paragraph", mvc.WithClass("mt-5")), Example(Example_Text_001),
		bs.Heading(3, "Lead Paragraph", mvc.WithClass("mt-5")), Example(Example_Text_002),
		bs.Heading(3, "Blockquote", mvc.WithClass("mt-5")), Example(Example_Text_003),
		bs.Heading(3, "Colored Paragraphs", mvc.WithClass("mt-5")), Example(Example_Text_004),
		bs.Heading(3, "Inline Text Styles", mvc.WithClass("mt-5")), Example(Example_Text_005),
		bs.Heading(3, "Links", mvc.WithClass("mt-5")), Example(Example_Text_006),
	)
}

func Example_Text_001() (mvc.View, string) {
	return bs.Para(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
		quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis 
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`), sourcecode()
}

func Example_Text_002() (mvc.View, string) {
	return bs.LeadPara(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
		quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis 
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`), sourcecode()
}

func Example_Text_003() (mvc.View, string) {
	return bs.Blockquote(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
		quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis 
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`), sourcecode()
}

func Example_Text_004() (mvc.View, string) {
	return bs.Container(
		bs.Para(bs.WithColor(bs.Primary), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Secondary), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Info), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Warning), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Success), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Danger), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		bs.Para(bs.WithColor(bs.Dark), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
	), sourcecode()
}

func Example_Text_005() (mvc.View, string) {
	return bs.Container(mvc.WithClass("my-2"),
		bs.Para(bs.Deleted("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
		bs.Para(bs.Highlighted("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
		bs.Para(bs.Strong("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
		bs.Para(bs.Smaller("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
		bs.Para(bs.Em("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
		bs.Para(bs.Code("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")),
	), sourcecode()
}

func Example_Text_006() (mvc.View, string) {
	return bs.Container(mvc.WithClass("my-2"),
		bs.Para(bs.Link("#link", "Default Link Color")),
		bs.Para(bs.Link("#link", "Secondary Link Color", bs.WithColor(bs.Secondary))),
		bs.Para(bs.Link("#link", "Danger Link Color", bs.WithColor(bs.Danger))),
	), sourcecode()
}
