package bs

import (
	"fmt"
	"strings"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Position defines the position for borders and alignment
type Position uint

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	Top Position = 1 << iota
	Bottom
	Start
	End
	Center
	Middle
	None Position = 0
)

const (
	// Combined positions
	All = Top | Bottom | Start | End
	X   = Start | End
	Y   = Top | Bottom
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (position Position) className(prefix string) string {
	switch position {
	case All:
		if prefix == "p" || prefix == "m" {
			return prefix + "-"
		}
		if prefix == "border" {
			return prefix
		}
		return prefix + "-all"
	case X:
		if prefix == "p" || prefix == "m" {
			return prefix + "x-"
		}
		return prefix + "-x"
	case Y:
		if prefix == "p" || prefix == "m" {
			return prefix + "y-"
		}
		return prefix + "-y"
	case Top:
		if prefix == "p" || prefix == "m" {
			return prefix + "t-"
		}
		return prefix + "-top"
	case Bottom:
		if prefix == "p" || prefix == "m" {
			return prefix + "b-"
		}
		return prefix + "-bottom"
	case Start:
		if prefix == "p" || prefix == "m" {
			return prefix + "s-"
		}
		return prefix + "-start"
	case End:
		if prefix == "p" || prefix == "m" {
			return prefix + "e-"
		}
		return prefix + "-end"
	case Center:
		return prefix + "-center"
	case Middle:
		return prefix + "-middle"
	default:
		return ""
	}
}

func (position Position) classNameWithSize(prefix string, size int) (string, error) {
	className := position.className(prefix)
	if className == "" {
		return "", fmt.Errorf("position.classNameWithSize: invalid position")
	}
	if size < 0 {
		return className + fmt.Sprintf("n%d", -size), nil
	}
	return className + fmt.Sprintf("%d", size), nil
}

func (position Position) classNames(prefix string) []string {
	className := position.className(prefix)
	if className == "" {
		panic("position.classNames: invalid position")
	}
	return []string{className}
}

func (position Position) allClassNames(prefix string) []string {
	classNames := make([]string, 0, 10)
	for p := Top; p <= End; p = p << 1 {
		classNames = append(classNames, p.className(prefix))
	}
	return append(classNames, All.className(prefix), X.className(prefix), Y.className(prefix))
}

func borderPrefixForView(name string) string {
	switch name {
	case ViewContainer, ViewBadge, ViewAlert, ViewCodeBlock, ViewNavbar:
		return "border"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithMargin adds margin to the specified position (Top, Bottom, Start, End, X, Y or All)
// with an optional size (-5 ... 5)
func WithMargin(position Position, size int) mvc.Opt {
	return func(o mvc.OptSet) error {
		// Remove existing margin classes for this specific position only
		classPrefix := position.className("m")
		if classPrefix == "" {
			return ErrInternalAppError.Withf("WithMargin: invalid position")
		}

		// Remove any existing classes with this exact prefix
		for _, class := range o.Classes() {
			if strings.HasPrefix(class, classPrefix) {
				if err := mvc.WithoutClass(class)(o); err != nil {
					return err
				}
			}
		}

		if size >= -5 && size <= 5 {
			if className, err := position.classNameWithSize("m", size); err != nil {
				return err
			} else {
				return mvc.WithClass(className)(o)
			}
		} else {
			return ErrInternalAppError.Withf("WithMargin: invalid size %d for view %q", size, o.Name())
		}
	}
}

// WithPadding adds padding to the specified position (Top, Bottom, Start, End, X, Y or All)
// Padding is an integer from 1 to 5
func WithPadding(position Position, size int) mvc.Opt {
	return func(o mvc.OptSet) error {
		// Remove existing padding classes for this specific position only
		classPrefix := position.className("p")
		if classPrefix == "" {
			return ErrInternalAppError.Withf("WithPadding: invalid position")
		}

		// Remove any existing classes with this exact prefix
		for _, class := range o.Classes() {
			if strings.HasPrefix(class, classPrefix) {
				if err := mvc.WithoutClass(class)(o); err != nil {
					return err
				}
			}
		}

		if size >= 1 && size <= 5 {
			if className, err := position.classNameWithSize("p", size); err != nil {
				return err
			} else {
				return mvc.WithClass(className)(o)
			}
		} else {
			return ErrInternalAppError.Withf("WithPadding: invalid size %d for view %q", size, o.Name())
		}
	}
}

// WithBorder adds a border to the specified position (Top, Bottom, Start, End, X, Y or All)
// with an optional color
func WithBorder(position Position, color ...Color) mvc.Opt {
	applyBorder := func(o mvc.OptSet, position Position) error {
		prefix := borderPrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithBorder: position: unsupported view %q", o.Name())
		}

		// Remove existing border classes for this specific position only
		classPrefix := position.className(prefix)
		if classPrefix == "" {
			return ErrInternalAppError.Withf("WithBorder: invalid position")
		}

		// Remove any existing classes with this exact prefix
		for _, class := range o.Classes() {
			if strings.HasPrefix(class, classPrefix) && (class == classPrefix || class[len(classPrefix)] == '-') {
				if err := mvc.WithoutClass(class)(o); err != nil {
					return err
				}
			}
		}

		// Add new border class
		if err := mvc.WithClass(position.classNames(prefix)...)(o); err != nil {
			return err
		}
		return nil
	}
	applyColor := func(o mvc.OptSet, color Color) error {
		prefix := "border"
		if prefix == "" {
			return ErrInternalAppError.Withf("WithBorder: color: unsupported view %q", o.Name())
		} else if err := mvc.WithoutClass(color.allClassNames(prefix)...)(o); err != nil {
			return err
		} else if err := mvc.WithClass(color.className(prefix))(o); err != nil {
			return err
		}

		// Add the 'border' class if not a NavBar
		if o.Name() != ViewNavbar {
			if err := mvc.WithClass("border")(o); err != nil {
				return err
			}
		}

		return nil
	}
	return func(o mvc.OptSet) error {
		// Apply border position
		if err := applyBorder(o, position); err != nil {
			return err
		}

		// Apply border color if specified
		if len(color) == 1 && color[0] != TRANSPARENT {
			if err := applyColor(o, color[0]); err != nil {
				return err
			}
		} else if len(color) > 1 {
			return ErrInternalAppError.Withf("WithBorder: multiple colors specified for view %q", o.Name())
		}

		// Return success
		return nil
	}
}

// WithAlign adds alignment to a view.
//
// Note: This function only supports views of type ViewFigureCaption and ViewContainer.
// Calling WithAlign on other view types will result in an error.
func WithAlign(position Position) mvc.Opt {
	return func(o mvc.OptSet) error {
		if o.Name() != ViewFigureCaption && o.Name() != ViewContainer {
			return fmt.Errorf("WithAlign: invalid view type %q", o.Name())
		}
		prefix := "text"
		if err := mvc.WithoutClass(position.allClassNames(prefix)...)(o); err != nil {
			return err
		}
		if err := mvc.WithClass(position.className(prefix))(o); err != nil {
			return err
		}
		return nil
	}
}
