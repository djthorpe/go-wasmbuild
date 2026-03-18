package carbon

import (
	"fmt"
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
var _ mvc.EnabledState = (*dropdown)(nil)
var _ mvc.ActiveGroup = (*dropdown)(nil)
var _ mvc.ActiveState = (*dropdownItem)(nil)

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
// helperText is shown below the dropdown; pass an empty string for none.
func Dropdown(helperText string, args ...any) *dropdown {
	if helperText != "" {
		args = append([]any{mvc.WithAttr("helper-text", helperText)}, args...)
	}
	dd := mvc.NewView(new(dropdown), ViewDropdown, "cds-dropdown", setView, args).(*dropdown)
	dd.applyThemePresentation()
	return dd
}

// DropdownItem returns a <cds-dropdown-item> web component.
func DropdownItem(args ...any) *dropdownItem {
	return mvc.NewView(new(dropdownItem), ViewDropdownItem, "cds-dropdown-item", setView, args).(*dropdownItem)
}

func dropdownTitleText(args ...any) dom.Element {
	return mvc.HTML("SPAN", append([]any{mvc.WithAttr("slot", "title-text")}, args...)...)
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

// Enabled reports whether the dropdown is enabled.
func (d *dropdown) Enabled() bool {
	return !d.Root().HasAttribute("disabled")
}

// SetEnabled enables or disables the dropdown.
func (d *dropdown) SetEnabled(enabled bool) *dropdown {
	if enabled {
		d.Root().RemoveAttribute("disabled")
	} else {
		d.Root().SetAttribute("disabled", "")
	}
	return d
}

// SetActive marks the specified items as selected and deselects all others.
// Also updates the dropdown's value to match the first active item.
// With no arguments, all items are deselected.
func (d *dropdown) SetActive(views ...mvc.View) {
	active := make(map[dom.Element]struct{}, len(views))
	for _, v := range views {
		if v != nil {
			active[v.Root()] = struct{}{}
		}
	}
	for _, child := range d.Root().Children() {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if item, ok := v.(*dropdownItem); ok {
				_, on := active[child]
				item.SetActive(on)
				if on {
					d.SetValue(item.Value())
				}
			}
		}
	}
}

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

// Label returns the dropdown's title-text slot content.
func (d *dropdown) Label() string {
	for _, child := range d.Root().Children() {
		if child.GetAttribute("slot") == "title-text" {
			return strings.TrimSpace(child.TextContent())
		}
	}
	return ""
}

// SetLabel sets the dropdown's title-text slot to the given string.
func (d *dropdown) SetLabel(label string) *dropdown {
	for _, child := range d.Root().Children() {
		if child.GetAttribute("slot") == "title-text" {
			d.Root().RemoveChild(child)
		}
	}
	if label != "" {
		d.Root().AppendChild(dropdownTitleText(label))
	}
	return d
}

// Content appends dropdown items, replacing any existing children.
// Panics if any argument is not a *dropdownItem.
func (d *dropdown) Content(args ...any) mvc.View {
	for _, arg := range args {
		if _, ok := arg.(*dropdownItem); !ok {
			panic(fmt.Sprintf("Dropdown.Content: expected *dropdownItem, got %T", arg))
		}
	}
	return d.View.Content(args...)
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
