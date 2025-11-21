package bootstrap

import (
	// Packages
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
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
	ViewToast      = "mvc-bs-toast"
	ViewToastGroup = "mvc-bs-toastgroup"
)

const (
	templateToast = `
		<div class="toast" role="alert" aria-live="assertive" aria-atomic="true">
			<slot name="header"></slot>
			<slot></slot>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewToast, newToastFromElement)
	mvc.RegisterView(ViewToastGroup, newToastGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Toast(name string, args ...any) *toast {
	return mvc.NewViewExEx(new(toast), ViewToast, templateToast, mvc.WithID(name), args).(*toast)
}

func ToastGroup(args ...any) *toastgroup {
	return mvc.NewView(new(toastgroup), ViewToastGroup, "DIV", mvc.WithClass("toast-container", "position-static"), args).(*toastgroup)
}

func newToastFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(toast), element)
}

func newToastGroupFromElement(element Element) mvc.View {
	if element.TagName() != "DIV" {
		return nil
	}
	return mvc.NewViewWithElement(new(toastgroup), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (toast *toast) SetView(view mvc.View) {
	toast.View = view
}

func (toastgroup *toastgroup) SetView(view mvc.View) {
	toastgroup.View = view
}

func (toast *toast) Header(args ...any) mvc.View {
	return toast.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("toast-header"), args))
}

func (toast *toast) Content(args ...any) mvc.View {
	return toast.ReplaceSlot("", mvc.HTML("DIV", mvc.WithClass("toast-body"), args))
}

func (toast *toast) Show() mvc.ViewWithVisibility {
	// TODO: Clear issues for non-wasm builds
	if j := js.GetProto("bootstrap.Toast"); j.IsUndefined() {
		panic("bootstrap.Toast is undefined")
	} else {
		j.New("#" + toast.ID()).Call("show")
	}
	return toast

}

func (toast *toast) Hide() mvc.ViewWithVisibility {
	// TODO: Clear issues for non-wasm builds
	if j := js.GetProto("bootstrap.Toast"); j.IsUndefined() {
		panic("bootstrap.Toast is undefined")
	} else {
		j.New("#" + toast.ID()).Call("hide")
	}
	return toast

}

func (toast *toast) Visible() bool {
	// TODO
	return false
}
