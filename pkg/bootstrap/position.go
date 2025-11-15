package bootstrap

import (
	"slices"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

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
	Sticky
	None Position = 0
)

const (
	// All border positions
	BorderAll = Top | Bottom | Start | End

	// All margin positions
	MarginAll = Top | Bottom | Start | End

	// All padding positions
	PaddingAll = Top | Bottom | Start | End

	// All Offcanvas positions
	OffcanvasAll = Start | End | Top | Bottom

	// All Navbar positions
	NavbarAll = Top | Bottom
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithFlex(position Position) mvc.Opt {
	return func(o mvc.OptSet) error {
		// Add or remove flex class
		if position != None {
			mvc.WithClass("d-flex")(o)
		} else {
			mvc.WithoutClass("d-flex")(o)
		}

		// Add or remove direction as row or column
		switch position {
		case Top, Bottom, Middle:
			mvc.WithoutClass("flex-row")(o)
			mvc.WithClass("flex-column")(o)
		case Start, End, Center:
			mvc.WithClass("flex-row")(o)
			mvc.WithoutClass("flex-column")(o)
		case None:
			mvc.WithoutClass("flex-row")(o)
			mvc.WithoutClass("flex-column")(o)
		default:
			return ErrInternalAppError.With("WithFlex: invalid position")
		}

		// TODO: Add alignment classes
		// justify-content-start, justify-content-center, justify-content-end

		// Return success
		return nil
	}

}

func WithPosition(position Position) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := positionPrefixForView(o.Name(), position)
		if prefix == "" {
			return ErrInternalAppError.Withf("WithPosition: unsupported view %q", o.Name())
		}

		// Remove all other classes
		classNames := allPositionClassNamesForView(o.Name())
		if err := mvc.WithoutClass(classNames...)(o); err != nil {
			return err
		}

		// If no position, return
		if position == None {
			return nil
		}

		// Add class for this position
		className := position.className(prefix)
		if !slices.Contains(classNames, className) {
			return ErrInternalAppError.Withf("WithPosition: invalid position %d for view %q", position, o.Name())
		}
		return mvc.WithClass(className)(o)
	}
}

func WithBorder(colors ...Color) mvc.Opt {
	// TODO: If there is one color, use it for all borders
	// If there are two, then use it for vertical and horizontal borders
	// If there are four, use it for each border individually (Top, Right, Bottom, Left)
	return func(o mvc.OptSet) error {
		// Add border class
		if err := mvc.WithClass("border")(o); err != nil {
			return err
		}

		// Remove all other border color classes
		prefix := borderPrefix()
		if err := mvc.WithoutClass(allColorClassNames(prefix)...)(o); err != nil {
			return err
		}

		// No border
		if len(colors) == 0 {
			return nil
		}

		// The single color use case
		if len(colors) == 1 {
			return mvc.WithClass(colors[0].className(prefix))(o)
		}

		// Not yet implemented
		return ErrInternalAppError.Withf("WithBorder: multi-border colors not yet implemented")
	}
}

func WithoutBorder() mvc.Opt {
	return func(o mvc.OptSet) error {
		// Remove border class
		if err := mvc.WithoutClass("border")(o); err != nil {
			return err
		}

		// Remove all other border color classes
		prefix := borderPrefix()
		if err := mvc.WithoutClass(allColorClassNames(prefix)...)(o); err != nil {
			return err
		}

		// Return success
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (position Position) Is(flag Position) bool {
	return position&flag == flag
}

func (position Position) className(prefix string) string {
	switch position & (Top | Bottom | Start | End | Center | Middle) {
	case Top:
		return prefix + "top"
	case Bottom:
		return prefix + "bottom"
	case Start:
		return prefix + "start"
	case End:
		return prefix + "end"
	case Center:
		return prefix + "center"
	case Middle:
		return prefix + "middle"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func borderPrefix() string {
	return "border"
}

func positionPrefixForView(name string, position Position) string {
	switch name {
	case ViewOffcanvas:
		return "offcanvas-"
	case ViewNavBar:
		if position.Is(Sticky) {
			return "sticky-"
		} else {
			return "fixed-"
		}
	case ViewText:
		return "text-"
	default:
		return ""
	}
}

func allPositionClassNamesForView(name string) []string {
	switch name {
	case ViewOffcanvas:
		return []string{
			Top.className("offcanvas-"),
			Bottom.className("offcanvas-"),
			Start.className("offcanvas-"),
			End.className("offcanvas-"),
		}
	case ViewNavBar:
		return []string{
			Top.className("fixed-"),
			Bottom.className("fixed-"),
			Top.className("sticky-"),
			Bottom.className("sticky-"),
		}
	case ViewText:
		return []string{
			Start.className("text-"),
			Center.className("text-"),
			End.className("text-"),
		}
	default:
		return nil
	}
}
