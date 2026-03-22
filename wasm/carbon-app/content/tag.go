package content

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var tagTypes = []carbon.Attr{
	carbon.TagRed,
	carbon.TagMagenta,
	carbon.TagPurple,
	carbon.TagBlue,
	carbon.TagCyan,
	carbon.TagTeal,
	carbon.TagGreen,
	carbon.TagGray,
	carbon.TagCoolGray,
	carbon.TagWarmGray,
	carbon.TagHighContrast,
	carbon.TagOutline,
}

var tagSizes = []carbon.Attr{carbon.SizeSmall, carbon.SizeMedium, carbon.SizeLarge}

func TagView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Tags")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicTagsStory(),
			iconTagsStory(),
			dismissibleTagsStory(),
			operationalTagsStory(),
			tagGroupStory(),
		),
	}
}

func basicTagsStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := carbon.SizeMedium

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		canvas.Content(tagMatrix(currentSize))
	}
	refresh()

	return storybook.Story(
		"Tag Types",
		"Tags are short status or metadata labels. Carbon's tag web components support a broad type scale and a compact size range, so the story focuses on the type system first and lets size/theme vary through controls.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, tagSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
	)
}

func iconTagsStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentType := carbon.TagBlue
	currentSize := carbon.SizeMedium

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		canvas.Content(
			tagRow("Decorative icons inherit the tag colour and are inserted into Carbon's icon slot.", currentSize,
				carbon.Tag(
					carbon.Icon(carbon.IconAdd),
					"New",
					carbon.With(currentType, currentSize),
				),
				carbon.DismissibleTag(
					"Beta",
					carbon.Icon(carbon.IconLaunch),
					carbon.With(currentType, currentSize),
				),
				carbon.OperationalTag(
					"Needs review",
					carbon.Icon(carbon.IconWarningFilled),
					carbon.With(currentType, currentSize),
				),
			),
		)
	}
	refresh()

	return storybook.Story(
		"Tag Icons",
		"Tags can render a leading Carbon icon. This story demonstrates why the wrapper normalizes passed icons into the `icon` slot and forces them to inherit the tag's current color.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Type", currentType, tagTypes, func(value carbon.Attr) {
			currentType = value
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, tagSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
	)
}

func tagGroupStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentType := carbon.TagTeal
	currentSize := carbon.SizeMedium

	group := carbon.TagGroup(
		carbon.Tag(carbon.Icon(carbon.IconAdd), "Draft", carbon.With(currentType, currentSize)),
		carbon.DismissibleTag("Beta", carbon.Icon(carbon.IconLaunch), carbon.With(currentType, currentSize)),
		carbon.OperationalTag("Review", carbon.Icon(carbon.IconWarningFilled), carbon.With(currentType, currentSize)),
	)
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
		group,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		group.Content(
			carbon.Tag(carbon.Icon(carbon.IconAdd), "Draft", carbon.With(currentType, currentSize)),
			carbon.DismissibleTag("Beta", carbon.Icon(carbon.IconLaunch), carbon.With(currentType, currentSize)),
			carbon.OperationalTag("Review", carbon.Icon(carbon.IconWarningFilled), carbon.With(currentType, currentSize)),
		)
	}
	refresh()

	return storybook.Story(
		"Tag Group",
		"TagGroup is a typed container for multiple tags. Because the child tag custom events bubble, the group can be observed once at the container level instead of wiring handlers on each individual tag.",
		canvas,
		group,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Type", currentType, tagTypes, func(value carbon.Attr) {
			currentType = value
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, tagSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
	)
}

func dismissibleTagsStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentType := carbon.TagGray
	currentSize := carbon.SizeMedium
	disabled := false

	tag := carbon.DismissibleTag("Release 1.8", carbon.With(currentType, currentSize), mvc.WithAttr("dismiss-tooltip-label", "Remove release tag"))
	reset := carbon.Button("Reset")
	reset.AddEventListener(carbon.EventClick, func(dom.Event) {
		tag.SetVisible(true)
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:flex;flex-wrap:wrap;align-items:center;gap:1rem"),
		tag,
		reset,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		tag.Apply(carbon.With(currentType, currentSize)...)
		tag.SetEnabled(!disabled)
		tag.SetVisible(true)
	}
	refresh()

	return storybook.Story(
		"Dismissible Tag",
		"Dismissible tags manage their own close affordance and emit close lifecycle events. The wrapper exposes the dedicated `cds-dismissible-tag` API directly, with visibility controlled through the shared MVC visible-state pattern.",
		canvas,
		tag,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Type", currentType, tagTypes, func(value carbon.Attr) {
			currentType = value
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, tagSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
		storybook.CheckboxGroup("State", "Disabled", disabled, func(value bool) {
			disabled = value
			refresh()
		}),
	)
}

func operationalTagsStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentType := carbon.TagBlue
	currentSize := carbon.SizeMedium
	disabled := false

	tag := carbon.OperationalTag("Sync ready", carbon.With(currentType, currentSize))
	reset := carbon.Button("Clear selection")
	reset.AddEventListener(carbon.EventClick, func(dom.Event) {
		tag.SetActive(false)
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:flex;flex-wrap:wrap;align-items:center;gap:1rem"),
		tag,
		reset,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		tag.Apply(carbon.With(currentType, currentSize)...)
		tag.SetEnabled(!disabled)
	}
	refresh()

	return storybook.Story(
		"Operational Tag",
		"Operational tags expose a selectable active state and emit a selected event when activated. In this wrapper they work best as compact one-step action/status chips, with the reset control clearing the active state explicitly.",
		canvas,
		tag,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Type", currentType, tagTypes, func(value carbon.Attr) {
			currentType = value
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, tagSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
		storybook.CheckboxGroup("State", "Disabled", disabled, func(value bool) {
			disabled = value
			refresh()
		}),
	)
}

func tagMatrix(size carbon.Attr) dom.Element {
	rows := make([]any, 0, 4)
	rows = append(rows,
		tagRow("Chromatic scale", size,
			carbon.Tag(tagLabel(carbon.TagRed), carbon.With(carbon.TagRed, size)),
			carbon.Tag(tagLabel(carbon.TagMagenta), carbon.With(carbon.TagMagenta, size)),
			carbon.Tag(tagLabel(carbon.TagPurple), carbon.With(carbon.TagPurple, size)),
			carbon.Tag(tagLabel(carbon.TagBlue), carbon.With(carbon.TagBlue, size)),
			carbon.Tag(tagLabel(carbon.TagCyan), carbon.With(carbon.TagCyan, size)),
			carbon.Tag(tagLabel(carbon.TagTeal), carbon.With(carbon.TagTeal, size)),
			carbon.Tag(tagLabel(carbon.TagGreen), carbon.With(carbon.TagGreen, size)),
		),
		tagRow("Neutral scale", size,
			carbon.Tag(tagLabel(carbon.TagGray), carbon.With(carbon.TagGray, size)),
			carbon.Tag(tagLabel(carbon.TagCoolGray), carbon.With(carbon.TagCoolGray, size)),
			carbon.Tag(tagLabel(carbon.TagWarmGray), carbon.With(carbon.TagWarmGray, size)),
			carbon.Tag(tagLabel(carbon.TagHighContrast), carbon.With(carbon.TagHighContrast, size)),
			carbon.Tag(tagLabel(carbon.TagOutline), carbon.With(carbon.TagOutline, size)),
		),
		mvc.HTML("DIV",
			mvc.WithStyle("display:grid;gap:0.35rem;max-width:44rem"),
			carbon.Compact("Use the colorful tags when the tag itself carries state meaning, and reserve the neutral tags for metadata, taxonomy, or low-emphasis context."),
		),
	)
	return mvc.HTML("DIV", append([]any{mvc.WithStyle("display:grid;gap:1rem")}, rows...)...)
}

func tagRow(title string, size carbon.Attr, tags ...mvc.View) dom.Element {
	children := make([]any, 0, len(tags))
	for _, tag := range tags {
		children = append(children, tag)
	}
	return mvc.HTML("DIV",
		mvc.WithStyle("display:grid;gap:0.5rem"),
		carbon.Compact(title),
		mvc.HTML("DIV", append([]any{mvc.WithStyle("display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center")}, children...)...),
	)
}

func tagLabel(value carbon.Attr) string {
	parts := strings.Split(string(value), "-")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}
