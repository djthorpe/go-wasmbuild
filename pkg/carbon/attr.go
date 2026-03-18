package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Attr is a typed attribute value applied to Carbon components
// (kind=, size=, data-carbon-theme=, etc.).
type Attr string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

// Button / general component kinds
const (
	KindPrimary        Attr = "primary"
	KindSecondary      Attr = "secondary"
	KindTertiary       Attr = "tertiary"
	KindGhost          Attr = "ghost"
	KindDanger         Attr = "danger"
	KindDangerTertiary Attr = "danger-tertiary"
	KindDangerGhost    Attr = "danger-ghost"
)

// Notification / status kinds (also used by Tag, InlineNotification, etc.)
const (
	KindSuccess    Attr = "success"
	KindInfo       Attr = "info"
	KindInfoSquare Attr = "info-square"
	KindWarning    Attr = "warning"
	KindWarningAlt Attr = "warning-alt"
	KindError      Attr = "error"
)

// Tag types.
const (
	TagRed          Attr = "red"
	TagMagenta      Attr = "magenta"
	TagPurple       Attr = "purple"
	TagBlue         Attr = "blue"
	TagCyan         Attr = "cyan"
	TagTeal         Attr = "teal"
	TagGreen        Attr = "green"
	TagGray         Attr = "gray"
	TagCoolGray     Attr = "cool-gray"
	TagWarmGray     Attr = "warm-gray"
	TagHighContrast Attr = "high-contrast"
	TagOutline      Attr = "outline"
)

// Component sizes
const (
	SizeExtraSmall Attr = "xs"
	SizeSmall      Attr = "sm"
	SizeMedium     Attr = "md"
	SizeLarge      Attr = "lg"
	SizeExtraLarge Attr = "xl"
	Size2XLarge    Attr = "2xl"
)

// Icon sizes.
const (
	IconSize16 Attr = "16"
	IconSize20 Attr = "20"
	IconSize24 Attr = "24"
	IconSize32 Attr = "32"
)

// Checkbox group orientations.
const (
	CheckboxOrientationHorizontal Attr = "horizontal"
	CheckboxOrientationVertical   Attr = "vertical"
)

// Carbon colour themes — applied as CSS class (.cds--white, .cds--g10, etc.).
const (
	ThemeWhite Attr = "white" // default light theme → .cds--white
	ThemeG10   Attr = "g10"   // light grey           → .cds--g10
	ThemeG90   Attr = "g90"   // dark grey             → .cds--g90
	ThemeG100  Attr = "g100"  // near-black            → .cds--g100
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE

// attrKey maps each Attr value to the HTML attribute it controls.
// Sizes → "size", checkbox orientations → "orientation", everything else → "kind".
// Themes are CSS classes — use ClassForTheme, not With.
var attrKey = func() map[Attr]string {
	m := make(map[Attr]string)
	for _, k := range []Attr{SizeExtraSmall, SizeSmall, SizeMedium, SizeLarge, SizeExtraLarge, Size2XLarge, IconSize16, IconSize20, IconSize24, IconSize32} {
		m[k] = "size"
	}
	for _, k := range []Attr{TagRed, TagMagenta, TagPurple, TagBlue, TagCyan, TagTeal, TagGreen, TagGray, TagCoolGray, TagWarmGray, TagHighContrast, TagOutline} {
		m[k] = "type"
	}
	for _, k := range []Attr{CheckboxOrientationHorizontal, CheckboxOrientationVertical} {
		m[k] = "orientation"
	}
	return m
}()

var themeAttrs = map[Attr]struct{}{
	ThemeWhite: {}, ThemeG10: {}, ThemeG90: {}, ThemeG100: {},
}

func keyForAttr(a Attr) string {
	if key, ok := attrKey[a]; ok {
		return key
	}
	return "kind"
}

///////////////////////////////////////////////////////////////////////////////
// PREDICATES

// IsComponentKind reports whether a is a component kind (kind= attribute).
func IsComponentKind(a Attr) bool { return keyForAttr(a) == "kind" }

// IsSize reports whether a is a size value (size= attribute).
func IsSize(a Attr) bool { return keyForAttr(a) == "size" }

// IsTheme reports whether a is a theme (applied as CSS class, not attribute).
func IsTheme(a Attr) bool {
	_, ok := themeAttrs[a]
	return ok
}

// ClassForTheme returns the CSS class name for a theme Attr (e.g. "cds--g100").
// Returns an empty string if a is not a theme Attr.
func ClassForTheme(a Attr) string {
	if _, ok := themeAttrs[a]; !ok {
		return ""
	}
	return "cds--" + string(a)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// WithBackground returns an mvc.Opt that sets a tile background to the given
// CSS colour value.
// The tile wrapper translates this into the --cds-layer custom property.
//
//	carbon.Tile(carbon.WithBackground("#d0e2ff"), carbon.Head(3, "Title"))
func WithBackground(color string) mvc.Opt {
	return mvc.WithAttr("data-carbon-layer", color)
}

// WithFill returns an mvc.Opt that makes a tile fill the width of its container.
func WithFill() mvc.Opt {
	return mvc.WithAttr("data-carbon-fill", "true")
}

// WithHeight returns an mvc.Opt that sets a fixed host height.
func WithHeight(height string) mvc.Opt {
	if height == "" {
		return mvc.WithoutAttr("data-carbon-height")
	}
	return mvc.WithAttr("data-carbon-height", height)
}

// With converts one or more Attr constants into mvc.Opt values and returns
// them as a []mvc.Opt slice. Because gatherOpts flattens []mvc.Opt, the
// result can be passed directly as an argument to any constructor:
//
//	carbon.Button(carbon.With(carbon.KindPrimary, carbon.SizeLarge), "Click me")
//	carbon.Tile(carbon.With(carbon.ThemeG90), "Dark tile")
//
// When a theme Attr is included, all other theme classes are removed first so
// that switching themes via Apply is always clean.
func With(attrs ...Attr) []mvc.Opt {
	opts := make([]mvc.Opt, 0, len(attrs))
	for _, a := range attrs {
		a := a
		if IsTheme(a) {
			for t := range themeAttrs {
				if t != a {
					opts = append(opts, mvc.WithoutClass(ClassForTheme(t)))
				}
			}
			opts = append(opts, mvc.WithClass(ClassForTheme(a)))
		} else {
			opts = append(opts, mvc.WithAttr(keyForAttr(a), string(a)))
		}
	}
	return opts
}
