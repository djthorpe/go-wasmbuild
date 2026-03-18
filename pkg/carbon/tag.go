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

type dismissibleTag struct{ base }

type operationalTag struct{ base }

var _ mvc.View = (*tag)(nil)
var _ mvc.View = (*dismissibleTag)(nil)
var _ mvc.View = (*operationalTag)(nil)
var _ mvc.EnabledState = (*tag)(nil)
var _ mvc.EnabledState = (*dismissibleTag)(nil)
var _ mvc.EnabledState = (*operationalTag)(nil)
var _ mvc.ActiveState = (*operationalTag)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	}, EventTagBeingClosed, EventTagClosed)

	mvc.RegisterView(ViewDismissibleTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(dismissibleTag), element, setView)
	}, EventTagDismissibleBeingClosed, EventTagDismissibleClosed)

	mvc.RegisterView(ViewOperationalTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(operationalTag), element, setView)
	}, EventTagOperationalBeingSelected, EventTagOperationalSelected)
}

// Tag returns a <cds-tag> web component.
func Tag(args ...any) *tag {
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewTag, "cds-tag", setView, args).(*tag)
}

// DismissibleTag returns a <cds-dismissible-tag> web component.
// An optional leading string sets the text attribute.
func DismissibleTag(args ...any) *dismissibleTag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(dismissibleTag), ViewDismissibleTag, "cds-dismissible-tag", setView, args).(*dismissibleTag)
}

// OperationalTag returns a <cds-operational-tag> web component.
// An optional leading string sets the text attribute.
func OperationalTag(args ...any) *operationalTag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(operationalTag), ViewOperationalTag, "cds-operational-tag", setView, args).(*operationalTag)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TAG

func (t *tag) Enabled() bool {
	return !tagBoolProperty(t.Root(), "disabled")
}

func (t *tag) SetEnabled(enabled bool) *tag {
	setTagBoolProperty(t.Root(), "disabled", !enabled)
	return t
}

func (t *tag) Text() string {
	return strings.TrimSpace(t.Root().TextContent())
}

func (t *tag) SetText(text string) *tag {
	t.Content(text)
	return t
}

func (t *tag) Type() Attr {
	if value := t.Root().GetAttribute("type"); value != "" {
		return Attr(value)
	}
	return TagGray
}

func (t *tag) SetType(value Attr) *tag {
	t.Root().SetAttribute("type", string(value))
	return t
}

func (t *tag) Size() Attr {
	if value := t.Root().GetAttribute("size"); value != "" {
		return Attr(value)
	}
	return SizeMedium
}

func (t *tag) SetSize(value Attr) *tag {
	t.Root().SetAttribute("size", string(value))
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - DISMISSIBLE TAG

func (t *dismissibleTag) Enabled() bool {
	return !tagBoolProperty(t.Root(), "disabled")
}

func (t *dismissibleTag) SetEnabled(enabled bool) *dismissibleTag {
	setTagBoolProperty(t.Root(), "disabled", !enabled)
	return t
}

func (t *dismissibleTag) Open() bool {
	return tagBoolProperty(t.Root(), "open")
}

func (t *dismissibleTag) SetOpen(open bool) *dismissibleTag {
	setTagBoolProperty(t.Root(), "open", open)
	return t
}

func (t *dismissibleTag) Text() string {
	return t.Root().GetAttribute("text")
}

func (t *dismissibleTag) SetText(text string) *dismissibleTag {
	t.Root().SetAttribute("text", text)
	return t
}

func (t *dismissibleTag) Title() string {
	return t.Root().GetAttribute("tag-title")
}

func (t *dismissibleTag) SetTitle(text string) *dismissibleTag {
	t.Root().SetAttribute("tag-title", text)
	return t
}

func (t *dismissibleTag) Type() Attr {
	if value := t.Root().GetAttribute("type"); value != "" {
		return Attr(value)
	}
	return TagGray
}

func (t *dismissibleTag) SetType(value Attr) *dismissibleTag {
	t.Root().SetAttribute("type", string(value))
	return t
}

func (t *dismissibleTag) Size() Attr {
	if value := t.Root().GetAttribute("size"); value != "" {
		return Attr(value)
	}
	return SizeMedium
}

func (t *dismissibleTag) SetSize(value Attr) *dismissibleTag {
	t.Root().SetAttribute("size", string(value))
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - OPERATIONAL TAG

func (t *operationalTag) Enabled() bool {
	return !tagBoolProperty(t.Root(), "disabled")
}

func (t *operationalTag) SetEnabled(enabled bool) *operationalTag {
	setTagBoolProperty(t.Root(), "disabled", !enabled)
	return t
}

func (t *operationalTag) Active() bool {
	return tagBoolProperty(t.Root(), "selected")
}

func (t *operationalTag) SetActive(active bool) *operationalTag {
	setTagBoolProperty(t.Root(), "selected", active)
	return t
}

func (t *operationalTag) Text() string {
	return t.Root().GetAttribute("text")
}

func (t *operationalTag) SetText(text string) *operationalTag {
	t.Root().SetAttribute("text", text)
	return t
}

func (t *operationalTag) Type() Attr {
	if value := t.Root().GetAttribute("type"); value != "" {
		return Attr(value)
	}
	return TagGray
}

func (t *operationalTag) SetType(value Attr) *operationalTag {
	t.Root().SetAttribute("type", string(value))
	return t
}

func (t *operationalTag) Size() Attr {
	if value := t.Root().GetAttribute("size"); value != "" {
		return Attr(value)
	}
	return SizeMedium
}

func (t *operationalTag) SetSize(value Attr) *operationalTag {
	t.Root().SetAttribute("size", string(value))
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
