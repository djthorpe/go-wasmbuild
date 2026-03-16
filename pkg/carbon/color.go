package carbon

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Theme represents a Carbon Design System colour palette.
// Themes work by scoping all Carbon CSS custom-property tokens to the element
// they are applied to, so every child component inherits the palette.
type Theme string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// ThemeWhite is the default light theme.
	ThemeWhite Theme = "cds--white"

	// ThemeG10 is a subtle light-gray variant of the light theme.
	ThemeG10 Theme = "cds--g10"

	// ThemeG90 is the standard dark theme.
	ThemeG90 Theme = "cds--g90"

	// ThemeG100 is the near-black dark theme, commonly used for the UI shell.
	ThemeG100 Theme = "cds--g100"
)

var allThemeClasses = []string{
	string(ThemeWhite),
	string(ThemeG10),
	string(ThemeG90),
	string(ThemeG100),
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTheme applies a Carbon colour theme to an element and all its descendants.
// Theme classes scope Carbon CSS custom properties so child components inherit
// the palette automatically.
//
//	cds.Header(cds.WithTheme(cds.ThemeG100), ...)
//	cds.SideNav(cds.WithTheme(cds.ThemeG90), ...)
func WithTheme(t Theme) mvc.Opt {
	return func(o mvc.OptSet) error {
		if err := mvc.WithoutClass(allThemeClasses...)(o); err != nil {
			return err
		}
		return mvc.WithClass(string(t))(o)
	}
}
