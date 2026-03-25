package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func ListView() []any {
	return []any{
		storybook.PageHeader("List", "List.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			unorderedListStory(),
			orderedListStory(),
		),
	}
}

func unorderedListStory() dom.Element {
	list := carbon.List(
		carbon.With(carbon.ListDisc),
		carbon.ListItem(carbon.Strong("Primary action"), carbon.Compact(" Validate the request before submission.")),
		carbon.ListItem(carbon.Strong("Secondary action"), carbon.Compact(" Keep the supporting detail concise.")),
		carbon.ListItem(carbon.Strong("Fallback"), carbon.Compact(" Preserve semantic list structure even with richer content.")),
	)
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		list,
	)

	return storybook.Story(
		"Unordered List",
		"Unordered lists are suited to grouped content where sequence does not carry meaning.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Bullet style", carbon.ListDisc, []carbon.Attr{carbon.ListDisc, carbon.ListCircle, carbon.ListSquare}, func(style carbon.Attr) {
			list.Apply(carbon.With(style)...)
		}),
	)
}

func orderedListStory() dom.Element {
	list := carbon.OrderedList(
		carbon.With(carbon.ListDecimal),
		carbon.ListItem("Create the page wrapper"),
		carbon.ListItem("Add the content stories"),
		carbon.ListItem("Link the docs in navigation"),
	)
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		list,
	)

	return storybook.Story(
		"Ordered List",
		"Ordered lists are suited to steps, ranking, and other sequences where position matters.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Marker style", carbon.ListDecimal, []carbon.Attr{carbon.ListDecimal, carbon.ListLowerAlpha, carbon.ListUpperAlpha, carbon.ListLowerRoman, carbon.ListUpperRoman}, func(style carbon.Attr) {
			list.Apply(carbon.With(style)...)
		}),
	)
}
