package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type overflowMenuItem struct{ base }

var _ mvc.View = (*overflowMenuItem)(nil)
var _ mvc.EnabledState = (*overflowMenuItem)(nil)

func init() {
	mvc.RegisterView(ViewOverflowItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(overflowMenuItem), element, setView)
	}, EventClick, EventOverflowMenuItemClick)
}

// OverflowMenuItem returns a <cds-overflow-menu-item> web component.
func OverflowMenuItem(args ...any) *overflowMenuItem {
	normalizeOverflowMenuArgs(args...)
	return mvc.NewView(new(overflowMenuItem), ViewOverflowItem, "cds-overflow-menu-item", setView, args...).(*overflowMenuItem)
}

func (i *overflowMenuItem) Apply(opts ...mvc.Opt) mvc.View {
	i.View.Apply(opts...)
	i.SetSize(i.Size())
	return i
}

func (i *overflowMenuItem) Enabled() bool {
	return !i.Root().HasAttribute("disabled")
}

func (i *overflowMenuItem) SetEnabled(enabled bool) mvc.View {
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
