package form

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

// checkable is a View whose checked state can be read back.
type checkable interface {
	mvc.View
	Active() bool
}

func CheckboxView() []any {
	return []any{
		storybook.PageHeader("Checkboxes", "Checkbox.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicCheckboxStory(),
			checkboxGroupStory(),
			checkboxGroupOrientationStory(),
		),
	}
}

func basicCheckboxStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	disabled := false

	chk := carbon.Checkbox("Enable notifications")
	chk.SetActive(true)

	status := carbon.Para(checkboxStatus(chk))

	// EventChange on the checkbox maps to Carbon's internal checkbox-specific
	// custom event, so we can still read Active() directly after each toggle.
	// so we can read Active() directly without ViewFromEventTarget.
	chk.AddEventListener(carbon.EventChange, func(dom.Event) {
		status.Content(checkboxStatus(chk))
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
		chk,
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		chk.SetEnabled(!disabled)
	}
	refresh()

	return storybook.Story(
		"Basic Checkbox",
		"A single checkbox toggles a boolean value and emits EventChange on every state transition. The status line below reads back the checked property after each event.",
		canvas,
		chk,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.CheckboxGroup("State", "Disabled", disabled, func(v bool) {
			disabled = v
			refresh()
		}),
	)
}

func checkboxGroupStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	disableAll := false

	type item struct {
		label   string
		state   carbon.CheckboxState
		enabled bool
	}
	items := []item{
		{"Checked", carbon.CheckboxStateTrue, true},
		{"Unchecked", carbon.CheckboxStateFalse, true},
		{"Indeterminate", carbon.CheckboxStateUndefined, true},
		{"Disabled checked", carbon.CheckboxStateTrue, false},
		{"Disabled unchecked", carbon.CheckboxStateFalse, false},
	}

	boxes := make([]checkable, len(items))
	args := make([]any, len(items))
	for i, it := range items {
		chk := carbon.Checkbox(it.label)
		chk.SetState(it.state)
		chk.SetEnabled(it.enabled)
		boxes[i] = chk
		args[i] = chk
	}

	labels := make([]string, len(items))
	for i, it := range items {
		labels[i] = it.label
	}

	status := carbon.Para(checkboxGroupStatus(labels, boxes))

	group := carbon.CheckboxGroup("Select one or more options.", args...).
		SetLabel("States")

	// One listener on the group catches bubbled EventChange from all
	// child checkboxes — more reliable than wiring N per-child listeners.
	group.AddEventListener(carbon.EventChange, func(dom.Event) {
		status.Content(checkboxGroupStatus(labels, boxes))
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1.5rem"),
		group,
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		if disableAll {
			group.SetEnabled()
		} else {
			// Re-apply per-item enabled state.
			enabled := make([]mvc.View, 0, len(items))
			for i, it := range items {
				if it.enabled {
					enabled = append(enabled, boxes[i])
				}
			}
			group.SetEnabled(enabled...)
		}
	}
	refresh()

	return storybook.Story(
		"Checkbox Group",
		"A checkbox group bundles related checkboxes under a shared legend and optional helper text. Items can be checked, unchecked, or indeterminate, and individual items or the whole group can be disabled.",
		canvas,
		group,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.CheckboxGroup("State", "Disable all", disableAll, func(v bool) {
			disableAll = v
			refresh()
		}),
	)
}

func checkboxGroupOrientationStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentOrientation := carbon.CheckboxOrientationVertical

	orientations := []carbon.Attr{
		carbon.CheckboxOrientationVertical,
		carbon.CheckboxOrientationHorizontal,
	}

	labels := []string{"Design", "Development", "Research", "Product"}
	initial := []bool{true, true, false, false}

	args := make([]any, len(labels))
	for i, label := range labels {
		chk := carbon.Checkbox(label)
		chk.SetActive(initial[i])
		args[i] = chk
	}

	group := carbon.CheckboxGroup("", args...).
		SetLabel("Disciplines").
		SetOrientation(currentOrientation)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem"),
		group,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		group.SetOrientation(currentOrientation)
	}
	refresh()

	return storybook.Story(
		"Orientation",
		"Checkbox groups lay out their children vertically by default. Switch to horizontal when space permits and the labels are short enough to fit on one line.",
		canvas,
		group,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Orientation", currentOrientation, orientations, func(o carbon.Attr) {
			currentOrientation = o
			refresh()
		}),
	)
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

func checkboxStatus(chk checkable) string {
	if chk.Active() {
		return "Checked: notifications enabled."
	}
	return "Unchecked: notifications disabled."
}

func checkboxGroupStatus(labels []string, boxes []checkable) string {
	var active []string
	for i, box := range boxes {
		if box.Active() {
			active = append(active, labels[i])
		}
	}
	if len(active) == 0 {
		return "No permissions selected."
	}
	return "Permissions: " + strings.Join(active, ", ") + "."
}
