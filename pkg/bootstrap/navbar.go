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

type navbar struct {
	mvc.ViewWithLabel
}

type navitem struct {
	mvc.View
}

var _ mvc.View = (*navbar)(nil)
var _ mvc.ViewWithLabel = (*navbar)(nil)
var _ mvc.View = (*navitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavBar  = "mvc-bs-navbar"
	ViewNavItem = "mvc-bs-navitem"
)

const (
	templateNavBar = `
		<nav class="navbar navbar-expand bg-primary">
			<div class="container-fluid">
				<slot name="label"><!-- Label --></slot>
				<slot name="toggle-button"><!-- Toggle Button --></slot>
				<div class="collapse navbar-collapse">
					<slot><!-- Body --></slot>
				</div>
			</div>
		</nav>
	`
	templateNavItem = `
		<li class="nav-item"><slot></slot></li>
	`
)

func init() {
	mvc.RegisterView(ViewNavBar, newNavBarFromElement)
	mvc.RegisterView(ViewNavItem, newNavItemFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(args ...any) *navbar {
	// Return the navbar
	return mvc.NewViewExEx(
		new(navbar), ViewNavBar, templateNavBar, args,
	).(*navbar)
}

func NavItem(href string, args ...any) *navitem {
	// Return the navitem
	return mvc.NewViewExEx(
		new(navitem), ViewNavItem, templateNavItem, args,
	).(*navitem)
}

func newNavBarFromElement(element Element) mvc.View {
	if element.TagName() != "NAV" {
		return nil
	}
	return mvc.NewViewWithElement(new(navbar), element)
}

func newNavItemFromElement(element Element) mvc.View {
	if element.TagName() != "LI" {
		return nil
	}
	return mvc.NewViewWithElement(new(navitem), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (navbar *navbar) SetView(view mvc.View) {
	navbar.ViewWithLabel = view.(mvc.ViewWithLabel)
}

func (navitem *navitem) SetView(view mvc.View) {
	navitem.View = view
}

func (navbar *navbar) Caption(children ...any) mvc.ViewWithCaption {
	// DEPRECATED: use Label instead
	return navbar.Label(children...)
}

func (navbar *navbar) Label(children ...any) mvc.ViewWithLabel {
	return navbar.ReplaceSlot("label", mvc.HTML("A", mvc.WithAttr("href", "#"), mvc.WithClass("navbar-brand"), children)).(mvc.ViewWithLabel)
}

func (navbar *navbar) Content(children ...any) mvc.View {
	items := []any{}
	for _, child := range children {
		switch child := child.(type) {
		case *navitem:
			items = append(items, child)
		default:
			panic(fmt.Sprintf("Content: invalid child type: %T", child))
		}
	}
	return navbar.ReplaceSlot("", mvc.HTML("ul", mvc.WithClass("navbar-nav me-auto"), items))
}

func (navitem *navitem) Content(children ...any) mvc.View {
	// TODO: Set href attribute on <a> tag
	return navitem.ReplaceSlot("", mvc.HTML("a", mvc.WithAttr("href", "#"), mvc.WithClass("nav-link"), children))
}
