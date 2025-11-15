package bootstrap

import (
	// Packages

	"fmt"
	"strings"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type form struct {
	mvc.View
}

type input struct {
	mvc.View
}

type inputgroup struct {
	mvc.View
}

type selectinput struct {
	mvc.View
}

type inputswitch struct {
	mvc.View
}

type inputoption struct {
	Name  string
	Value string
}

var _ mvc.View = (*form)(nil)
var _ mvc.View = (*inputgroup)(nil)
var _ mvc.View = (*input)(nil)
var _ mvc.View = (*selectinput)(nil)
var _ mvc.View = (*inputswitch)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewForm          = "mvc-bs-form"
	ViewInput         = "mvc-bs-input"
	ViewInputGroup    = "mvc-bs-inputgroup"
	ViewSelect        = "mvc-bs-select"
	ViewRadioGroup    = "mvc-bs-radiogroup"
	ViewCheckboxGroup = "mvc-bs-checkboxgroup"

	// Class used to indicate inline groups
	classInlineGroup = "mvc-bs-inlinegroup"
)

const (
	templateInput = `
		<div>
			<slot name="label"></slot>
			<slot></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewForm, newFormFromElement)
	mvc.RegisterView(ViewInput, newInputFromElement)
	mvc.RegisterView(ViewInputGroup, newInputGroupFromElement)
	mvc.RegisterView(ViewSelect, newSelectFromElement)
	mvc.RegisterView(ViewRadioGroup, newRadioGroupFromElement)
	mvc.RegisterView(ViewCheckboxGroup, newCheckboxGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Form(name string, args ...any) *form {
	return mvc.NewView(new(form), ViewForm, "FORM", mvc.WithAttr("name", name), args).(*form)
}

func Input(name string, args ...any) *input {
	// Make the base input view
	view := mvc.NewViewExEx(new(input), ViewInput, templateInput).(*input)

	// Replace the content body with an input element
	return view.ReplaceSlot("", mvc.HTML("INPUT", mvc.WithAttr("id", name), mvc.WithClass("form-control"), args)).(*input)
}

func PasswordInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "password"), args)
}

func NumberInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "number"), args)
}

func RangeInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "range"), mvc.WithClass("form-range"), mvc.WithoutClass("form-control"), args)
}

func SearchInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "search"), mvc.WithClass("form-control"), args)
}

func InputGroup(args ...any) *inputgroup {
	return mvc.NewView(new(inputgroup), ViewInputGroup, "DIV", mvc.WithClass("input-group"), args).(*inputgroup)
}

func Textarea(name string, args ...any) *input {
	// Make the base input view
	view := mvc.NewViewExEx(new(input), ViewInput, templateInput).(*input)

	// Replace the content body with an input element
	return view.ReplaceSlot("", mvc.HTML("TEXTAREA", mvc.WithAttr("id", name), mvc.WithClass("form-control"), args)).(*input)
}

func Select(name string, args ...any) *selectinput {
	return mvc.NewView(new(selectinput), ViewSelect, "SELECT", mvc.WithAttr("id", name), mvc.WithClass("form-select"), args).(*selectinput)
}

func MultiSelect(name string, args ...any) *selectinput {
	return mvc.NewView(new(selectinput), ViewSelect, "SELECT", mvc.WithAttr("id", name), mvc.WithAttr("multiple", "multiple"), mvc.WithClass("form-select"), args).(*selectinput)
}

func RadioGroup(name string, args ...any) *inputswitch {
	return mvc.NewView(new(inputswitch), ViewRadioGroup, "DIV", mvc.WithAttr("id", name), args).(*inputswitch)
}

func InlineRadioGroup(name string, args ...any) *inputswitch {
	return RadioGroup(name, mvc.WithClass(classInlineGroup), args)
}

func CheckboxGroup(name string, args ...any) *inputswitch {
	return mvc.NewView(new(inputswitch), ViewCheckboxGroup, "DIV", mvc.WithAttr("id", name), args).(*inputswitch)
}

func InlineCheckboxGroup(name string, args ...any) *inputswitch {
	return CheckboxGroup(name, mvc.WithClass(classInlineGroup), args)
}

func SwitchGroup(name string, args ...any) *inputswitch {
	return CheckboxGroup(name, mvc.WithClass("form-switch"), args)
}

func InlineSwitchGroup(name string, args ...any) *inputswitch {
	return SwitchGroup(name, mvc.WithClass(classInlineGroup), args)
}

func Option(name, value string) *inputoption {
	return &inputoption{Name: name, Value: value}
}

func newFormFromElement(element Element) mvc.View {
	if element.TagName() != "FORM" {
		return nil
	}
	return mvc.NewViewWithElement(new(form), element)
}

func newInputFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(input), element)
}

func newInputGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(inputgroup), element)
}

func newSelectFromElement(element Element) mvc.View {
	if element.TagName() != "SELECT" {
		return nil
	}
	return mvc.NewViewWithElement(new(selectinput), element)
}

func newRadioGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(inputswitch), element)
}

func newCheckboxGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(inputswitch), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (form *form) SetView(view mvc.View) {
	form.View = view
}

func (input *input) SetView(view mvc.View) {
	input.View = view
}

func (inputgroup *inputgroup) SetView(view mvc.View) {
	inputgroup.View = view
}

func (selectinput *selectinput) SetView(view mvc.View) {
	selectinput.View = view
}

func (inputswitch *inputswitch) SetView(view mvc.View) {
	inputswitch.View = view
}

func (input *input) Label(children ...any) mvc.View {
	if elem := input.Slot(""); elem == nil || (elem.TagName() != "INPUT" && elem.TagName() != "TEXTAREA") {
		panic("Label: input body slot is not INPUT or TEXTAREA" + fmt.Sprintf("%v", input))
	} else {
		input.ReplaceSlot("label", mvc.HTML("LABEL", mvc.WithClass("form-label"), mvc.WithAttr("for", elem.ID()), children))
	}
	return input
}

func (inputgroup *inputgroup) Content(children ...any) mvc.View {
	// Wrap all text children in span with class "input-group-text"
	for _, child := range children {
		switch child.(type) {
		case string:
			col := mvc.HTML("SPAN", mvc.WithClass("input-group-text"))
			col.AppendChild(mvc.NodeFromAny(child))
			inputgroup.View.Append(col)
		default:
			inputgroup.View.Append(child)
		}
	}
	return inputgroup
}

func (inputswitch *inputswitch) Content(children ...any) mvc.View {
	isInline := inputswitch.Root().ClassList().Contains(classInlineGroup)

	// Factory function to create switch element
	switchFactory := func(index int, opt *inputoption) Element {
		classes := []string{"form-check"}
		if isInline {
			classes = append(classes, "form-check-inline", "mx-3")
		}
		div := mvc.HTML("DIV", mvc.WithClass(classes...))
		id := fmt.Sprintf("%s-%d", inputswitch.Root().ID(), index)
		var input Element
		switch inputswitch.Name() {
		case ViewRadioGroup:
			input = mvc.HTML("INPUT",
				mvc.WithID(id),
				mvc.WithClass("form-check-input"),
				mvc.WithAttr("type", "radio"),
				mvc.WithAttr("name", inputswitch.Root().ID()),
				mvc.WithAttr("value", opt.Value),
			)
		case ViewCheckboxGroup:
			input = mvc.HTML("INPUT",
				mvc.WithID(id),
				mvc.WithClass("form-check-input"),
				mvc.WithAttr("type", "checkbox"),
				mvc.WithAttr("value", opt.Value),
			)
		default:
			panic("Append: unsupported inputswitch type")
		}
		label := mvc.HTML("LABEL",
			mvc.WithClass("form-check-label"),
			mvc.WithAttr("for", id),
		)
		label.AppendChild(mvc.NodeFromAny(opt.Name))
		div.AppendChild(input)
		div.AppendChild(label)
		return div
	}

	// Wrap all text children in span with class "input-group-text"
	for i, child := range children {
		switch child := child.(type) {
		case string:
			inputswitch.View.Append(switchFactory(i, &inputoption{
				Name:  child,
				Value: child,
			}))
		case *inputoption:
			inputswitch.View.Append(switchFactory(i, child))
		default:
			panic("Append: unsupported child type for select input")
		}
	}
	return inputswitch
}

func (selectinput *selectinput) Append(children ...any) mvc.View {
	// Wrap all text children in option elements
	for _, child := range children {
		switch child := child.(type) {
		case string:
			opt := mvc.HTML("OPTION")
			opt.AppendChild(mvc.NodeFromAny(child))
			selectinput.View.Append(opt)
		case *inputoption:
			opt := mvc.HTML("OPTION", mvc.WithAttr("value", child.Value))
			opt.AppendChild(mvc.NodeFromAny(child.Name))
			selectinput.View.Append(opt)
		default:
			panic("Append: unsupported child type for select input")
		}
	}
	return selectinput
}

func (input *input) Value() string {
	return input.Root().Value()
}

func (selectinput *selectinput) Value() string {
	return selectinput.Root().Value()
}

func (input *input) Set(value string) mvc.View {
	input.Root().SetValue(value)
	return input
}

func (selectinput *selectinput) Set(value string) mvc.View {
	selectinput.Root().SetValue(value)
	return selectinput
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithPlaceholder(placeholder string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" && o.Name() != "TEXTAREA" {
			return ErrInternalAppError.Withf("WithPlaceholder: not supported for view type %q", o.Name())
		}
		if err := mvc.WithAttr("placeholder", placeholder)(o); err != nil {
			return err
		}
		if err := mvc.WithAttr("aria-label", placeholder)(o); err != nil {
			return err
		}
		return nil
	}
}

func WithRequired() mvc.Opt {
	return mvc.WithAttr("required", "required")
}

func WithAutocomplete(tokens ...string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" && o.Name() != "TEXTAREA" {
			return ErrInternalAppError.Withf("WithAutocomplete: not supported for view type %q", o.Name())
		}
		if len(tokens) == 0 {
			return mvc.WithAttr("autocomplete", "on")(o)
		}
		return mvc.WithAttr("autocomplete", strings.Join(tokens, " "))(o)
	}
}

func WithoutAutocomplete() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" && o.Name() != "TEXTAREA" {
			return ErrInternalAppError.Withf("WithoutAutocomplete: not supported for view type %q", o.Name())
		}
		return mvc.WithAttr("autocomplete", "off")(o)
	}
}

func WithMinMax(min, max int) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" && o.Name() != "TEXTAREA" && o.Name() != ViewProgress {
			return ErrInternalAppError.Withf("WithMinMax: not supported for view type %q", o.Name())
		}
		if min > max {
			return ErrBadParameter.Withf("WithMinMax: min (%d) must be less than or equal to max (%d)", min, max)
		}
		switch o.Name() {
		case ViewProgress:
			if err := mvc.WithAttr("aria-valuemin", fmt.Sprintf("%d", min))(o); err != nil {
				return err
			}
			return mvc.WithAttr("aria-valuemax", fmt.Sprintf("%d", max))(o)
		default:
			if err := mvc.WithAttr("min", fmt.Sprintf("%d", min))(o); err != nil {
				return err
			}
			return mvc.WithAttr("max", fmt.Sprintf("%d", max))(o)
		}
	}
}
