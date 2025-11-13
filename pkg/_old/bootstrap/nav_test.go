package bootstrap

import (
	"strings"
	"testing"

	dom "github.com/djthorpe/go-wasmbuild"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

func TestNav_Basic(t *testing.T) {
	nav := Nav()

	// Check element type
	if tagName := nav.Element().TagName(); tagName != "NAV" {
		t.Errorf("Expected tag name 'NAV', got '%s'", tagName)
	}

	// Check class
	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav") {
		t.Errorf("Expected class to contain 'nav', got '%s'", class)
	}

	// Check no children initially
	children := nav.Element().ChildNodes()
	if len(children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(children))
	}
}

func TestNav_WithTabs(t *testing.T) {
	nav := Nav(WithTabs())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-tabs") {
		t.Errorf("Expected class to contain 'nav-tabs', got '%s'", class)
	}
}

func TestNav_WithPills(t *testing.T) {
	nav := Nav(WithPills())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-pills") {
		t.Errorf("Expected class to contain 'nav-pills', got '%s'", class)
	}
}

func TestNav_WithUnderline(t *testing.T) {
	nav := Nav(WithUnderline())

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-underline") {
		t.Errorf("Expected class to contain 'nav-underline', got '%s'", class)
	}
}

func TestNav_WithClass(t *testing.T) {
	nav := Nav(WithClass("justify-content-center"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "justify-content-center") {
		t.Errorf("Expected class to contain 'justify-content-center', got '%s'", class)
	}
}

func TestNav_Vertical(t *testing.T) {
	nav := Nav(WithClass("flex-column"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "flex-column") {
		t.Errorf("Expected class to contain 'flex-column', got '%s'", class)
	}
}

func TestNav_Fill(t *testing.T) {
	nav := Nav(WithPills(), WithClass("nav-fill"))

	class := nav.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-fill") {
		t.Errorf("Expected class to contain 'nav-fill', got '%s'", class)
	}
	if !strings.Contains(class, "nav-pills") {
		t.Errorf("Expected class to contain 'nav-pills', got '%s'", class)
	}
}

func TestNav_Component(t *testing.T) {
	nav := Nav()

	// Check component name
	if name := nav.Name(); name != "nav" {
		t.Errorf("Expected component name 'nav', got '%s'", name)
	}

	// Check that Element() returns the same as root
	if elem := nav.Element(); elem != nav.root {
		t.Error("Element() should return the root element")
	}
}

func TestNav_WithAttribute(t *testing.T) {
	nav := Nav(WithAttribute("data-bs-theme", "dark"))

	// Check custom attribute
	theme := nav.Element().GetAttribute("data-bs-theme")
	if theme != "dark" {
		t.Errorf("Expected data-bs-theme 'dark', got '%s'", theme)
	}
}

func TestNav_WithMultipleAttributes(t *testing.T) {
	nav := Nav(
		WithAttribute("data-bs-theme", "dark"),
		WithAttribute("role", "navigation"),
		WithAttribute("id", "main-nav"),
	)

	theme := nav.Element().GetAttribute("data-bs-theme")
	if theme != "dark" {
		t.Errorf("Expected data-bs-theme 'dark', got '%s'", theme)
	}

	role := nav.Element().GetAttribute("role")
	if role != "navigation" {
		t.Errorf("Expected role 'navigation', got '%s'", role)
	}

	id := nav.Element().GetAttribute("id")
	if id != "main-nav" {
		t.Errorf("Expected id 'main-nav', got '%s'", id)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVBAR TESTS

func TestNavBar_Basic(t *testing.T) {
	navbar := NavBar()

	// Check element type
	if tagName := navbar.Element().TagName(); tagName != "NAV" {
		t.Errorf("Expected tag name 'NAV', got '%s'", tagName)
	}

	// Check class
	class := navbar.Element().GetAttribute("class")
	if !strings.Contains(class, "navbar") {
		t.Errorf("Expected class to contain 'navbar', got '%s'", class)
	}

	// Check container-fluid exists
	children := navbar.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (container-fluid), got %d", len(children))
	}

	container := children[0]
	if container.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element node")
	}

	containerClass := container.(Element).GetAttribute("class")
	if !strings.Contains(containerClass, "container-fluid") {
		t.Errorf("Expected container class to contain 'container-fluid', got '%s'", containerClass)
	}
}

func TestNavBar_WithExpand(t *testing.T) {
	navbar := NavBar(WithResponsive(BreakpointLarge))

	class := navbar.Element().GetAttribute("class")
	if !strings.Contains(class, "navbar-expand-lg") {
		t.Errorf("Expected class to contain 'navbar-expand-lg', got '%s'", class)
	}
}

func TestNavBar_Header(t *testing.T) {
	navbar := NavBar().Header("My Brand")

	// Check that brand was added to container
	container := navbar.Element().ChildNodes()[0].(Element)
	children := container.ChildNodes()

	// Should have 3 children: brand, toggler, collapse
	if len(children) != 3 {
		t.Errorf("Expected 3 elements (brand, toggler, collapse), got %d", len(children))
	}

	// First child should be brand (inserted at top)
	brand := children[0].(Element)
	if brand.TagName() != "A" {
		t.Errorf("Expected brand tag 'A', got '%s'", brand.TagName())
	}

	brandClass := brand.GetAttribute("class")
	if !strings.Contains(brandClass, "navbar-brand") {
		t.Errorf("Expected brand class to contain 'navbar-brand', got '%s'", brandClass)
	}

	// Check text content
	if len(brand.ChildNodes()) != 1 {
		t.Error("Expected brand to have 1 text node")
	}
}

func TestNavBar_Structure(t *testing.T) {
	navbar := NavBar()

	// Check that toggler and collapse were automatically created
	container := navbar.Element().ChildNodes()[0].(Element)
	children := container.ChildNodes()

	if len(children) != 2 {
		t.Errorf("Expected 2 children (toggler + collapse), got %d", len(children))
	}

	// Check toggler
	toggler := children[0].(Element)
	if toggler.TagName() != "BUTTON" {
		t.Errorf("Expected toggler tag 'BUTTON', got '%s'", toggler.TagName())
	}

	togglerClass := toggler.GetAttribute("class")
	if !strings.Contains(togglerClass, "navbar-toggler") {
		t.Errorf("Expected toggler class to contain 'navbar-toggler', got '%s'", togglerClass)
	}

	// Check toggler data attributes
	target := toggler.GetAttribute("data-bs-target")
	if target != "#navbarNav" {
		t.Errorf("Expected data-bs-target '#navbarNav', got '%s'", target)
	}

	// Check toggler icon exists
	icon := toggler.ChildNodes()[0].(Element)
	iconClass := icon.GetAttribute("class")
	if !strings.Contains(iconClass, "navbar-toggler-icon") {
		t.Errorf("Expected icon class to contain 'navbar-toggler-icon', got '%s'", iconClass)
	}

	// Check collapse
	collapse := children[1].(Element)
	if collapse.TagName() != "DIV" {
		t.Errorf("Expected collapse tag 'DIV', got '%s'", collapse.TagName())
	}

	collapseClass := collapse.GetAttribute("class")
	if !strings.Contains(collapseClass, "collapse") {
		t.Errorf("Expected collapse class to contain 'collapse', got '%s'", collapseClass)
	}
	if !strings.Contains(collapseClass, "navbar-collapse") {
		t.Errorf("Expected collapse class to contain 'navbar-collapse', got '%s'", collapseClass)
	}

	collapseId := collapse.GetAttribute("id")
	if collapseId != "navbarNav" {
		t.Errorf("Expected collapse id 'navbarNav', got '%s'", collapseId)
	}

	// The collapse element is where navbar content is appended
	// (verified by other tests like TestNavBar_Append)
}

func TestNavBar_Append(t *testing.T) {
	homeItem := NavItem("#", "Home")
	homeItem.Active(true)

	navbar := NavBar().Append(
		homeItem,
		NavItem("#", "About"),
	)

	// Check that items were added to the collapse body
	container := navbar.Element().ChildNodes()[0].(Element)
	collapse := container.ChildNodes()[1].(Element) // collapse is second child after toggler
	navList := collapse.ChildNodes()[0].(Element)   // navbar-nav UL inside collapse
	children := navList.ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 nav items, got %d", len(children))
	}

	// Check first item is active
	firstItem := children[0].(Element)
	if !strings.Contains(firstItem.GetAttribute("class"), "active") {
		t.Error("Expected first item to be active")
	}
}

