package bootstrap

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type modal struct {
	mvc.ViewWithHeaderFooter
}

var _ mvc.View = (*modal)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewModal = "mvc-bs-modal"
)

func init() {
	mvc.RegisterView(ViewModal, newModalFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Modal(id string, args ...any) *modal {
	// TODO:
	// <div class="modal fase">
	//  <div class="modal-dialog">
	//  <div class="modal-content">
	//    <div class="modal-header">...</div>
	//    <div class="modal-body">...</div>
	//    <div class="modal-footer">...</div>
	//  </div>
	//  </div>
	// </div>
	header := mvc.HTML("DIV", mvc.WithClass("modal-header"))
	body := mvc.HTML("DIV", mvc.WithClass("modal-dialog"))
	footer := mvc.HTML("DIV", mvc.WithClass("modal-footer"))
	return mvc.NewViewEx(
		new(modal), ViewModal, "DIV",
		header, body, footer, nil,
		mvc.WithAttr("id", id), mvc.WithClass("modal", "fade"), mvc.WithAttr("tabindex", "-1"),
		args,
	).(*modal)
}

func newModalFromElement(element Element) mvc.View {
	tagName := element.TagName()
	if tagName != "DIV" {
		panic(fmt.Sprintf("newModalFromElement: invalid tag name %q", tagName))
	}
	return mvc.NewViewWithElement(new(modal), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (modal *modal) SetView(view mvc.View) {
	modal.ViewWithHeaderFooter = view.(mvc.ViewWithHeaderFooter)
}

func WithModal(id string) mvc.Opt {
	return func(opts mvc.OptSet) error {
		if opts.Name() != ViewButton {
			return fmt.Errorf("WithModal: invalid view type %q", opts.Name())
		}
		if err := mvc.WithAttr("data-bs-toggle", "modal")(opts); err != nil {
			return err
		}
		if err := mvc.WithAttr("data-bs-target", "#"+id)(opts); err != nil {
			return err
		}
		return mvc.WithAttr("aria-controls", id)(opts)
	}
}
