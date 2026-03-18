package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func HeadingView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Headings")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeG10),
			headingScaleStory(),
		),
	}
}

func headingScaleStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:0.75rem"),
		carbon.Head(1, "Heading 1"),
		carbon.Head(2, "Heading 2"),
		carbon.Head(3, "Heading 3"),
		carbon.Head(4, "Heading 4"),
		carbon.Head(5, "Heading 5"),
		carbon.Head(6, "Heading 6"),
	)

	return storybook.Story(
		"Heading Scale",
		"Carbon headings map semantic levels 1 through 6 onto the design system heading scale.",
		canvas,
		nil,
	)
}
