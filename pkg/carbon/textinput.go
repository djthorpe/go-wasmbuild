package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type input struct {
	mvc.View
}

var _ mvc.View = (*input)(nil)

// InputType controls the HTML input type used by cds-text-input.
type InputType string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	InputText   InputType = "text"
	InputEmail  InputType = "email"
	InputTel    InputType = "tel"
	InputURL    InputType = "url"
	InputSearch InputType = "search"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewInput, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(input), element, func(self, child mvc.View) {
			self.(*input).View = child
		})
	})
}

func newInput(tag string, args ...any) *input {
	return mvc.NewView(new(input), ViewInput, tag, func(self, child mvc.View) {
		self.(*input).View = child
	}, args).(*input)
}

func slottedInput(tag, label, helper string, args ...any) *input {
	children := make([]any, 0, len(args)+2)
	if label != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "label-text"), label))
	}
	if helper != "" {
		children = append(children, mvc.HTML("span", mvc.WithAttr("slot", "helper-text"), helper))
	}
	children = append(children, args...)
	return newInput(tag, children)
}

func inputWithTrailingIcon(field mvc.View, iconName string) mvc.View {
	return Section(
		mvc.WithAttr("style", "position:relative;display:block;"),
		field,
		mvc.HTML("span",
			mvc.WithAttr("style", "position:absolute;right:0.875rem;top:3rem;transform:translateY(-50%);display:inline-flex;align-items:center;pointer-events:none;color:var(--cds-icon-secondary,#525252);"),
			Icon(iconName, mvc.WithAttr("size", "16")),
		),
	)
}

// Input returns a <cds-text-input> with slotted label and helper text.
// Placeholder, value, type, disabled, and readonly are controlled by options.
func Input(label, helper string, args ...any) mvc.View {
	return slottedInput("cds-text-input", label, helper, args...)
}

// EmailInput returns a text input specialized for email addresses.
func EmailInput(label, helper string, args ...any) mvc.View {
	args = append(args, WithInputType(InputEmail))
	return inputWithTrailingIcon(slottedInput("cds-text-input", label, helper, args...), "email--new")
}

// TelInput returns a text input specialized for telephone numbers.
func TelInput(label, helper string, args ...any) mvc.View {
	args = append(args, WithInputType(InputTel))
	return inputWithTrailingIcon(slottedInput("cds-text-input", label, helper, args...), "phone--filled")
}

// SecureInput returns a <cds-password-input> with Carbon's built-in reveal icon.
func SecureInput(label, helper string, args ...any) mvc.View {
	return slottedInput("cds-password-input", label, helper, args...)
}

// RangeInput returns a <cds-number-input> with increment/decrement steppers.
func RangeInput(label, helper string, args ...any) mvc.View {
	return slottedInput("cds-number-input", label, helper, args...)
}

// SearchInput returns a <cds-search>. Carbon search does not expose helper text
// via the same slots, so helper text is rendered directly beneath the control.
func SearchInput(label, helper string, args ...any) mvc.View {
	searchArgs := make([]any, 0, len(args)+1)
	if label != "" {
		searchArgs = append(searchArgs, label)
	}
	searchArgs = append(searchArgs, args...)
	search := newInput("cds-search", searchArgs...)
	if helper == "" {
		return search
	}
	return Section(
		search,
		HelperText(
			mvc.WithAttr("style", "margin-top:var(--cds-spacing-03,0.5rem);display:block;"),
			helper,
		),
	)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithInputPlaceholder sets the placeholder shown when the input is empty.
func WithInputPlaceholder(value string) mvc.Opt {
	return mvc.WithAttr("placeholder", value)
}

// WithInputValue sets the input value.
func WithInputValue(value string) mvc.Opt {
	return mvc.WithAttr("value", value)
}

// WithInputType sets the underlying HTML input type.
func WithInputType(value InputType) mvc.Opt {
	return mvc.WithAttr("type", string(value))
}

// WithInputDisabled makes the control non-interactive.
func WithInputDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// WithInputReadOnly allows focus and selection but prevents editing.
func WithInputReadOnly() mvc.Opt {
	return mvc.WithAttr("readonly", "")
}

// WithInputMin sets the minimum numeric value for RangeInput.
func WithInputMin(value string) mvc.Opt {
	return mvc.WithAttr("min", value)
}

// WithInputMax sets the maximum numeric value for RangeInput.
func WithInputMax(value string) mvc.Opt {
	return mvc.WithAttr("max", value)
}

// WithInputStep sets the numeric increment for RangeInput.
func WithInputStep(value string) mvc.Opt {
	return mvc.WithAttr("step", value)
}

// Compatibility aliases while the examples migrate to the new naming.
type TextInputType = InputType

const (
	TextInputText   = InputText
	TextInputEmail  = InputEmail
	TextInputTel    = InputTel
	TextInputURL    = InputURL
	TextInputSearch = InputSearch
)

func TextInput(label, helper string, args ...any) mvc.View { return Input(label, helper, args...) }
func WithTextInputPlaceholder(value string) mvc.Opt        { return WithInputPlaceholder(value) }
func WithTextInputValue(value string) mvc.Opt              { return WithInputValue(value) }
func WithTextInputType(value InputType) mvc.Opt            { return WithInputType(value) }
func WithTextInputDisabled() mvc.Opt                       { return WithInputDisabled() }
func WithTextInputReadOnly() mvc.Opt                       { return WithInputReadOnly() }
