package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type icon struct{ base }

// IconName identifies an icon bundled into the Carbon example registry.
// The values match the underlying Carbon icon names.
type IconName string

// IconSize is the rendered size of a Carbon icon.
type IconSize = Attr

const carbonIconLookup = "goWasmBuildCarbonIcon"

// Common icon name constants for convenience.
const (
	IconAdd           IconName = "add"
	IconArrowRight    IconName = "arrow-right"
	IconCheckmark     IconName = "checkmark"
	IconClose         IconName = "close"
	IconFavorite      IconName = "favorite"
	IconLaunch        IconName = "launch"
	IconLogoGithub    IconName = "logo--github"
	IconSearch        IconName = "search"
	IconSettings      IconName = "settings"
	IconUserAvatar    IconName = "user--avatar"
	IconWarningFilled IconName = "warning--filled"
)

var _ mvc.View = (*icon)(nil)
var _ mvc.ValueState = (*icon)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewIcon, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(icon), element, setView)
	})
}

// Icon returns a <cds-icon> web component backed by the bundled icon registry.
// Only the icons exported by this package are guaranteed to resolve.
func Icon(name IconName, args ...any) *icon {
	i := mvc.NewView(new(icon), ViewIcon, "cds-icon", setView, args).(*icon)
	i.SetValue(string(name))
	return i
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (i *icon) Apply(opts ...mvc.Opt) mvc.View {
	i.View.Apply(opts...)
	if name := IconName(i.Value()); name != "" {
		setIconProperty(i.Root(), iconProperty(name, currentIconSize(i.Root())))
	}
	return i
}

// Value returns the current bundled icon name.
func (i *icon) Value() string {
	return i.Root().GetAttribute("data-carbon-icon")
}

// SetValue updates the icon glyph using a bundled icon name.
func (i *icon) SetValue(value string) mvc.View {
	root := i.Root()
	name := IconName(value)
	if name == "" {
		root.RemoveAttribute("data-carbon-icon")
		setIconProperty(root, js.Undefined())
		return i
	}
	root.SetAttribute("data-carbon-icon", string(name))
	setIconProperty(root, iconProperty(name, currentIconSize(root)))
	return i
}

// Label returns the icon's accessible name (aria-label).
func (i *icon) Label() string {
	return i.Root().GetAttribute("aria-label")
}

// SetLabel sets both the accessible name (aria-label) and tooltip text on the icon.
// It also removes aria-hidden so the icon is announced by assistive technology.
func (i *icon) SetLabel(label string) *icon {
	if label == "" {
		i.Apply(mvc.WithoutAttr("aria-label"), mvc.WithoutAttr("tooltip-text"))
		if i.Root().GetAttribute("slot") == "icon" {
			i.Root().SetAttribute("aria-hidden", "true")
		}
	} else {
		i.Apply(mvc.WithAriaLabel(label), mvc.WithAttr("tooltip-text", label), mvc.WithoutAttr("aria-hidden"))
	}
	return i
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func normalizeIconSize(size IconSize) IconSize {
	switch size {
	case IconSize16, IconSize20, IconSize24, IconSize32:
		return size
	default:
		return IconSize16
	}
}

func currentIconSize(root dom.Element) IconSize {
	return normalizeIconSize(IconSize(root.GetAttribute("size")))
}

func iconProperty(name IconName, size IconSize) js.Value {
	return js.Global().Call(carbonIconLookup, string(name), iconSizeNumber(size))
}

func iconSizeNumber(size IconSize) int {
	switch normalizeIconSize(size) {
	case IconSize20:
		return 20
	case IconSize24:
		return 24
	case IconSize32:
		return 32
	default:
		return 16
	}
}

func setIconProperty(element dom.Element, icon js.Value) {
	if node, ok := element.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Set("icon", icon)
	}
}
