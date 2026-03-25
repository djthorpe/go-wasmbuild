package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type checkboxGroup struct{ base }

var _ mvc.View = (*checkboxGroup)(nil)
var _ mvc.ActiveGroup = (*checkboxGroup)(nil)
var _ mvc.EnabledGroup = (*checkboxGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewCheckboxGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(checkboxGroup), element, setView)
	}, EventChange)
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
// PUBLIC METHODS

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
	for _, view := range views {
		if view != nil {
			active[view.Root()] = struct{}{}
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
