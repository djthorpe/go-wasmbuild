package bootstrap

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type modal struct {
	mvc.View
	instance js.Value
}

var _ mvc.View = (*modal)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	templateModal = `
		<div class="modal fade" tabindex="-1">
			<div class="modal-dialog">
				<div class="modal-content">
					<slot name="header"></slot>
					<slot></slot>
					<slot name="footer"></slot>
				</div>
			</div>
		</div>
	`
)

func init() {
	mvc.RegisterView(ViewModal, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(modal), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Modal(id string, args ...any) *modal {
	return mvc.NewView(new(modal), ViewModal, templateModal, setView, mvc.WithAttr("id", id), mvc.WithClass("modal-dialog-scrollable"), args).(*modal)
}

func StickyModal(id string, args ...any) *modal {
	// When modal is set to sticky. modal will not close when clicking outside of it.
	return Modal(id, mvc.WithAttr("data-bs-backdrop", "static"), mvc.WithAttr("data-bs-keyboard", "false"), args)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (modal *modal) Header(children ...any) *modal {
	modal.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("modal-header"), children))
	return modal
}

func (modal *modal) Footer(children ...any) *modal {
	modal.ReplaceSlot("footer", mvc.HTML("DIV", mvc.WithClass("modal-footer"), children))
	return modal
}

func (modal *modal) Content(children ...any) mvc.View {
	modal.ReplaceSlot("", mvc.HTML("DIV", mvc.WithClass("modal-body"), children))
	return modal
}

func (modal *modal) Show() *modal {
	if modal.instance.IsUndefined() {
		modal.instance = js.GetProto("bootstrap.Modal").New("#" + modal.ID())
	}
	modal.instance.Call("show")
	return modal
}

func (modal *modal) Hide() *modal {
	if modal.instance.IsUndefined() {
		modal.instance = js.GetProto("bootstrap.Modal").New("#" + modal.ID())
	}
	modal.instance.Call("hide")
	return modal
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithModal returns an option which configures a button to open the modal with the given ID
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
