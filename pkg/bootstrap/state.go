package bootstrap

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithDisabled adds a disabled attribute to a view
func WithDisabled(disabled bool) mvc.Opt {
	return func(o mvc.OptSet) error {
		switch o.Name() {
		case ViewButton:
			if disabled {
				return mvc.WithAttr("disabled", "disabled")(o)
			} else {
				return mvc.WithoutAttr("disabled")(o)
			}
		case ViewPaginationItem:
			if disabled {
				return mvc.WithClass("disabled")(o)
			} else {
				return mvc.WithoutClass("disabled")(o)
			}
		default:
			return fmt.Errorf("WithDisabled: invalid view type %q", o.Name())
		}
	}
}

// WithActive adds an active attribute to a view
func WithActive(active bool) mvc.Opt {
	return func(o mvc.OptSet) error {
		switch o.Name() {
		case ViewButton, ViewPaginationItem:
			if active {
				return mvc.WithClass("active")(o)
			} else {
				return mvc.WithoutClass("active")(o)
			}
		default:
			return fmt.Errorf("WithActive: invalid view type %q", o.Name())
		}
	}
}

// WithActiveToggle adds an active attribute to a view, and allows for toggling
// of the state
func WithActiveToggle() mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewButton {
			return fmt.Errorf("WithActiveToggle: invalid view type %q", o.Name())
		}
		return mvc.WithAttr("data-bs-toggle", "button")(o)
	}
}