func TestNavBar_Insert(t *testing.T) {
	navbar := NavBar().Append(NavItem("#", "First"))

	// Insert at the top
	newFirst := NavItem("#", "New First")
	newFirst.Active(true)
	navbar.Insert(newFirst)

	// Check that new item is first
	container := navbar.Element().ChildNodes()[0].(Element)
	collapse := container.ChildNodes()[1].(Element)
	navList := collapse.ChildNodes()[0].(Element) // navbar-nav UL inside collapse
	children := navList.ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 nav items, got %d", len(children))
	}

	firstItem := children[0].(Element)
	if !strings.Contains(firstItem.GetAttribute("class"), "active") {
		t.Error("Expected first item to be active (newly inserted)")
	}
}

func TestNavBar_AppendWithSpacer(t *testing.T) {
	navbar := NavBar().Append(
		NavItem("#", "Home"),
		NavSpacer(),
		NavItem("#", "Login"),
	)

	// Check that all items were added
	container := navbar.Element().ChildNodes()[0].(Element)
	collapse := container.ChildNodes()[1].(Element)
	navList := collapse.ChildNodes()[0].(Element) // navbar-nav UL inside collapse
	children := navList.ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children (2 items + spacer), got %d", len(children))
	}

	// Check middle item is spacer
	spacer := children[1].(Element)
	if spacer.GetAttribute("data-component") != string(NavSpacerComponent) {
		t.Error("Expected middle child to be NavSpacer")
	}
}

