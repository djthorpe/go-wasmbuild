package bootstrap

import (
	// Namespace imports
	"slices"

	. "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

// Size defines button sizes
type Size string

const (
	SizeDefault Size = ""
	SizeSmall   Size = "sm"
	SizeMedium  Size = "md"
	SizeLarge   Size = "lg"
	SizeXLarge  Size = "xl"
	SizeXXLarge Size = "xxl"
)

var (
	allSizes = []Size{
		SizeSmall,
		SizeMedium,
		SizeLarge,
		SizeXLarge,
		SizeXXLarge,
		"fluid",
	}
	allButtonSizes = []Size{
		SizeSmall,
		SizeLarge,
	}
)

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithSize(size Size) mvc.Opt {
	return func(o mvc.OptSet) error {
		prefix := sizePrefixForView(o.Name())
		if prefix == "" {
			return ErrInternalAppError.Withf("WithSize: unsupported view %q", o.Name())
		}

		// Remove all other size classes
		if err := mvc.WithoutClass(size.allClassNames(o.Name())...)(o); err != nil {
			return err
		}

		// Add class for this size
		if !slices.Contains(allSizesForView(o.Name()), size) {
			return ErrInternalAppError.Withf("WithSize: invalid size %q for view %q", size, o.Name())
		} else if err := mvc.WithClass(size.className(prefix))(o); err != nil {
			return err
		}

		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (size Size) className(prefix string) string {
	if size == SizeDefault {
		return prefix
	}
	return prefix + "-" + string(size)
}

func (size Size) allClassNames(name string) []string {
	// Get prefix
	prefix := sizePrefixForView(name)
	if prefix == "" {
		return nil
	}

	// Create class names
	classNames := make([]string, 0, 10)

	// For containers, include the base class since it gets replaced by size variants
	if name == ViewContainer {
		classNames = append(classNames, prefix)
	}

	// Add all size-specific classes
	for _, s := range allSizesForView(name) {
		classNames = append(classNames, s.className(prefix))
	}
	return classNames
}

func sizePrefixForView(name string) string {
	switch name {
	case ViewContainer:
		return "container"
	//	case ViewButton:
	//		return "btn"
	//	case ViewButtonGroup:
	//		return "btn-group"
	//	case ViewNavbar:
	//		return "navbar-expand"
	default:
		return ""
	}
}

func allSizesForView(name string) []Size {
	//	if name == ViewButton || name == ViewButtonGroup {
	//		return allButtonSizes
	//	}
	//if name == ViewNavbar {
	// Include SizeDefault for navbar
	//	return append([]Size{SizeDefault}, allSizes...)
	//}
	return allSizes
}
