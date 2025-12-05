package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type toast struct {
	mvc.View
}

type toastgroup struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateToast = `
		<div class="toast" role="alert" aria-live="assertive" aria-atomic="true">
			<script data-slot="header"></script>
			<div class="toast-body" data-slot="body"></div>
		</div>
	`
	templateToastHeader = `
		<div class="toast-header">
			<strong class="me-auto" data-slot="title"></strong>
			<small class="text-body-secondary" data-slot="subtitle"></small>
			<button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewToast, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(toast), element, setView)
	})
	mvc.RegisterView(ViewToastGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(toastgroup), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Toast(args ...any) *toast {
	return mvc.NewView(new(toast), ViewToast, templateToast, setView, args).(*toast)
}

func ToastGroup(args ...any) *toastgroup {
	return mvc.NewView(new(toastgroup), ViewToastGroup, "DIV", setView, mvc.WithClass("toast-container"), args).(*toastgroup)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (toast *toast) Header(args ...any) *toast {
	toast.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("toast-header"), args))
	return toast
}

// Show displays the toast
func (toast *toast) Show() {
	jsinstance(toast.Root(), "bootstrap.Toast").Call("show")
}

// Hide hides the toast
func (toast *toast) Hide() {
	jsinstance(toast.Root(), "bootstrap.Toast").Call("hide")
}
