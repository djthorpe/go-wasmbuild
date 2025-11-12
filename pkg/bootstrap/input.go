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

type textarea struct {
	mvc.View
}

type rangeinput struct {
	mvc.View
}

var _ mvc.View = (*form)(nil)
var _ mvc.View = (*inputgroup)(nil)
var _ mvc.ViewWithValue = (*textarea)(nil)
var _ mvc.ViewWithValue = (*rangeinput)(nil)
var _ mvc.ViewWithValue = (*input)(nil)
var _ mvc.ViewWithCaption = (*input)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewForm       = "mvc-bs-form"
	ViewInput      = "mvc-bs-input"
	ViewInputGroup = "mvc-bs-inputgroup"
	ViewTextarea   = "mvc-bs-textarea"
	ViewRange      = "mvc-bs-range"
)

func init() {
	mvc.RegisterView(ViewForm, newFormFromElement)
	mvc.RegisterView(ViewInput, newInputFromElement)
	mvc.RegisterView(ViewInputGroup, newInputGroupFromElement)
	mvc.RegisterView(ViewTextarea, newTextareaFromElement)
	mvc.RegisterView(ViewRange, newRangeFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Form(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("name", name)}, opt...)
	return mvc.NewView(new(form), ViewForm, "FORM", append([]mvc.Opt{mvc.WithClass("form")}, opt...)...)
}

func Input(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("id", name)}, opt...)
	label := mvc.HTML("LABEL", mvc.WithClass("form-label"), mvc.WithAttr("for", name))
	return mvc.NewViewEx(new(input), ViewInput, "INPUT", nil, nil, nil, label, append([]mvc.Opt{mvc.WithClass("form-control")}, opt...)...)
}

func Password(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("type", "password")}, opt...)
	return Input(name, opt...)
}

func Number(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("type", "number")}, opt...)
	return Input(name, opt...)
}

func InputGroup(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(inputgroup), ViewInputGroup, "DIV", append([]mvc.Opt{mvc.WithClass("input-group")}, opt...)...)
}

func Textarea(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("id", name)}, opt...)
	return mvc.NewView(new(textarea), ViewTextarea, "TEXTAREA", append([]mvc.Opt{mvc.WithClass("form-control")}, opt...)...)
}

func Range(name string, opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithAttr("id", name), mvc.WithAttr("type", "range")}, opt...)
	return mvc.NewView(new(rangeinput), ViewRange, "INPUT", append([]mvc.Opt{mvc.WithClass("form-range")}, opt...)...)
}

func newFormFromElement(element Element) mvc.View {
	if element.TagName() != "FORM" {
		return nil
	}
	return mvc.NewViewWithElement(new(form), element)
}

func newInputFromElement(element Element) mvc.View {
	if element.TagName() != "INPUT" {
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

func newTextareaFromElement(element Element) mvc.View {
	if element.TagName() != "TEXTAREA" {
		return nil
	}
	return mvc.NewViewWithElement(new(textarea), element)
}

func newRangeFromElement(element Element) mvc.View {
	if element.TagName() != "INPUT" {
		return nil
	}
	return mvc.NewViewWithElement(new(rangeinput), element)
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

func (textarea *textarea) SetView(view mvc.View) {
	textarea.View = view
}

func (rangeinput *rangeinput) SetView(view mvc.View) {
	rangeinput.View = view
}

func (input *input) Append(children ...any) mvc.View {
	// TODO: This should be the input "value" and should only accept text
	panic("Append: not supported for input")
}

func (input *input) Caption(children ...any) mvc.ViewWithCaption {
	// TODO: Create the caption element
	input.View.(mvc.ViewWithCaption).Caption(children...)
	//<label for="inputPassword5" class="form-label">Password</label>
	return input
}

func (inputgroup *inputgroup) Append(children ...any) mvc.View {
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

func (input *input) Value() string {
	return input.Root().Value()
}

func (textarea *textarea) Value() string {
	return textarea.Root().TextContent()
}

func (rangeinput *rangeinput) Value() string {
	return rangeinput.Root().Value()
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithPlaceholder(placeholder string) mvc.Opt {
	return func(o mvc.OptSet) error {
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
		if o.Name() != ViewInput && o.Name() != ViewTextarea {
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
		if o.Name() != ViewInput && o.Name() != ViewTextarea {
			return ErrInternalAppError.Withf("WithAutocomplete: not supported for view type %q", o.Name())
		}
		return mvc.WithAttr("autocomplete", "off")(o)
	}
}

func WithMinMax(min, max int) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewInput && o.Name() != ViewRange {
			return ErrInternalAppError.Withf("WithMinMax: not supported for view type %q", o.Name())
		}
		if min > max {
			return ErrBadParameter.Withf("WithMinMax: min (%d) must be less than or equal to max (%d)", min, max)
		}
		if err := mvc.WithAttr("min", fmt.Sprintf("%d", min))(o); err != nil {
			return err
		}
		return mvc.WithAttr("max", fmt.Sprintf("%d", max))(o)
	}
}
