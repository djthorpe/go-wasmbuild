package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func init() {
	mvc.RegisterView(ViewOperationalTag, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tag), element, setView)
	}, EventTagOperationalSelected)
}

// OperationalTag returns a <cds-operational-tag> web component.
// An optional leading string sets the text attribute.
func OperationalTag(args ...any) *tag {
	args = normalizeTagTextArgs(args...)
	normalizeTagArgs(args...)
	return mvc.NewView(new(tag), ViewOperationalTag, "cds-operational-tag", setView, args).(*tag)
}
