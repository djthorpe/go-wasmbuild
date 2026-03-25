package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type buttonGroup struct{ base }

var _ mvc.View = (*buttonGroup)(nil)
var _ mvc.EnabledGroup = (*buttonGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewButtonGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(buttonGroup), element, setView)
	}, EventClick, EventHover, EventNoHover, EventFocus, EventNoFocus)
}

// ButtonGroup returns a <cds-button-group> web component that arranges
// buttons horizontally with correct Carbon spacing.
func ButtonGroup(args ...any) *buttonGroup {
	return mvc.NewView(new(buttonGroup), ViewButtonGroup, "cds-button-group", setView, args).(*buttonGroup)
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

func (g *buttonGroup) Enabled() []mvc.View {
	enabled := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok && b.Enabled() {
			enabled = append(enabled, b)
		}
	}
	return enabled
}

func (g *buttonGroup) SetEnabled(views ...mvc.View) mvc.View {
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok {
			// Compare by root element identity rather than interface identity so
			// callers can pass equivalent view wrappers for the same DOM element.
			on := false
			for _, view := range views {
				if view != nil && child.Root().Equals(view.Root()) {
					on = true
					break
				}
			}
			b.SetEnabled(on)
		}
	}
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Content appends buttons to the group. Panics if any arg is not a *button.
func (g *buttonGroup) Content(args ...any) mvc.View {
	children := make([]any, 0, len(args))
	for _, arg := range args {
		if b, ok := arg.(*button); ok {
			children = append(children, b)
		} else {
			panic(fmt.Sprintf("ButtonGroup.Content: expected *button, got %T", arg))
		}
	}
	return g.View.Content(children...)
}