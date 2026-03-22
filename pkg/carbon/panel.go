package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type panel struct{ base }

var _ mvc.View = (*panel)(nil)
var _ mvc.VisibleState = (*panel)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewHeaderPanel, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(panel), element, setView)
	})
}

// HeaderPanel returns a <cds-header-panel> UI shell right panel.
func HeaderPanel(args ...any) *panel {
	return mvc.NewView(new(panel), ViewHeaderPanel, "cds-header-panel", setView, args...).(*panel)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Visible reports whether the panel is expanded (visible).
func (p *panel) Visible() bool {
	return tagBoolProperty(p.Root(), "expanded")
}

// SetVisible shows or hides the panel by setting the expanded property.
func (p *panel) SetVisible(visible bool) mvc.View {
	setTagBoolProperty(p.Root(), "expanded", visible)
	return p
}

