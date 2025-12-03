package bootstrap

import (
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

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
			<input type="text" class="form-control" data-slot="input"></input>
		</span>
	`
	templateLabel = `<label class="form-label"></label>`
)

func init() {
	mvc.RegisterView(ViewInput, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(input), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Input(name string, args ...any) *input {
	return mvc.NewView(new(input), ViewInput, templateInput, setView).ReplaceSlotChildren("input", args...).(*input)
}

func SearchInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "search"), args)
}

func RangeInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "range"), mvc.WithClass("form-range"), mvc.WithoutClass("form-control"), args)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (input *input) Label(children ...any) mvc.View {
	return input.ReplaceSlot("label", mvc.HTML(templateLabel, children...))
	/*
		if elem := input.Slot(""); elem == nil || (elem.TagName() != "INPUT" && elem.TagName() != "TEXTAREA") {
			panic("Label: input body slot is not INPUT or TEXTAREA" + fmt.Sprintf("%v", input))
		} else {
			input.ReplaceSlot("label", mvc.HTML("LABEL", mvc.WithClass("form-label"), mvc.WithAttr("for", elem.ID()), children))
		}
	*/
}

func (input *input) Value() any {
	elem := input.Slot("input")
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

func WithMinMax(min, max int) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" {
			return dom.ErrInternalAppError.Withf("WithMinMax: not supported for view type %q", o.Name())
		}
		if min > max {
			return dom.ErrBadParameter.Withf("WithMinMax: min (%d) must be less than or equal to max (%d)", min, max)
		}
		if err := mvc.WithAttr("min", fmt.Sprintf("%d", min))(o); err != nil {
			return err
		}
		return mvc.WithAttr("max", fmt.Sprintf("%d", max))(o)
	}
}
