package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct {
	mvc.View
}

var _ mvc.View = (*button)(nil)

// ButtonKind controls the visual style of a Carbon button.
type ButtonKind string

// ButtonSize controls the size of a Carbon button.
type ButtonSize string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// ButtonKind values — maps to the cds-button `kind` attribute.
	ButtonPrimary        ButtonKind = "primary"          // default: blue fill
	ButtonSecondary      ButtonKind = "secondary"        // gray fill
	ButtonTertiary       ButtonKind = "tertiary"         // outlined
	ButtonGhost          ButtonKind = "ghost"            // text-only
	ButtonDanger         ButtonKind = "danger"           // red fill (danger--primary)
	ButtonDangerTertiary ButtonKind = "danger--tertiary" // outlined danger
	ButtonDangerGhost    ButtonKind = "danger--ghost"    // ghost danger
)

const (
	// ButtonSize values — maps to the cds-button `size` attribute.
	ButtonSM  ButtonSize = "sm"  // 32px
	ButtonMD  ButtonSize = "md"  // 40px (default)
	ButtonLG  ButtonSize = "lg"  // 48px
	ButtonXL  ButtonSize = "xl"  // 64px
	ButtonXXL ButtonSize = "2xl" // 80px
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewButton, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(button), element, func(self, child mvc.View) {
			self.(*button).View = child
		})
	})
}

// Button returns a <cds-button> with kind="primary" by default.
// Pass ButtonKind options, ButtonSize options, icons, or text as args.
// Icons passed as children automatically receive slot="icon".
//
//	cds.Button("Save")
//	cds.Button("Delete", cds.WithButtonKind(cds.ButtonDanger))
//	cds.Button(cds.Icon("add"), "New item")
func Button(args ...any) *button {
	autoSlotIcons(args)
	return mvc.NewView(new(button), ViewButton, "cds-button", func(self, child mvc.View) {
		self.(*button).View = child
	}, args).(*button)
}

// autoSlotIcons recursively scans args and sets slot="icon" on any icon views.
func autoSlotIcons(args []any) {
	for _, arg := range args {
		switch v := arg.(type) {
		case []any:
			autoSlotIcons(v)
		case mvc.View:
			if v.Name() == ViewIcon {
				v.Root().SetAttribute("slot", "icon")
			}
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithButtonKind sets the visual style of a cds-button.
func WithButtonKind(k ButtonKind) mvc.Opt {
	return mvc.WithAttr("kind", string(k))
}

// WithButtonSize sets the size of a cds-button.
func WithButtonSize(s ButtonSize) mvc.Opt {
	return mvc.WithAttr("size", string(s))
}

// WithDisabled marks a cds-button as disabled.
func WithDisabled() mvc.Opt {
	return mvc.WithAttr("disabled", "")
}

// ButtonSet groups buttons horizontally in a flex row.
// Unlike <cds-button-set>, this preserves each button's individual kind and size.
func ButtonSet(args ...any) mvc.View {
	return mvc.NewView(new(button), ViewButton, "DIV", func(self, child mvc.View) {
		self.(*button).View = child
	}, mvc.WithClass("cds--button-group"), args)
}
