package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type navbar struct {
	mvc.View
}

type navitem struct {
	mvc.View
}

type navdropdown struct {
	mvc.View
}

var _ mvc.View = (*navbar)(nil)
var _ mvc.View = (*navitem)(nil)
var _ mvc.View = (*navdropdown)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewNavBar      = "mvc-bs-navbar"
	ViewNavItem     = "mvc-bs-navitem"
	ViewNavDropdown = "mvc-bs-navdropdown"
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
	templateNavDivider = `
		<li><hr class="dropdown-divider"></li>
	`
	templateNavDropdown = `
		<li class="nav-item dropdown">
			<a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false"><slot name="label"><!-- Label --></slot></a>
			<slot></slot>
		</li>
	`
	templateToggleButton = `
		<button class="navbar-toggler" type="button" data-bs-toggle="collapse" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
	    </button>
	`
	dataAttrNavHref = mvc.DataComponentAttrKey + "-href"
)

func init() {
	mvc.RegisterView(ViewNavBar, newNavBarFromElement)
	mvc.RegisterView(ViewNavItem, newNavItemFromElement)
	mvc.RegisterView(ViewNavDropdown, newNavDropdownFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NavBar(id string, args ...any) *navbar {
	// Create the navbar
	navbar := mvc.NewViewExEx(new(navbar), ViewNavBar, templateNavBar, args).(*navbar)

	// Replace the toggle button slot
	navbar.ReplaceSlot("toggle-button", mvc.HTML(templateToggleButton, mvc.WithAttr("data-bs-target", "#"+id), mvc.WithAttr("aria-controls", id)))

	// Set the target for the toggle
	if collapse := navbar.Root().GetElementsByClassName("collapse"); len(collapse) > 0 {
		collapse[0].SetAttribute("id", id)
	}

	// Return the navbar
	return navbar
}

func NavItem(href string, args ...any) *navitem {
	// Ensure href is not empty
	if href == "" {
		href = "#"
	}

	// Return the navitem
	item := mvc.NewViewExEx(
		new(navitem), ViewNavItem, templateNavItem, mvc.WithAttr(dataAttrNavHref, href), args,
	).(*navitem)
	return item
}

func NavDivider() *navitem {
	return mvc.NewViewExEx(new(navitem), ViewNavItem, templateNavDivider).(*navitem)
}

func NavDropdown(args ...any) *navdropdown {
	// Return the navdropdown
	return mvc.NewViewExEx(
		new(navdropdown), ViewNavDropdown, templateNavDropdown, args,
	).(*navdropdown)
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

func newNavDropdownFromElement(element Element) mvc.View {
	if element.TagName() != "LI" {
		return nil
	}
	return mvc.NewViewWithElement(new(navdropdown), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (navbar *navbar) SetView(view mvc.View) {
	navbar.View = view
}

func (navitem *navitem) SetView(view mvc.View) {
	navitem.View = view
}

func (navdropdown *navdropdown) SetView(view mvc.View) {
	navdropdown.View = view
}

func (navbar *navbar) Label(children ...any) mvc.View {
	return navbar.ReplaceSlot("label", mvc.HTML("A", mvc.WithAttr("href", "#"), mvc.WithClass("navbar-brand"), children))
}

func (navbar *navbar) Apply(opts ...mvc.Opt) mvc.View {
	// Apply the options first
	navbar.View.Apply(opts...)

	// Determine if navbar is positioned at bottom
	isBottomFixed := navbar.Root().ClassList().Contains("fixed-bottom") ||
		navbar.Root().ClassList().Contains("sticky-bottom")

	// Update all nav-item dropdown/dropup classes
	items := navbar.Root().GetElementsByClassName("nav-item")
	for _, item := range items {
		classList := item.ClassList()
		if isBottomFixed {
			if classList.Contains("dropdown") {
				classList.Remove("dropdown")
				classList.Add("dropup")
			}
		} else {
			if classList.Contains("dropup") {
				classList.Remove("dropup")
				classList.Add("dropdown")
			}
		}
	}

	return navbar
}

func (navbar *navbar) Content(children ...any) mvc.View {
	items := []any{}
	for _, child := range children {
		switch child := child.(type) {
		case *navitem:
			items = append(items, child)
		case *navdropdown:
			items = append(items, child)
		default:
			panic(fmt.Sprintf("Content[navbar]: invalid child type: %T", child))
		}
	}
	return navbar.ReplaceSlot("", mvc.HTML("ul", mvc.WithClass("navbar-nav"), items))
}

func (navdropdown *navdropdown) Content(children ...any) mvc.View {
	items := []any{}
	for _, child := range children {
		switch child := child.(type) {
		case *navitem:
			items = append(items, child)
		default:
			panic(fmt.Sprintf("Content[navdropdown]: invalid child type: %T", child))
		}
	}
	return navdropdown.ReplaceSlot("", mvc.HTML("ul", mvc.WithClass("dropdown-menu"), items))
}

func (navitem *navitem) Content(children ...any) mvc.View {
	href := navitem.Root().GetAttribute(dataAttrNavHref)
	if href == "" {
		href = "#"
	}
	return navitem.ReplaceSlot("", mvc.HTML("a", mvc.WithAttr("href", href), mvc.WithClass("nav-link", "text-nowrap"), children))
}
