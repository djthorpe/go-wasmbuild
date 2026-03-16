package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type textarea struct {
	mvc.View
}

var _ mvc.View = (*textarea)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTextArea, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(textarea), element, func(self, child mvc.View) {
			self.(*textarea).View = child
		})
	})
}

// TextArea returns a <cds-textarea> with slotted label and helper text.
// Placeholder, value, rows, readonly, and disabled are controlled by options.
func TextArea(label, helper string, args ...any) mvc.View {
	children := make([]any, 0, len(args)+2)
	if label != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "label-text"), label))
	}
	if helper != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "helper-text"), helper))
	}
	children = append(children, args...)
	return mvc.NewView(new(textarea), ViewTextArea, "cds-textarea", func(self, child mvc.View) {
		self.(*textarea).View = child
	}, children)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTextAreaPlaceholder sets the placeholder shown when the field is empty.
func WithTextAreaPlaceholder(value string) mvc.Opt {
	return mvc.WithAttr("placeholder", value)
}

// WithTextAreaValue sets the textarea value.
func WithTextAreaValue(value string) mvc.Opt {
	return mvc.WithAttr("value", value)
}

// WithTextAreaRows sets the number of visible rows.
func WithTextAreaRows(rows string) mvc.Opt {
	return mvc.WithAttr("rows", rows)
}

// WithTextAreaDisabled makes the control non-interactive.
func WithTextAreaDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithTextAreaReadOnly allows selection but prevents editing.
func WithTextAreaReadOnly() mvc.Opt {
	return mvc.WithAttr("readonly", "")
}
