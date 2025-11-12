package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TextExamples() mvc.View {
	return bs.Container(
		ParaExample(),
		bs.HRule(),
		LeadExample(),
		bs.HRule(),
		ColorParaExample(),
	)
}

func ParaExample() mvc.View {
	return bs.Grid().Content(
		bs.Container(
			mvc.WithClass("my-2"),
			bs.Para("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		), bs.Container(
			bs.Para(`Paragraphs are created with the bs.Para() method.`),
			bs.CodeBlock(bs.WithColor(bs.Light), mvc.WithClass("p-3"), mvc.WithClass("border", "border-dark-subtle")).Content(
				`bs.Para("Lorem ipsum dolor sit amet, consectetur ... laborum.")`,
			),
		),
	)
}

func LeadExample() mvc.View {
	return bs.Grid().Content(
		bs.Container(mvc.WithClass("my-2")).Content(
			bs.LeadPara("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		), bs.Container().Content(
			bs.Para(
				`Use the LeadPara variation for lead paragraphs.`,
			),
			bs.CodeBlock(bs.WithColor(bs.Light), mvc.WithClass("p-3"), mvc.WithClass("border", "border-dark-subtle")).Content(
				`bs.LeadPara("Lorem ipsum dolor sit amet, consectetur ... laborum.")`,
			),
		),
	)
}

func ColorParaExample() mvc.View {
	return bs.Grid().Content(
		bs.Container(mvc.WithClass("my-2")).Content(
			bs.Para(bs.WithColor(bs.Primary), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Secondary), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Info), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Warning), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Success), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Danger), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			bs.Para(bs.WithColor(bs.Dark), "Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
		), bs.Container().Content(
			bs.Para(
				`Colored paragraphs can be created with the bs.WithColor option.`,
			),
			bs.CodeBlock(bs.WithColor(bs.Light), mvc.WithClass("p-3"), mvc.WithClass("border", "border-dark-subtle")).Content(
				`bs.Para(bs.WithColor(bs.Primary), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Secondary), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Info), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Warning), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Success), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Danger), "Lorem ipsum dolor sit amet...."),
bs.Para(bs.WithColor(bs.Dark), "Lorem ipsum dolor sit amet...."),`,
			),
		),
	)
}
