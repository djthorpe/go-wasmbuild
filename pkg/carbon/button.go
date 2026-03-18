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
// ACTIVE STATE

var _ mvc.ActiveState = (*button)(nil)

// Active reports whether the button is in its pressed/active state.
func (b *button) Active() bool {
	return b.Root().GetAttribute("aria-pressed") == "true"
}

// SetActive sets the pressed/active state of the button via aria-pressed.
func (b *button) SetActive(active bool) *button {
	if active {
		b.Root().SetAttribute("aria-pressed", "true")
	} else {
		b.Root().RemoveAttribute("aria-pressed")
	}
	return b
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

var _ mvc.EnabledState = (*button)(nil)

func (b *button) Enabled() bool {
	return !b.Root().HasAttribute("disabled")
}

func (b *button) SetEnabled(enabled bool) *button {
	if enabled {
		b.Root().RemoveAttribute("disabled")
	} else {
		b.Root().SetAttribute("disabled", "")
	}
	return b
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

// Label returns the button's accessible name (aria-label).
func (b *button) Label() string {
	return b.Root().GetAttribute("aria-label")
}

// SetLabel sets both the accessible name (aria-label) and tooltip text on the button.
// Use this for icon-only buttons where the button itself carries the accessible name.
func (b *button) SetLabel(label string) *button {
	if label == "" {
		b.Apply(mvc.WithoutAttr("aria-label"), mvc.WithoutAttr("tooltip-text"))
	} else {
		b.Apply(mvc.WithAriaLabel(label), mvc.WithAttr("tooltip-text", label))
	}
	return b
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON GROUP

type buttonGroup struct{ base }

var _ mvc.View = (*buttonGroup)(nil)
var _ mvc.ActiveGroup = (*buttonGroup)(nil)
var _ mvc.EnabledGroup = (*buttonGroup)(nil)

func init() {
	mvc.RegisterView(ViewButtonGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(buttonGroup), element, setView)
	}, EventClick, EventHover, EventNoHover, EventFocus, EventNoFocus)
}

// ButtonGroup returns a <cds-button-group> web component that arranges
// buttons horizontally with correct Carbon spacing.
func ButtonGroup(args ...any) *buttonGroup {
	return mvc.NewView(new(buttonGroup), ViewButtonGroup, "cds-button-group", setView, args).(*buttonGroup)
}

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

// AddEventListener registers a button-like event handler on the group root.
// Hover and focus events are mapped to bubbling equivalents so child button
// interactions are observed at the group level while keeping e.Target() on the
// individual button.
func (g *buttonGroup) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	g.View.AddEventListener(buttonGroupEvent(event), handler)
	return g
}

// RemoveEventListener removes a previously registered button-like event handler.
func (g *buttonGroup) RemoveEventListener(event string) mvc.View {
	g.View.RemoveEventListener(buttonGroupEvent(event))
	return g
}

// SetActive marks the specified buttons active and clears the rest.
// With no arguments, all buttons are deactivated.
func (g *buttonGroup) SetActive(views ...mvc.View) {
	active := make(map[mvc.View]bool, len(views))
	for _, v := range views {
		active[v] = true
	}
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok {
			b.SetActive(active[child])
		}
	}
}

// SetEnabled enables the specified buttons and disables all others in the group.
// With no arguments, all buttons are disabled.
func (g *buttonGroup) SetEnabled(views ...mvc.View) {
	enabled := make(map[mvc.View]bool, len(views))
	for _, v := range views {
		enabled[v] = true
	}
	for _, child := range g.Children() {
		if b, ok := child.(*button); ok {
			b.SetEnabled(enabled[child])
		}
	}
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

func buttonGroupEvent(event string) string {
	switch event {
	case EventHover:
		return EventHoverBubbled
	case EventNoHover:
		return EventNoHoverBubbled
	case EventFocus:
		return EventFocusBubbled
	default:
		return event
	}
}
