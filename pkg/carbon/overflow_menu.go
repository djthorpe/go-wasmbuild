package carbon

import (
	"strconv"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type overflowMenu struct{ base }

type overflowMenuItem struct{ base }

var _ mvc.View = (*overflowMenu)(nil)
var _ mvc.View = (*overflowMenuItem)(nil)
var _ mvc.EnabledState = (*overflowMenu)(nil)
var _ mvc.VisibleState = (*overflowMenu)(nil)
var _ mvc.EnabledState = (*overflowMenuItem)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewOverflowMenu, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(overflowMenu), element, setView)
	}, EventClick)

	mvc.RegisterView(ViewOverflowItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(overflowMenuItem), element, setView)
	}, EventClick, EventOverflowMenuItemClick)
}

// OverflowMenu returns a <cds-overflow-menu> web component.
func OverflowMenu(args ...any) *overflowMenu {
	normalizeOverflowMenuArgs(args...)
	return mvc.NewView(new(overflowMenu), ViewOverflowMenu, "cds-overflow-menu", setView, args...).(*overflowMenu)
}

// OverflowMenuItem returns a <cds-overflow-menu-item> web component.
func OverflowMenuItem(args ...any) *overflowMenuItem {
	normalizeOverflowMenuArgs(args...)
	return mvc.NewView(new(overflowMenuItem), ViewOverflowItem, "cds-overflow-menu-item", setView, args...).(*overflowMenuItem)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - OVERFLOW MENU

func (m *overflowMenu) Apply(opts ...mvc.Opt) mvc.View {
	m.View.Apply(opts...)
	m.syncPresentation()
	return m
}

func (m *overflowMenu) Enabled() bool {
	return !tagBoolProperty(m.Root(), "disabled")
}

func (m *overflowMenu) SetEnabled(enabled bool) *overflowMenu {
	setTagBoolProperty(m.Root(), "disabled", !enabled)
	return m
}

func (m *overflowMenu) Visible() bool {
	return tagBoolProperty(m.Root(), "open")
}

func (m *overflowMenu) SetVisible(visible bool) mvc.View {
	setTagBoolProperty(m.Root(), "open", visible)
	return m
}

func (m *overflowMenu) SetFlipped(flipped bool) *overflowMenu {
	setTagBoolProperty(m.Root(), "flipped", flipped)
	if body := m.menuBody(); body != nil {
		setTagBoolProperty(body, "flipped", flipped)
	}
	return m
}

func (m *overflowMenu) Size() Attr {
	return normalizeOverflowMenuSize(Attr(m.Root().GetAttribute("size")))
}

func (m *overflowMenu) SetSize(size Attr) *overflowMenu {
	size = normalizeOverflowMenuSize(size)
	if size == "" {
		m.Root().RemoveAttribute("size")
	} else {
		m.Root().SetAttribute("size", string(size))
	}
	if body := m.menuBody(); body != nil {
		if size == "" {
			body.RemoveAttribute("size")
		} else {
			body.SetAttribute("size", string(size))
		}
	}
	return m
}

// Label returns the tooltip-content slot text.
func (m *overflowMenu) Label() string {
	if child := m.slotChild("tooltip-content"); child != nil {
		return strings.TrimSpace(child.TextContent())
	}
	return ""
}

// SetLabel sets the trigger tooltip/aria label via the tooltip-content slot.
func (m *overflowMenu) SetLabel(label string) *overflowMenu {
	if child := m.slotChild("tooltip-content"); child != nil {
		child.Remove()
	}
	if label != "" {
		m.Root().AppendChild(mvc.HTML("SPAN", mvc.WithAttr("slot", "tooltip-content"), label))
	}
	return m
}

// Content appends root children such as trigger icons directly to the host and
// routes overflow-menu items into the menu body.
func (m *overflowMenu) Content(args ...any) mvc.View {
	items := make([]any, 0, len(args))
	firstEnabledIndex := 0
	for _, arg := range args {
		if item, ok := arg.(*overflowMenuItem); ok {
			items = append(items, arg)
			if firstEnabledIndex == 0 && item.Enabled() {
				firstEnabledIndex = len(items)
			}
			continue
		}
		m.Root().AppendChild(mvc.NodeFromAny(arg))
	}
	if len(items) == 0 {
		return m
	}
	body := mvc.HTML("cds-overflow-menu-body")
	setTagBoolProperty(body, "flipped", tagBoolProperty(m.Root(), "flipped"))
	for _, arg := range items {
		body.AppendChild(mvc.NodeFromAny(arg))
	}
	if existing := m.menuBody(); existing != nil {
		existing.Remove()
	}
	m.Root().AppendChild(body)
	if firstEnabledIndex > 0 && m.Root().GetAttribute("index") == "" {
		m.Root().SetAttribute("index", toString(firstEnabledIndex))
	}
	return m
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - OVERFLOW MENU ITEM

func (i *overflowMenuItem) Apply(opts ...mvc.Opt) mvc.View {
	i.View.Apply(opts...)
	i.SetSize(i.Size())
	return i
}

func (i *overflowMenuItem) Enabled() bool {
	return !i.Root().HasAttribute("disabled")
}

func (i *overflowMenuItem) SetEnabled(enabled bool) *overflowMenuItem {
	if enabled {
		i.Root().RemoveAttribute("disabled")
	} else {
		i.Root().SetAttribute("disabled", "")
	}
	return i
}

func (i *overflowMenuItem) Value() string {
	if value := i.Root().Value(); value != "" {
		return value
	}
	return i.Root().GetAttribute("value")
}

func (i *overflowMenuItem) SetValue(value string) *overflowMenuItem {
	i.Root().SetValue(value)
	i.Root().SetAttribute("value", value)
	return i
}

func (i *overflowMenuItem) SetLabel(label string) *overflowMenuItem {
	i.Root().SetInnerHTML(label)
	return i
}

func (i *overflowMenuItem) SetDanger(danger bool) *overflowMenuItem {
	setTagBoolProperty(i.Root(), "danger", danger)
	return i
}

func (i *overflowMenuItem) SetDivider(divider bool) *overflowMenuItem {
	setTagBoolProperty(i.Root(), "divider", divider)
	return i
}

func (i *overflowMenuItem) Size() Attr {
	return normalizeOverflowMenuSize(Attr(i.Root().GetAttribute("size")))
}

func (i *overflowMenuItem) SetSize(size Attr) *overflowMenuItem {
	size = normalizeOverflowMenuSize(size)
	if size == "" {
		i.Root().RemoveAttribute("size")
	} else {
		i.Root().SetAttribute("size", string(size))
	}
	return i
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func normalizeOverflowMenuArgs(args ...any) {
	for _, arg := range args {
		switch value := arg.(type) {
		case *icon:
			applyOverflowMenuIconSlot(value)
		case []any:
			normalizeOverflowMenuArgs(value...)
		}
	}
}

func applyOverflowMenuIconSlot(icon *icon) {
	if icon == nil {
		return
	}
	applyButtonIconSlot(icon)
	icon.Root().ClassList().Add("cds--overflow-menu__icon")
}

func (m *overflowMenu) syncPresentation() {
	m.SetFlipped(tagBoolProperty(m.Root(), "flipped"))
	m.SetSize(m.Size())
}

func normalizeOverflowMenuSize(size Attr) Attr {
	switch size {
	case SizeExtraSmall, SizeSmall, SizeMedium, SizeLarge:
		return size
	default:
		return SizeMedium
	}
}

func (m *overflowMenu) menuBody() dom.Element {
	for _, child := range m.Root().Children() {
		if strings.EqualFold(child.TagName(), "cds-overflow-menu-body") {
			return child
		}
	}
	return nil
}

func (m *overflowMenu) slotChild(slot string) dom.Element {
	for _, child := range m.Root().Children() {
		if child.GetAttribute("slot") == slot {
			return child
		}
	}
	return nil
}

func toString(value int) string {
	return strconv.Itoa(value)
}
