package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type dropdownItem struct{ base }

var _ mvc.View = (*dropdownItem)(nil)
var _ mvc.ActiveState = (*dropdownItem)(nil)

func init() {
	mvc.RegisterView(ViewDropdownItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dropdownItem), element, setView)
	})
}

// DropdownItem returns a <cds-dropdown-item> web component.
func DropdownItem(args ...any) *dropdownItem {
	return mvc.NewView(new(dropdownItem), ViewDropdownItem, "cds-dropdown-item", setView, args).(*dropdownItem)
}

// Value returns the item value.
func (i *dropdownItem) Value() string {
	if value := i.Root().Value(); value != "" {
		return value
	}
	return i.Root().GetAttribute("value")
}

// SetValue sets the item value.
func (i *dropdownItem) SetValue(value string) *dropdownItem {
	i.Root().SetValue(value)
	i.Root().SetAttribute("value", value)
	return i
}

// Active reports whether the item is marked as the active selection.
func (i *dropdownItem) Active() bool {
	return i.Root().HasAttribute("selected")
}

// SetActive marks the item as active or inactive.
func (i *dropdownItem) SetActive(active bool) mvc.View {
	if active {
		i.Root().SetAttribute("selected", "")
	} else {
		i.Root().RemoveAttribute("selected")
	}
	return i
}
