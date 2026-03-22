package form

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func IntroductionView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Introduction")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			introductionStory(),
		),
	}
}

func introductionStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	required := true
	initialValue := ""

	input := carbon.Input().
		SetLabel("Input").
		SetValue(initialValue)
	input.Root().SetAttribute("placeholder", "Quarterly reporting workspace")
	checkbox := carbon.Checkbox("Checkbox")
	checkbox.SetValue("checked")

	group := carbon.FormGroup(input, checkbox).SetLabel("Form Group")
	validate := carbon.Button(carbon.With(carbon.KindSecondary), "Check validity")
	clear := carbon.Button(carbon.With(carbon.KindGhost), "Clear")

	status := carbon.Para(formIntroductionStatus("Ready", input.Value(), checkbox.Active(), required))
	formView := carbon.Form(
		mvc.WithStyle("display:grid;gap:1rem;max-width:28rem"),
		group,
		carbon.ButtonGroup(validate, clear),
	)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		formStage("32rem", formView),
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		input.SetRequired(required)
		if !required || strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
	}
	refresh()

	updateStatus := func(source string) {
		status.Content(formIntroductionStatus(source, input.Value(), checkbox.Active(), required))
	}
	formView.AddEventListener(carbon.EventInput, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInput))
	})
	formView.AddEventListener(carbon.EventChange, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventChange))
	})
	formView.AddEventListener(carbon.EventFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventFocus))
	})
	formView.AddEventListener(carbon.EventNoFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventNoFocus))
	})
	formView.AddEventListener(carbon.EventInvalid, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInvalid))
	})

	validate.AddEventListener(carbon.EventClick, func(dom.Event) {
		if input.CheckValidity() {
			updateStatus("checkValidity")
		}
	})
	clear.AddEventListener(carbon.EventClick, func(dom.Event) {
		input.SetValue("")
		input.SetCustomValidity("")
		checkbox.SetActive(false)
		updateStatus("cleared")
	})

	return storybook.Story(
		"Introduction",
		"This story composes the new form wrappers into a minimal fieldset: a form, one form group, one text input, one checkbox, and two action buttons. The form itself now collects focus, input, change, and invalid signals so the event panel reflects the whole group rather than a single field.",
		canvas,
		formView,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
			updateStatus("theme")
		}),
		storybook.CheckboxGroup("Validation", "Required", required, func(value bool) {
			required = value
			refresh()
			updateStatus("validation")
		}),
	)
}

func formIntroductionStatus(source, value string, checked, required bool) string {
	value = strings.TrimSpace(value)
	if value == "" {
		value = "(empty)"
	}
	state := "unchecked"
	if checked {
		state = "checked"
	}
	validation := "optional"
	if required {
		validation = "required"
	}
	return "Last signal: " + source + ". Validation: " + validation + ". Input value: " + value + ". Checkbox: " + state + "."
}

func formStage(maxWidth string, child mvc.View) dom.Element {
	style := "width:100%"
	if maxWidth != "" {
		style += ";max-width:" + maxWidth
	}
	return mvc.HTML("DIV", mvc.WithStyle(style), child)
}
