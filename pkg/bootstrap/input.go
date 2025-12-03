package bootstrap

import (
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
			<slot name="label"></slot>
			<input type="text" class="form-control"></input>
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
	view := mvc.NewView(new(input), ViewInput, templateInput, setView)
	// TODO: Set attributes on input element
	return view.(*input)
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

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithPlaceholder(placeholder string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != "INPUT" && o.Name() != "TEXTAREA" && o.Name() != ViewInput {
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
