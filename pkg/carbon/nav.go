package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type nav struct {
	base
	items []*navitem
}

type navglobal struct{ base }

type navitem struct {
	base
	items []*navitem
}

var _ mvc.View = (*nav)(nil)
var _ mvc.View = (*navglobal)(nil)
var _ mvc.View = (*navitem)(nil)
var _ mvc.ActiveGroup = (*nav)(nil)
var _ mvc.ActiveState = (*navitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateShellHeader = `
		<cds-header>
			<cds-header-menu-button></cds-header-menu-button>
			<div data-slot="name"></div>
			<cds-header-nav data-slot="body"></cds-header-nav>
			<div class="cds--header__global" data-slot="global"></div>
		</cds-header>
	`

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
		return mvc.NewViewWithElement(new(nav), element, setView)
	}, EventClick, EventHoverBubbled, EventNoHoverBubbled, EventFocusBubbled, EventNoFocus, EventSectionToggling, EventSectionToggle)
	mvc.RegisterView(ViewNavGlobal, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(navglobal), element, setView)
	})
	mvc.RegisterView(ViewNavItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(navitem), element, setView)
	}, EventSectionToggling, EventSectionToggle)
}

// Header returns a <cds-header> UI shell header.
func Header(args ...any) *nav {
	opts, body, global := splitHeaderArgs(args...)
	n := mvc.NewView(new(nav), ViewNav, templateShellHeader, setView, opts).(*nav)
	n.items = navItems(body...)
	n.ReplaceSlotChildren("body", body...)
	if global != nil {
		n.ReplaceSlot("global", global)
	}
	bindNavItemClicks(n.items, func(item *navitem) {
		n.SetActive(item)
	})
	return n
}

// HeaderNavGlobal returns the right-aligned global actions container for a header.
func HeaderNavGlobal(args ...any) *navglobal {
	adapted := adaptHeaderGlobalArgs(args...)
	return mvc.NewView(new(navglobal), ViewNavGlobal, "div", setView, append([]any{mvc.WithClass("cds--header__global")}, adapted...)...).(*navglobal)
}

// HeaderNavItem returns a <cds-header-nav-item> link for the header menu bar.
func HeaderNavItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-header-nav-item", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}

// HeaderGlobalAction returns a button that adapts to a header global action when
// placed inside HeaderNavGlobal.
func HeaderGlobalAction(args ...any) *button {
	return Button(args...)
}

// SideNav returns a <cds-side-nav> shell panel.
func SideNav(args ...any) *nav {
	n := mvc.NewView(new(nav), ViewNav, templateShellSideNav, setView, args).(*nav)
	n.items = navItems(args...)
	return n
}

// SideNavLink returns a <cds-side-nav-link> top-level navigation entry.
func SideNavLink(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-link", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}

// SideNavSection returns a <cds-side-nav-menu> collapsible navigation group.
func SideNavSection(title string, args ...any) *navitem {
	n := mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu", setView, append([]any{mvc.WithAttr("title", title), mvc.WithAttr("expanded", "")}, args...)...).(*navitem)
	if strings.TrimSpace(n.Root().Value()) == "" && !n.Root().HasAttribute("value") {
		n.Root().SetValue(title)
		n.Root().SetAttribute("value", title)
	}
	n.items = navItems(args...)
	return n
}

// SideNavItem returns a <cds-side-nav-menu-item> for a SideNavSection.
func SideNavItem(href string, args ...any) *navitem {
	return mvc.NewView(new(navitem), ViewNavItem, "cds-side-nav-menu-item", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*navitem)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - NAV

// Label sets the header name slot with a <cds-header-name> link.
func (n *nav) Label(href, prefix string, args ...any) *nav {
	el := mvc.HTML("cds-header-name", append([]any{mvc.WithAttr("href", href), mvc.WithAttr("prefix", prefix)}, args...)...)
	n.ReplaceSlot("name", el)
	aria := strings.TrimSpace(strings.Join([]string{strings.TrimSpace(prefix), strings.TrimSpace(el.TextContent())}, " "))
	if aria == "" {
		n.Root().RemoveAttribute("aria-label")
	} else {
		n.Root().SetAttribute("aria-label", aria)
	}
	return n
}

// SetActive marks the supplied nav items active and clears the rest.
func (n *nav) SetActive(views ...mvc.View) {
	active := make(map[dom.Element]struct{}, len(views))
	for _, view := range views {
		if view != nil {
			active[view.Root()] = struct{}{}
		}
	}
	for _, item := range n.items {
		setNavItemActive(item, active)
	}
}

// OnSectionToggle adds a listener for side-nav section toggle completion.
func (n *nav) OnSectionToggle(handler func(dom.Event)) *nav {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventTarget(evt) != nil {
				handler(evt)
			}
		})
	}
	return n
}

// OnSectionExpanded adds a listener for side-nav sections after expansion.
func (n *nav) OnSectionExpanded(handler func(dom.Event)) *nav {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventExpanded(evt) {
				handler(evt)
			}
		})
	}
	return n
}

// OnSectionCollapsed adds a listener for side-nav sections after collapse.
func (n *nav) OnSectionCollapsed(handler func(dom.Event)) *nav {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventCollapsed(evt) {
				handler(evt)
			}
		})
	}
	return n
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - NAV ITEM

// SetActive marks the navigation item active or inactive.
func (n *navitem) SetActive(active bool) {
	setNavItemActiveElement(n.Root(), active)
}

