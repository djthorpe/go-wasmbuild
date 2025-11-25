package main

import (
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type Row struct {
	Name        string
	Description string
	Label       bool
	Header      bool
	Footer      bool
}

var (
	views = []Row{
		{"Heading(level, any)", "A heading view, with level 1-6", false, false, false},
		{"Para(any)", "A simple paragraph view", false, false, false},
		{"LeadPara(any)", "Lead paragraph view", false, false, false},
		{"Blockquote(any)", "Blockquote view with optional label", true, false, false},
		{"Deleted(any)", "Inline deleted text view", false, false, false},
		{"Highlighted(any)", "Inline highlighted text view", false, false, false},
		{"Smaller(any)", "Inline smaller text view", false, false, false},
		{"Strong(any)", "Inline strong text view", false, false, false},
		{"Em(any)", "Inline emphasized text view", false, false, false},
		{"Code(any)", "An inline code view", false, false, false},
		{"CodeBlock(any)", "A block of code view", false, false, false},
		{"Link(href, any)", "A link view", false, false, false},
		{"List(any)", "An ordered list view", false, false, false},
		{"BulletList(any)", "A bullet list view", false, false, false},
		{"DefinitionList([]Option)", "A definition list view", false, false, false},
		{"UnstyledList(any)", "An unstyled list view", false, false, false},
		{"Badge(any)", "A badge view", false, false, false},
		{"PillBadge(any)", "A pill (rounded) badge view", false, false, false},
		{"Button(any)", "A button view", true, false, false},
		{"OutlineButton(any)", "An outline button view", true, false, false},
		{"CloseButton()", "A close button view, cannot contain children", false, false, false},
		{"ButtonGroup([]Button)", "A button group view, can only contain buttons", false, false, false},
		{"VButtonGroup([]Button)", "A vertical button group view, can only contain buttons", false, false, false},
		{"ButtonToolbar([]ButtonGroup)", "A button toolbar view, can only contain button groups", false, false, false},
		{"Card(any)", "A card view", true, true, true},
		{"CardGroup([]Card)", "A group of cards. Can only contain cards", false, false, false},
		{"Container(any)", "A container view", false, false, false},
		{"FluidContainer(any)", "A fluid container view", false, false, false},
		{"Icon(name)", "An icon view", false, false, false},
		{"Image(href, any)", "An image view", false, false, false},
		{"Form(name, any)", "A form view", false, false, false},
		{"Input(name, any)", "A text input view", true, false, false},
		{"PasswordInput(name, any)", "A password input view", true, false, false},
		{"NumberInput(name, any)", "A number input view", true, false, false},
		{"RangeInput(name, any)", "A number input view as a slider", true, false, false},
		{"SearchInput(name, any)", "A text input view as a search box", true, false, false},
		{"TextArea(name, any)", "A text area view", true, false, false},
		{"Select(name, []Option)", "A select view with options", true, false, false},
		{"MultiSelect(name, []Option)", "A select view which allows for selecting more than one option", true, false, false},
		{"RadioGroup(name, []Option)", "A radio group (select one from many)", true, false, false},
		{"InlineRadioGroup(name, []Option)", "A horizontal set of radio buttons", true, false, false},
		{"CheckboxGroup(name, []Option)", "A checkbox group (select one or more)", true, false, false},
		{"InlineCheckboxGroup(name, []Option)", "An inline checkbox group", true, false, false},
		{"SwitchGroup(name, []Option)", "A switch group", true, false, false},
		{"InlineSwitchGroup(name, []Option)", "An inline switch group", true, false, false},
	}
)

func Views() mvc.View {
	rows := []any{}
	boolToIcon := func(b bool) mvc.View {
		if b {
			return bs.Icon("check-lg", bs.WithColor(bs.Success))
		}
		return bs.Para()
	}
	for _, view := range views {
		rows = append(rows, bs.Row(
			bs.Code(view.Name),
			view.Description,
			boolToIcon(view.Label),
			boolToIcon(view.Header),
			boolToIcon(view.Footer),
		))
	}
	return bs.Container(
		bs.Table(mvc.WithClass("table-bordered"), rows).Header("View", "Description", "Label", "Header", "Footer"),
	)
}
