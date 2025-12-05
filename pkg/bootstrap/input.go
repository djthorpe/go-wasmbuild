package bootstrap

import (
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type form struct {
	mvc.View
}

type input struct {
	mvc.View
}

var _ mvc.View = (*input)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateInput = `
		<span>		
			<script data-slot="label"></script>
			<input type="text" class="form-control" data-slot="body"></input>
		</span>
	`
	templateLabel = `<label class="form-label"></label>`
)

func init() {
	mvc.RegisterView(ViewForm, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(form), element, setView)
	})
	mvc.RegisterView(ViewInput, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(input), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Form(name string, args ...any) *form {
	return mvc.NewView(new(form), ViewForm, "FORM", setView, mvc.WithAttr("name", name), args).(*form)
}

func Input(name string, args ...any) *input {
	return mvc.NewView(new(input), ViewInput, templateInput, setView).ReplaceSlotChildren("body", args...).(*input)
}

func SearchInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "search"), args)
}

func SecureInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "password"), args)
}

func RangeInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "range"), mvc.WithClass("form-range"), mvc.WithoutClass("form-control"), args)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (input *input) Label(children ...any) mvc.View {
	// TODO: Add "for" attribute to label
	return input.ReplaceSlot("label", mvc.HTML(templateLabel, children...))
}

func (input *input) Value() any {
	elem := input.Slot("body")
	if elem == nil || elem.TagName() != "INPUT" {
		panic("Value: input slot is not INPUT" + fmt.Sprintf("%v", input))
	}
	return elem.Value()
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithPlaceholder(placeholder string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" {
			return dom.ErrInternalAppError.Withf("WithPlaceholder: not supported for view type %q", o.Name())
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

// Range inputs can set a minimum and maximum value
func WithMinMax(min, max int) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" {
			return dom.ErrInternalAppError.Withf("WithMinMax: not supported for view type %q", o.Name())
		}
		if min >= max {
			return dom.ErrBadParameter.Withf("WithMinMax: min (%d) must be less than max (%d)", min, max)
		}
		if err := mvc.WithAttr("min", fmt.Sprintf("%d", min))(o); err != nil {
			return err
		}
		return mvc.WithAttr("max", fmt.Sprintf("%d", max))(o)
	}
}
