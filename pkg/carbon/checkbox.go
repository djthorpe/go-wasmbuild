package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type checkbox struct{ base }

type checkboxGroup struct {
	base
}

// CheckboxState represents the tri-state value of a checkbox.
//
// `undefined` maps to Carbon's indeterminate state.
type CheckboxState string

// CheckboxOrientation determines how a checkbox group lays out its children.
// It is an Attr so it can be applied with With().
type CheckboxOrientation = Attr

const (
	CheckboxStateUndefined CheckboxState = "undefined"
	CheckboxStateFalse     CheckboxState = "false"
	CheckboxStateTrue      CheckboxState = "true"
)

var _ mvc.View = (*checkbox)(nil)
var _ mvc.View = (*checkboxGroup)(nil)
var _ mvc.ActiveState = (*checkbox)(nil)
var _ mvc.ActiveGroup = (*checkboxGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewCheckbox, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(checkbox), element, setView)
	}, EventChange)
	mvc.RegisterView(ViewCheckboxGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(checkboxGroup), element, setView)
	}, EventChange)
}

// Checkbox returns a <cds-checkbox> web component. An optional leading string
// argument sets the label-text attribute.
//
//	carbon.Checkbox("Enabled")
func Checkbox(args ...any) *checkbox {
	if len(args) > 0 {
		if label, ok := args[0].(string); ok {
			args = append([]any{mvc.WithAttr("label-text", label)}, args[1:]...)
		}
	}
	return mvc.NewView(new(checkbox), ViewCheckbox, "cds-checkbox", setView, args).(*checkbox)
}

// AddEventListener registers an event handler on the checkbox.
// EventChange is mapped to Carbon's cds-checkbox-changed custom event.
func (c *checkbox) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	c.View.AddEventListener(checkboxEvent(event), handler)
	return c
}

// RemoveEventListener removes an event handler from the checkbox.
func (c *checkbox) RemoveEventListener(event string) mvc.View {
	c.View.RemoveEventListener(checkboxEvent(event))
	return c
}

// CheckboxGroup returns a <cds-checkbox-group> web component.
// helperText is shown below the group; pass an empty string for none.
func CheckboxGroup(helperText string, args ...any) *checkboxGroup {
	if helperText != "" {
		args = append([]any{mvc.WithAttr("helper-text", helperText)}, args...)
	}
	return mvc.NewView(new(checkboxGroup), ViewCheckboxGroup, "cds-checkbox-group", setView, args).(*checkboxGroup)
}

// AddEventListener registers an event handler on the checkbox group.
// EventChange is mapped to Carbon's cds-checkbox-changed custom event.
func (g *checkboxGroup) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	g.View.AddEventListener(checkboxEvent(event), handler)
	return g
}

// RemoveEventListener removes an event handler from the checkbox group.
func (g *checkboxGroup) RemoveEventListener(event string) mvc.View {
	g.View.RemoveEventListener(checkboxEvent(event))
	return g
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

var _ mvc.EnabledState = (*checkbox)(nil)
var _ mvc.EnabledGroup = (*checkboxGroup)(nil)

func (c *checkbox) Enabled() bool {
	return !boolProperty(c.Root(), "disabled")
}

func (c *checkbox) SetEnabled(enabled bool) mvc.View {
	setBoolProperty(c.Root(), "disabled", !enabled)
	return c
}

// SetEnabled enables the specified checkboxes and disables the rest.
// Calling SetEnabled with no arguments disables all members.
func (g *checkboxGroup) Enabled() []mvc.View {
	enabled := make([]mvc.View, 0)
	for _, child := range g.Root().Children() {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if chk, ok := v.(*checkbox); ok && chk.Enabled() {
				enabled = append(enabled, chk)
			}
		}
	}
	return enabled
}

func (g *checkboxGroup) SetEnabled(views ...mvc.View) mvc.View {
	for _, child := range g.Root().Children() {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if chk, ok := v.(*checkbox); ok {
				// Use Equals (JS object identity) rather than a map keyed on
				// dom.Element interface values — Children() creates fresh wrapper
				// objects each call, so pointer equality always fails.
				on := false
				for _, ev := range views {
					if ev != nil && child.Equals(ev.Root()) {
						on = true
						break
					}
				}
				chk.SetEnabled(on)
			}
		}
	}
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CHECKBOX

// State returns the tri-state value of the checkbox.
func (c *checkbox) State() CheckboxState {
	if boolProperty(c.Root(), "indeterminate") {
		return CheckboxStateUndefined
	}
	if boolProperty(c.Root(), "checked") {
		return CheckboxStateTrue
	}
	return CheckboxStateFalse
}

