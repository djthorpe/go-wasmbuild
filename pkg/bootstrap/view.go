package bootstrap

import (
	"fmt"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// The view names for Bootstrap components
const (
	ViewAccordion      = "mvc-bs-accordion"
	ViewAccordionItem  = "mvc-bs-accordionitem"
	ViewBadge          = "mvc-bs-badge"
	ViewButton         = "mvc-bs-button"
	ViewButtonGroup    = "mvc-bs-buttongroup"
	ViewButtonToolbar  = "mvc-bs-buttontoolbar"
	ViewCard           = "mvc-bs-card"
	ViewCardGroup      = "mvc-bs-cardgroup"
	ViewCarousel       = "mvc-bs-carousel"
	ViewCarouselItem   = "mvc-bs-carouselitem"
	ViewCheckboxGroup  = "mvc-bs-checkboxgroup"
	ViewCodeBlock      = "mvc-bs-codeblock"
	ViewContainer      = "mvc-bs-container"
	ViewDefinitionList = "mvc-bs-deflist"
	ViewForm           = "mvc-bs-form"
	ViewGrid           = "mvc-bs-grid"
	ViewHeading        = "mvc-bs-heading"
	ViewIcon           = "mvc-bs-icon"
	ViewImage          = "mvc-bs-img"
	ViewInput          = "mvc-bs-input"
	ViewInputGroup     = "mvc-bs-inputgroup"
	ViewLink           = "mvc-bs-link"
	ViewList           = "mvc-bs-list"
	ViewListGroup      = "mvc-bs-listgroup"
	ViewMarkdown       = "mvc-bs-markdown"
	ViewModal          = "mvc-bs-modal"
	ViewMedia          = "mvc-bs-media"
	ViewNavBar         = "mvc-bs-navbar"
	ViewNavDropdown    = "mvc-bs-navdropdown"
	ViewNavItem        = "mvc-bs-navitem"
	ViewOffcanvas      = "mvc-bs-offcanvas"
	ViewPagination     = "mvc-bs-pagination"
	ViewPaginationItem = "mvc-bs-paginationitem"
	ViewProgress       = "mvc-bs-progress"
	ViewRadioGroup     = "mvc-bs-radiogroup"
	ViewRule           = "mvc-bs-rule"
	ViewSelect         = "mvc-bs-select"
	ViewTable          = "mvc-bs-table"
	ViewTableRow       = "mvc-bs-tablerow"
	ViewText           = "mvc-bs-text"
	ViewToast          = "mvc-bs-toast"
	ViewToastGroup     = "mvc-bs-toastgroup"
)

const (
	// The prefix class for outline buttons
	viewOutlineButtonClassPrefix = "btn-outline"
)

// Set the view element's child view
func setView(self mvc.View, child mvc.View) {
	switch list := self.(type) {
	case *button:
		list.View = child
	case *buttongroup:
		list.View = child
	case *buttontoolbar:
		list.View = child
	case *carousel:
		list.View = child
	case *carouselitem:
		list.View = child
	case *list:
		list.View = child
	case *deflist:
		list.View = child
	case *badge:
		list.View = child
	case *img:
		list.View = child
	case *input:
		list.View = child
	default:
		panic(fmt.Sprintf("setView: unsupported view type %T", self))
	}
}
