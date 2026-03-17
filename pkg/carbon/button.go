package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct{ base }

var _ mvc.View = (*button)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewButton, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(button), element, setView)
	}, EventClick, EventHover, EventNoHover, EventFocus, EventNoFocus)
}

// Button returns a <cds-button> web component.
// Use With() to apply kind, size, and other attributes:
//
//	carbon.Button(carbon.With(carbon.KindPrimary, carbon.SizeLarge), "Save")
func Button(args ...any) *button {
	normalizeButtonArgs(args...)
	b := mvc.NewView(new(button), ViewButton, "cds-button", setView, args).(*button)
	applyIconOnlyDefaultKind(b)
	return b
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

var _ mvc.EnabledState = (*button)(nil)

func (b *button) Enabled() bool {
	return !b.Root().HasAttribute("disabled")
}

func (b *button) SetEnabled(enabled bool) {
	if enabled {
		b.Root().RemoveAttribute("disabled")
	} else {
		b.Root().SetAttribute("disabled", "")
	}
}

///////////////////////////////////////////////////////////////////////////////
// VALUE

// Value returns the value attribute of the button, or empty string if not set.
func (b *button) Value() string {
	return b.Root().GetAttribute("value")
}

// SetValue sets the value attribute on the button element.
// This value is accessible from event handlers via e.Target().(dom.Element).GetAttribute("value").
func (b *button) SetValue(value string) *button {
	b.Root().SetAttribute("value", value)
	return b
}

// AddIcon appends an icon to the button's dedicated icon slot.
func (b *button) AddIcon(icon *icon) *button {
	if icon == nil {
		return b
	}
	applyButtonIconSlot(icon)
	b.Root().AppendChild(icon.Root())
	applyIconOnlyDefaultKind(b)
	return b
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func normalizeButtonArgs(args ...any) {
	for _, arg := range args {
		switch value := arg.(type) {
		case *icon:
			applyButtonIconSlot(value)
		case []any:
			normalizeButtonArgs(value...)
		}
	}
}

func applyIconOnlyDefaultKind(b *button) {
	if b == nil {
		return
	}
	root := b.Root()
	if root == nil || root.GetAttribute("kind") != "" {
		return
	}
	if strings.TrimSpace(root.TextContent()) != "" {
		return
	}
	for _, child := range root.Children() {
		if child.GetAttribute("slot") == "icon" {
			root.SetAttribute("kind", string(KindGhost))
			return
		}
	}
}

func applyButtonIconSlot(icon *icon) {
	if icon == nil {
		return
	}
	root := icon.Root()
	root.SetAttribute("slot", "icon")
	if root.GetAttribute("aria-hidden") == "" {
		root.SetAttribute("aria-hidden", "true")
	}
	style := strings.TrimSpace(root.GetAttribute("style"))
	if !strings.Contains(style, "color:") {
		if style == "" {
			root.SetAttribute("style", "color:currentColor")
		} else {
			root.SetAttribute("style", strings.TrimRight(style, "; ")+";color:currentColor")
		}
	}
}
