package bootstrap

import (
	// Packages
	"fmt"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type accordion struct {
	BootstrapView
}

type accordionitem struct {
	BootstrapView
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewAccordion     = "mvc-bs-accordion"
	ViewAccordionItem = "mvc-bs-accordionitem"
)

const (
	templateAccordionItem = `
		<div class="accordion-item">
			<h2 class="accordion-header">
				<button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#replace" data-slot="header"></button>
			</h2>
			<div class="accordion-collapse collapse">
				<div class="accordion-body" data-slot=""></div>
			</div>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewAccordion, newAccordionFromElement)
	mvc.RegisterView(ViewAccordionItem, newAccordionItemFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Accordion(name string, args ...any) *accordion {
	a := new(accordion)
	a.BootstrapView.View = mvc.NewView(a, ViewAccordion, "DIV", mvc.WithClass("accordion"), mvc.WithID(name), args)
	return a
}

func FlushAccordion(name string, args ...any) *accordion {
	a := new(accordion)
	a.BootstrapView.View = mvc.NewView(a, ViewAccordion, "DIV", mvc.WithClass("accordion", "accordion-flush"), mvc.WithID(name), args)
	return a
}

func AccordionItem(args ...any) *accordionitem {
	a := new(accordionitem)
	a.BootstrapView.View = mvc.NewViewExEx(a, ViewAccordionItem, templateAccordionItem, args)
	return a
}

func newAccordionFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	a := new(accordion)
	a.BootstrapView.View = mvc.NewViewWithElement(a, element)
	return a
}
func newAccordionItemFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	a := new(accordionitem)
	a.BootstrapView.View = mvc.NewViewWithElement(a, element)
	return a
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (accordion *accordion) Self() mvc.View {
	return accordion
}

func (accordionitem *accordionitem) Self() mvc.View {
	return accordionitem
}

func (accordion *accordion) Content(args ...any) *accordion {
	for i, arg := range args {
		switch arg := arg.(type) {
		case *accordionitem:
			item := fmt.Sprint(accordion.ID(), "-", i)
			show := false

			// Show the first item in the accordion
			if i == 0 {
				show = true
			}

			// Set header attributes
			header := arg.Slot("header")
			if header != nil {
				header.SetAttribute("data-bs-target", "#"+item)
				header.SetAttribute("aria-controls", item)
				header.SetAttribute("aria-expanded", fmt.Sprint(show))
				if !show {
					header.ClassList().Add("collapsed")
				}
			}

			// Set body attributes
			body := arg.Slot("").ParentElement()
			if body != nil {
				body.SetID(item)
				body.SetAttribute("data-bs-parent", "#"+accordion.ID())
				if show {
					body.ClassList().Add("show")
				}
			}
		default:
			panic(ErrInternalAppError.Withf("Content[accordionitem] unexpected argument '%T'", arg))
		}
	}
	accordion.ReplaceSlot("body", wrapChildren(args...))
	return accordion
}