func TestNavBar_AppendWithDivider(t *testing.T) {
	navbar := NavBar().Append(
		NavItem("#", "Home"),
		NavDivider(true),
		NavItem("#", "About"),
	)

	// Check that all items were added
	container := navbar.Element().ChildNodes()[0].(Element)
	collapse := container.ChildNodes()[1].(Element)
	navList := collapse.ChildNodes()[0].(Element) // navbar-nav UL inside collapse
	children := navList.ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children (2 items + divider), got %d", len(children))
	}

	// Check middle item is divider
	divider := children[1].(Element)
	if divider.GetAttribute("data-component") != string(NavDividerComponent) {
		t.Error("Expected middle child to be NavDivider")
	}
}

func TestNavBar_Complete(t *testing.T) {
	// Build a complete navbar
	homeItem := NavItem("#", "Home")
	homeItem.Active(true)

	navbar := NavBar(WithResponsive(BreakpointLarge), WithClass("bg-body-tertiary")).
		Header("Navbar").
		Append(
			homeItem,
			NavItem("#", "About"),
		)

	// Verify structure
	root := navbar.Element()
	if root.TagName() != "NAV" {
		t.Errorf("Expected root tag 'NAV', got '%s'", root.TagName())
	}

	// Check classes
	rootClass := root.GetAttribute("class")
	if !strings.Contains(rootClass, "navbar") {
		t.Error("Expected navbar class")
	}
	if !strings.Contains(rootClass, "navbar-expand-lg") {
		t.Error("Expected navbar-expand-lg class")
	}
	if !strings.Contains(rootClass, "bg-body-tertiary") {
		t.Error("Expected bg-body-tertiary class")
	}

	// Check container has brand, toggler, and collapse
	container := root.ChildNodes()[0].(Element)
	containerChildren := container.ChildNodes()
	if len(containerChildren) != 3 {
		t.Errorf("Expected 3 children (brand, toggler, collapse), got %d", len(containerChildren))
	}

	// Check collapse has nav items
	collapse := containerChildren[2].(Element)
	navList := collapse.ChildNodes()[0].(Element) // navbar-nav UL inside collapse
	navItems := navList.ChildNodes()
	if len(navItems) != 2 {
		t.Errorf("Expected 2 nav items in collapse, got %d", len(navItems))
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVITEM TESTS

func TestNavItem_Basic(t *testing.T) {
	navItem := NavItem("/home", "Home")

	// Check element type
	if tagName := navItem.Element().TagName(); tagName != "A" {
		t.Errorf("Expected tag name 'A', got '%s'", tagName)
	}

	// Check class
	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-link") {
		t.Errorf("Expected class to contain 'nav-link', got '%s'", class)
	}

	// Check href
	href := navItem.Element().GetAttribute("href")
	if href != "/home" {
		t.Errorf("Expected href '/home', got '%s'", href)
	}

	// Check text content
	children := navItem.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (text node), got %d", len(children))
	}
}

func TestNavItem_Active(t *testing.T) {
	navItem := NavItem("#", "Home")
	navItem.Active(true)

	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	ariaCurrent := navItem.Element().GetAttribute("aria-current")
	if ariaCurrent != "page" {
		t.Errorf("Expected aria-current 'page', got '%s'", ariaCurrent)
	}
}

