package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type dropdown struct {
	mvc.View
}

type dropdownitem struct {
	mvc.View
}

var _ mvc.View = (*dropdown)(nil)
var _ mvc.View = (*dropdownitem)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewDropdown, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dropdown), element, func(self, child mvc.View) {
			self.(*dropdown).View = child
		})
	})
	mvc.RegisterView(ViewDropdownItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dropdownitem), element, func(self, child mvc.View) {
			self.(*dropdownitem).View = child
		})
	})
}

// Dropdown returns a <cds-dropdown> with slotted label and helper text.
func Dropdown(label, helper string, args ...any) mvc.View {
	children := make([]any, 0, len(args)+2)
	if label != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "label-text"), label))
	}
	if helper != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "helper-text"), helper))
	}
	children = append(children, args...)
	return mvc.NewView(new(dropdown), ViewDropdown, "cds-dropdown", func(self, child mvc.View) {
		self.(*dropdown).View = child
	}, children)
}

// DropdownItem returns a <cds-dropdown-item> option.
func DropdownItem(label, value string, args ...any) mvc.View {
	all := []any{mvc.WithAttr("value", value), label}
	all = append(all, args...)
	return mvc.NewView(new(dropdownitem), ViewDropdownItem, "cds-dropdown-item", func(self, child mvc.View) {
		self.(*dropdownitem).View = child
	}, all)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithDropdownPlaceholder sets the trigger content shown before a selection is made.
func WithDropdownPlaceholder(value string) mvc.Opt {
	return mvc.WithAttr("trigger-content", value)
}

// WithDropdownValue sets the current value.
func WithDropdownValue(value string) mvc.Opt {
	return mvc.WithAttr("value", value)
}

// WithDropdownDisabled makes the dropdown non-interactive.
func WithDropdownDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithDropdownOpen renders the menu in the open state.
func WithDropdownOpen() mvc.Opt {
	return mvc.WithAttr("open", "")
}

// WithDropdownRequired marks the field as required.
func WithDropdownRequired() mvc.Opt {
	return mvc.WithAttr("required", "")
}

// WithDropdownItemDisabled disables an option.
func WithDropdownItemDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithDropdownItemSelected marks an option as selected on first render.
func WithDropdownItemSelected() mvc.Opt {
	return mvc.WithAttr("selected", "")
}
