package bootstrap

import (

	// Packages
	"slices"

	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Color defines the color for components and backgrounds
type Color string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	Transparent     Color = ""
	Primary         Color = "primary"
	PrimarySubtle   Color = "primary-subtle"
	Secondary       Color = "secondary"
	SecondarySubtle Color = "secondary-subtle"
	Success         Color = "success"
	SuccessSubtle   Color = "success-subtle"
	Danger          Color = "danger"
	DangerSubtle    Color = "danger-subtle"
	Warning         Color = "warning"
	WarningSubtle   Color = "warning-subtle"
	Info            Color = "info"
	InfoSubtle      Color = "info-subtle"
	Light           Color = "light"
	LightSubtle     Color = "light-subtle"
	Dark            Color = "dark"
	DarkSubtle      Color = "dark-subtle"
	White           Color = "white"
	Black           Color = "black"
)

var (
	allColors = []Color{
		Primary,
		PrimarySubtle,
		Secondary,
		SecondarySubtle,
		Success,
		SuccessSubtle,
		Danger,
		DangerSubtle,
		Warning,
		WarningSubtle,
		Info,
		InfoSubtle,
		Light,
		LightSubtle,
		Dark,
		DarkSubtle,
		White,
		Black,
	}
	themeColors = []Color{
		Light,
		Dark,
	}
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithColor(color Color) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := colorPrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithColor: unsupported view %q", o.Name())
		}

		if o.Name() == ViewButton {
			// For outline buttons, adjust prefix
			if slices.Contains(o.Classes(), viewOutlineButtonClassPrefix) {
				prefix = viewOutlineButtonClassPrefix
			}
		}

		// Remove all other color classes
		if err := mvc.WithoutClass(allColorClassNames(prefix)...)(o); err != nil {
			return err
		}

		// Add class for this color
		if err := mvc.WithClass(color.className(prefix))(o); err != nil {
			return err
		}

		return nil
	}
}

func WithTheme(color Color) mvc.Opt {
	return func(o mvc.OptSet) error {
		if !slices.Contains(themeColors, color) {
			return ErrInternalAppError.Withf("WithThemeColor: invalid theme %q", color)
		}
		return mvc.WithAttr("data-bs-theme", string(color))(o)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (color Color) className(prefix string) string {
	if color == Transparent {
		return prefix
	}
	return prefix + "-" + string(color)
}

func allColorClassNames(prefix string) []string {
	classNames := make([]string, 0, len(allColors))
	for _, c := range allColors {
		classNames = append(classNames, c.className(prefix))
	}
	return classNames
}

func colorPrefixForView(name string) string {
	switch name {
	case ViewBadge:
		return "text-bg"
	case ViewCard:
		return "text-bg"
	case ViewLink:
		return "link"
	case ViewButton:
		return "btn"
	case ViewIcon:
		return "text"
	case ViewCodeBlock:
		return "bg"
	case ViewText:
		return "text"
	case ViewProgress:
		return "progress-bar"
	case ViewContainer:
		return "bg"
	case ViewTable, ViewTableRow:
		return "table"
	//case ViewAlert:
	//	return "alert"
	//case ViewNavbar:
	//		return "text"
	//	case ViewAlert:
	//		return "alert"
	//	case ViewNavbar:
	default:
		return ""
	}
}
