package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var tileBackgrounds = []carbon.Attr{
	carbon.Attr("Default"),
	carbon.Attr("Layer 02"),
	carbon.Attr("Accent"),
	carbon.Attr("Info Tint"),
	carbon.Attr("Success Tint"),
	carbon.Attr("Warning Tint"),
}

func TileView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Tiles")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicTilesStory(),
			decoratorTilesStory(),
			layoutTilesStory(),
		),
	}
}

func basicTilesStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentBackground := tileBackgrounds[0]

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
	)

	render := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		canvas.Content(
			carbon.Grid(
				carbon.Col8(tileCard("Default tile", currentBackground,
					carbon.Compact("Tiles are quiet content containers. Use them when you need separation and structure without introducing new interaction semantics."),
				)),
				carbon.Col8(tileCard("Supporting details", currentBackground,
					carbon.Compact("The wrapper keeps tile presentation opt-driven, so fill, height, and background can be configured directly at construction time."),
				)),
			),
		)
	}
	render()

	return storybook.Story(
		"Basic Tiles",
		"Tiles are structural surfaces for grouping related content. They are intentionally simple: no selection model, no component-specific events, just a Carbon container with a few presentational opts. The background control includes both neutral surfaces and tinted variants so the effect is obvious.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Background", currentBackground, tileBackgrounds, func(background carbon.Attr) {
			currentBackground = background
			render()
		}),
	)
}

func decoratorTilesStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentBackground := tileBackgrounds[1]

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
	)

	render := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		canvas.Content(
			carbon.Grid(
				carbon.Col8(tileCard("Plain tile", currentBackground,
					carbon.Compact("This is the same base tile without any slotted decorator content. The title and body are identical so the top-right treatment is the only meaningful difference."),
				)),
				carbon.Col8(tileDecoratedCard("Decorated tile", currentBackground,
					tileDecoratorBadge("AI"),
					carbon.Compact("Carbon tiles use a decorator slot for small top-right adornments such as AI, status, or feature badges. This is a slot, not a separate visual tile variant."),
				)),
			),
		)
	}
	render()

	return storybook.Story(
		"Decorator Slot",
		"The tile decorator is visible only when you provide slotted decorator content. This story compares a plain tile with the same tile carrying a top-right decorator badge so the difference is explicit.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Background", currentBackground, tileBackgrounds, func(background carbon.Attr) {
			currentBackground = background
			render()
		}),
	)
}

func layoutTilesStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentBackground := tileBackgrounds[2]

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
	)

	render := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		canvas.Content(
			carbon.Grid(
				carbon.Col4(tileMetricCard("Errors", "12", currentBackground)),
				carbon.Col4(tileMetricCard("Warnings", "34", currentBackground)),
				carbon.Col4(tileMetricCard("Skipped", "5", currentBackground)),
				carbon.Col4(tileMetricCard("Queued", "19", currentBackground)),
			),
		)
	}
	render()

	return storybook.Story(
		"Layout Surface",
		"Fill and height are presentational opts for building consistent tile rows. This pattern is useful for dashboards, summary strips, and story canvases where several tiles should align visually, and where tinted surfaces can help separate categories without changing the component model.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Background", currentBackground, tileBackgrounds, func(background carbon.Attr) {
			currentBackground = background
			render()
		}),
	)
}

func tileCard(title string, background carbon.Attr, body ...any) mvc.View {
	children := make([]any, 0, len(body)+1)
	children = append(children, carbon.Head(4, title))
	children = append(children, body...)
	return carbon.Tile(
		carbon.WithFill(),
		carbon.WithBackground(tileBackgroundValue(background)),
		mvc.WithStyle("padding:1rem"),
		children,
	)
}

func tileDecoratedCard(title string, background carbon.Attr, decorator any, body ...any) mvc.View {
	children := make([]any, 0, len(body)+2)
	children = append(children, decorator)
	children = append(children, carbon.Head(4, title))
	children = append(children, body...)
	return carbon.Tile(
		carbon.WithFill(),
		carbon.WithBackground(tileBackgroundValue(background)),
		mvc.WithStyle("padding:1rem"),
		children,
	)
}

func tileDecoratorBadge(label string) dom.Element {
	return mvc.HTML("SPAN",
		mvc.WithAttr("slot", "decorator"),
		mvc.WithStyle("display:inline-grid;place-items:center;min-width:1.5rem;height:1.5rem;padding:0 0.375rem;border-radius:999px;background:var(--cds-support-info,#0043ce);color:#ffffff;font-size:0.75rem;font-weight:600;line-height:1"),
		label,
	)
}

func tileMetricCard(label, value string, background carbon.Attr) mvc.View {
	return carbon.Tile(
		carbon.WithFill(),
		carbon.WithHeight("9rem"),
		carbon.WithBackground(tileBackgroundValue(background)),
		mvc.WithStyle("padding:1rem;display:grid;align-content:space-between;gap:0.5rem"),
		carbon.Compact(label),
		carbon.Head(3, value),
	)
}

func tileBackgroundValue(background carbon.Attr) string {
	switch background {
	case carbon.Attr("Layer 02"):
		return "var(--cds-layer-02,#e0e0e0)"
	case carbon.Attr("Accent"):
		return "var(--cds-layer-accent-01,#e0e0e0)"
	case carbon.Attr("Info Tint"):
		return "color-mix(in srgb, var(--cds-layer-01,#f4f4f4) 82%, var(--cds-support-info,#4589ff) 18%)"
	case carbon.Attr("Success Tint"):
		return "color-mix(in srgb, var(--cds-layer-01,#f4f4f4) 82%, var(--cds-support-success,#42be65) 18%)"
	case carbon.Attr("Warning Tint"):
		return "color-mix(in srgb, var(--cds-layer-01,#f4f4f4) 80%, var(--cds-support-warning,#f1c21b) 20%)"
	default:
		return "var(--cds-layer-01,#f4f4f4)"
	}
}
