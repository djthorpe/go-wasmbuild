package form

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var inputSizes = []carbon.Attr{
	carbon.SizeSmall,
	carbon.SizeMedium,
	carbon.SizeLarge,
}

var inputPresets = []carbon.Attr{
	carbon.Attr(""),
	carbon.Attr("Quarterly report"),
	carbon.Attr("Platform migration"),
	carbon.Attr("User research"),
}

type inputValue interface {
	Value() string
}

func InputView() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Inputs")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicInputStory(),
			validationInputStory(),
			secureInputStory(),
			numberInputStory(),
		),
	}
}

func basicInputStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := carbon.SizeMedium
	currentValue := inputPresets[1]
	required := false

	input := carbon.Input().
		SetLabel("Project name").
		SetHelperText("When required, leaving the field empty triggers validation on blur.").
		SetInvalidText("Enter a project name before leaving the field.").
		SetValue(string(currentValue))
	input.Root().SetAttribute("placeholder", "Type a project name")

	status := carbon.Para(inputStatus("Ready", input, required))
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		inputStage("24rem", input),
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		input.Apply(carbon.With(currentSize)...)
		input.SetRequired(required)
		input.SetValue(string(currentValue))
		if !required || strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
	}
	refresh()

	updateStatus := func(source string) {
		status.Content(inputStatus(source, input, required))
	}
	validateOnBlur := func() {
		if required && strings.TrimSpace(input.Value()) == "" {
			input.SetCustomValidity(input.InvalidText())
			input.CheckValidity()
			return
		}
		input.SetCustomValidity("")
		input.CheckValidity()
	}
	input.AddEventListener(carbon.EventFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventFocus))
	})
	input.AddEventListener(carbon.EventNoFocus, func(dom.Event) {
		validateOnBlur()
		updateStatus(carbon.GoName(carbon.EventNoFocus))
	})
	input.AddEventListener(carbon.EventInput, func(dom.Event) {
		if strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
		updateStatus(carbon.GoName(carbon.EventInput))
	})
	input.AddEventListener(carbon.EventChange, func(dom.Event) {
		currentValue = carbon.Attr(input.Value())
		if strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
		updateStatus(carbon.GoName(carbon.EventChange))
	})
	input.AddEventListener(carbon.EventInvalid, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInvalid))
	})

	return storybook.Story(
		"Basic Input",
		"Text inputs expose direct focus, blur, input, change, and invalid signals. This story also validates on blur when the field is required and empty, while still letting you vary theme, size, required state, and a preset value.",
		canvas,
		input,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
			updateStatus("theme")
		}),
		storybook.Dropdown("Size", currentSize, inputSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
			updateStatus("size")
		}),
		storybook.Dropdown("Preset", currentValue, inputPresets, func(value carbon.Attr) {
			currentValue = value
			refresh()
			updateStatus("preset")
		}),
		storybook.CheckboxGroup("Validation", "Required", required, func(value bool) {
			required = value
			refresh()
			updateStatus("validation")
		}),
	)
}

func validationInputStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	required := true

	input := carbon.Input().
		SetLabel("Validation target").
		SetHelperText("Click Check validity to force the invalid state.").
		SetInvalidText("Validation target is required.")
	input.Root().SetAttribute("placeholder", "Leave blank, then validate")
	validate := carbon.Button(carbon.With(carbon.KindSecondary), "Check validity")
	clear := carbon.Button(carbon.With(carbon.KindGhost), "Clear")

	status := carbon.Para(inputStatus("Ready", input, required))
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		inputStage("24rem", carbon.Form(
			mvc.WithStyle("display:grid;gap:1rem;max-width:24rem"),
			carbon.FormGroup(input).SetLabel("Validation"),
			carbon.ButtonGroup(validate, clear),
		)),
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
		status.Content(inputStatus(source, input, required))
	}
	input.AddEventListener(carbon.EventInvalid, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInvalid))
	})
	input.AddEventListener(carbon.EventChange, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventChange))
	})

	validate.AddEventListener(carbon.EventClick, func(dom.Event) {
		if input.CheckValidity() {
			updateStatus("checkValidity")
		}
	})
	clear.AddEventListener(carbon.EventClick, func(dom.Event) {
		input.SetValue("")
		input.SetCustomValidity("")
		updateStatus("cleared")
	})

	return storybook.Story(
		"Validation",
		"This story isolates validation behavior. Keep the field required and blank, then click Check validity to surface EventInvalid and the invalid UI. Clear resets both the field value and the invalid state.",
		canvas,
		input,
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

func secureInputStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := carbon.SizeMedium
	currentValue := carbon.Attr("")
	required := true

	input := carbon.SecureInput().
		SetLabel("Password").
		SetHelperText("Leaving the field empty triggers validation as soon as focus leaves the control.").
		SetInvalidText("Enter a password before leaving the field.").
		SetValue(string(currentValue))
	input.Root().SetAttribute("placeholder", "Enter a secure password")

	status := carbon.Para(inputStatus("Ready", input, required))
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		inputStage("24rem", input),
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		input.Apply(carbon.With(currentSize)...)
		input.SetRequired(required)
		input.SetValue(string(currentValue))
		if !required || strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
	}
	refresh()

	updateStatus := func(source string) {
		status.Content(inputStatus(source, input, required))
	}
	validateOnBlur := func() {
		if required && strings.TrimSpace(input.Value()) == "" {
			input.SetCustomValidity(input.InvalidText())
			input.CheckValidity()
			return
		}
		input.SetCustomValidity("")
		input.CheckValidity()
	}
	input.AddEventListener(carbon.EventFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventFocus))
	})
	input.AddEventListener(carbon.EventNoFocus, func(dom.Event) {
		validateOnBlur()
		updateStatus(carbon.GoName(carbon.EventNoFocus))
	})
	input.AddEventListener(carbon.EventInput, func(dom.Event) {
		if strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
		updateStatus(carbon.GoName(carbon.EventInput))
	})
	input.AddEventListener(carbon.EventChange, func(dom.Event) {
		currentValue = carbon.Attr(input.Value())
		if strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
		updateStatus(carbon.GoName(carbon.EventChange))
	})
	input.AddEventListener(carbon.EventInvalid, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInvalid))
	})

	return storybook.Story(
		"Secure Input",
		"Secure inputs wrap Carbon's password field while keeping the same input-style API. This story starts empty so the required validation is visible: tab into the field and then leave it empty to trigger invalid state as soon as focus leaves the control.",
		canvas,
		input,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
			updateStatus("theme")
		}),
		storybook.Dropdown("Size", currentSize, inputSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
			updateStatus("size")
		}),
		storybook.CheckboxGroup("Validation", "Required", required, func(value bool) {
			required = value
			refresh()
			updateStatus("validation")
		}),
	)
}

func numberInputStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := carbon.SizeMedium
	required := false
	allowEmpty := false
	hideSteppers := false

	input := carbon.NumberInput().
		SetLabel("Seats").
		SetHelperText("Choose between 1 and 12 seats.").
		SetInvalidText("Seat count must stay between 1 and 12.").
		SetMin("1").
		SetMax("12").
		SetStep("1").
		SetValue("3")
	input.Root().SetAttribute("placeholder", "Enter a seat count")

	status := carbon.Para(numberInputStatus("Ready", input, required))
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%"),
		inputStage("24rem", input),
		status,
	)

	refresh := func() {
		canvas.Apply(carbon.With(currentTheme)...)
		input.Apply(carbon.With(currentSize)...)
		input.SetRequired(required)
		input.SetAllowEmpty(allowEmpty)
		input.SetHideSteppers(hideSteppers)
		if !required || allowEmpty || strings.TrimSpace(input.Value()) != "" {
			input.SetCustomValidity("")
		}
	}
	refresh()

	updateStatus := func(source string) {
		status.Content(numberInputStatus(source, input, required))
	}
	input.AddEventListener(carbon.EventFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventFocus))
	})
	input.AddEventListener(carbon.EventNoFocus, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventNoFocus))
	})
	input.AddEventListener(carbon.EventInput, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInput))
	})
	input.AddEventListener(carbon.EventChange, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventChange))
	})
	input.AddEventListener(carbon.EventInvalid, func(dom.Event) {
		updateStatus(carbon.GoName(carbon.EventInvalid))
	})

	return storybook.Story(
		"Number Input",
		"Number inputs normalize Carbon's custom cds-number-input signal onto the same EventInput and EventChange vocabulary used elsewhere. This story exposes min, max, stepper visibility, and empty-value handling while the status line reads back the current numeric value and constraints.",
		canvas,
		input,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
			updateStatus("theme")
		}),
		storybook.Dropdown("Size", currentSize, inputSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
			updateStatus("size")
		}),
		storybook.CheckboxGroup("Validation", "Required", required, func(value bool) {
			required = value
			refresh()
			updateStatus("validation")
		}),
		storybook.CheckboxGroup("Behavior", "Allow empty", allowEmpty, func(value bool) {
			allowEmpty = value
			refresh()
			updateStatus("allow-empty")
		}),
		storybook.CheckboxGroup("Behavior", "Hide steppers", hideSteppers, func(value bool) {
			hideSteppers = value
			refresh()
			updateStatus("hide-steppers")
		}),
	)
}

func inputStatus(source string, input inputValue, required bool) string {
	value := strings.TrimSpace(input.Value())
	if value == "" {
		value = "(empty)"
	}
	validation := "optional"
	if required {
		validation = "required"
	}
	return "Last signal: " + source + ". Validation: " + validation + ". Input value: " + value + "."
}

func numberInputStatus(source string, input inputValue, required bool) string {
	value := strings.TrimSpace(input.Value())
	if value == "" {
		value = "(empty)"
	}
	validation := "optional"
	if required {
		validation = "required"
	}
	return "Last signal: " + source + ". Validation: " + validation + ". Number value: " + value + "."
}

func inputStage(maxWidth string, child mvc.View) dom.Element {
	style := "width:100%"
	if maxWidth != "" {
		style += ";max-width:" + maxWidth
	}
	return mvc.HTML("DIV", mvc.WithStyle(style), child)
}
