package bootstrap

import (
	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTabs applies the nav-tabs style (tabbed interface).
// Can only be used with Nav components.
func WithTabs() Opt {
	return func(o *opts) error {
		// Error if this isn't a nav component
		if o.name != NavComponent {
			return ErrBadParameter.Withf("Cannot use WithTabs with component of type %q", o.name)
		}
		o.classList.Add("nav-tabs")
		return nil
	}
}

// WithPills applies the nav-pills style (pill-shaped items).
// Can only be used with Nav components.
func WithPills() Opt {
	return func(o *opts) error {
		// Error if this isn't a nav component
		if o.name != NavComponent {
			return ErrBadParameter.Withf("Cannot use WithPills with component of type %q", o.name)
		}
		o.classList.Add("nav-pills")
		return nil
	}
}

// WithUnderline applies the nav-underline style (underlined items).
// Can only be used with Nav components.
func WithUnderline() Opt {
	return func(o *opts) error {
		// Error if this isn't a nav component
		if o.name != NavComponent {
			return ErrBadParameter.Withf("Cannot use WithUnderline with component of type %q", o.name)
		}
		o.classList.Add("nav-underline")
		return nil
	}
}
