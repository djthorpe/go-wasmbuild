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

// Expanded reports whether the header panel is open.
func (p *panel) Expanded() bool {
	return p.Root().HasAttribute("expanded")
}

// SetExpanded opens or closes the header panel.
func (p *panel) SetExpanded(expanded bool) *panel {
	if expanded {
		p.Root().SetAttribute("expanded", "")
	} else {
		p.Root().RemoveAttribute("expanded")
	}
	return p
}

// Open expands the header panel.
func (p *panel) Open() *panel {
	return p.SetExpanded(true)
}

// Close collapses the header panel.
func (p *panel) Close() *panel {
	return p.SetExpanded(false)
}

// Toggle flips the expanded state of the header panel.
func (p *panel) Toggle() *panel {
	return p.SetExpanded(!p.Expanded())
}
