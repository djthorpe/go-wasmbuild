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
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithPosition(position Position) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := positionPrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithPosition: unsupported view %q", o.Name())
		}

		// Remove all other classes
		classNames := allPositionClassNamesForView(o.Name())
		if err := mvc.WithoutClass(classNames...)(o); err != nil {
			return err
		}

		// Add class for this position
		className := position.className(prefix)
		if !slices.Contains(classNames, className) {
			return ErrInternalAppError.Withf("WithPosition: invalid position %d for view %q", position, o.Name())
		}

		return mvc.WithClass(className)(o)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func positionPrefixForView(name string) string {
	switch name {
	case ViewOffcanvas:
		return "offcanvas-"
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
	default:
		return nil
	}
}

func (position Position) className(prefix string) string {
	switch position {
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