// SetState updates the checkbox tri-state value.
func (c *checkbox) SetState(state CheckboxState) *checkbox {
	switch state {
	case CheckboxStateUndefined:
		setBoolProperty(c.Root(), "checked", false)
		setBoolProperty(c.Root(), "indeterminate", true)
	case CheckboxStateTrue:
		setBoolProperty(c.Root(), "indeterminate", false)
		setBoolProperty(c.Root(), "checked", true)
	default:
		setBoolProperty(c.Root(), "indeterminate", false)
		setBoolProperty(c.Root(), "checked", false)
	}
	return c
}

// Active reports whether the checkbox is checked.
func (c *checkbox) Active() bool {
	return c.State() == CheckboxStateTrue
}

// SetActive checks or unchecks the checkbox.
func (c *checkbox) SetActive(active bool) mvc.View {
	if active {
		c.SetState(CheckboxStateTrue)
	} else {
		c.SetState(CheckboxStateFalse)
	}
	return c
}

// Label returns the label-text attribute when explicitly set.
func (c *checkbox) Label() string {
	return c.Root().GetAttribute("label-text")
}

// SetLabel sets the label-text attribute.
func (c *checkbox) SetLabel(label string) *checkbox {
	c.Root().SetAttribute("label-text", label)
	return c
}

// Value returns the checkbox value attribute.
func (c *checkbox) Value() string {
	if value := c.Root().Value(); value != "" {
		return value
	}
	return c.Root().GetAttribute("value")
}

// SetValue sets the checkbox value attribute.
func (c *checkbox) SetValue(value string) *checkbox {
	c.Root().SetValue(value)
	c.Root().SetAttribute("value", value)
	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - CHECKBOX GROUP

// Content appends one or more checkboxes to the group, replacing any existing children.
// Panics if any argument is not a *checkbox.
func (g *checkboxGroup) Content(args ...any) mvc.View {
	for _, arg := range args {
		if _, ok := arg.(*checkbox); !ok {
			panic(fmt.Sprintf("CheckboxGroup.Content: expected *checkbox, got %T", arg))
		}
	}
	return g.View.Content(args...)
}

// SetActive marks the specified checkboxes active and deactivates the rest.
// Calling SetActive with no arguments deactivates all members.
func (g *checkboxGroup) Active() []mvc.View {
	active := make([]mvc.View, 0)
	for _, child := range g.Root().Children() {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if chk, ok := v.(*checkbox); ok && chk.Active() {
				active = append(active, chk)
			}
		}
	}
	return active
}

func (g *checkboxGroup) SetActive(views ...mvc.View) mvc.View {
	active := make(map[dom.Element]struct{}, len(views))
	for _, v := range views {
		if v != nil {
			active[v.Root()] = struct{}{}
		}
	}
	for _, child := range g.Root().Children() {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if chk, ok := v.(*checkbox); ok {
				_, on := active[child]
				chk.SetActive(on)
			}
		}
	}
	return g
}

// Label returns the group's legend text.
func (g *checkboxGroup) Label() string {
	return g.Root().GetAttribute("legend-text")
}

// SetLabel sets the group's legend text.
func (g *checkboxGroup) SetLabel(text string) *checkboxGroup {
	g.Root().SetAttribute("legend-text", text)
	return g
}

// Orientation returns the group's orientation.
func (g *checkboxGroup) Orientation() CheckboxOrientation {
	if value := g.Root().GetAttribute("orientation"); value != "" {
		return CheckboxOrientation(value)
	}
	return CheckboxOrientationVertical
}

// SetOrientation sets the group's orientation.
func (g *checkboxGroup) SetOrientation(orientation CheckboxOrientation) *checkboxGroup {
	g.Root().SetAttribute("orientation", string(orientation))
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func setBoolAttr(element dom.Element, name string, value bool) {
	if value {
		element.SetAttribute(name, "")
	} else {
		element.RemoveAttribute(name)
	}
}

func boolProperty(element dom.Element, name string) bool {
	if value, ok := element.JSValue().(js.Value); ok && !value.IsUndefined() && !value.IsNull() {
		return value.Get(name).Bool()
	}
	return element.HasAttribute(name)
}

func setBoolProperty(element dom.Element, name string, value bool) {
	if node, ok := element.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Set(name, value)
	}
	setBoolAttr(element, name, value)
}

func checkboxEvent(event string) string {
	if event == EventChange {
		return checkboxChangeEvent
	}
	return event
}
