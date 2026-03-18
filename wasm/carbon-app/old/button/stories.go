package button

import (
	"strings"

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

// Stories returns the button-related stories for the Carbon example app.
func Stories() []dom.Element {
	return []dom.Element{
		basicButtonStory(),
		iconButtonStory(),
		iconOnlyButtonStory(),
	}
}

func basicButtonStory() dom.Element {
	btn := carbon.Button(carbon.With(carbon.KindPrimary), mvc.WithAttr("value", "example-button"), "Example button")
	canvas := mvc.HTML("DIV", mvc.WithClass("canvas"), btn)

	return storybook.Story(
		"Button",
		"Use the controls below to change the theme, kind and size of the button.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(canvas, a)
		}),
		storybook.Dropdown("Kind", carbon.KindPrimary, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.CheckboxGroup("Options", "Enabled", true, func(checked bool) {
			btn.SetEnabled(checked)
		}),
	)
}

func iconButtonStory() dom.Element {
	icon := carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize16))
	btn := carbon.Button(carbon.With(carbon.KindPrimary), mvc.WithAttr("value", "icon-button"), "Launch", icon)
	canvas := mvc.HTML("DIV", mvc.WithClass("canvas"), btn)

	return storybook.Story(
		"Button With Icon",
		"Carbon buttons accept an icon in the dedicated icon slot. This story keeps the button interactive while letting you swap the icon, theme, kind, and size.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(canvas, a)
		}),
		storybook.Dropdown("Kind", carbon.KindPrimary, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.IconDropdown("Icon", carbon.IconLaunch, buttonIcons, func(name carbon.IconName) {
			icon.SetIcon(name)
		}),
	)
}

func iconOnlyButtonStory() dom.Element {
	label := buttonIconLabel(carbon.IconSettings)
	icon := carbon.Icon(carbon.IconSettings, carbon.With(carbon.IconSize16))
	btn := carbon.Button(
		mvc.WithAttr("value", "icon-only-button"),
		mvc.WithAriaLabel(label),
		mvc.WithAttr("tooltip-text", label),
		mvc.WithAttr("title", label),
		icon,
	)
	canvas := mvc.HTML("DIV", mvc.WithClass("canvas"), btn)

	return storybook.Story(
		"Icon Only Button",
		"Icon-only Carbon buttons still need an accessible name. This story keeps the button label, tooltip text, and icon selection aligned while you change theme, kind, and size.",
		canvas,
		btn,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(canvas, a)
		}),
		storybook.Dropdown("Kind", carbon.KindGhost, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		storybook.IconDropdown("Icon", carbon.IconSettings, buttonIcons, func(name carbon.IconName) {
			label := buttonIconLabel(name)
			icon.SetIcon(name)
			btn.Apply(mvc.WithAriaLabel(label), mvc.WithAttr("tooltip-text", label), mvc.WithAttr("title", label))
		}),
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

func buttonIconLabel(name carbon.IconName) string {
	if label, ok := buttonIconLabels[name]; ok {
		return label
	}

	replacer := strings.NewReplacer("--", " ", "-", " ")
	return strings.TrimSpace(replacer.Replace(string(name)))
}
