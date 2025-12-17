package bootstrap

import (
	"fmt"

	// Packages
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
		<div>		
			<script data-slot="label"></script>
			<input type="text" class="form-control" data-slot="input"></input>
		</div>
	`
	templateRangeInput = `
		<div>		
			<script data-slot="label"></script>
			<input type="range" class="form-range" data-slot="input"></input>
		</div>
	`
	templateSecureInput = `
		<div>
			<script data-slot="label"></script>
			<div class="position-relative">
				<input type="password" class="form-control pe-5" data-slot="input">			
				<i class="bi-key-fill fs-5 position-absolute top-50 end-0 translate-middle-y me-3 text-muted"></i>
			</div>
		</div>
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
	return mvc.NewView(new(input), ViewInput, templateInput, setView, withInputID(name), args).(*input)
}

func SearchInput(name string, args ...any) *input {
	return Input(name, mvc.WithAttr("type", "search"), withInputID(name), args)
}

func SecureInput(name string, args ...any) *input {
	return mvc.NewView(new(input), ViewInput, templateSecureInput, setView, withInputID(name), args).(*input)
}

func RangeInput(name string, args ...any) *input {
	return mvc.NewView(new(input), ViewInput, templateRangeInput, setView, withInputID(name), args).(*input)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (input *input) Label(children ...any) mvc.View {
	// TODO: Use the input slot
	id := input.ID()
	if id == "" {
		panic("Label: input has no ID" + fmt.Sprintf("%v", input))
	}
	return input.ReplaceSlot("label", mvc.HTML(templateLabel, mvc.WithAttr("for", id), children))
}

func (input *input) Value() any {
	elem := input.Slot("input")
	if elem == nil || elem.TagName() != "INPUT" {
		panic("Value: input slot is not INPUT" + fmt.Sprintf("%v", input))
	}
	return elem.Value()
}

func (form *form) Value() any {
	// TODO: Get the form data
	return form.Root().Data()
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func withInputID(name string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewInput {
			return dom.ErrInternalAppError.Withf("WithInputID: not supported for view type %q", o.Name())
		}
		if err := mvc.WithSlotAttr("input", "name", name)(o); err != nil {
			return err
		}
		if err := mvc.WithSlotAttr("input", "id", mvc.Counter(name))(o); err != nil {
			return err
		}
		return nil
	}
}

func WithPlaceholder(placeholder string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewInput {
			return dom.ErrInternalAppError.Withf("WithPlaceholder: not supported for view type %q", o.Name())
		}
		if err := mvc.WithSlotAttr("input", "placeholder", placeholder)(o); err != nil {
			return err
		}
		if err := mvc.WithSlotAttr("input", "aria-label", placeholder)(o); err != nil {
			return err
		}
		return nil
	}
}

// Range inputs can set a minimum and maximum value
func WithMinMax(min, max int) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewInput {
			return dom.ErrInternalAppError.Withf("WithMinMax: not supported for view type %q", o.Name())
		}
		if min >= max {
			return dom.ErrBadParameter.Withf("WithMinMax: min (%d) must be less than max (%d)", min, max)
		}
		if err := mvc.WithSlotAttr("input", "min", fmt.Sprintf("%d", min))(o); err != nil {
			return err
		}
		return mvc.WithSlotAttr("input", "max", fmt.Sprintf("%d", max))(o)
	}
}
