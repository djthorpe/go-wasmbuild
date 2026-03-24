package form

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

type dropdownChoice interface {
	mvc.View
	Value() string
}

var dropdownSizes = []carbon.Attr{
	carbon.SizeSmall,
	carbon.SizeMedium,
	carbon.SizeLarge,
}

var dropdownStates = []carbon.Attr{
	carbon.Attr("Draft"),
	carbon.Attr("Review"),
	carbon.Attr("Published"),
	carbon.Attr("Archived"),
}

var dropdownAssignees = []carbon.Attr{
	carbon.Attr("Ada Lovelace"),
	carbon.Attr("Grace Hopper"),
	carbon.Attr("Margaret Hamilton"),
	carbon.Attr("Katherine Johnson"),
}

func DropdownView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Dropdowns")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicDropdownStory(),
			helperTextDropdownStory(),
			widthDropdownStory(),
		),
	}
}

func basicDropdownStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := carbon.SizeMedium
	currentValue := dropdownStates[1]
	enabled := true

	items, args := makeDropdownItems(dropdownStates)
	dd := carbon.Dropdown("", args...).
		SetLabel("Status").
		SetValue(string(currentValue))
	setDropdownSelection(dd, items, currentValue)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		dropdownStage("20rem", dd),
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		dd.Apply(carbon.With(currentSize)...)
		dd.SetEnabled(enabled)
		setDropdownSelection(dd, items, currentValue)
	}
	refresh()

	return storybook.Story(
		"Basic Dropdown",
		"Dropdowns expose a selected value and a selected event. This story keeps the component simple: one label, one selected value, one enabled state, and theme/size controls around it.",
		canvas,
		dd,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, dropdownSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
		storybook.Dropdown("Selected", currentValue, dropdownStates, func(value carbon.Attr) {
			currentValue = value
			refresh()
		}),
		storybook.CheckboxGroup("Enabled", "Enabled", enabled, func(value bool) {
			enabled = value
			refresh()
		}),
	)
}

func helperTextDropdownStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentValue := dropdownAssignees[0]
	items, args := makeDropdownItems(dropdownAssignees)
	dd := carbon.Dropdown("The helper text gives the field more context without adding extra layout chrome.", args...).
		SetLabel("Assignee").
		SetValue(string(currentValue))
	setDropdownSelection(dd, items, currentValue)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		dropdownStage("24rem", dd),
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		setDropdownSelection(dd, items, currentValue)
	}
	refresh()

	return storybook.Story(
		"Helper Text",
		"The wrapper accepts helper text at construction time and title text through `SetLabel`. This is the common form-field shape for explaining what the dropdown controls.",
		canvas,
		dd,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Selected", currentValue, dropdownAssignees, func(value carbon.Attr) {
			currentValue = value
			refresh()
		}),
	)
}

func widthDropdownStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentValue := dropdownStates[2]
	fullWidth := true
	items, args := makeDropdownItems(dropdownStates)
	dd := carbon.Dropdown("", args...).
		SetLabel("Publishing state").
		SetValue(string(currentValue))
	setDropdownSelection(dd, items, currentValue)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		dropdownStage("32rem", dd),
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		if fullWidth {
			dd.Apply(mvc.WithStyle("width:100%"))
		} else {
			dd.Apply(mvc.WithoutAttr("style"))
		}
		setDropdownSelection(dd, items, currentValue)
	}
	refresh()

	return storybook.Story(
		"Field Width",
		"Dropdowns often need to expand to the available form width. This story shows the same control in intrinsic width and full-width modes without changing any component semantics.",
		canvas,
		dd,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Selected", currentValue, dropdownStates, func(value carbon.Attr) {
			currentValue = value
			refresh()
		}),
		storybook.CheckboxGroup("Layout", "Full width", fullWidth, func(value bool) {
			fullWidth = value
			refresh()
		}),
	)
}

func makeDropdownItems(values []carbon.Attr) ([]dropdownChoice, []any) {
	items := make([]dropdownChoice, 0, len(values))
	args := make([]any, 0, len(values))
	for _, value := range values {
		item := dropdownChoice(carbon.DropdownItem(string(value)).SetValue(string(value)))
		items = append(items, item)
		args = append(args, item)
	}
	return items, args
}

func setDropdownSelection(dd mvc.ActiveGroup, items []dropdownChoice, selected carbon.Attr) {
	active := make([]mvc.View, 0, 1)
	for _, item := range items {
		if item.Value() == string(selected) {
			active = append(active, item)
			break
		}
	}
	dd.SetActive(active...)
}

func dropdownStage(maxWidth string, child mvc.View) dom.Element {
	style := "width:100%"
	if maxWidth != "" {
		style += ";max-width:" + maxWidth
	}
	return mvc.HTML("DIV", mvc.WithStyle(style), child)
}
