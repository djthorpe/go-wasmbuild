package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Text() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("text_examples.md"),
		bs.HRule(),
		bs.Heading(3, "Text Examples"),
		bs.Heading(4, "Paragraph", mvc.WithClass("mt-4")), Example(Example_Text_001),
		bs.Heading(4, "Lead Paragraph", mvc.WithClass("mt-4")), Example(Example_Text_002),
		bs.Heading(4, "Blockquote", mvc.WithClass("mt-4")), Example(Example_Text_003),
		bs.Heading(4, "Code Blocks", mvc.WithClass("mt-4")), Example(Example_Text_009),
		bs.Heading(4, "Headings", mvc.WithClass("mt-4")), Example(Example_Text_010),
		bs.Heading(4, "Color", mvc.WithClass("mt-4")), Example(Example_Text_004),
		bs.Heading(4, "Inline Styles", mvc.WithClass("mt-4")), Example(Example_Text_005),
		bs.Heading(4, "Markdown", mvc.WithClass("mt-4")), Example(Example_Text_006),
		bs.Heading(4, "Position", mvc.WithClass("mt-4")), Example(Example_Text_007),
		bs.Heading(4, "Link", mvc.WithClass("mt-4")), Example(Example_Text_008),
	)
}

func Example_Text_001() (mvc.View, string) {
	return bs.Para(
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
		quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis 
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	), sourcecode()
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
		aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	).Label(
		"Said someone very important",
	), sourcecode()
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
		bs.Markdown("This is some _markdown_ formatted content with ~~deleted~~ and **strong** text"),
	), sourcecode()
}

func Example_Text_007() (mvc.View, string) {
	return bs.Container(mvc.WithClass("my-2"),
		bs.Heading(5, "Start Aligned", bs.WithPosition(bs.Start)),
		bs.Para("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", bs.WithPosition(bs.Start)),
		bs.Heading(5, "Center Aligned", bs.WithPosition(bs.Center)),
		bs.Para("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", bs.WithPosition(bs.Center)),
		bs.Heading(5, "End Aligned", bs.WithPosition(bs.End)),
		bs.Para("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", bs.WithPosition(bs.End)),
	), sourcecode()
}

func Example_Text_008() (mvc.View, string) {
	return bs.Container(mvc.WithClass("my-2"),
		bs.Para(bs.Link("#link", "Default Link Color")),
		bs.Para(bs.Link("#link", "Secondary Link Color", bs.WithColor(bs.Secondary))),
		bs.Para(bs.Link("#link", "Danger Link Color", bs.WithColor(bs.Danger))),
		bs.Para(bs.IconLink("#link", bs.Icon("link"), "Icon Link")),
		bs.Para(bs.IconLink("#link", "Icon Link", bs.Icon("arrow-right"), bs.WithColor(bs.Danger))),
	), sourcecode()
}

func Example_Text_009() (mvc.View, string) {
	const codeBlock = `<html>
<head>
  <title>Example</title>
</head>
<body>
  <p>This is a paragraph.</p>
</body>
</html>`
	return bs.Container(mvc.WithClass("my-2"),
		bs.CodeBlock(codeBlock),
	), sourcecode()
}

func Example_Text_010() (mvc.View, string) {
	return bs.Container(mvc.WithClass("my-2"),
		bs.Heading(1, "Heading 1"),
		bs.Heading(2, "Heading 2"),
		bs.Heading(3, "Heading 3"),
		bs.Heading(4, "Heading 4"),
		bs.Heading(5, "Heading 5"),
		bs.Heading(6, "Heading 6"),
	), sourcecode()
}
