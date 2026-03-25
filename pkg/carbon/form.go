package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type form struct{ base }

var _ mvc.View = (*form)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewForm, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(form), element, setView)
	}, EventInput, EventChange, EventInvalid, EventFocusBubbled, EventNoFocus)
}

// Form returns a <cds-form> web component.
func Form(args ...any) *form {
	return mvc.NewView(new(form), ViewForm, "cds-form", setView, args).(*form)
}

// AddEventListener registers an event handler on the form.
// EventInvalid is bridged because descendant invalid events do not bubble.
func (f *form) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	if event == EventChange {
		f.View.AddEventListener(checkboxChangeEvent, handler)
	}
	if event == EventInput || event == EventChange {
		f.View.AddEventListener(numberInputChangeEvent, handler)
	}
	if event == EventInvalid {
		ensureFormInvalidBridge(f)
	}
	f.View.AddEventListener(formContainerEvent(event), handler)
	return f
}

// RemoveEventListener removes an event handler from the form.
func (f *form) RemoveEventListener(event string) mvc.View {
	if event == EventChange {
		f.View.RemoveEventListener(checkboxChangeEvent)
	}
	if event == EventInput || event == EventChange {
		f.View.RemoveEventListener(numberInputChangeEvent)
	}
	f.View.RemoveEventListener(formContainerEvent(event))
	return f
}

func formContainerEvent(event string) string {
	switch event {
	case EventFocus:
		return EventFocusBubbled
	default:
		return event
	}
}
