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
	ViewSection             = "mvc-cds-section"
	ViewForm                = "mvc-cds-form"
	ViewFormGroup           = "mvc-cds-form-group"
	ViewInput               = "mvc-cds-input"
	ViewSecureInput         = "mvc-cds-secure-input"
	ViewNumberInput         = "mvc-cds-number-input"
	ViewLink                = "mvc-cds-link"
	ViewText                = "mvc-cds-text"
	ViewList                = "mvc-cds-list"
	ViewMarkdown            = "mvc-cds-markdown"
	ViewButton              = "mvc-cds-button"
	ViewButtonGroup         = "mvc-cds-button-group"
	ViewIcon                = "mvc-cds-icon"
	ViewNav                 = "mvc-cds-nav"
	ViewNavGlobal           = "mvc-cds-nav-global"
	ViewHeaderPanel         = "mvc-cds-header-panel"
	ViewOverflowMenu        = "mvc-cds-overflow-menu"
	ViewOverflowItem        = "mvc-cds-overflow-menu-item"
	ViewNavItem             = "mvc-cds-navitem"
	ViewGrid                = "mvc-cds-grid"
	ViewTile                = "mvc-cds-tile"
	ViewStructuredList      = "mvc-cds-structured-list"
	ViewStructuredListHead  = "mvc-cds-structured-list-head"
	ViewStructuredListRow   = "mvc-cds-structured-list-row"
	ViewStructuredListCell  = "mvc-cds-structured-list-cell"
	ViewStructuredListTH    = "mvc-cds-structured-list-header-cell"
	ViewTable               = "mvc-cds-table"
	ViewTableHeader         = "mvc-cds-table-header"
	ViewTableRow            = "mvc-cds-table-row"
	ViewTableToolbar        = "mvc-cds-table-toolbar"
	ViewTableToolbarContent = "mvc-cds-table-toolbar-content"
	ViewTableToolbarSearch  = "mvc-cds-table-toolbar-search"
	ViewPagination          = "mvc-cds-pagination"
	ViewCheckbox            = "mvc-cds-checkbox"
	ViewCheckboxGroup       = "mvc-cds-checkbox-group"
	ViewDropdown            = "mvc-cds-dropdown"
	ViewDropdownItem        = "mvc-cds-dropdown-item"
	ViewCodeSnippet         = "mvc-cds-code-snippet"
	ViewTag                 = "mvc-cds-tag"
	ViewTagGroup            = "mvc-cds-tag-group"
	ViewDismissibleTag      = "mvc-cds-dismissible-tag"
	ViewOperationalTag      = "mvc-cds-operational-tag"
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
