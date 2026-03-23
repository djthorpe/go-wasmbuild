package carbon

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tag struct{ base }

type tagGroup struct{ base }

var _ mvc.View = (*tag)(nil)
var _ mvc.View = (*tagGroup)(nil)
var _ mvc.EnabledState = (*tag)(nil)
var _ mvc.ActiveState = (*tag)(nil)
var _ mvc.VisibleState = (*tag)(nil)
var _ mvc.EnabledGroup = (*tagGroup)(nil)
var _ mvc.ActiveGroup = (*tagGroup)(nil)
var _ mvc.VisibleGroup = (*tagGroup)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	})

	mvc.RegisterView(ViewTagGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tagGroup), element, setView)
	}, EventTagDismissibleClosed, EventTagOperationalSelected)

	mvc.RegisterView(ViewDismissibleTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	}, EventTagDismissibleClosed)

	mvc.RegisterView(ViewOperationalTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	}, EventTagOperationalSelected)
}

// Tag returns a <cds-tag> web component.
func Tag(args ...any) *tag {
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewTag, "cds-tag", setView, args).(*tag)
}

// TagGroup returns a container for one or more tags.
// Child tag events bubble to the group, allowing group-level observation.
func TagGroup(args ...any) *tagGroup {
	args = append([]any{mvc.WithStyle("display:flex;flex-wrap:wrap;align-items:center;gap:0.75rem")}, args...)
	return mvc.NewView(new(tagGroup), ViewTagGroup, "DIV", setView, args).(*tagGroup)
}

// DismissibleTag returns a <cds-dismissible-tag> web component.
// An optional leading string sets the text attribute.
func DismissibleTag(args ...any) *tag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewDismissibleTag, "cds-dismissible-tag", setView, args).(*tag)
}

// OperationalTag returns a <cds-operational-tag> web component.
// An optional leading string sets the text attribute.
func OperationalTag(args ...any) *tag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewOperationalTag, "cds-operational-tag", setView, args).(*tag)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TAG

func (t *tag) Enabled() bool {
	return !tagBoolProperty(t.Root(), "disabled")
}

func (t *tag) SetEnabled(enabled bool) mvc.View {
	setTagBoolProperty(t.Root(), "disabled", !enabled)
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DISMISSIBLE / OPERATIONAL TAGS

func (t *tag) Visible() bool {
	return tagBoolProperty(t.Root(), "open")
}

func (t *tag) SetVisible(visible bool) mvc.View {
	setTagBoolProperty(t.Root(), "open", visible)
	return t
}

func (t *tag) Active() bool {
	return tagBoolProperty(t.Root(), "selected")
}

func (t *tag) SetActive(active bool) mvc.View {
	setTagBoolProperty(t.Root(), "selected", active)
	return t
}

// Label returns the visible label for any tag variant.
// Plain tags use child content; dismissible and operational tags use the text attribute.
func (t *tag) Label() string {
	if value := strings.TrimSpace(t.Root().GetAttribute("text")); value != "" {
		return value
	}
	return strings.TrimSpace(t.Root().TextContent())
}

// SetLabel updates the visible label for any tag variant.
// Plain tags use child content; dismissible and operational tags use the text attribute.
func (t *tag) SetLabel(label string) *tag {
	if t.Name() == ViewDismissibleTag || t.Name() == ViewOperationalTag {
		t.Root().SetAttribute("text", label)
		return t
	}
	t.Content(label)
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TAG GROUP

// Content appends tags to the group. Panics if any arg is not a *tag.
func (g *tagGroup) Content(args ...any) mvc.View {
	children := make([]any, 0, len(args))
	for _, arg := range args {
		if t, ok := arg.(*tag); ok {
			children = append(children, t)
		} else {
			panic(fmt.Sprintf("TagGroup.Content: expected *tag, got %T", arg))
		}
	}
	return g.View.Content(children...)
}

// SetActive marks the specified tags active and deactivates the rest.
// With no arguments, all children are deactivated.
func (g *tagGroup) Active() []mvc.View {
	active := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Active() {
			active = append(active, t)
		}
	}
	return active
}

func (g *tagGroup) SetActive(views ...mvc.View) mvc.View {
	active := make(map[mvc.View]bool, len(views))
	for _, v := range views {
		active[v] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetActive(active[child])
		}
	}
	return g
}

// SetEnabled enables the specified tags and disables the rest.
// With no arguments, all children are disabled.
func (g *tagGroup) Enabled() []mvc.View {
	enabled := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Enabled() {
			enabled = append(enabled, t)
		}
	}
	return enabled
}

func (g *tagGroup) SetEnabled(views ...mvc.View) mvc.View {
	enabled := make(map[mvc.View]bool, len(views))
	for _, v := range views {
		enabled[v] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetEnabled(enabled[child])
		}
	}
	return g
}

// SetVisible makes the specified tags visible and hides the rest.
// With no arguments, all children are hidden.
func (g *tagGroup) Visible() []mvc.View {
	visible := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Visible() {
			visible = append(visible, t)
		}
	}
	return visible
}

func (g *tagGroup) SetVisible(views ...mvc.View) mvc.View {
	visible := make(map[mvc.View]bool, len(views))
	for _, v := range views {
		visible[v] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetVisible(visible[child])
		}
	}
	return g
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func normalizeTagTextArgs(args ...any) []any {
	if len(args) == 0 {
		return args
	}
	text, ok := args[0].(string)
	if !ok {
		return args
	}
	return append([]any{mvc.WithAttr("text", text)}, args[1:]...)
}

func normalizeTagArgs(args ...any) {
	for _, arg := range args {
		switch value := arg.(type) {
		case *icon:
			applyTagIconSlot(value)
		case []any:
			normalizeTagArgs(value...)
		}
	}
}

func applyTagIconSlot(icon *icon) {
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

func tagBoolProperty(element dom.Element, name string) bool {
	if value, ok := element.JSValue().(js.Value); ok && !value.IsUndefined() && !value.IsNull() {
		return value.Get(name).Bool()
	}
	return element.HasAttribute(name)
}

func setTagBoolProperty(element dom.Element, name string, value bool) {
	if node, ok := element.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Set(name, value)
	}
	if value {
		element.SetAttribute(name, "")
	} else {
		element.RemoveAttribute(name)
	}
}
