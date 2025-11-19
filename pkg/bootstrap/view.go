package bootstrap

import (
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
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
	return b.replaceSlot("", children...)
}

// Header sets the header slot content. Components must override this method.
func (b *BootstrapView) Header(children ...any) *BootstrapView {
	return b.replaceSlot("header", children...)
}

// Footer sets the footer slot content. Components must override this method.
func (b *BootstrapView) Footer(children ...any) *BootstrapView {
	return b.replaceSlot("footer", children...)
}

// Label sets the label slot content. Components must override this method.
func (b *BootstrapView) Label(children ...any) *BootstrapView {
	return b.replaceSlot("label", children...)
}

// SetContent updates the default slot on any Bootstrap-compatible view.
func SetContent(view mvc.View, children ...any) mvc.View {
	if view == nil {
		return nil
	}
	view.ReplaceSlot("", wrapChildren(children...))
	return view
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (b *BootstrapView) replaceSlot(name string, children ...any) *BootstrapView {
	if b == nil || b.View == nil {
		panic(fmt.Sprintf("BootstrapView %q has no underlying view", b.Name()))
	}
	b.View.ReplaceSlot(name, wrapChildren(children...))
	return b
}

// wrapChildren wraps multiple children in a container element if needed
func wrapChildren(children ...any) dom.Element {
	switch len(children) {
	case 0:
		return mvc.HTML("div")
	case 1:
		switch child := children[0].(type) {
		case dom.Element:
			return child
		case mvc.View:
			return child.Root()
		default:
			div := mvc.HTML("div")
			div.AppendChild(mvc.NodeFromAny(child))
			return div
		}
	default:
		div := mvc.HTML("div")
		for _, child := range children {
			div.AppendChild(mvc.NodeFromAny(child))
		}
		return div
	}
}
