package buttons

import (
	dom "github.com/djthorpe/go-wasmbuild"
	"github.com/djthorpe/go-wasmbuild/pkg/carbon"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
	"github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
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

func View() mvc.View {
	return carbon.Section(
		basicButtonStory(),
	)
}

func basicButtonStory() dom.Element {
	btn := carbon.Button(carbon.With(carbon.KindPrimary), "Example button").SetValue("example-button")

	return storybook.Story(
		"Basic Button",
		"Use the controls below to change the theme, kind and size of the button.",
		mvc.HTML("DIV", mvc.WithClass("canvas"), btn),
		btn,
		Dropdown("Theme", carbon.ThemeG10, storybook.DefaultThemes, func(theme carbon.Attr) {
			btn.Parent().Apply(carbon.With(theme)...)
		}),
		Dropdown("Kind", carbon.KindDanger, buttonKinds, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
		Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn.Apply(carbon.With(a)...)
		}),
	)
}

// Dropdown builds a Carbon dropdown for a set of Attr options.
func Dropdown(label string, selected carbon.Attr, options []carbon.Attr, onChange func(carbon.Attr)) dom.Element {
	// TODO: Select the default option
	//onChange(selected)

	// Build the dropdown items
	items := make([]any, 0, len(options)+1)
	items = append(items, carbon.DropdownTitleText(label))
	for _, option := range options {
		item := carbon.DropdownItem(mvc.WithAttr("value", string(option)), string(option))
		if option == selected {
			item.SetSelected(true)
		}
		items = append(items, item)
	}

	dd := carbon.Dropdown(append([]any{
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
	}, items...)...)
	dd.SetValue(string(selected))
	dd.AddEventListener(carbon.EventSelected, func(dom.Event) {
		onChange(carbon.Attr(dd.Value()))
	})
	return dd.Root()
}
