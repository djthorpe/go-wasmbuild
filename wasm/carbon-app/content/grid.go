package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func GridView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Grid")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			gridSpansStory(),
			gridVariantsStory(),
			gridLayoutStory(),
		),
	}
}

const (
	gridVariantStandard = carbon.Attr("Standard")
	gridVariantFull     = carbon.Attr("Full width")
	gridVariantNarrow   = carbon.Attr("Narrow")
	gridVariantDense    = carbon.Attr("Condensed")

	gridCellBaseStyle = "display:grid;gap:0.5rem;align-content:start;box-sizing:border-box;height:100%;min-height:7rem;margin-inline:calc(var(--cds-grid-mode-start,0rem)*-1) calc(var(--cds-grid-mode-end,0rem)*-1);padding-block:1rem;padding-inline:calc(1rem + var(--cds-grid-mode-start,0rem)) calc(1rem + var(--cds-grid-mode-end,0rem));border:1px solid var(--cds-border-subtle-01,#c6c6c6);"
	gridCellStyleA    = gridCellBaseStyle + "background:var(--cds-layer-01,#f4f4f4);"
	gridCellStyleB    = gridCellBaseStyle + "background:var(--cds-layer-02,#e0e0e0);"
	gridCanvasStyle   = "display:grid;gap:1rem"
	gridLabelStyle    = "color:var(--cds-text-secondary,#525252);margin:0;max-width:56rem"
	gridVariantStyle  = "display:grid;gap:0.75rem;padding:0;align-items:stretch;justify-content:stretch;min-height:auto"
	gridVariantBody   = "display:grid;gap:1rem;padding:1.25rem 1.5rem 1.5rem"
	gridVariantStage  = "min-width:120rem;padding:1rem 0 1.5rem;background:linear-gradient(180deg, var(--cds-layer-01,#f4f4f4) 0%, var(--cds-layer,#ffffff) 100%)"
	gridVariantFrame  = "overflow-x:auto;border-block:1px solid var(--cds-border-subtle-01,#c6c6c6)"
	gridVariantHint   = "color:var(--cds-text-secondary,#525252);margin:0;padding:0 1.5rem"
	gridWrapperStyle  = "display:grid;gap:1.5rem"
	gridSidebarStyle  = gridCellBaseStyle + "background:var(--cds-layer-accent-01,#e8f1ff);"
	gridContentStyle  = gridCellBaseStyle + "background:var(--cds-layer-01,#f4f4f4);"
	gridUtilityStyle  = gridCellBaseStyle + "background:var(--cds-layer-02,#e0e0e0);min-height:5rem;"
)

func gridSpansStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle(gridCanvasStyle),
		carbon.With(carbon.ThemeWhite),
		carbon.Grid(
			gridCell(4, true, "4 columns", "Quarter-width block for supporting content, filters, or metadata."),
			gridCell(4, false, "4 columns", "Pair several equal spans to build balanced dashboards or card rows."),
			gridCell(4, true, "4 columns", "Use matching spans when rhythm matters more than hierarchy."),
			gridCell(4, false, "4 columns", "The four-column cadence makes the 16-column system obvious."),
			gridCell(8, false, "8 columns", "Half-width regions work well for side-by-side comparisons or two-panel layouts."),
			gridCell(8, true, "8 columns", "Wider spans give dense content enough room to breathe."),
			gridCell(4, true, "4-column rail", "A narrow rail can hold navigation, filters, or status summaries."),
			gridCell(12, false, "12-column body", "The main content area stays dominant while still aligning to the same grid."),
			gridCell(16, true, "16-column section", "Use full-width rows for page-level banners, tables, or long-form content."),
		),
	)

	return storybook.Story(
		"Column Spans",
		"The base grid story should make the 16-column system legible at a glance and show where common span combinations fit.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func gridVariantsStory() dom.Element {
	description := carbon.Compact(mvc.WithStyle(gridLabelStyle))
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle(gridVariantStyle),
		carbon.With(carbon.ThemeWhite),
	)

	render := func(name string) {
		description.Content(gridVariantDescription(name))
		canvas.Content(
			mvc.HTML("DIV",
				mvc.WithStyle(gridVariantBody),
				description,
				carbon.Compact(
					"The preview stage below is intentionally 120rem wide so the container width and gutter changes are visible instead of being masked by the story frame.",
					mvc.WithStyle(gridVariantHint),
				),
				mvc.HTML("DIV",
					mvc.WithStyle(gridVariantFrame),
					mvc.HTML("DIV",
						mvc.WithStyle(gridVariantStage),
						gridVariant(name),
					),
				),
			),
		)
	}

	render("Standard")

	return storybook.Story(
		"Grid Variants",
		"Full-width, narrow, and condensed grids change the page rhythm. This story lets you compare the same content structure across variants.",
		canvas,
		nil,
		storybook.Dropdown("Variant", gridVariantStandard, []carbon.Attr{gridVariantStandard, gridVariantFull, gridVariantNarrow, gridVariantDense}, func(selected carbon.Attr) {
			switch selected {
			case gridVariantFull:
				render("Full width")
			case gridVariantNarrow:
				render("Narrow")
			case gridVariantDense:
				render("Condensed")
			default:
				render("Standard")
			}
		}),
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func gridVariantDescription(name string) string {
	switch name {
	case "Full width":
		return "Full width removes the default maximum page width so the grid can stretch edge to edge. Use it for immersive layouts, wide dashboards, or sections that should align with the full viewport instead of the standard Carbon content frame."
	case "Narrow":
		return "Narrow reduces the starting gutter so columns sit tighter against the layout edge. Use it when the default outer spacing feels too loose and you want content to feel more tightly packed without collapsing the overall column structure."
	case "Condensed":
		return "Condensed reduces grid gutters to nearly zero, producing a denser layout. Use it for highly structured surfaces like data-heavy dashboards or utility panels where preserving horizontal space matters more than generous separation."
	default:
		return "Standard uses Carbon's default content width and gutter rhythm. It is the baseline page grid for most application surfaces and the safest choice when you want familiar spacing and balanced readability."
	}
}

func gridVariant(name string) mvc.View {
	children := []any{
		gridCell(4, true, "Navigation", "Compact supporting content kept aligned with the main layout."),
		gridCell(8, false, "Primary content", "A broader central region demonstrates the gutter width for denser content."),
		gridCell(4, true, "Utilities", "Secondary actions and metadata stay visually separate without leaving the grid."),
	}
	switch name {
	case "Full width":
		return carbon.GridFullWidth(children...)
	case "Narrow":
		return carbon.GridNarrow(children...)
	case "Condensed":
		return carbon.GridCondensed(children...)
	default:
		return carbon.Grid(children...)
	}
}

func gridLayoutStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle(gridWrapperStyle),
		carbon.With(carbon.ThemeWhite),
		carbon.Grid(
			carbon.Col4(
				gridPanel(
					gridSidebarStyle,
					carbon.Head(4, "Sidebar"),
					carbon.Compact("Filters, secondary navigation, and status blocks can live in a narrower rail without fighting the main column."),
				),
			),
			carbon.Col12(
				gridPanel(
					gridContentStyle,
					carbon.Head(4, "Main content"),
					carbon.Para("Primary workflow content benefits from the wider span, especially once text, forms, or tables are involved."),
				),
			),
		),
		carbon.Grid(
			carbon.Col10(
				gridPanel(
					gridContentStyle,
					carbon.Head(4, "Feature area"),
					carbon.Compact("A 10-column content region leaves room for contextual tools without collapsing the reading width."),
				),
			),
			carbon.Col6(
				gridPanel(
					gridUtilityStyle,
					carbon.Head(4, "Context panel"),
					carbon.Compact("Use the remaining span for activity, inspectors, or supporting actions."),
				),
			),
		),
	)

	return storybook.Story(
		"Page Composition",
		"A grid example is more useful when it shows page-level composition patterns instead of just span arithmetic.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func gridCell(span int, accent bool, title, copy string) mvc.View {
	style := gridCellStyleA
	if !accent {
		style = gridCellStyleB
	}
	return carbon.ColSpan(span,
		gridPanel(
			style,
			carbon.Strong(title),
			carbon.Compact(copy),
		),
	)
}

func gridPanel(style string, children ...any) dom.Element {
	args := make([]any, 0, len(children)+1)
	args = append(args, mvc.WithAttr("style", style))
	args = append(args, children...)
	return mvc.HTML("DIV", args...)
}
