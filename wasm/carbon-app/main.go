package main

import (
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, ` +
		`sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, ` +
		`quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`
)

// Application displays examples of IBM Carbon Design System components
func main() {
	// Shell — position:fixed header and side nav, attached to body independently
	mvc.New(shell())

	// Content — router switches pages based on URL hash
	mvc.New(
		mvc.WithClass("cds--content"),
		router(),
	).Run()
}

func shell() mvc.View {
	return cds.Section(
		cds.Header(
			cds.WithTheme(cds.ThemeG90),
			cds.HeaderNavItem("#headings", "Headings"),
			cds.HeaderNavItem("#text", "Text"),
			cds.HeaderNavItem("#links", "Links"),
			cds.HeaderNavItem("#code", "Code"),
			cds.HeaderNavItem("#icons", "Icons"),
			cds.HeaderNavItem("#buttons", "Buttons"),
			cds.HeaderNavItem("#grid", "Grid"),
		).Label("/", "IBM", "Carbon Examples"),
		sideNav,
	)
}

// sideNav and its items are package-level so router() can reference them.
var (
	navHeadings      = cds.SideNavItem("#headings", "Headings")
	navText          = cds.SideNavItem("#text", "Text")
	navLinks         = cds.SideNavItem("#links", "Links")
	navCode          = cds.SideNavItem("#code", "Code")
	navIcons         = cds.SideNavItem("#icons", "Icons")
	navForms         = cds.SideNavItem("#forms", "Form basics")
	navInput         = cds.SideNavItem("#input", "Input")
	navSelect        = cds.SideNavItem("#select", "Select")
	navDropdown      = cds.SideNavItem("#dropdown", "Dropdown")
	navButtons       = cds.SideNavItem("#buttons", "Buttons")
	navTags          = cds.SideNavItem("#tags", "Tags")
	navNotifications = cds.SideNavItem("#notifications", "Notifications")
	navAccordion     = cds.SideNavItem("#accordion", "Accordion")
	navTabs          = cds.SideNavItem("#tabs", "Tabs")
	navGrid          = cds.SideNavLink("#grid", "Grid")
	sideNav          = cds.SideNav(
		cds.WithTheme(cds.ThemeG90),
		cds.SideNavSection("Typography", navText, navHeadings, navLinks, navCode, navIcons),
		cds.SideNavSection("Form", navForms, navInput, navSelect, navDropdown),
		cds.SideNavSection("Components", navButtons, navTags, navNotifications, navAccordion, navTabs),
		navGrid,
	)
)

func router() mvc.View {
	return mvc.Router().
		Selectable(sideNav).
		Page("#headings", HeadingExamples(), navHeadings).
		Page("#text", TextExamples(), navText).
		Page("#links", LinkExamples(), navLinks).
		Page("#code", CodeExamples(), navCode).
		Page("#icons", IconExamples(), navIcons).
		Page("#forms", FormsExamples(), navForms).
		Page("#input", InputExamples(), navInput).
		Page("#select", SelectExamples(), navSelect).
		Page("#dropdown", DropdownExamples(), navDropdown).
		Page("#buttons", ButtonExamples(), navButtons).
		Page("#tags", TagExamples(), navTags).
		Page("#notifications", NotificationExamples(), navNotifications).
		Page("#accordion", AccordionExamples(), navAccordion).
		Page("#tabs", TabExamples(), navTabs).
		Page("#grid", GridExamples(), navGrid)
}
