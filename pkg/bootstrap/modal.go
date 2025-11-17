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
	mvc.View
}

var _ mvc.View = (*modal)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewModal = "mvc-bs-modal"
)

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
	mvc.RegisterView(ViewModal, newModalFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Modal(id string, args ...any) *modal {
	return mvc.NewViewExEx(new(modal), ViewModal, templateModal, mvc.WithAttr("id", id), mvc.WithClass("modal-dialog-scrollable"), args).(*modal)
}

func StickyModal(id string, args ...any) *modal {
	// When modal is set to sticky. modal will not close when clicking outside of it.
	return Modal(id, mvc.WithAttr("data-bs-backdrop", "static"), mvc.WithAttr("data-bs-keyboard", "false"), args)
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
	modal.View = view
}

func (modal *modal) Header(children ...any) mvc.View {
	return modal.View.ReplaceSlot("header", mvc.HTML("DIV", mvc.WithClass("modal-header"), children))
}

func (modal *modal) Footer(children ...any) mvc.View {
	return modal.View.ReplaceSlot("footer", mvc.HTML("DIV", mvc.WithClass("modal-footer"), children))
}

func (modal *modal) Content(children ...any) mvc.View {
	return modal.View.ReplaceSlot("", mvc.HTML("DIV", mvc.WithClass("modal-body"), children))
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

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
