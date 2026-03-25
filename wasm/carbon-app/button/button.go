package button

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var (
	buttonKinds = []carbon.Attr{
		carbon.KindPrimary,
		carbon.KindSecondary,
		carbon.KindTertiary,
		carbon.KindGhost,
		carbon.KindDanger,
	}
	buttonSizes = []carbon.Attr{
		carbon.SizeLarge,
		carbon.SizeSmall,
		carbon.SizeMedium,
		carbon.SizeExtraLarge,
		carbon.Size2XLarge,
	}
	buttonIcons = []carbon.IconName{
		carbon.IconAdd,
		carbon.IconSearch,
		carbon.IconSettings,
		carbon.IconArrowRight,
		carbon.IconFavorite,
		carbon.IconUserAvatar,
		carbon.IconWarningFilled,
		carbon.IconLaunch,
	}
	buttonIconLabels = map[carbon.IconName]string{
		carbon.IconAdd:           "Add",
		carbon.IconSearch:        "Search",
		carbon.IconSettings:      "Settings",
		carbon.IconArrowRight:    "Arrow right",
		carbon.IconFavorite:      "Favorite",
		carbon.IconUserAvatar:    "User avatar",
		carbon.IconWarningFilled: "Warning filled",
		carbon.IconLaunch:        "Launch",
	}
)

func View() []any {
	return []any{
		storybook.PageHeader("Button", "Button.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeG10),
			basicButtonStory(),
			iconButtonStory(),
			iconOnlyButtonStory(),
		),
	}
}

func basicButtonStory() dom.Element {
	// Create the button and the canvas
	btn := carbon.Button(carbon.With(carbon.KindPrimary), "Example button")
	btn.SetValue("example-button")
	canvas := carbon.Section(mvc.WithClass("canvas"), btn)

	// Return the story
	return storybook.Story(
		"Basic Button",
		"Use the controls below to change the theme, kind and size of the button.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeG10, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Kind", carbon.KindDanger, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.CheckboxGroup("Enabled", "Enabled", true, func(a bool) {
			btn.SetEnabled(a)
		}),
	)
}

func iconButtonStory() dom.Element {
	// Create the icon, button and the canvas
	icon := carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize16))
	btn := carbon.Button(carbon.With(carbon.KindPrimary), "Icon button", icon)
	btn.SetValue("icon-button")
	canvas := carbon.Section(mvc.WithClass("canvas"), btn)

	return storybook.Story(
		"Button With Icon",
		"Carbon buttons accept an icon in the dedicated icon slot. This story keeps the button interactive while letting you swap the icon, theme, kind, and size.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeG10, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Kind", carbon.KindPrimary, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeExtraLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.IconDropdown("Icon", carbon.IconLaunch, buttonIcons, func(name carbon.IconName) {
			icon.SetValue(string(name))
		}),
	)
}

func iconOnlyButtonStory() dom.Element {
	// Create the icon, button and the canvas
	icon := carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize16))
	btn := carbon.Button(carbon.With(carbon.KindPrimary), icon)
	btn.SetValue("icon-only-button")
	btn.SetLabel(buttonIconLabels[carbon.IconLaunch])
	canvas := carbon.Section(mvc.WithClass("canvas"), btn)

	return storybook.Story(
		"Icon Only Button",
		"Icon-only Carbon buttons still need an accessible name. This story keeps the button label, tooltip text, and icon selection aligned while you change theme, kind, and size.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Kind", carbon.KindSecondary, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.IconDropdown("Icon", carbon.IconSettings, buttonIcons, func(name carbon.IconName) {
			icon.SetValue(string(name))
			btn.SetLabel(buttonIconLabels[name])
		}),
	)
}
