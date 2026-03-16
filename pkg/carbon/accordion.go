package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type accordion struct {
	mvc.View
}

type accordionItem struct {
	mvc.View
}

var _ mvc.View = (*accordion)(nil)
var _ mvc.View = (*accordionItem)(nil)

// AccordionSize controls the row height of accordion items.
type AccordionSize string

// AccordionAlign controls where the expand/collapse chevron is placed.
type AccordionAlign string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	AccordionSM AccordionSize = "sm"
	AccordionMD AccordionSize = "md" // default
	AccordionLG AccordionSize = "lg"
)

const (
	// AccordionAlignEnd places the chevron on the right (default).
	AccordionAlignEnd AccordionAlign = "END"
	// AccordionAlignStart places the chevron on the left, flush with content.
	AccordionAlignStart AccordionAlign = "start"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewAccordion, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(accordion), element, func(self, child mvc.View) {
			self.(*accordion).View = child
		})
	})
	mvc.RegisterView(ViewAccordionItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(accordionItem), element, func(self, child mvc.View) {
			self.(*accordionItem).View = child
		})
	})
}

// Accordion returns a <cds-accordion> container.
// Pass AccordionItem children and any options as args.
//
//	cds.Accordion(
//	    cds.AccordionItem("Section 1", "Content here"),
//	    cds.AccordionItem("Section 2", "More content", cds.WithAccordionOpen()),
//	)
func Accordion(args ...any) *accordion {
	return mvc.NewView(new(accordion), ViewAccordion, "cds-accordion", func(self, child mvc.View) {
		self.(*accordion).View = child
	}, args).(*accordion)
}

// AccordionItem returns a <cds-accordion-item> with the given header title.
// Body content and options are passed as further args.
//
//	cds.AccordionItem("Getting started", cds.Para("Read the docs first."), cds.WithAccordionOpen())
func AccordionItem(title string, args ...any) *accordionItem {
	args = append([]any{mvc.WithAttr("title", title)}, args...)
	return mvc.NewView(new(accordionItem), ViewAccordionItem, "cds-accordion-item", func(self, child mvc.View) {
		self.(*accordionItem).View = child
	}, args).(*accordionItem)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithAccordionSize sets the row height: sm (32px), md (40px, default), or lg (48px).
func WithAccordionSize(s AccordionSize) mvc.Opt {
	return mvc.WithAttr("size", string(s))
}

// WithAccordionAlign sets chevron placement: end (right, default) or start (left).
func WithAccordionAlign(a AccordionAlign) mvc.Opt {
	return mvc.WithAttr("alignment", string(a))
}

// WithAccordionOpen expands an AccordionItem by default.
func WithAccordionOpen() mvc.Opt {
	return mvc.WithAttr("open", "")
}
