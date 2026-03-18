package carbon

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

// The view names for Carbon Design System components
const (
	ViewSection       = "mvc-cds-section"
	ViewText          = "mvc-cds-text"
	ViewMarkdown      = "mvc-cds-markdown"
	ViewButton        = "mvc-cds-button"
	ViewButtonGroup   = "mvc-cds-button-group"
	ViewIcon          = "mvc-cds-icon"
	ViewNav           = "mvc-cds-nav"
	ViewNavGlobal     = "mvc-cds-nav-global"
	ViewHeaderPanel   = "mvc-cds-header-panel"
	ViewNavItem       = "mvc-cds-navitem"
	ViewGrid          = "mvc-cds-grid"
	ViewTile          = "mvc-cds-tile"
	ViewCheckbox      = "mvc-cds-checkbox"
	ViewCheckboxGroup = "mvc-cds-checkbox-group"
	ViewDropdown      = "mvc-cds-dropdown"
	ViewDropdownItem  = "mvc-cds-dropdown-item"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// base is embedded by all Carbon view structs. Embedding it automatically
// satisfies the viewSetter interface, so setView requires no type switch.
type base struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// viewSetter is satisfied for free by any struct that embeds *base.
type viewSetter interface {
	setMVCView(mvc.View)
}

func (b *base) setMVCView(v mvc.View) { b.View = v }

// setView wires the inner mvc.View (child) onto the outer concrete type (self).
// Works for any type that embeds base — no per-type case needed.
func setView(self mvc.View, child mvc.View) {
	s, ok := self.(viewSetter)
	if !ok {
		panic(fmt.Sprintf("setView: unsupported view type %T", self))
	}
	s.setMVCView(child)
}