// OnSectionToggle adds a listener for side-nav section toggle completion.
func (n *navitem) OnSectionToggle(handler func(dom.Event)) *navitem {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventTarget(evt) != nil {
				handler(evt)
			}
		})
	}
	return n
}

// OnSectionExpanded adds a listener for side-nav sections after expansion.
func (n *navitem) OnSectionExpanded(handler func(dom.Event)) *navitem {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventExpanded(evt) {
				handler(evt)
			}
		})
	}
	return n
}

// OnSectionCollapsed adds a listener for side-nav sections after collapse.
func (n *navitem) OnSectionCollapsed(handler func(dom.Event)) *navitem {
	if handler != nil {
		n.AddEventListener(EventSectionToggle, func(evt dom.Event) {
			if sectionEventCollapsed(evt) {
				handler(evt)
			}
		})
	}
	return n
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func navItems(args ...any) []*navitem {
	items := make([]*navitem, 0, len(args))
	for _, arg := range args {
		if item, ok := arg.(*navitem); ok && item != nil {
			items = append(items, item)
		}
	}
	return items
}

func sectionEventExpanded(evt dom.Event) bool {
	target := sectionEventTarget(evt)
	return target != nil && target.HasAttribute("expanded")
}

func sectionEventCollapsed(evt dom.Event) bool {
	target := sectionEventTarget(evt)
	return target != nil && !target.HasAttribute("expanded")
}

func sectionEventTarget(evt dom.Event) dom.Element {
	target, ok := evt.Target().(dom.Element)
	if !ok || target == nil || !target.HasAttribute(mvc.DataComponentAttrKey) {
		return nil
	}
	return target
}

func splitHeaderArgs(args ...any) ([]mvc.Opt, []any, *navglobal) {
	var (
		opts   []mvc.Opt
		body   []any
		global *navglobal
	)
	for _, arg := range args {
		switch value := arg.(type) {
		case []any:
			childOpts, childBody, childGlobal := splitHeaderArgs(value...)
			opts = append(opts, childOpts...)
			body = append(body, childBody...)
			if childGlobal != nil {
				global = childGlobal
			}
		case []mvc.Opt:
			opts = append(opts, value...)
		case mvc.Opt:
			opts = append(opts, value)
		case *navglobal:
			global = value
		default:
			body = append(body, value)
		}
	}
	return opts, body, global
}

func adaptHeaderGlobalArgs(args ...any) []any {
	adapted := make([]any, 0, len(args))
	for _, arg := range args {
		switch value := arg.(type) {
		case []any:
			adapted = append(adapted, adaptHeaderGlobalArgs(value...)...)
		case *button:
			adapted = append(adapted, adaptHeaderGlobalButton(value))
		default:
			adapted = append(adapted, value)
		}
	}
	return adapted
}

func adaptHeaderGlobalButton(b *button) *button {
	if b == nil {
		return nil
	}
	if strings.EqualFold(b.Root().TagName(), "cds-header-global-action") {
		return b
	}
	args := make([]any, 0, len(b.Root().Attributes())+len(b.Root().Children())+2)
	if id := b.Root().ID(); id != "" {
		args = append(args, mvc.WithID(id))
	}
	if classes := b.Root().ClassList().Values(); len(classes) > 0 {
		args = append(args, mvc.WithClass(classes...))
	}
	for _, attr := range b.Root().Attributes() {
		switch name := attr.Name(); name {
		case mvc.DataComponentAttrKey, "id", "class", "kind", "size":
			continue
		default:
			args = append(args, mvc.WithAttr(name, attr.Value()))
		}
	}
	if text := strings.TrimSpace(b.Root().TextContent()); text != "" {
		args = append(args, text)
	}
	for _, child := range b.Root().Children() {
		args = append(args, child)
	}
	adapted := mvc.NewView(new(button), ViewButton, "cds-header-global-action", setView, args...).(*button)
	setView(b, adapted)
	return b
}

func bindNavItemClicks(items []*navitem, onClick func(*navitem)) {
	for _, item := range items {
		if item == nil {
			continue
		}
		localItem := item
		localItem.AddEventListener(EventClick, func(dom.Event) {
			onClick(localItem)
		})
		bindNavItemClicks(localItem.items, onClick)
	}
}

func setNavItemActive(item *navitem, active map[dom.Element]struct{}) bool {
	if item == nil {
		return false
	}
	_, selfActive := active[item.Root()]
	childActive := false
	for _, child := range item.items {
		if setNavItemActive(child, active) {
			childActive = true
		}
	}
	activeState := selfActive || childActive
	item.SetActive(activeState)
	return activeState
}

func setNavItemActiveElement(element dom.Element, active bool) {
	if element == nil {
		return
	}
	isHeaderNavItem := strings.EqualFold(element.TagName(), "cds-header-nav-item")
	for _, attr := range navItemActiveAttrs(element.TagName()) {
		if active {
			element.SetAttribute(attr, "")
		} else {
			element.RemoveAttribute(attr)
		}
	}
	if isHeaderNavItem {
		element.RemoveAttribute("aria-current")
		return
	}
	if active {
		element.SetAttribute("aria-current", "page")
	} else {
		element.RemoveAttribute("aria-current")
	}
}

func navItemActiveAttrs(tagName string) []string {
	if strings.EqualFold(tagName, "cds-header-nav-item") {
		return []string{"is-active"}
	}
	return []string{"active"}
}
