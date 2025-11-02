package bs

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithAlt adds an alt attribute to the image view
func WithAlt(alt string) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewImage {
			return fmt.Errorf("WithAlt: invalid view type %q", o.Name())
		}
		return mvc.WithAttr("alt", alt)(o)
	}
}
