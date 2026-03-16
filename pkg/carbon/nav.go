package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type nav struct {
	mvc.View
}

type navitem struct {
	mvc.View
}

var _ mvc.View = (*nav)(nil)
var _ mvc.View = (*navitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// Header template: <cds-header-menu-button> is the hamburger toggle on small
	// screens and auto-connects to <cds-side-nav> via document-level events.
	// <div data-slot="name"> is replaced by .Named(); the default body slot
	// is <cds-header-nav> so HeaderNavItems are appended directly inside it.
	templateShellHeader = `
		<cds-header>
			<cds-header-menu-button></cds-header-menu-button>
			<div data-slot="name"></div>
			<cds-header-nav data-slot="body"></cds-header-nav>
		</cds-header>
	`

	// SideNav template: collapse-mode="responsive" means expanded on large
	// screens (≥1056px) and collapsed on small screens, toggled by the
	// <cds-header-menu-button> in the header.
	templateShellSideNav = `
		<cds-side-nav collapse-mode="responsive">
			<cds-side-nav-items data-slot="body"></cds-side-nav-items>
		</cds-side-nav>
	`
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewNav, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(nav), element, func(self, child mvc.View) {
			self.(*nav).View = child
		})
	})
	mvc.RegisterView(ViewNavItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(navitem), element, func(self, child mvc.View) {
			self.(*navitem).View = child
		})
	})
}

// Header returns a <cds-header> UI shell header bar.
// Pass HeaderNavItem children as args; call .Named() to set the product name link.
func Header(args ...any) *nav {
	return mvc.NewView(new(nav), ViewNav, templateShellHeader, func(self, child mvc.View) {
		self.(*nav).View = child
	}, args).(*nav)
}

// Label inserts a <cds-header-name> product link into the header name slot and
// sets the header's aria-label to "prefix name" for screen readers. Typically
// called fluently: cds.Header(...).Label("/", "IBM", "Product Name").
func (n *nav) Label(href, prefix string, args ...any) *nav {
	el := mvc.HTML("cds-header-name",
		mvc.WithAttr("href", href),
		mvc.WithAttr("prefix", prefix),
		args)
	n.ReplaceSlot("name", el)
	// Derive aria-label from prefix + visible name content
	n.Root().SetAttribute("aria-label", prefix+" "+el.InnerHTML())
	return n
}

// HeaderNavItem returns a <cds-header-nav-item> link for the header menu bar.
func HeaderNavItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-header-nav-item", func(self, child mvc.View) {
		self.(*navitem).View = child
	}, mvc.WithAttr("href", href), args).(*navitem)
}

// SideNav returns a <cds-side-nav> panel.
// Pass SideNavLink and SideNavSection children as args.
func SideNav(args ...any) *nav {
	return mvc.NewView(new(nav), ViewNav, templateShellSideNav, func(self, child mvc.View) {
		self.(*nav).View = child
	}, args).(*nav)
}

// SideNavLink returns a <cds-side-nav-link> top-level navigation entry.
func SideNavLink(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-link", func(self, child mvc.View) {
		self.(*navitem).View = child
	}, mvc.WithAttr("href", href), args).(*navitem)
}

// SideNavSection returns a <cds-side-nav-menu> collapsible category.
// The section starts expanded by default. The web component manages subsequent
// expand/collapse; pass SideNavItem children as args.
func SideNavSection(title string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu", func(self, child mvc.View) {
		self.(*navitem).View = child
	}, mvc.WithAttr("title", title), mvc.WithAttr("expanded", ""), args).(*navitem)
}

// SideNavItem returns a <cds-side-nav-menu-item> for use inside a SideNavSection.
func SideNavItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu-item", func(self, child mvc.View) {
		self.(*navitem).View = child
	}, mvc.WithAttr("href", href), args).(*navitem)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - SIDENAV

// Select implements mvc.Selectable. It clears the active attribute from every
// <cds-side-nav-link>, <cds-side-nav-menu-item>, and <cds-side-nav-menu> in
// the side nav, then sets active on each of the supplied views and on any
// parent <cds-side-nav-menu> that contains an active item. Pass no arguments
// to deselect all.
func (n *nav) Select(views ...mvc.View) {
	for _, tag := range []string{"cds-side-nav-link", "cds-side-nav-menu-item", "cds-side-nav-menu"} {
		for _, el := range n.Root().GetElementsByTagName(tag) {
			el.RemoveAttribute("active")
		}
	}
	for _, v := range views {
		v.Root().SetAttribute("active", "")
		// Also activate the parent section if this item is inside a cds-side-nav-menu
		if p := v.Root().ParentElement(); p != nil && p.TagName() == "cds-side-nav-menu" {
			p.SetAttribute("active", "")
		}
	}
}
