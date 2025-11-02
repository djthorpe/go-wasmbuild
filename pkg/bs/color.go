package bs

import (
	// Packages

	"slices"

	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

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
	TRANSPARENT      Color = ""
	PRIMARY          Color = "primary"
	PRIMARY_SUBTLE   Color = "primary-subtle"
	SECONDARY        Color = "secondary"
	SECONDARY_SUBTLE Color = "secondary-subtle"
	SUCCESS          Color = "success"
	SUCCESS_SUBTLE   Color = "success-subtle"
	DANGER           Color = "danger"
	DANGER_SUBTLE    Color = "danger-subtle"
	WARNING          Color = "warning"
	WARNING_SUBTLE   Color = "warning-subtle"
	INFO             Color = "info"
	INFO_SUBTLE      Color = "info-subtle"
	LIGHT            Color = "light"
	LIGHT_SUBTLE     Color = "light-subtle"
	DARK             Color = "dark"
	DARK_SUBTLE      Color = "dark-subtle"
	WHITE            Color = "white"
	BLACK            Color = "black"
)

var (
	allColors = []Color{
		PRIMARY,
		PRIMARY_SUBTLE,
		SECONDARY,
		SECONDARY_SUBTLE,
		SUCCESS,
		SUCCESS_SUBTLE,
		DANGER,
		DANGER_SUBTLE,
		WARNING,
		WARNING_SUBTLE,
		INFO,
		INFO_SUBTLE,
		LIGHT,
		LIGHT_SUBTLE,
		DARK,
		DARK_SUBTLE,
		WHITE,
		BLACK,
	}
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (color Color) className(prefix string) string {
	if color == TRANSPARENT {
		return prefix
	}
	return prefix + "-" + string(color)
}

func (color Color) allClassNames(prefix string) []string {
	classNames := make([]string, 0, len(allColors))
	for _, c := range allColors {
		classNames = append(classNames, c.className(prefix))
	}
	return classNames
}

func colorPrefixForView(name string) string {
	switch name {
	case ViewHeading, ViewContainer:
		return "text"
	case ViewText:
		return "text"
	case ViewBadge:
		return "text-bg"
	case ViewButton:
		return "btn"
	case ViewLink:
		return "link"
	case ViewAlert:
		return "alert"
		//	case ViewNavbar:
		//		return "bg"
	default:
		return ""
	}
}

func backgroundColorPrefixForView(name string) string {
	switch name {
	//	case ViewContainer:
	//		return "bg"
	case ViewBadge:
		return "text-bg"
	case ViewButton:
		return "btn"
	case ViewAlert:
		return "alert"
		//	case ViewNavbar:
		//		return "bg"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithColor(color Color) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := colorPrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithColor: unsupported view %q", o.Name())
		} else if o.Name() == ViewButton {
			// For outline buttons, adjust prefix
			if slices.Contains(o.Classes(), viewOutlineButtonClassPrefix) {
				prefix = viewOutlineButtonClassPrefix
			}
		}

		// Remove all other color classes
		if err := mvc.WithoutClass(color.allClassNames(prefix)...)(o); err != nil {
			return err
		}

		// Add class for this color
		if err := mvc.WithClass(color.className(prefix))(o); err != nil {
			return err
		}

		return nil
	}
}

func WithBackground(color Color) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := backgroundColorPrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithBackground: unsupported view %q", o.Name())
		} else if o.Name() == ViewButton {
			// For outline buttons, adjust prefix
			if slices.Contains(o.Classes(), viewOutlineButtonClassPrefix) {
				prefix = viewOutlineButtonClassPrefix
			}
		}

		// Remove all other background color classes
		if err := mvc.WithoutClass(color.allClassNames(prefix)...)(o); err != nil {
			return err
		}

		// Add class for this background color
		if err := mvc.WithClass(color.className(prefix))(o); err != nil {
			return err
		}

		return nil
	}
}

func WithTheme(color Color) mvc.Opt {
	return func(o mvc.OptSet) error {
		//		if o.Name() != ViewNavbar && o.Name() != ViewNavDropdown {
		//			return ErrInternalAppError.Withf("WithTheme: unsupported view %q", o.Name())
		//		}
		return mvc.WithAttr("data-bs-theme", string(color))(o)
	}
}
