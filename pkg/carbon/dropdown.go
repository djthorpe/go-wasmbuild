package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type dropdown struct{ base }

type dropdownItem struct{ base }

const dropdownThemeWhiteStyle = "--cds-field:#ffffff;--cds-field-hover:#e8e8e8;--cds-layer:#ffffff;--cds-layer-hover:#e8e8e8;--cds-layer-selected:#e0e0e0;--cds-layer-selected-hover:#d1d1d1;--cds-border-subtle:#e0e0e0"

var _ mvc.View = (*dropdown)(nil)
var _ mvc.View = (*dropdownItem)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewDropdown, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dropdown), element, setView)
	}, EventSelected)
	mvc.RegisterView(ViewDropdownItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dropdownItem), element, setView)
	})
}

// Dropdown returns a <cds-dropdown> web component.
//
//	carbon.Dropdown(
//		carbon.DropdownTitleText("Theme"),
//		carbon.DropdownItem(mvc.WithAttr("value", "one"), "One"),
//	)
func Dropdown(args ...any) *dropdown {
	dd := mvc.NewView(new(dropdown), ViewDropdown, "cds-dropdown", setView, args).(*dropdown)
	dd.applyThemePresentation()
	return dd
}

// DropdownItem returns a <cds-dropdown-item> web component.
func DropdownItem(args ...any) *dropdownItem {
	return mvc.NewView(new(dropdownItem), ViewDropdownItem, "cds-dropdown-item", setView, args).(*dropdownItem)
}

// DropdownTitleText returns a node slotted into the dropdown's title-text slot.
func DropdownTitleText(args ...any) dom.Element {
	return mvc.HTML("SPAN", append([]any{mvc.WithAttr("slot", "title-text")}, args...)...)
}

// DropdownHelperText returns a node slotted into the dropdown's helper-text slot.
func DropdownHelperText(args ...any) dom.Element {
	return mvc.HTML("SPAN", append([]any{mvc.WithAttr("slot", "helper-text")}, args...)...)
}

func (d *dropdown) Apply(opts ...mvc.Opt) mvc.View {
	d.View.Apply(opts...)
	d.applyThemePresentation()
	return d
}

func (d *dropdown) applyThemePresentation() {
	root := d.Root()
	style := strings.TrimSpace(strings.ReplaceAll(root.GetAttribute("style"), dropdownThemeWhiteStyle, ""))
	style = strings.Trim(style, "; ")
	if root.ClassList().Contains(ClassForTheme(ThemeWhite)) && !hasExplicitDropdownSurfaceStyle(style) {
		style = appendStyle(style, dropdownThemeWhiteStyle)
	}
	if style == "" {
		root.RemoveAttribute("style")
	} else {
		root.SetAttribute("style", style)
	}
}

func hasExplicitDropdownSurfaceStyle(style string) bool {
	return strings.Contains(style, "--cds-field:") || strings.Contains(style, "--cds-layer:") || strings.Contains(style, "--cds-border-subtle:")
}

func appendStyle(base, extra string) string {
	base = strings.Trim(base, "; ")
	extra = strings.Trim(extra, "; ")
	switch {
	case base == "":
		return extra
	case extra == "":
		return base
	default:
		return base + ";" + extra
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DROPDOWN

// Value returns the dropdown's selected value.
func (d *dropdown) Value() string {
	if value := d.Root().Value(); value != "" {
		return value
	}
	return d.Root().GetAttribute("value")
}

// SetValue sets the dropdown's selected value.
func (d *dropdown) SetValue(value string) *dropdown {
	d.Root().SetValue(value)
	d.Root().SetAttribute("value", value)
	return d
}

// Label sets the dropdown's title-text slot to the given string label.
func (d *dropdown) Label(label string) *dropdown {
	d.Root().AppendChild(DropdownTitleText(label))
	return d
}

// TitleText appends content to the dropdown's title-text slot.
func (d *dropdown) TitleText(args ...any) *dropdown {
	d.Root().AppendChild(DropdownTitleText(args...))
	return d
}

// HelperText appends content to the dropdown's helper-text slot.
func (d *dropdown) HelperText(args ...any) *dropdown {
	d.Root().AppendChild(DropdownHelperText(args...))
	return d
}

// AddItem appends a dropdown item as a child of the dropdown.
func (d *dropdown) AddItem(item *dropdownItem) *dropdown {
	if item != nil {
		d.Root().AppendChild(item.Root())
	}
	return d
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DROPDOWN ITEM

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
func (i *dropdownItem) SetActive(active bool) *dropdownItem {
	if active {
		i.Root().SetAttribute("selected", "")
	} else {
		i.Root().RemoveAttribute("selected")
	}
	return i
}