func TestNavItem_Disabled(t *testing.T) {
	navItem := NavItem("#", "Disabled")
	navItem.Disabled(true)

	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "disabled") {
		t.Errorf("Expected class to contain 'disabled', got '%s'", class)
	}

	ariaDisabled := navItem.Element().GetAttribute("aria-disabled")
	if ariaDisabled != "true" {
		t.Errorf("Expected aria-disabled 'true', got '%s'", ariaDisabled)
	}
}

func TestNavItem_WithContent(t *testing.T) {
	icon := Icon("home")
	navItem := NavItem("/dashboard", icon, " Dashboard")
	navItem.Active(true)

	// Check active state
	class := navItem.Element().GetAttribute("class")
	if !strings.Contains(class, "active") {
		t.Errorf("Expected class to contain 'active', got '%s'", class)
	}

	// Check children (icon + text)
	children := navItem.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (icon + text), got %d", len(children))
	}

	// Check that first child is the icon element
	if children[0].NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element (icon)")
	}
}

func TestNavItem_Component(t *testing.T) {
	navItem := NavItem("#", "Test")
	if navItem.component.name != NavItemComponent {
		t.Errorf("Expected component type %s, got %s", NavItemComponent, navItem.component.name)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVSPACER TESTS

func TestNavSpacer_Basic(t *testing.T) {
	spacer := NavSpacer()

	// Check element type
	if tagName := spacer.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check default class
	class := spacer.Element().GetAttribute("class")
	if !strings.Contains(class, "me-auto") {
		t.Errorf("Expected class to contain 'me-auto', got '%s'", class)
	}
}

func TestNavSpacer_WithCustomClass(t *testing.T) {
	spacer := NavSpacer(WithClass("ms-auto"))

	class := spacer.Element().GetAttribute("class")
	if !strings.Contains(class, "ms-auto") {
		t.Errorf("Expected class to contain 'ms-auto', got '%s'", class)
	}
	// Should still have the default class unless explicitly overridden
	if !strings.Contains(class, "me-auto") {
		t.Errorf("Expected class to contain 'me-auto', got '%s'", class)
	}
}

func TestNavSpacer_ComponentInterface(t *testing.T) {
	spacer := NavSpacer()
	if spacer.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDIVIDER TESTS

func TestNavDivider_Vertical(t *testing.T) {
	divider := NavDivider(true)

	// Check element type for vertical divider
	if tagName := divider.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check vertical divider class
	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "vr") {
		t.Errorf("Expected class to contain 'vr', got '%s'", class)
	}
}

func TestNavDivider_Horizontal(t *testing.T) {
	divider := NavDivider(false)

	// Check element type for horizontal divider
	if tagName := divider.Element().TagName(); tagName != "HR" {
		t.Errorf("Expected tag name 'HR', got '%s'", tagName)
	}

	// Check horizontal divider class
	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "my-2") {
		t.Errorf("Expected class to contain 'my-2', got '%s'", class)
	}
}

func TestNavDivider_WithCustomClass(t *testing.T) {
	divider := NavDivider(true, WithClass("mx-3"))

	class := divider.Element().GetAttribute("class")
	if !strings.Contains(class, "mx-3") {
		t.Errorf("Expected class to contain 'mx-3', got '%s'", class)
	}
	// Should still have the default class
	if !strings.Contains(class, "vr") {
		t.Errorf("Expected class to contain 'vr', got '%s'", class)
	}
}

func TestNavDivider_ComponentInterface(t *testing.T) {
	divider := NavDivider(true)
	if divider.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}

	divider2 := NavDivider(false)
	if divider2.Element() == nil {
		t.Error("Expected Element() to return non-nil")
	}
}

///////////////////////////////////////////////////////////////////////////////
// INTEGRATION TESTS

