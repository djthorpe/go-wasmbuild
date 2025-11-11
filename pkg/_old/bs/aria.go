package bs

import (
	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithAriaLabel adds an aria-label attribute to a view
func WithAriaLabel(label string) mvc.Opt {
	return func(o mvc.OptSet) error {
		return mvc.WithAttr("aria-label", label)(o)
	}
}
