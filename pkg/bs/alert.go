package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type alert struct {
	View
}

var _ ViewWithVisibility = (*alert)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewAlert = "mvc-bs-alert"
)

func init() {
	RegisterView(ViewAlert, newAlertFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Alert(opt ...Opt) View {
	opt = append([]Opt{WithClass("alert", "alert-primary", "fade", "show"), WithAttr("role", "alert")}, opt...)
	return NewView(new(alert), ViewAlert, "DIV", opt...).Self()
}

func DismissableAlert(opt ...Opt) View {
	opt = append([]Opt{WithClass("alert", "alert-primary", "alert-dismissible", "fade", "show"), WithAttr("role", "alert")}, opt...)

	// Create a span for the body
	body := mvc.Tag("SPAN")

	// Set the body as the alert content
	view := NewView(new(alert), ViewAlert, "DIV", opt...).Body(body.Root())

	// Return the view
	return view
}

func newAlertFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return NewViewWithElement(new(alert), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (alert *alert) SetView(view View) {
	alert.View = view
}

func (alert *alert) Visible() bool {
	return alert.Root().ClassList().Contains("show")
}

func (alert *alert) Show() ViewWithVisibility {
	alert.Root().ClassList().Add("show")
	return alert
}

func (alert *alert) Hide() ViewWithVisibility {
	alert.Root().ClassList().Remove("show")
	return alert
}