func TestNav_WithSpacerAndDivider(t *testing.T) {
	// Test that NavSpacer and NavDivider work correctly within a Nav component
	homeItem := NavItem("#", "Home")
	homeItem.Active(true)

	nav := Nav().
		Append(homeItem).
		Append(NavDivider(true)).
		Append(NavItem("#", "About")).
		Append(NavSpacer()).
		Append(NavItem("#", "Login"))

	// Check that nav has the expected number of children
	children := nav.Element().ChildNodes()
	expectedChildren := 5 // NavItem + NavDivider + NavItem + NavSpacer + NavItem
	if len(children) != expectedChildren {
		t.Errorf("Expected %d children, got %d", expectedChildren, len(children))
	}

	// Check specific child types and classes
	// First child should be NavItem (anchor)
	if children[0].NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected first child to be an element")
	}

	// Second child should be NavDivider (div with vr class)
	dividerEl := children[1]
	if dividerEl.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected second child to be an element")
	}
	if tagName := dividerEl.(dom.Element).TagName(); tagName != "DIV" {
		t.Errorf("Expected divider tag name 'DIV', got '%s'", tagName)
	}

	// Fourth child should be NavSpacer (div with me-auto class)
	spacerEl := children[3]
	if spacerEl.NodeType() != dom.ELEMENT_NODE {
		t.Error("Expected fourth child to be an element")
	}
	if tagName := spacerEl.(dom.Element).TagName(); tagName != "DIV" {
		t.Errorf("Expected spacer tag name 'DIV', got '%s'", tagName)
	}
}

///////////////////////////////////////////////////////////////////////////////
// NAVDROPDOWN TESTS

func TestNavDropdown_Basic(t *testing.T) {
	dropdown := NavDropdown("More",
		NavItem("#", "Action"),
		NavItem("#", "Another action"),
	)

	// Check element type
	if tagName := dropdown.Element().TagName(); tagName != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tagName)
	}

	// Check dropdown container class
	class := dropdown.Element().GetAttribute("class")
	if !strings.Contains(class, "nav-item") || !strings.Contains(class, "dropdown") {
		t.Errorf("Expected class to contain 'nav-item dropdown', got '%s'", class)
	}

	// Check children (toggle + menu)
	children := dropdown.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (toggle + menu), got %d", len(children))
	}

	// Check dropdown toggle
	toggle := children[0].(dom.Element)
	if tagName := toggle.TagName(); tagName != "A" {
		t.Errorf("Expected toggle tag name 'A', got '%s'", tagName)
	}
	toggleClass := toggle.GetAttribute("class")
	if !strings.Contains(toggleClass, "nav-link") || !strings.Contains(toggleClass, "dropdown-toggle") {
		t.Errorf("Expected toggle class to contain 'nav-link dropdown-toggle', got '%s'", toggleClass)
	}

	// Check dropdown menu
	menu := children[1].(dom.Element)
	if tagName := menu.TagName(); tagName != "UL" {
		t.Errorf("Expected menu tag name 'UL', got '%s'", tagName)
	}
	menuClass := menu.GetAttribute("class")
	if !strings.Contains(menuClass, "dropdown-menu") {
		t.Errorf("Expected menu class to contain 'dropdown-menu', got '%s'", menuClass)
	}
}

func TestNavDropdownDivider(t *testing.T) {
	divider := NavDropdownDivider()

	// Check element type
	if tagName := divider.Element().TagName(); tagName != "LI" {
		t.Errorf("Expected tag name 'LI', got '%s'", tagName)
	}

	// Check hr inside li
	children := divider.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (hr), got %d", len(children))
	}

	hr := children[0].(dom.Element)
	if tagName := hr.TagName(); tagName != "HR" {
		t.Errorf("Expected hr tag name 'HR', got '%s'", tagName)
	}

	class := hr.GetAttribute("class")
	if !strings.Contains(class, "dropdown-divider") {
		t.Errorf("Expected class to contain 'dropdown-divider', got '%s'", class)
	}
}

func TestNavDropdownHeader(t *testing.T) {
	header := NavDropdownHeader("Settings")

	// Check element type
	if tagName := header.Element().TagName(); tagName != "LI" {
		t.Errorf("Expected tag name 'LI', got '%s'", tagName)
	}

	// Check h6 inside li
	children := header.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (h6), got %d", len(children))
	}

	h6 := children[0].(dom.Element)
	if tagName := h6.TagName(); tagName != "H6" {
		t.Errorf("Expected h6 tag name 'H6', got '%s'", tagName)
	}

	class := h6.GetAttribute("class")
	if !strings.Contains(class, "dropdown-header") {
		t.Errorf("Expected class to contain 'dropdown-header', got '%s'", class)
	}
}

