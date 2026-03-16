package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tag struct {
	mvc.View
}

var _ mvc.View = (*tag)(nil)

// TagType controls the colour of a Carbon tag.
type TagType string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	TagGray     TagType = "gray"
	TagCoolGray TagType = "cool-gray"
	TagWarmGray TagType = "warm-gray"
	TagRed      TagType = "red"
	TagMagenta  TagType = "magenta"
	TagPurple   TagType = "purple"
	TagBlue     TagType = "blue" // default
	TagCyan     TagType = "cyan"
	TagTeal     TagType = "teal"
	TagGreen    TagType = "green"
	TagOutline  TagType = "outline"
	// Semantic
	TagHighContrast TagType = "high-contrast"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, func(self, child mvc.View) {
			self.(*tag).View = child
		})
	})
}

// Tag returns a <cds-tag> with the given type and label text.
//
//	cds.Tag("Draft")
//	cds.Tag("Error", cds.WithTagType(cds.TagRed))
//	cds.Tag("v1.0", cds.WithTagType(cds.TagGreen), cds.WithTagFilter())
func Tag(args ...any) *tag {
	return mvc.NewView(new(tag), ViewTag, "cds-tag", func(self, child mvc.View) {
		self.(*tag).View = child
	}, args).(*tag)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTagType sets the colour palette of a cds-tag.
func WithTagType(t TagType) mvc.Opt {
	return mvc.WithAttr("type", string(t))
}

// WithTagFilter makes the tag interactive (clickable/dismissible).
// The tag will emit a cds-tag-closed event when the × button is clicked.
func WithTagFilter() mvc.Opt {
	return mvc.WithAttr("filter", "")
}

// WithTagSize sets the tag to the small variant.
func WithTagSmall() mvc.Opt {
	return mvc.WithAttr("size", "sm")
}
