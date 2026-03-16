package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type selectfield struct {
	mvc.View
}

type selectitem struct {
	mvc.View
}

type selectgroup struct {
	mvc.View
}

var _ mvc.View = (*selectfield)(nil)
var _ mvc.View = (*selectitem)(nil)
var _ mvc.View = (*selectgroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewSelect, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(selectfield), element, func(self, child mvc.View) {
			self.(*selectfield).View = child
		})
	})
	mvc.RegisterView(ViewSelectItem, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(selectitem), element, func(self, child mvc.View) {
			self.(*selectitem).View = child
		})
	})
	mvc.RegisterView(ViewSelectGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(selectgroup), element, func(self, child mvc.View) {
			self.(*selectgroup).View = child
		})
	})
}

// Select returns a <cds-select> with slotted label and helper text.
func Select(label, helper string, args ...any) mvc.View {
	children := make([]any, 0, len(args)+2)
	if label != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "label-text"), label))
	}
	if helper != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "helper-text"), helper))
	}
	children = append(children, args...)
	return mvc.NewView(new(selectfield), ViewSelect, "cds-select", func(self, child mvc.View) {
		self.(*selectfield).View = child
	}, children)
}

// SelectItem returns a <cds-select-item> option.
func SelectItem(label, value string, args ...any) mvc.View {
	all := []any{mvc.WithAttr("label", label), mvc.WithAttr("value", value), label}
	all = append(all, args...)
	return mvc.NewView(new(selectitem), ViewSelectItem, "cds-select-item", func(self, child mvc.View) {
		self.(*selectitem).View = child
	}, all)
}

// SelectGroup returns a <cds-select-item-group> option group.
func SelectGroup(label string, args ...any) mvc.View {
	all := []any{mvc.WithAttr("label", label)}
	all = append(all, args...)
	return mvc.NewView(new(selectgroup), ViewSelectGroup, "cds-select-item-group", func(self, child mvc.View) {
		self.(*selectgroup).View = child
	}, all)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithSelectPlaceholder sets the prompt shown when no value is selected.
func WithSelectPlaceholder(value string) mvc.Opt {
	return mvc.WithAttr("placeholder", value)
}

// WithSelectValue sets the current value.
func WithSelectValue(value string) mvc.Opt {
	return mvc.WithAttr("value", value)
}

// WithSelectDisabled makes the select non-interactive.
func WithSelectDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithSelectReadOnly prevents changing the selected option.
func WithSelectReadOnly() mvc.Opt {
	return mvc.WithAttr("readonly", "")
}

// WithSelectItemDisabled disables an option or option group.
func WithSelectItemDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithSelectItemSelected marks an option as selected on first render.
func WithSelectItemSelected() mvc.Opt {
	return mvc.WithAttr("selected", "")
}