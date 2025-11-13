package bootstrap

import mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithTooltip(name string) mvc.Opt {
	return func(o mvc.OptSet) error {
		// Remove all other tooltips
		if err := mvc.WithoutAttr("data-bs-toggle", "data-bs-title")(o); err != nil {
			return err
		}
		// Add tooltip attributes
		if err := mvc.WithAttr("data-bs-toggle", "tooltip")(o); err != nil {
			return err
		}
		if err := mvc.WithAttr("data-bs-title", name)(o); err != nil {
			return err
		}
		return nil
	}
}
