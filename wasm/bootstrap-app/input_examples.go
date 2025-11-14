package main

import (
	// Packages

	"fmt"
	"strings"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func InputExamples() mvc.View {
	rangevalue := mvc.Text("")
	return bs.Container().Content(
		bs.Heading(1).Content("Input Examples"),
		bs.HRule(),
		bs.Form("input").Content(
			bs.Card().Header(
				bs.Heading(4).Content("Enter your details"),
			).Content(
				bs.InputGroup(mvc.WithClass("my-2")).Content(
					bs.Input("username", bs.WithPlaceholder("Enter username"), bs.WithRequired(), bs.WithAutocomplete("user", "email")),
					"@",
					bs.Input("domain", bs.WithPlaceholder("Enter domain here"), bs.WithRequired(), bs.WithAutocomplete("domain")),
				),
				bs.Password("password", bs.WithPlaceholder("Enter password here"), mvc.WithClass("my-2"), bs.WithRequired(), bs.WithoutAutocomplete()),
				bs.Number(
					"number", bs.WithMinMax(-5, 5), bs.WithPlaceholder("Enter number here"), mvc.WithClass("my-2"), bs.WithRequired(), bs.WithoutAutocomplete(),
				).Caption(
					"Number of times",
				),
				bs.Textarea("description", bs.WithPlaceholder("Enter description here"), mvc.WithClass("my-2")),
				bs.InputGroup(mvc.WithClass("my-2")).Content(
					bs.Range("range", bs.WithMinMax(-5, 5)).AddEventListener("input", func(e Event) {
						r := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
						if r != nil {
							rangevalue.SetData(r.Value())
						}
					}),
					rangevalue,
				),
			).(mvc.ViewWithHeaderFooter).Footer(
				bs.Button(bs.WithColor(bs.Primary), bs.WithSubmit()).Content("Submit"),
			),
		),
		bs.HRule(),
		bs.Heading(5, "Radio Group Input"),
		ExampleInputRadioGroup(),
		bs.HRule(),
		bs.Heading(5, "Checkbox Group Input"),
		ExampleInputCheckboxGroup(),
		bs.HRule(),
		bs.Heading(5, "Switch Group Input"),
		ExampleInputSwitchGroup(),
		bs.HRule(),
		bs.Heading(5, "Single Selection Input"),
		ExampleInputSelect(),
		bs.HRule(),
		bs.Heading(5, "Multiple Selection Input"),
		ExampleInputMultiSelect(),
		bs.HRule(),
		bs.Heading(5, "Selection Input With Options"),
		ExampleInputMultiSelectOption(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputSelectCode = `bs.Select(
  "single-select", "a", "b", "c"
).AddEventListener("input", func(e Event) {
  s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
  if s != nil {
    response.Content("You selected: " + s.Value())
  }
})`

func ExampleInputSelect() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))

	example := func() mvc.View {
		return bs.Form("select-input").Content(
			bs.Select("single-select", "a", "b", "c").AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(exampleInputSelectCode, bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputMultiSelectCode = `bs.MultiSelect(
  "multi-select", "a", "b", "c"
).AddEventListener("input", func(e Event) {
  s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
  if s != nil {
    response.Content("You selected: " + s.Value())
  }
})`

func ExampleInputMultiSelect() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))

	example := func() mvc.View {
		return bs.Form("multi-select-input").Content(
			bs.MultiSelect("multi-select", "a", "b", "c").AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(exampleInputMultiSelectCode, bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputMultiSelectOptionCode = `bs.MultiSelect(
	"multi-select-option",
	bs.Option("Option A", "a"),
	bs.Option("Option B", "b"),
	bs.Option("Option C", "c"),
).SetValue("a").AddEventListener("input", func(e Event) {
	s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
	if s != nil {
		response.Content("You selected: " + s.Value())
	}
}),
`

func ExampleInputMultiSelectOption() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))

	example := func() mvc.View {
		return bs.Form("multi-select-option-input").Content(
			bs.MultiSelect("multi-select-option",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
				bs.Option("Option D", "d"),
			).SetValue("a").AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(exampleInputMultiSelectOptionCode, bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputRadioGroupCode = `
bs.RadioGroup(
	"radiogroup",
	bs.Option("Option A", "a"),
	bs.Option("Option B", "b"),
	bs.Option("Option C", "c"),
).AddEventListener("input", func(e Event) {
	s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
	if s != nil {
		response.Content("You selected: " + s.Value())
	}
})
`

func ExampleInputRadioGroup() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))

	example := func() mvc.View {
		return bs.Form(
			"radiogroup-input",
			bs.RadioGroup(
				"radiogroup",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				fmt.Println(s)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			bs.Para("...or inline version with ", bs.Code("InlineRadioGroup"), ":", mvc.WithClass("mt-3")),
			bs.InlineRadioGroup(
				"radiogroup-inline",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				fmt.Println(s)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(strings.TrimSpace(exampleInputRadioGroupCode), bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputCheckboxGroupCode = `
bs.CheckboxGroup(
	"checkboxgroup",
	bs.Option("Option A", "a"),
	bs.Option("Option B", "b"),
	bs.Option("Option C", "c"),
).AddEventListener("input", func(e Event) {
	s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
	if s != nil {
		response.Content("You selected: " + s.Value())
	}
})
`

func ExampleInputCheckboxGroup() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))
	example := func() mvc.View {
		return bs.Form(
			"checkboxgroup-input",
			bs.CheckboxGroup(
				"checkboxgroup",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				fmt.Println(s)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			bs.Para("...or inline version with ", bs.Code("InlineCheckboxGroup"), ":", mvc.WithClass("mt-3")),
			bs.InlineCheckboxGroup(
				"checkboxgroup-inline",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				fmt.Println(s)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(strings.TrimSpace(exampleInputCheckboxGroupCode), bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}

/////////////////////////////////////////////////////////////////////

const exampleInputSwitchGroupCode = `
bs.SwitchGroup(
	"switchgroup",
	bs.Option("Option A", "a"),
	bs.Option("Option B", "b"),
	bs.Option("Option C", "c"),
).AddEventListener("input", func(e Event) {
	s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
	if s != nil {
		response.Content("You selected: " + s.Value())
	}
})
`

func ExampleInputSwitchGroup() mvc.View {
	var response mvc.View = bs.Para("", mvc.WithClass("mt-2"))
	example := func() mvc.View {
		return bs.Form(
			"switchgroup-input",
			bs.SwitchGroup(
				"switchgroup",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			bs.Para("...or inline version with ", bs.Code("InlineSwitchGroup"), ":", mvc.WithClass("mt-3")),
			bs.InlineSwitchGroup(
				"switchgroup-inline",
				bs.Option("Option A", "a"),
				bs.Option("Option B", "b"),
				bs.Option("Option C", "c"),
			).AddEventListener("input", func(e Event) {
				s := mvc.ViewFromEvent(e).(mvc.ViewWithValue)
				if s != nil {
					response.Content("You selected: " + s.Value())
				}
			}),
			response,
		)
	}

	code := func() mvc.View {
		return bs.CodeBlock(strings.TrimSpace(exampleInputSwitchGroupCode), bs.WithColor(bs.Light), mvc.WithClass("p-2"), mvc.WithClass("border"))
	}
	return bs.Grid(
		example(), code(),
	)
}
