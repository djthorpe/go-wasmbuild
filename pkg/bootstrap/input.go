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

func init() {
	mvc.RegisterView(ViewInput, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(input), element, func(self, child mvc.View) {
			self.(*input).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Input(name string, args ...any) mvc.View {
	return mvc.NewView(new(input), ViewInput, "INPUT", func(self, child mvc.View) {
		self.(*input).View = child
	}, mvc.WithClass("form-control"), mvc.WithAttr("name", name), args)
}

func SearchInput(name string, args ...any) mvc.View {
	return Input(name, mvc.WithAttr("type", "search"), args)
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
