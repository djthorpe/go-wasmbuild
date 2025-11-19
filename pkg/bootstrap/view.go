package bootstrap

import (
	"fmt"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// BootstrapView wraps mvc.View and adds Bootstrap-specific slot methods
type BootstrapView struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - Bootstrap-specific

// Content sets the body slot content. Components must override this method.
func (b *BootstrapView) Content(children ...any) *BootstrapView {
	panic(fmt.Sprintf("Content() not implemented in view %q", b.Name()))
}

// Header sets the header slot content. Components must override this method.
func (b *BootstrapView) Header(children ...any) *BootstrapView {
	panic(fmt.Sprintf("Header() not implemented in view %q", b.Name()))
}

// Footer sets the footer slot content. Components must override this method.
func (b *BootstrapView) Footer(children ...any) *BootstrapView {
	panic(fmt.Sprintf("Footer() not implemented in view %q", b.Name()))
}

// Label sets the label slot content. Components must override this method.
func (b *BootstrapView) Label(children ...any) *BootstrapView {
	panic(fmt.Sprintf("Label() not implemented in view %q", b.Name()))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// wrapChildren wraps multiple children in a container element if needed
func wrapChildren(children ...any) any {
	if len(children) == 0 {
		return mvc.HTML("div")
	}
	if len(children) == 1 {
		return children[0]
	}
	// Multiple children - wrap in a div
	div := mvc.HTML("div")
	for _, child := range children {
		div.AppendChild(mvc.NodeFromAny(child))
	}
	return div
}
