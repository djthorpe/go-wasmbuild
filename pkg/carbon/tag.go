package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tag struct{ base }

var _ mvc.View = (*tag)(nil)
var _ mvc.EnabledState = (*tag)(nil)
var _ mvc.ActiveState = (*tag)(nil)
var _ mvc.VisibleState = (*tag)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	})
}

// Tag returns a <cds-tag> web component.
func Tag(args ...any) *tag {
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewTag, "cds-tag", setView, args).(*tag)
}

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
		property := value.Get(name)
		if !property.IsUndefined() && !property.IsNull() {
			return property.Bool()
		}
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
