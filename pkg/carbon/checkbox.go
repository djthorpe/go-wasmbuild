package carbon

import (
	// Packages
	"fmt"

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
var _ mvc.ActiveGroup = (*checkboxGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewCheckbox, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(checkbox), element, setView)
	}, EventCheckboxChanged)
	mvc.RegisterView(ViewCheckboxGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(checkboxGroup), element, setView)
	})
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

// CheckboxGroup returns a <cds-checkbox-group> web component.
func CheckboxGroup(args ...any) *checkboxGroup {
	return mvc.NewView(new(checkboxGroup), ViewCheckboxGroup, "cds-checkbox-group", setView, args).(*checkboxGroup)
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

var _ mvc.EnabledState = (*checkbox)(nil)
var _ mvc.EnabledState = (*checkboxGroup)(nil)

func (c *checkbox) Enabled() bool {
	return !boolProperty(c.Root(), "disabled")
}

func (c *checkbox) SetEnabled(enabled bool) {
	setBoolProperty(c.Root(), "disabled", !enabled)
}

func (g *checkboxGroup) Enabled() bool {
	return !boolProperty(g.Root(), "disabled")
}

func (g *checkboxGroup) SetEnabled(enabled bool) {
	setBoolProperty(g.Root(), "disabled", !enabled)
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
func (c *checkbox) SetState(state CheckboxState) {
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
}

// Active reports whether the checkbox is checked.
func (c *checkbox) Active() bool {
	return c.State() == CheckboxStateTrue
}

// SetActive checks or unchecks the checkbox and returns the receiver for chaining.
func (c *checkbox) SetActive(active bool) *checkbox {
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
func (c *checkbox) SetValue(value string) {
	c.Root().SetValue(value)
	c.Root().SetAttribute("value", value)
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
func (g *checkboxGroup) SetActive(views ...mvc.View) {
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
}

// HelperText returns the group's helper text.
func (g *checkboxGroup) HelperText() string {
	return g.Root().GetAttribute("helper-text")
}

// SetHelperText sets the group's helper text.
func (g *checkboxGroup) SetHelperText(text string) {
	g.Root().SetAttribute("helper-text", text)
}

// LegendText returns the group's legend text.
func (g *checkboxGroup) LegendText() string {
	return g.Root().GetAttribute("legend-text")
}

// SetLegendText sets the group's legend text.
func (g *checkboxGroup) SetLegendText(text string) {
	g.Root().SetAttribute("legend-text", text)
}

// Label returns the group's legend text.
func (g *checkboxGroup) Label() string {
	return g.LegendText()
}

// SetLabel sets the group's legend text and returns the receiver for chaining.
func (g *checkboxGroup) SetLabel(text string) *checkboxGroup {
	g.SetLegendText(text)
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
func (g *checkboxGroup) SetOrientation(orientation CheckboxOrientation) {
	g.Root().SetAttribute("orientation", string(orientation))
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
