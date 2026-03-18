package icon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var (
	icons = []carbon.IconName{
		carbon.IconAdd,
		carbon.IconSearch,
		carbon.IconSettings,
		carbon.IconArrowRight,
		carbon.IconFavorite,
		carbon.IconUserAvatar,
		carbon.IconWarningFilled,
		carbon.IconLaunch,
	}
	iconSizes = []carbon.IconSize{
		carbon.IconSize16,
		carbon.IconSize20,
		carbon.IconSize24,
		carbon.IconSize32,
	}
)

// Stories returns the icon-related stories for the Carbon example app.
func Stories() []dom.Element {
	return []dom.Element{
		iconStory(),
		galleryStory(),
	}
}

func iconStory() dom.Element {
	previewIcon := carbon.Icon(carbon.IconAdd, carbon.With(carbon.IconSize32), mvc.WithAttr("style", "color:currentColor"), mvc.WithAriaLabel("add"))
	iconName := carbon.Para("add", mvc.WithAttr("style", "margin:0;opacity:0.72"))
	canvas := mvc.HTML("DIV", mvc.WithClass("canvas"),
		mvc.HTML("DIV", mvc.WithAttr("style", "display:flex;flex-direction:column;align-items:center;gap:0.75rem;color:var(--cds-text-primary,#161616)"),
			previewIcon.Root(),
			iconName,
		),
	)

	return storybook.Story(
		"Icon",
		"Use the controls below to swap the bundled Carbon icon, change its size, and preview it against the Carbon themes.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(canvas, a)
		}),
		storybook.IconDropdown("Icon", carbon.IconAdd, icons, func(name carbon.IconName) {
			previewIcon.SetIcon(name)
			previewIcon.Apply(mvc.WithAriaLabel(string(name)))
			iconName.Root().SetInnerHTML(string(name))
		}),
		storybook.IconSizeDropdown("Size", carbon.IconSize32, iconSizes, func(size carbon.IconSize) {
			previewIcon.Apply(carbon.With(size)...)
		}),
	)
}

func galleryStory() dom.Element {
	type sample struct {
		name  carbon.IconName
		label string
		color string
	}

	samples := []sample{
		{carbon.IconAdd, "Add", "var(--cds-icon-primary,#161616)"},
		{carbon.IconSearch, "Search", "var(--cds-link-primary,#0f62fe)"},
		{carbon.IconSettings, "Settings", "var(--cds-support-info,#0f62fe)"},
		{carbon.IconWarningFilled, "Warning", "var(--cds-support-warning,#f1c21b)"},
		{carbon.IconFavorite, "Favorite", "var(--cds-support-error,#da1e28)"},
		{carbon.IconUserAvatar, "User avatar", "var(--cds-support-success,#198038)"},
	}

	cols := make([]any, 0, len(samples))
	for _, sample := range samples {
		icon := carbon.Icon(sample.name, carbon.With(carbon.IconSize24), mvc.WithAriaLabel(sample.label), mvc.WithAttr("style", "color:"+sample.color))
		cols = append(cols, carbon.Col4(
			carbon.Tile(
				mvc.WithAttr("style", "height:100%"),
				mvc.HTML("DIV", mvc.WithAttr("style", "display:flex;flex-direction:column;gap:0.75rem;align-items:flex-start"),
					icon.Root(),
					mvc.HTML("SPAN", mvc.WithClass("cds--caption-01"), sample.label),
				),
			),
		))
	}

	return storybook.Story(
		"Icon Gallery",
		"A small set of bundled icons styled with Carbon semantic colours, following the same spirit as the Bootstrap icon examples.",
		carbon.Grid(append([]any{mvc.WithAttr("style", "row-gap:1rem")}, cols...)...).Root(),
		nil,
	)
}

func applyTheme(canvas dom.Element, theme carbon.Attr) {
	cl := canvas.ClassList()
	for _, t := range storybook.DefaultThemes {
		cl.Remove(carbon.ClassForTheme(t))
	}
	cl.Add(carbon.ClassForTheme(theme))
	canvas.SetClassName(cl.Value())
}
