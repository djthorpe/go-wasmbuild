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
	mvc.View
}

type accordionitem struct {
	mvc.View
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
	return mvc.NewView(new(accordion), ViewAccordion, "DIV", mvc.WithClass("accordion"), mvc.WithID(name), args).(*accordion)
}

func FlushAccordion(name string, args ...any) *accordion {
	return mvc.NewView(new(accordion), ViewAccordion, "DIV", mvc.WithClass("accordion", "accordion-flush"), mvc.WithID(name), args).(*accordion)
}

func AccordionItem(args ...any) *accordionitem {
	return mvc.NewViewExEx(new(accordionitem), ViewAccordionItem, templateAccordionItem, args).(*accordionitem)
}

func newAccordionFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(accordion), element)
}
func newAccordionItemFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(accordionitem), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (accordion *accordion) SetView(view mvc.View) {
	accordion.View = view
}

func (accordionitem *accordionitem) SetView(view mvc.View) {
	accordionitem.View = view
}

func (accordion *accordion) Content(args ...any) mvc.View {
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
	return accordion.View.Content(args...)
}