func TestNavDropdown_Integration(t *testing.T) {
	// Test a complete dropdown in a nav
	homeItem := NavItem("#", "Home")
	homeItem.Active(true)

	nav := Nav().
		Append(homeItem).
		Append(NavDropdown("More",
			NavDropdownHeader("Account"),
			NavItem("/profile", "Profile"),
			NavItem("/settings", "Settings"),
			NavDropdownDivider(),
			NavDropdownHeader("Help"),
			NavItem("/help", "Documentation"),
			NavItem("/contact", "Contact"),
		)).
		Append(NavItem("#", "About"))

	// Check nav has expected children
	children := nav.Element().ChildNodes()
	expectedChildren := 3 // NavItem + NavDropdown + NavItem
	if len(children) != expectedChildren {
		t.Errorf("Expected %d children, got %d", expectedChildren, len(children))
	}

	// Check that middle child is the dropdown
	dropdownEl := children[1].(dom.Element)
	class := dropdownEl.GetAttribute("class")
	if !strings.Contains(class, "dropdown") {
		t.Errorf("Expected middle child to have dropdown class, got '%s'", class)
	}
}

func TestNavDropdown_InNavBar(t *testing.T) {
	// Test dropdown can be added to navbar
	navbar := NavBar(WithResponsive(BreakpointLarge))

	// Create dropdown and add to navbar
	dropdown := NavDropdown("Services",
		NavItem("/web", "Web Design"),
		NavItem("/mobile", "Mobile Apps"),
	)

	navbar.Append(dropdown)

	// Verify the dropdown was added correctly
	container := navbar.Element().ChildNodes()[0].(Element)
	collapse := container.ChildNodes()[1].(Element)
	navList := collapse.ChildNodes()[0].(Element) // navbar-nav UL inside collapse
	children := navList.ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child in navbar body, got %d", len(children))
	}

	dropdownEl := children[0].(dom.Element)
	class := dropdownEl.GetAttribute("class")
	if !strings.Contains(class, "nav-item") || !strings.Contains(class, "dropdown") {
		t.Errorf("Expected dropdown class 'nav-item dropdown', got '%s'", class)
	}
}

func TestNavDropdown_Active(t *testing.T) {
	// Test Active method
	dropdown := NavDropdown("Menu",
		NavItem("#", "Item 1"),
		NavItem("#", "Item 2"),
		NavItem("#", "Item 3"),
	)

	// Initially no items should be active
	activeIndices := dropdown.Active()
	if len(activeIndices) != 0 {
		t.Errorf("Expected 0 active items initially, got %d", len(activeIndices))
	}

	// Set item 1 as active
	activeIndices = dropdown.Active(1)
	if len(activeIndices) != 1 || activeIndices[0] != 1 {
		t.Errorf("Expected item 1 to be active, got %v", activeIndices)
	}

	// Check that -1 doesn't change state, just returns current
	activeIndices = dropdown.Active(-1)
	if len(activeIndices) != 1 || activeIndices[0] != 1 {
		t.Errorf("Expected item 1 to still be active after -1 call, got %v", activeIndices)
	}

	// Set item 0 as active (doesn't clear previous)
	dropdown.Active(0)
	activeIndices = dropdown.Active()
	if len(activeIndices) != 2 {
		t.Errorf("Expected 2 active items, got %d: %v", len(activeIndices), activeIndices)
	}
}

func TestNavDropdown_Disabled(t *testing.T) {
	// Test Disabled method
	dropdown := NavDropdown("Menu",
		NavItem("#", "Item 1"),
		NavItem("#", "Item 2"),
		NavItem("#", "Item 3"),
	)

	// Initially no items should be disabled
	disabledIndices := dropdown.Disabled()
	if len(disabledIndices) != 0 {
		t.Errorf("Expected 0 disabled items initially, got %d", len(disabledIndices))
	}

	// Set item 2 as disabled
	disabledIndices = dropdown.Disabled(2)
	if len(disabledIndices) != 1 || disabledIndices[0] != 2 {
		t.Errorf("Expected item 2 to be disabled, got %v", disabledIndices)
	}

	// Check that -1 doesn't change state, just returns current
	disabledIndices = dropdown.Disabled(-1)
	if len(disabledIndices) != 1 || disabledIndices[0] != 2 {
		t.Errorf("Expected item 2 to still be disabled after -1 call, got %v", disabledIndices)
	}

	// Set item 0 as disabled
	dropdown.Disabled(0)
	disabledIndices = dropdown.Disabled()
	if len(disabledIndices) != 2 {
		t.Errorf("Expected 2 disabled items, got %d: %v", len(disabledIndices), disabledIndices)
	}
}
