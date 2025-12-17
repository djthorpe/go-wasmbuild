package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type alert struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateAlert = `
		<div class="alert alert-dismissible fade show" role="alert">
			<slot name="header"></slot>
			<slot name="body"></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewAlert, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(alert), element, setView)
	}, "close.bs.alert", "closed.bs.alert")
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Alert(args ...any) *alert {
	return mvc.NewView(new(alert), ViewAlert, templateAlert, setView, args).(*alert)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS
