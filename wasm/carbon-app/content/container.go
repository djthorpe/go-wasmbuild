package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func SectionView() []any {
	return []any{
		storybook.PageHeader("Section", "Section.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			sectionStory(),
		),
	}
}

func PageView() []any {
	return []any{
		storybook.PageHeader("Page", "Page.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			pageStory(),
		),
	}
}

func sectionStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		carbon.Head(2, "Overview"),
		carbon.Para("Section adds Carbon's standard cds--content wrapper so content aligns with the page rhythm used across the design system."),
		carbon.Grid(
			carbon.Col4(sectionPanel("Summary", "Compact supporting content inside a content section.")),
			carbon.Col12(sectionPanel("Primary content", "Wider content still sits inside the same cds--content section wrapper.")),
		),
	)

	return storybook.Story(
		"Section",
		"Section is the Carbon content container used for standard page spacing and content alignment.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func pageStory() dom.Element {
	page := carbon.Page(
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem;border:1px dashed var(--cds-border-subtle,#c6c6c6)"),
		carbon.Head(2, "Page wrapper"),
		carbon.Compact("Page does not apply cds--content. It stays neutral so route-level layouts can define their own spacing."),
		carbon.Grid(
			carbon.Col6(sectionPanel("Navigation rail", "A page can place its own layout primitives without inheriting Section spacing.")),
			carbon.Col10(sectionPanel("Workspace", "This area demonstrates custom page composition inside the plain wrapper.")),
		),
	)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		page,
	)

	return storybook.Story(
		"Page",
		"Page is a neutral wrapper used when the route or story defines its own layout instead of relying on cds--content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func sectionPanel(title, copy string) dom.Element {
	return mvc.HTML("DIV",
		mvc.WithStyle("display:grid;gap:0.5rem;height:100%;padding:1rem;border:1px solid var(--cds-border-subtle,#c6c6c6);background:var(--cds-layer,#ffffff)"),
		carbon.Strong(title),
		carbon.Compact(copy),
	)
}
