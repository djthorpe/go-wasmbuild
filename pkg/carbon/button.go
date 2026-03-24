package carbon

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type button struct{ base }
type buttonGroup struct{ base }

var _ mvc.View = (*button)(nil)
var _ mvc.View = (*buttonGroup)(nil)
var _ mvc.EnabledState = (*button)(nil)
var _ mvc.EnabledGroup = (*buttonGroup)(nil)
var _ mvc.LabelState = (*button)(nil)
var _ mvc.ValueState = (*button)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	// Button and CloseButton
	mvc.RegisterView(ViewButton, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(button), element, setView)
	}, EventClick, EventHover, EventNoHover, EventFocus, EventNoFocus)

	// ButtonGroup
	mvc.RegisterView(ViewButtonGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(buttonGroup), element, setView)
	}, EventClick, EventHover, EventNoHover, EventFocus, EventNoFocus)
}

// Button returns a <cds-button> web component.
// Use With() to apply kind, size, and other attributes
func Button(args ...any) *button {
	normalizeButtonArgs(args...)
	return mvc.NewView(new(button), ViewButton, "cds-button", setView, args).(*button)
}

// CloseButton returns a ghost icon-only button with a close (X) icon.
func CloseButton(args ...any) *button {
	b := Button(append([]any{Icon(IconClose, With(IconSize20)), With(KindGhost), mvc.WithAriaLabel("Close")}, args...)...)
	b.AddEventListener(EventClick, func(evt dom.Event) {
		view := mvc.ViewFromEvent(evt)
		for {
			if view == nil {
				return
			}
			if vs, ok := view.(mvc.VisibleState); ok {
				vs.SetVisible(false)
				return
			}
			view = view.Parent()
		}
	})
	return b
}

// ButtonGroup returns a <cds-button-group> web component that arranges
// buttons horizontally with correct Carbon spacing.
func ButtonGroup(args ...any) *buttonGroup {
	return mvc.NewView(new(buttonGroup), ViewButtonGroup, "cds-button-group", setView, args).(*buttonGroup)
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

func (b *button) Enabled() bool {
	return !b.Root().HasAttribute("disabled")
}

func (b *button) SetEnabled(enabled bool) mvc.View {
	if enabled {
		b.Root().RemoveAttribute("disabled")
	} else {
		b.Root().SetAttribute("disabled", "")
	}
	return b
}

func (g *buttonGroup) Enabled() []mvc.View {
	enabled := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok && b.Enabled() {
			enabled = append(enabled, b)
		}
	}
	return enabled
}

func (g *buttonGroup) SetEnabled(views ...mvc.View) mvc.View {
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok {
			// Compare by root element identity rather than interface identity so
			// callers can pass equivalent view wrappers for the same DOM element.
			on := false
			for _, view := range views {
				if view != nil && child.Root().Equals(view.Root()) {
					on = true
					break
				}
			}
			b.SetEnabled(on)
		}
	}
	return g
}

///////////////////////////////////////////////////////////////////////////////
// VALUE STATE

// Value returns the value attribute of the button, or empty string if not set.
func (b *button) Value() string {
	return b.Root().GetAttribute("value")
}

// SetValue sets the value attribute on the button element.
// This value is accessible from event handlers via e.Target().(dom.Element).GetAttribute("value").
func (b *button) SetValue(value string) mvc.View {
	b.Root().SetAttribute("value", value)
	return b
}

///////////////////////////////////////////////////////////////////////////////
// LABEL STATE

// Label returns the button's accessible name (aria-label).
func (b *button) Label() string {
	return b.Root().GetAttribute("aria-label")
}

// SetLabel sets both the accessible name (aria-label) and tooltip text on the button.
// Use this for icon-only buttons where the button itself carries the accessible name.
func (b *button) SetLabel(label string) mvc.View {
	if label == "" {
		b.Apply(mvc.WithoutAttr("aria-label"), mvc.WithoutAttr("tooltip-text"))
	} else {
		b.Apply(mvc.WithAriaLabel(label), mvc.WithAttr("tooltip-text", label))
	}
	return b
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Content appends buttons to the group. Panics if any arg is not a *button.
func (g *buttonGroup) Content(args ...any) mvc.View {
	children := make([]any, 0, len(args))
	for _, arg := range args {
		if b, ok := arg.(*button); ok {
			children = append(children, b)
		} else {
			panic(fmt.Sprintf("ButtonGroup.Content: expected *button, got %T", arg))
		}
	}
	return g.View.Content(children...)
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

func applyButtonIconSlot(icon *icon) {
	if icon == nil {
		return
	}
	root := icon.Root()
	root.SetAttribute("slot", "icon")
	if root.GetAttribute("aria-hidden") == "" && root.GetAttribute("aria-label") == "" {
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
