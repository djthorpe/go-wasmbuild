package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func init() {
	mvc.RegisterView(ViewSecureInput, func(element dom.Element) mvc.View {
		i := mvc.NewViewWithElement(new(input), element, setView).(*input)
		initializeInput(i)
		return i
	}, EventInput, EventChange, EventInvalid, EventFocus, EventNoFocus)
}

// SecureInput returns a <cds-password-input> web component.
// It shares the same public methods as Input.
func SecureInput(args ...any) *input {
	i := mvc.NewView(new(input), ViewSecureInput, "cds-password-input", setView, args).(*input)
	initializeInput(i)
	return i
}
