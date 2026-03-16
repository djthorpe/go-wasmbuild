package carbon

import (
	"fmt"

	// Package imports
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type text struct {
	mvc.View
}

type heading struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var headingLevels = map[int]struct {
	tag   string
	class string
}{
	1: {"H1", "cds--heading-06"},
	2: {"H2", "cds--heading-05"},
	3: {"H3", "cds--heading-04"},
	4: {"H4", "cds--heading-03"},
	5: {"H5", "cds--heading-02"},
	6: {"H6", "cds--heading-01"},
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewText, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(text), element, func(self, child mvc.View) {
			self.(*text).View = child
		})
	})
	mvc.RegisterView(ViewHeading, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(heading), element, func(self, child mvc.View) {
			self.(*heading).View = child
		})
	})
}

// Heading returns a heading at the given level (1–6) using the corresponding
// Carbon productive-heading type token.
func Heading(level int, args ...any) mvc.View {
	h, exists := headingLevels[level]
	if !exists {
		panic(fmt.Sprintf("Heading: invalid level %d", level))
	}
	return mvc.NewView(new(heading), ViewHeading, h.tag, func(self, child mvc.View) {
		self.(*heading).View = child
	}, mvc.WithClass(h.class), args)
}

// Para returns a standard body paragraph using the Carbon body-01 type style
// (14px, line-height 1.43 — suitable for reading longer passages).
func Para(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--body-01"), args).(*text)
}

// LeadPara returns a larger body paragraph using the Carbon body-02 type style
// (16px, line-height 1.5 — for introductory or prominent text).
func LeadPara(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--body-02"), args).(*text)
}

// CompactPara returns a compact body paragraph using the Carbon body-compact-01
// type style (14px, line-height 1.29 — for inline UI labels and tight layouts).
func CompactPara(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--body-compact-01"), args).(*text)
}

// Caption returns a caption using the Carbon caption-01 type style
// (12px — for image captions and secondary annotations).
func Caption(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--caption-01"), args).(*text)
}

// HelperText returns helper text using the Carbon helper-text-01 type style
// (12px italic — for form field hints and contextual guidance).
func HelperText(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--helper-text-01"), args).(*text)
}

// LabelText returns a label using the Carbon label-01 type style
// (12px — for field labels and tight UI text).
func LabelText(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "P", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--label-01"), args).(*text)
}

// Deleted wraps content in a <del> element.
func Deleted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "DEL", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

// Highlighted wraps content in a <mark> element.
func Highlighted(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "MARK", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

// Smaller wraps content in a <small> element.
func Smaller(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "SMALL", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

// Strong wraps content in a <strong> element.
func Strong(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "STRONG", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

// Em wraps content in an <em> element.
func Em(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "EM", func(self, child mvc.View) {
		self.(*text).View = child
	}, args).(*text)
}

// Code wraps content in a <code> element using the Carbon code-01 type style.
func Code(args ...any) *text {
	return mvc.NewView(new(text), ViewText, "CODE", func(self, child mvc.View) {
		self.(*text).View = child
	}, mvc.WithClass("cds--code-01"), args).(*text)
}
