package bootstrap

import (
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Position defines the position for borders and alignment
type Position uint

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	TOP Position = 1 << iota
	BOTTOM
	START
	END
	CENTER
	MIDDLE
	NONE Position = 0
)

const (
	// All border positions
	BorderAll = TOP | BOTTOM | START | END

	// All margin positions
	MarginAll = TOP | BOTTOM | START | END

	// All padding positions
	PaddingAll = TOP | BOTTOM | START | END
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (position Position) borderClassNames() []string {
	prefix := "border"

	// In the case of all borders, return just the prefix
	if position == BorderAll {
		return []string{prefix}
	}

	// Return a list of class names
	classNames := []string{}
	for i := TOP; i <= END; i = i << 1 {
		if position&i != 0 {
			classNames = append(classNames, i.borderClassName(prefix))
		}
	}
	return classNames
}

func (position Position) className(prefix string) string {
	switch position {
	case TOP:
		return prefix + "top"
	case BOTTOM:
		return prefix + "bottom"
	case START:
		return prefix + "start"
	case END:
		return prefix + "end"
	default:
		return ""
	}
}

func (position Position) borderClassName(prefix string) string {
	suffix := position.className("")
	if suffix == "" {
		return ""
	}
	return prefix + "-" + suffix
}

func (position Position) marginClassNames(size int) []string {
	// Return a list of class names
	classNames := []string{}

	// Handle special case for all margins
	if position == MarginAll {
		return []string{formatMarginSize("m", size)}
	}

	// Track which positions have been handled by shorthands
	handledPositions := Position(0)

	// Handle shorthand for vertical (top & bottom)
	if (position & (TOP | BOTTOM)) == (TOP | BOTTOM) {
		classNames = append(classNames, formatMarginSize("my", size))
		handledPositions |= TOP | BOTTOM
	}

	// Handle shorthand for horizontal (start & end)
	if (position & (START | END)) == (START | END) {
		classNames = append(classNames, formatMarginSize("mx", size))
		handledPositions |= START | END
	}

	// Handle individual positions that weren't covered by shorthands
	for i := TOP; i <= END; i = i << 1 {
		if (position&i != 0) && (handledPositions&i == 0) {
			classNames = append(classNames, i.marginClassName(size))
		}
	}
	return classNames
}

func (position Position) marginClassName(size int) string {
	switch position {
	case TOP:
		return formatMarginSize("mt", size)
	case BOTTOM:
		return formatMarginSize("mb", size)
	case START:
		return formatMarginSize("ms", size)
	case END:
		return formatMarginSize("me", size)
	default:
		return ""
	}
}

// formatMarginSize formats a margin/padding size, handling negative values
// For negative values, uses "n" prefix (e.g., "mt-n1" instead of "mt--1")
func formatMarginSize(prefix string, size int) string {
	if size < 0 {
		return prefix + "-n" + strconv.Itoa(-size)
	}
	return prefix + "-" + strconv.Itoa(size)
}

func (position Position) paddingClassNames(size int) []string {
	// Return a list of class names
	classNames := []string{}

	// Handle special case for all padding
	if position == PaddingAll {
		return []string{formatPaddingSize("p", size)}
	}

	// Track which positions have been handled by shorthands
	handledPositions := Position(0)

	// Handle shorthand for vertical (top & bottom)
	if (position & (TOP | BOTTOM)) == (TOP | BOTTOM) {
		classNames = append(classNames, formatPaddingSize("py", size))
		handledPositions |= TOP | BOTTOM
	}

	// Handle shorthand for horizontal (start & end)
	if (position & (START | END)) == (START | END) {
		classNames = append(classNames, formatPaddingSize("px", size))
		handledPositions |= START | END
	}

	// Handle individual positions that weren't covered by shorthands
	for i := TOP; i <= END; i = i << 1 {
		if (position&i != 0) && (handledPositions&i == 0) {
			classNames = append(classNames, i.paddingClassName(size))
		}
	}
	return classNames
}

func (position Position) paddingClassName(size int) string {
	switch position {
	case TOP:
		return formatPaddingSize("pt", size)
	case BOTTOM:
		return formatPaddingSize("pb", size)
	case START:
		return formatPaddingSize("ps", size)
	case END:
		return formatPaddingSize("pe", size)
	default:
		return ""
	}
}

// formatPaddingSize formats a padding size (only positive values allowed)
func formatPaddingSize(prefix string, size int) string {
	return prefix + "-" + strconv.Itoa(size)
}
